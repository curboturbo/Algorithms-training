package main


func subarraySum(nums []int, k int) int {
    mapa := make(map[int]int)
    mapa[0] = 1
    prefix := 0
    cnt := 0
    for i := 0; i < len(nums); i++ {
        prefix += nums[i]
        if val, ok := mapa[prefix-k]; ok {
            cnt += val
        }
        mapa[prefix]++
    }
    return cnt
}

