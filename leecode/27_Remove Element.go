package main

func removeElement(nums []int, val int) int {
	newslice := []int{}
	for _, v := range nums {
		if v != val {
			newslice = append(newslice, v)
		}
	}
	copy(nums[:], newslice)
	return len(newslice)
}
