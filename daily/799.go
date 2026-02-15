package main

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp:= make([][]float64,query_row+2)
	for i := range dp {
        dp[i] = make([]float64, query_row+2)
    }
	dp[0][0] = float64(poured)
	for row:=0;row<=query_row;row++{
		for c:=0;c<=row;c++{
			if dp[row][c]>1{
				profit:= (dp[row][c]-1.0)/2.0
				dp[row+1][c] += profit
				dp[row+1][c+1] += profit
			}
		}
	}
	if dp[query_row][query_glass] > 1{return 1}else{return dp[query_row][query_glass]}
}
