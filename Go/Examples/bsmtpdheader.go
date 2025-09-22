/*
Package bsmtpdheader handles translating the many headers emitted by
bsmtp into proper actions.
*/
package bsmtpdheader

// re-read the repdef if the nsrep.dat file changes

import (
	"errors"
	"log/slog"
	"strings"

	"github.com/barracuda-internal/bnsf-artfuldice/data"
	"github.com/barracuda-internal/bnsf-artfuldice/data/perrecip"
	"github.com/barracuda-internal/bnsf-artfuldice/db"
	"github.com/barracuda-internal/bnsf-artfuldice/esmtp/headers"
	"github.com/barracuda-internal/bnsf-artfuldice/esmtp/message"
	"github.com/barracuda-internal/bnsf-artfuldice/esmtp/protosupport"
	"github.com/barracuda-internal/bnsf-artfuldice/jobs"
	"github.com/barracuda-internal/bnsf-artfuldice/jobs/disposition"
	"github.com/barracuda-internal/bnsf-artfuldice/mail"
	"github.com/barracuda-internal/bnsf-artfuldice/metare"
	"github.com/barracuda-internal/bnsf-artfuldice/metrics"
	"github.com/barracuda-internal/bnsf-artfuldice/userlist"
	"github.com/thejerf/cm"
)

const (
	// LogUnknownReason is logged if the reason in the parens is
	// unknown. This is also a coding error and should never be
	// encountered in production.
	LogUnknownReason = "unknown reason text"

	// LogUnknownAction is logged if the hit type could not be
	// determined
	LogUnknownAction = "unknown action"

	// LogDBError is logged at error if we couldn't get per-user
	// policies.
	LogDBError = "error accessing per-user policies"

	// LogFiltered is logged at debug when a rule is skipped
	// due to not containing the requisite text.
	LogFiltered = "rule skipped because filter clause did not match"

	// LogHeaderHit is logged at info when a header is found.
	LogHeaderHit = "hit on a header"

	// LogInvalidHeaders is logged at Error when invalid headers
	// are found.
	LogInvalidHeaders = "invalid list headers found"

	// LogBlockingOnNotAllowListed is logged at Info when the
	// message is blocked due to BlockIfNotWhitelisted and there
	// were no allowlists
	LogBlockingOnNotAllowListed = "received BlockIfNotWhitelisted header, non-allowed recipient"

	// LogInvalidRE is logged at Error if a user's regular expression is invalid.
	LogInvalidRE = "invalid regular expression"

	// LogInvalidREPostTrim is logged at Error if a user's regular
	// expression couldn't be compiled.
	LogInvalidREPostTrim = "invalid regular expression (after slash trim)"

	// LogAllowListed is logged when an email is passed through
	// because it is allow listed.
	LogAllowListed = "email allowlisted"

	// LogBlocked is logged when a mail is blocked by the
	// per-domain IP block.
	LogBlocked = "email blocked"

	// LogQuarantined is logged when a mail is blocked by per-domain IP block
	LogQuarantined = "email quarantined"

	// LogTagged is logged when a mail is tagged by a per-domain IP block
	LogTagged = "email tagged"

	// LogListsChecked is logged when a mail is not affected by
	// any lists.
	LogListsChecked = "email not affected by any block/allow lists"

	// LogRBLHit is logged at info when we have found an RBL header.
	LogRBLHit = "rbl hit"

	// LogRBLNotConfigured is logged at error if bsmtpd claims an
	// RBL hit that we don't have configuration for.
	LogRBLNotConfigured = "rbl does not have an action configured"

	// LogBRBLHitFound is logged at info if there is a BRBL hit.
	LogBRBLHitFound = "BRBL hit"

	// in the X-ASG-Whitelist header, this marks a recipient allow
	// based on regular experessions.
	recipientRegexpMarker = "Recipient"

	// LogDatabaseError is logged at Error when there is an error
	// using the text file databases.
	LogDatabaseError = "error access allow/block databases"
)
const (
	// LogNoSPFResult is logged it info if there was no SPF
	// header.
	LogNoSPFResult = "no Received-SPF header so no SPF handling"

	// LogUnknownSPFResult is logged at error if the SPF result
	// could not have its first word trimmed off. This shouldn't
	// be possible.
	LogUnknownSPFResult = "unknown Received-SPF header"

	// LogKnownButUnusedResult is logged at Debug if the SPF
	// result is known, but not one we take action on.
	LogKnownButUnusedResult = "known, but unused, SPF result"

	// LogValidSPFFound is logged at info if the SPF result is a
	// pass.
	LogValidSPFFound = "Received-SPF has a passing SPF check"
)

