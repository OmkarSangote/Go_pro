func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

OR

func twoSum (nums []int, target int) []int {
	mymap := make(map[int]int)

	for i, num := range nums {
		diff := target - num 
		if idx, ok := mymap[diff]; ok {
			return []int{idx, i}
		}
		mymap[num] = i
	}
	return nil
}

OR 

func twoSum (nums []int, target int) []int {
	left := 0
	right := len(nums)-1
	currsum := 0

	sort.Ints(arr)

	for left < right {
		currsum = nums[left] + nums[right]
		if currsum > target {
			right = right - 1
		}

		else if currsum < target {
			left = left + 1
		}

		else return []int{left + 1, right + 1}
	}
func twoSum (nums []int, target int) []int {
	left := 0
	right := len(nums)-1
	currsum := 0

	for left < right {
		currsum = nums[left] + nums[right]
		if currsum > target {
			right = right - 1
		}

		else if currsum < target {
			left = left + 1
		}

		else return []int{left + 1, right + 1}
	}

func twoSum (nums []int, target int) []int {
	left := 0
	right := len(nums)-1
	currsum := 0
	sort.Ints(nums)

	for left < right {
		currsum = nums[left] + nums[right]
		if currsum > target {
			right--
		}

		else if currsum < target {
			left++
		}

		else return []int{left + 1, right + 1}
	}
return []int{}

}