package main

func maximumSubarraySum(nums []int, k int) int64 {
	var ans, sum int
	cnt := make(map[int]int)
	for _, num := range nums[:k-1] {
		cnt[num]++
		sum += num
	}
	for i := k - 1; i < len(nums); i++ {
		in, out := nums[i], nums[i-k+1]
		cnt[in]++
		sum += in
		if len(cnt) == k {
			ans = max(ans, sum)
		}
		cnt[out]--
		if cnt[out] == 0 {
			delete(cnt, out)
		}
		sum -= out
	}
	return int64(ans)
}