var (
	spfHeader = headers.MustCanonical("Received-SPF")

	// SPFFailKey is set to true if the check soft or hard fails.
	SPFFailKey = SPFKey{"spfkey"}
)

// SPFKey is a key type for storing the results of the SPF check.
type SPFKey struct {
	data.Key[bool]
}

var (
	allowHeader            = headers.MustCanonical("X-ASG-Whitelist")
	encryptHeader          = headers.MustCanonical("X-Barracuda-Encrypt")
	fingerprintBlockHeader = headers.MustCanonical("X-Barracuda-Fingerprint-Block")
	quarantineHeader       = headers.MustCanonical("X-ASG-Quarantine")
	rblBlockHeader         = headers.MustCanonical("X-Barracuda-RBL-Block")
	recipientAllow         = headers.MustCanonical("X-ASG-Recipient-Whitelist")
	redirectHeader         = headers.MustCanonical("X-Barracuda-Redirected")
	registryHeader         = headers.MustCanonical("X-Barracuda-Registry")
	repBlockHeader         = headers.MustCanonical("X-Barracuda-Reputation-Block")
	tagHeader1             = headers.MustCanonical("X-ASG-Tag")
	// see mod_content in bsmtpd
	tagHeader2          = headers.MustCanonical("X-ASG")
	userWhitelistHeader = headers.MustCanonical("X-Barracuda-User-Whitelist")
	rblRestriction      = headers.MustCanonical("X-ASG-RBL-Restriction")
	quarantineRBL       = headers.MustCanonical("X-ASG-Quarantine-RBL")
	quarantineBRL       = headers.MustCanonical("X-Barracuda-BRL-Quarantine")
	tagBRL              = headers.MustCanonical("X-Barracuda-BRL-Tag")
	// These headers seem to have no way for bsmtpd to emit them,
	// so I do not know what their contents may be. Amavis is
	// copied in order to match if there are devices in the field
	// that may emit these.
	blockHeader = headers.MustCanonical("X-ASG-Block")
	ipddHeader  = headers.MustCanonical("X-Barracuda-IPDD")

	// PerUserBypassedKey is used by a later job to generate a
	// X-Barracuda-Spam-Score-Disabled header
	PerUserBypassedKey = RecipKey[bool]{"per_user_bypassed"}

	// BlockIfNotWhitelistKey stores whether the given recipient
	// was blocked due to the BlockIfNotWhitelisted header.
	BlockIfNotWhitelistKey = RecipKey[bool]{"block_if_not_whitelist"}

	// since we can't use coverage testing on the rules, this
	// records for testing purposes what headers we have not yet
	// seen. The tests can then assert that this is empty, meaning
	// we've positively tested them all.
	untestedHeaders = cm.SetFromSlice([]headers.CanonicalHeader{
		allowHeader,
		encryptHeader,
		fingerprintBlockHeader,
		quarantineHeader,
		rblBlockHeader,
		recipientAllow,
		redirectHeader,
		registryHeader,
		tagHeader1,
		tagHeader2,
		userWhitelistHeader,
		//  quarantineRBL,
		// quarantineBRL,
		// tagBRL,

		// the non-bsmtpd headers
		blockHeader,
		ipddHeader,
	})

	headerCounter = metrics.NewCounterVec(
		"bsmtpd_header_handled",
		"bsmtpd headers handled for bsmtpd-driven dispositions",
		[]string{"header"})

	// These next four variables started life as a map, but in
	// order to pass the precedence tests, these need to be deterministic.
	allowHeaders = []headerReason{
		{headers.MustCanonical("X-ASG-Recipient-Whitelist"), disposition.ReasonList},
		{headers.MustCanonical("X-ASG-PD-Recipient-Whitelist"), disposition.ReasonList},
		{headers.MustCanonical("X-Barracuda-User-Whitelist"), disposition.ReasonList},
		{headers.MustCanonical("X-Barracuda-PD-Sender-Whitelist"), disposition.ReasonPDSender},
		{headers.MustCanonical("X-Barracuda-PD-IP-Whitelist"), disposition.ReasonList},
		{headers.MustCanonical("X-Barracuda-PD-Recipient-Whitelist"), disposition.ReasonPDRecip},
	}

	blockHeaders = []headerReason{
		{headers.MustCanonical("X-Barracuda-User-Blacklist"), disposition.ReasonList},
		{headers.MustCanonical("X-Barracuda-PD-IP-Block"), disposition.ReasonList},
		{headers.MustCanonical("X-Barracuda-PD-Recipient-Block"), disposition.ReasonPDRecip},
		{headers.MustCanonical("X-Barracuda-PD-Sender-Blacklist"), disposition.ReasonPDSender},
	}

	quarantineHeaders = []headerReason{
		{headers.MustCanonical("X-Barracuda-PD-IP-Quarantine"), disposition.ReasonList},
		{headers.MustCanonical("X-Barracuda-PD-Recipient-Quarantine"), disposition.ReasonPDRecip},
		{headers.MustCanonical("X-Barracuda-PD-Sender-Quarantine"), disposition.ReasonPDSender},
	}

	tagHeaders = []headerReason{
		{headers.MustCanonical("X-Barracuda-PD-IP-Tag"), disposition.ReasonList},
		{headers.MustCanonical("X-Barracuda-PD-Recipient-Tag"), disposition.ReasonPDRecip},
		{headers.MustCanonical("X-Barracuda-PD-Sender-Tag"), disposition.ReasonPDSender},
		{headers.MustCanonical("X-Barracuda-Emailcat"), disposition.ReasonEmailCat},
	}

	// Special things that don't work like the above headers.
	blockIfNotWhitelistedHeader = headers.MustCanonical("X-Barracuda-BlockIfNotWhitelisted")
	asgWhitelistHeader          = headers.MustCanonical("X-ASG-Whitelist")

	allowedCounter = metrics.NewCounterVec(
		"lists_emails_allow_listed",
		"the number of emails passed through due to allow listing",
		[]string{"list"})
	blockedCounter = metrics.NewCounterVec(
		"lists_emails_block_listed",
		"the number of emails blocked due to block listing",
		[]string{"list"})
	quarantinedCounter = metrics.NewCounterVec(
		"lists_emails_quarantined",
		"the number of emails quarantined",
		[]string{"list"})
	tagCounter = metrics.NewCounterVec(
		"lists_emails_tagged",
		"the number of emails tagged",
		[]string{"list"})
)

