// 4. Shared Metadata in Database Models
// In a database management system, you may want all records to include common metadata, like timestamps for when the record was created or updated.

// Example: Tracking Timestamps in Database Models

package main

import (
	"fmt"
	"time"
)

// Define a common struct for tracking timestamps
type Timestamps struct {
	CreatedAt time.Time // from the time package, using the Time struct
	UpdatedAt time.Time
}

// User struct embeds Timestamps for common metadata
type User struct {
	ID   int
	Name string
	Timestamps
}

// Product struct embeds Timestamps for common metadata
type Product struct {
	ID    int
	Name  string
	Price float64
	Timestamps
}

func main() {
	user := User{
		ID:   1,
		Name: "John Doe",
		Timestamps: Timestamps{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	fmt.Printf("User %s created at %v\n", user.Name, user.CreatedAt)
}

//Use Case: In an application managing users and products, both entities can share common metadata (like CreatedAt and UpdatedAt) by embedding the Timestamps struct.has context menu
