package main


func isTrionic(nums []int) bool {
    n := len(nums)
    if n < 3 { return false }
    i := 1
    for i < n && nums[i] > nums[i-1] { i++ }
    if i == 1 || i == n { return false }

    for i < n && nums[i] < nums[i-1] { i++ }
    if i == n { return false }

    for i < n && nums[i] > nums[i-1] { i++ }
    return i == n
}