type headerReason struct {
	Header headers.CanonicalHeader
	Reason disposition.Reason
}

func init() {
	jobs.HeaderRegistry.Register(&Job{})
}

// RecipKey is a per-recipient dictionary key.
type RecipKey[T any] struct {
	perrecip.Key[T]
}

// Job implements the BSMTD header processing.
type Job struct {
	// Block/Accept -> IP Reputation -> Custom External RBLs
	ExternalRBLs map[mail.Domain]disposition.Action `yaml:"external_rbls"`

	BRBLAction disposition.Action `yaml:"brbl_action"`
	SPFEnabled bool               `yaml:"enabled_spf"`
	HardFail   disposition.Action `yaml:"hard_fail_action"`
	SoftFail   disposition.Action `yaml:"soft_fail_action"`
}

// New returns a new instance of the job.
func (j *Job) New() jobs.Job[message.HeadersInfo] { return &Job{} }

// Type returns "bsmtpd_header".
func (j *Job) Type() string { return "bsmtpd_header" }

// These rules are tried in order on the email in order to figure out
// what to do with the bsmtp-originated headers.
//
// For each recipient, once a rule is hit, that recipient will not be
// searched any more, so these are in priority order.
var headerRules = []headerRule{
	// X-ASG-Recipient-Whitelist is not in this rule list because
	// of the special handling for matching on the recipient name,
	// but it conceptually appears here, allowing any recipient
	// that appears in it and skipping consulting the rest of this
	// list.

	// X-ASG-Whitelist, but "Barracuda Reputation" doesn't stop
	// outbound email.
	{
		Header: allowHeader,
		Extraction: action{disposition.Whitelist,
			reason{disposition.ReasonBarracudaReputation, null{}}},
		Direction: protosupport.Inbound,
		Match: func(val string) bool {
			return strings.Contains(val, "Barracuda Reputation")
		},
	},
	// global allow header
	{
		Header:     allowHeader,
		Extraction: action{disposition.Whitelist, reasonAndExtra{}},
		Direction:  protosupport.InboundAndOutbound,
		Match: func(val string) bool {
			return !strings.Contains(val, "Barracuda Reputation")
		},
	},
	// quarantine header, when it is by Sender
	{
		Header:     quarantineHeader,
		Extraction: action{disposition.Quarantine, reasonAndExtra{}},
		Match: func(val string) bool {
			return strings.Contains(val, "Sender")
		},
		Direction: protosupport.InboundAndOutbound,
	},
	// quarantine header otherwise
	{
		Header:     quarantineHeader,
		Extraction: action{disposition.Quarantine, reasonAndExtra{}},
		Direction:  protosupport.InboundAndOutbound,
	},
	// X-Barracuda-Registry header
	{
		Header: registryHeader,
		Match: func(val string) bool {
			return strings.Contains(val, "Level2")
		},
		Extraction: reason{disposition.ReasonBarracudaRegistry, null{}},
		Direction:  protosupport.Inbound,
	},
	// X-Barracuda-IPDD
	{
		Header:     ipddHeader,
		Extraction: reason{disposition.ReasonIPDD, valueHitRegexp{}},
		Direction:  protosupport.InboundAndOutbound,
	},
	// Block header from amavis. bsmtpd doesn't seem to mention
	// this, it seems to just block, but this is useful for unit
	// testing if nothing else because of the way Block will
	// overwrite any non-block result.
	{
		Header:     blockHeader,
		Extraction: action{disposition.Block, reasonAndExtra{}},
		Direction:  protosupport.InboundAndOutbound,
	},
	// RBL block header
	{
		Header: rblBlockHeader,
		Extraction: action{disposition.Block,
			reason{disposition.ReasonRBL,
				valueReasonExtra{}}},
		Direction: protosupport.InboundAndOutbound,
	},
	// fingerprint block header
	{
		Header: fingerprintBlockHeader,
		Extraction: action{disposition.Block,
			reason{disposition.ReasonFingerprint,
				valueReasonExtra{}}},
		Direction: protosupport.InboundAndOutbound,
	},
	// X-Barracuda-Encrypt
	{
		Header: encryptHeader,
		Extraction: action{disposition.Encrypt,
			reasonAndExtra{}},
		Direction: protosupport.Outbound,
	},
	// X-Barracuda-Redirect
	{
		Header: redirectHeader,
		Extraction: action{disposition.Redirect,
			reasonAndExtra{}},
		Direction: protosupport.Outbound,
	},
	// X-Barracuda-BRL-Quarantine
	{
		Header: quarantineBRL,
		Extraction: action{disposition.Quarantine,
			reason{disposition.ReasonBRL,
				headerValIsReasonExtra{}}},
		Direction: protosupport.InboundAndOutbound,
	},
	// X-Barracuda-BRL-Tag
	{
		Header: tagBRL,
		Extraction: action{disposition.Tag,
			reason{disposition.ReasonBRL,
				headerValIsReasonExtra{}}},
		Direction: protosupport.InboundAndOutbound,
	},
	// X-ASG-Tag
	{
		Header: tagHeader1,
		Extraction: action{disposition.Tag,
			reasonAndExtra{}},
		Direction: protosupport.Inbound,
	},
	// X-ASG - must have been the first header, has no "tag" in
	// it.
	{
		Header: tagHeader2,
		Extraction: action{disposition.Tag,
			reasonAndExtra{}},
		Direction: protosupport.Inbound,
	},
}

