package main


func binary_search(nums []int, left int, right int, k int) int {
    for left <= right {
        mid := (left + right) / 2
        if k == nums[mid] {
            return mid
        }
        if k > nums[mid] {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}

func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }
    left, right := 0, len(nums)-1
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    pivot := left
    if target >= nums[pivot] && target <= nums[len(nums)-1] {
        return binary_search(nums, pivot, len(nums)-1, target)
    } else {
        return binary_search(nums, 0, pivot-1, target)
    }
}
