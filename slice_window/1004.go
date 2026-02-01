package main

func mx(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func longestOnes(nums []int, k int) int {
	ans := -120
	cnt := 0
	cnt_zero := 0
	left := 0
	for i := 0; i < len(nums); i++ {
		ans = mx(ans, cnt)
		if nums[i] == 0 {
			if cnt_zero+1 <= k {
				cnt++
				cnt_zero++
			} else {
                cnt_zero++
                cnt++
				for cnt_zero > k {
					if nums[left] == 0 {
						cnt_zero--
					}
					left++
					cnt--
				}
			}
		} else {
			cnt++
		}
	}

	ans = mx(ans, cnt)
	return ans
}