type headerRule struct {
	// The header to check
	Header headers.CanonicalHeader

	// Direction this can affect.
	Direction protosupport.Direction

	// Trigger only if this function returns true; undef is
	// assumed true.
	Match func(s string) bool

	// How to convert the header into the various bits we care
	// about for logging.
	Extraction headerExtraction
}

// RBLHit is something that can be set as the additional data for the asgheader job.
type RBLHit string

// ProcessMail implements the decisions yielded by headers from
// bsmtpd.
func (j *Job) ProcessMail(
	mc *jobs.MailContext,
	info message.HeadersInfo,
) error {
	headers := info.ReadHeaders(mc.Logger)
	env := info.Envelope(mc.Logger)

	recipientsWithDir := env.RecipientsWithDirection(protosupport.InboundAndOutbound, mc.Config.LocalDomains)
	recipientsRemaining := cm.SetFromSlice(recipientsWithDir)

	if mc.Config.QuarantineAdminPerUser {
		mailboxes := recipientsWithDir.MailboxSlice()
		pups, err := mc.DB.PoliciesForUsers(mailboxes...)
		if err != nil {
			mc.Error(LogDBError,
				slog.String("error", err.Error()))
		}

		for _, recip := range recipientsWithDir {
			if pups[recip.Mailbox()].BypassQuarantine {
				perrecip.Set(mc, recip.Address, PerUserBypassedKey, true)
			}
		}
	}

	// This is basically a long series of checks of the various
	// headers we may be getting, in order of priority in the
	// code.

	// This is the only recipient-specific header currently
	// showing in bsmtpd. amavisd shows a Recipient Quarantine but
	// it doesn't seem exist anymore.
	recipientsAllowed := append(headers.Values(recipientAllow),
		headers.Values(userWhitelistHeader)...)
	for _, recip := range recipientsAllowed {
		addr := mail.AddressFromEmail(strings.TrimSpace(recip.String()))
		disposition.Set(
			j,
			mc.Logger,
			mc,
			addr,
			disposition.Whitelist,
			"",
			disposition.ReasonPerUserSender,
			"",
		)
		untestedHeaders.Remove(recipientAllow)
		untestedHeaders.Remove(userWhitelistHeader)
	}

	// Walk through the table of headers above and implement the results.
	for ruleIdx, rule := range headerRules {
		headerValues := headers.Values(rule.Header)
		for _, headerU := range headerValues {
			headerContent := strings.TrimSpace(headerU.String())

			headerCounter.With(rule.Header.String()).Inc()

			headerLogger := mc.Logger.With(
				slog.String("header", rule.Header.String()),
				slog.Any("values", headerContent))

			if rule.Match != nil && !rule.Match(strings.TrimSpace(headerContent)) {
				headerLogger.Debug(LogFiltered)
				continue
			}

			headerLogger.Debug(LogHeaderHit,
				slog.Int("rule_idx", ruleIdx))

			// Check for  SPF handling.
			if j.SPFEnabled && strings.Contains(headerContent, "SPF") {
				spfResult := headers.GetDefault(spfHeader, "")
				if spfResult.String() == "" {
					mc.Debug(LogNoSPFResult)
					continue
				}

				result, _, ok := strings.Cut(spfResult.String(), " ")
				if !ok {
					mc.Info(LogUnknownSPFResult,
						slog.String("value", spfResult.String()))
					continue
				}

				switch strings.ToLower(result) {
				case "pass":
					mc.Info(LogValidSPFFound)
				case "softfail":
					data.Set(mc, SPFFailKey, true)
					for _, recip := range recipientsWithDir {
						if rule.Direction.Contains(recip.Direction) {
							untestedHeaders.Remove(rule.Header)
							disposition.Set(
								j,
								mc.Logger,
								mc,
								recip.Address,
								j.SoftFail,
								"",
								disposition.ReasonSPF,
								spfResult.String(),
							)
							recipientsRemaining.Remove(recip)
						}
					}
				case "fail":
					data.Set(mc, SPFFailKey, true)
					for _, recip := range recipientsWithDir {
						if rule.Direction.Contains(recip.Direction) {
							untestedHeaders.Remove(rule.Header)
							disposition.Set(
								j,
								mc.Logger,
								mc,
								recip.Address,
								j.HardFail,
								"",
								disposition.ReasonSPF,
								spfResult.String(),
							)
							recipientsRemaining.Remove(recip)
						}
					}
				case "unknown", "none", "neutral", "error":
					mc.Info(LogKnownButUnusedResult,
						slog.String("result", result),
						slog.String("value", spfResult.String()))
				default:
					mc.Info(LogUnknownSPFResult,
						slog.String("value", spfResult.String()))
				}
			}

			// Proceed with the existing header processing.
			action, err := rule.Extraction.Action(headerContent)
			if err != nil {
				headerLogger.Error(LogUnknownAction,
					slog.String("error", err.Error()))
			}
			if strings.Contains(headerContent, "ClientPTR") {
				for _, recip := range recipientsWithDir {
					disposition.Set(
						j,
						mc.Logger,
						mc,
						recip.Address,
						action,
						"",
						disposition.ReasonClientPTR,
						strings.Split(headerContent, "-")[len(strings.Split(headerContent, "."))-1],
					)
				}
			} else {
				reason, err := rule.Extraction.Reason(headerContent)
				if err != nil {
					headerLogger.Error(LogUnknownReason,
						slog.String("error", err.Error()))
				}
				hitRegexp := rule.Extraction.HitRegexp(headerContent)
				reasonExtra := rule.Extraction.ReasonExtra(headerContent)

				for _, recip := range recipientsWithDir {
					if rule.Direction.Contains(recip.Direction) {
						untestedHeaders.Remove(rule.Header)
						disposition.Set(
							j,
							mc.Logger,
							mc,
							recip.Address,
							action,
							hitRegexp,
							reason,
							reasonExtra,
						)
						recipientsRemaining.Remove(recip)
					}
				}
			}
		}
	}
	// For handling the Quarantine RBL
	quarantineRBL, have := headers.Get(quarantineRBL)
	if have {
		mc.Info(LogRBLHit,
			slog.String("rbl", quarantineRBL.String()))
		domain := mail.NewDomain(quarantineRBL.String())
		action, haveAction := j.ExternalRBLs[domain]
		if !haveAction {
			mc.Error(LogRBLNotConfigured,
				slog.String("rbl", quarantineRBL.String()))
		} else {
			for _, recip := range recipientsWithDir {
				disposition.SetWithAdditional(
					j,
					mc.Logger,
					mc,
					recip.Address,
					action,
					"",
					disposition.ReasonRBL,
					quarantineRBL.String(),
					RBLHit(quarantineRBL.String()),
				)

				recipientsRemaining.Remove(recip)
			}
		}
	}
	// Check RBL retrictions; this can't be a rule because we have
	// to lookup what to do based on what hits.
	rblRestriction, have := headers.Get(rblRestriction)
	if have {
		mc.Info(LogRBLHit,
			slog.String("rbl", rblRestriction.String()))
		domain := mail.NewDomain(rblRestriction.String())
		action, haveAction := j.ExternalRBLs[domain]
		if !haveAction {
			mc.Error(LogRBLNotConfigured,
				slog.String("rbl", rblRestriction.String()))
		} else {
			for _, recip := range recipientsWithDir {
				disposition.SetWithAdditional(
					j,
					mc.Logger,
					mc,
					recip.Address,
					action,
					"",
					disposition.ReasonRBL,
					rblRestriction.String(),
					RBLHit(rblRestriction.String()),
				)

				recipientsRemaining.Remove(recip)
			}
		}
	}

	// Check the BRBL header
	brblHit, have := headers.Get(repBlockHeader)
	if have && j.BRBLAction != disposition.None {
		mc.Info(LogBRBLHitFound,
			slog.String("hit", brblHit.String()))
		for _, recip := range recipientsWithDir {
			disposition.SetWithAdditional(
				j,
				mc.Logger,
				mc,
				recip.Address,
				j.BRBLAction,
				"",
				disposition.ReasonBRL,
				brblHit.String(),
				RBLHit(brblHit.String()),
			)
			recipientsRemaining.Remove(recip)
		}
	}

	lists, err := getLists(mc.Logger, info.ReadHeaders(mc.Logger))
	if err != nil {
		// Log the issues but proceed ahead as best we can.
		mc.Error(LogInvalidHeaders, slog.Any("error", err))
	}

	recipientMailboxes := recipientsWithDir.MailboxSlice()

	allowBlockLists, err := mc.DB.AllowBlockListsForUsers(recipientMailboxes...)
	if err != nil {
		mc.Error(LogDatabaseError, "error", err.Error())
	}

	var retErr error
	for _, recipient := range recipientsWithDir {
		j.processForRecipient(
			mc,
			info,
			lists,
			recipient,
			allowBlockLists,
		)
	}

	return retErr
}

