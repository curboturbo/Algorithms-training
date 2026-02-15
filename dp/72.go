package main
func m(x int, y int) int {
    if x <= y {
        return x
    } else {
        return y
    }
}

func minDistance(word1 string, word2 string) int {
	w1:=" "+word1
	w2:=" "+word2
	if word1 == "" {return len([]rune(word2))}
	if word2 == "" {return len([]rune(word1))}
	a:=[]rune(w1) // БИБА
	b:=[]rune(w2) // БОБА
	dp:=[][]int{}
	for i:=0;i<len(b);i++{
		z:=make([]int,len(a))
		z[0] = i
		dp = append(dp, z)
	}
	for j:=0;j<len(a);j++{
		dp[0][j] =j
	}
	for i:=1;i<len(b);i++{
		for j:=1;j<len(a);j++{
			if b[i] == a[j]{
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = m(
					dp[i-1][j]+1,
					dp[i-1][j-1]+1,
				)
				dp[i][j] = m(
					dp[i][j],
					dp[i][j-1]+1,
				)
			}
		}
	}
	return dp[len(b)-1][len(a)-1]
}