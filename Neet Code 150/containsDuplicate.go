func containsDuplicate (nums []int) bool {
	mymap := make(map[int]bool)

	for _, i := range nums {
		if mymap[i] {
			return true
		}
		mymap[i] = true
	}
	return false
}


OR

func containsDuplicate (nums []int) bool {
	for i := 0; i <= len(nums) ; i++ {
		for j := i+1; j <= len(nums); j++ {
			if nums[i] == nums[j]
			return true
		}
	}
	return false
}