func (j *Job) processForRecipient(
	mc *jobs.MailContext,
	info message.HeadersInfo,
	lists *listHeaders,
	recipient mail.AddressDirection,
	allowBlock map[mail.Mailbox]db.AllowBlockList,
) {
	mailboxAddr := recipient.Mailbox()
	envelope := info.Envelope(mc.Logger)
	sender := envelope.From().MailboxFromLocalDomains(mc.Config.LocalDomains)
	perRecipData := mc.PerRecip(recipient.Address)
	recipData := mc.PerRecipHandle.PerRecip(recipient.Address)
	userAllowBlockList, haveUserAllowBlock := allowBlock[mailboxAddr]

	// There are many ways BSMTP can label things as being
	// allowed, blocked, tagged, or quarantined, plu a
	// BlockIfNotAllowListed flag it can set. So, first, we start
	// by seeing if this is allowlisted as that dominates everything.

	// allowed by user specifically?
	if haveUserAllowBlock {
		if userAllowBlockList.Allowed(sender) {
			mc.Info(LogAllowListed,
				slog.String("allow_list", "recipient"),
				slog.String("recipient", recipient.Address.String()),
				slog.String("source", "wblist"))
			allowedCounter.With("recipient_wblist").Inc()
			disposition.SetR(j, mc.Logger, perRecipData,
				disposition.Whitelist,
				"",
				disposition.ReasonNone,
				"")
			return
		}
	}
	// Allowed by any of the other stable of headers?
	if have, reason := lists.AllowList.InList(mailboxAddr); have {
		mc.Info(LogAllowListed,
			slog.String("recipient", recipient.Address.String()))
		allowedCounter.With("bsmtpd_header").Inc()
		disposition.SetR(j, mc.Logger, perRecipData,
			disposition.Whitelist,
			"",
			reason,
			"")
		return
	}
	// Allowed by regexp list?
	allowed, by := lists.RegexpAllow.InListAllowedBy(mailboxAddr)
	if allowed {
		mc.Info(LogAllowListed,
			slog.String("allow_list", "recipient"),
			slog.String("recipient", recipient.Address.String()),
			slog.String("pattern", (*metare.RE)(by).String()))
		// go ahead and deliver this email
		allowedCounter.With("regexp_list").Inc()
		disposition.SetR(j, mc.Logger, perRecipData,
			disposition.Whitelist,
			"",
			disposition.ReasonList,
			"")
		return
	}

	// We have now run through all the reasons we might
	// allow. Now, if the BlockIfNotWhitelist is set, we need to block.
	if lists.BlockIfNotAllowListed {
		mc.Info(LogBlockingOnNotAllowListed,
			slog.String("recipient", recipient.Address.String()))
		disposition.SetR(j, mc.Logger, recipData,
			disposition.Block,
			"",
			// Yes, this is what amavis said. Maybe this
			// is only emitted by smtpd in the first place if
			// there is an attachment?
			disposition.ReasonAttachment,
			"")
		// Accept message, but report it was blocked.
		envelope.SetSMTPResponse(
			protosupport.ESMTPReply{
				Code:    250,
				Class:   5,
				Subject: 7,
				Detail:  1,
				Lines:   []string{"Message contained a banned attachment type"},
			},
			recipient.Address,
		)
		perrecip.SetR(recipData, BlockIfNotWhitelistKey, true)

		blockedCounter.With("blockifnotwhitelisted").Inc()
		return
	}

	// Any "normal" reason to block?
	if have, reason := lists.BlockList.InList(mailboxAddr); have {
		mc.Info(LogBlocked,
			slog.String("recipient", recipient.Address.String()))
		disposition.SetR(j, mc.Logger, recipData,
			disposition.Block,
			"",
			reason,
			"")
		blockedCounter.With("bsmtpd_header").Inc()
		return
	}
	if haveUserAllowBlock {
		if userAllowBlockList.Blocked(sender) {
			mc.Info(LogBlocked,
				slog.String("block_list", "recipient"),
				slog.String("recpient", recipient.Address.String()),
				slog.String("source", "wblist"))
			blockedCounter.With("recipient_wblist").Inc()
			disposition.SetR(j, mc.Logger, perRecipData,
				disposition.Block,
				"",
				disposition.ReasonSender,
				"")
			return
		}
	}

	// Any reason to quarantine?
	if have, reason := lists.QuarantineList.InList(mailboxAddr); have {
		mc.Info(LogQuarantined,
			slog.String("recipient", recipient.Address.String()))
		disposition.SetR(j, mc.Logger, recipData,
			disposition.Quarantine,
			"",
			reason,
			"")
		quarantinedCounter.With("bsmtpd_header").Inc()
		return
	}

	// Any reason to tag?
	if have, reason := lists.TagList.InList(mailboxAddr); have {
		mc.Info(LogTagged,
			slog.String("recipient", recipient.Address.String()))
		disposition.SetR(j, mc.Logger, recipData,
			disposition.Tag,
			"",
			reason,
			"")
		tagCounter.With("bsmtpd_header").Inc()
		return
	}
}

// getLists gathers up all the relevant headers and turns their
// contents into userlist.ListMembers, allowing them to be tested for.
func getLists(l *slog.Logger, h *headers.ReadHandle) (*listHeaders, error) {
	var retErr error
	ret := &listHeaders{}

	ret.BlockIfNotAllowListed = !h.GetDefault(blockIfNotWhitelistedHeader, "").EqUTF8("")

	// These headers come from bsmtpd, and despite their names
	// suggesting that they represent configuration, the actually
	// represent *hits*. An X-Barracuda-PD-IP-Tag header is not
	// instructing how to do per-domain tagging based on IP, it is
	// indicating that the per-domain IP filters have hit with a
	// "tag" result.
	//
	// It was not entirely clear to me (jbowers) which headers
	// have specific email addresses and which have domains. It
	// didn't seem to match what the source code for bsmtpd said,
	// based on my reading. However, it seems save to let
	// MailboxOrDomain just go ahead and handle both, since their
	// representations should be entirely disjoint there can be no conflict.

	mergeHeaders := func(headersReason []headerReason) listsWithReason {
		lwr := listsWithReason{}
		for _, headerReason := range headersReason {
			thisList, err := headers.GetValuesAs[*mail.MailboxOrDomain](h, headerReason.Header)
			retErr = errors.Join(retErr, err)
			lwr = append(lwr, listWithReason{Reason: headerReason.Reason, ListList: thisList})
		}
		return lwr
	}

	ret.AllowList = mergeHeaders(allowHeaders)
	ret.BlockList = mergeHeaders(blockHeaders)
	ret.QuarantineList = mergeHeaders(quarantineHeaders)
	ret.TagList = mergeHeaders(tagHeaders)

	wlHeaders := h.Values(asgWhitelistHeader)
	regexpAllows := userlist.ListList[*userlist.RegexpList]{}
	for i, headerVal := range wlHeaders {
		headerVal = headerVal.TrimSpace()
		// This header comes in with a lot of different first
		// words, so just skip the ones we're not interested in
		if !strings.HasPrefix(headerVal.String(), recipientRegexpMarker) {
			continue
		}
		// it seems something is filtering this to just
		// regular expressions before it gets to us
		reStr := strings.TrimSpace(headerVal.String()[len(recipientRegexpMarker):])
		if len(reStr) < 2 || reStr[0] != '/' || reStr[len(reStr)-1] != '/' {
			l.Info(LogInvalidRE,
				"regex", reStr)
			continue
		}
		reStr = reStr[1 : len(reStr)-1]

		re, err := metare.New(reStr, "asg whitelist header", i+1)
		if err != nil {
			l.Info(LogInvalidREPostTrim,
				"regex", reStr,
				"error", err)
			continue
		}

		reAllow := userlist.RegexpList(re)
		regexpAllows = append(regexpAllows, &reAllow)
	}
	ret.RegexpAllow = regexpAllows

	return ret, retErr
}

type listWithReason struct {
	disposition.Reason
	userlist.ListList[*mail.MailboxOrDomain]
}

type listsWithReason []listWithReason

// disposition.ReasonNone = not in list
func (lwr listsWithReason) InList(m mail.Mailbox) (bool, disposition.Reason) {
	for _, listAndReason := range lwr {
		if listAndReason.ListList.InList(m) {
			return true, listAndReason.Reason
		}
	}
	return false, disposition.ReasonNone
}

// This collects all the headers of interest into one place. Think
// mise en place.
type listHeaders struct {
	AllowList      listsWithReason
	BlockList      listsWithReason
	QuarantineList listsWithReason
	TagList        listsWithReason

	RegexpAllow userlist.ListList[*userlist.RegexpList]

	BlockIfNotAllowListed bool
}
