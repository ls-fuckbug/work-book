package main

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				}
			}
		}

	}
	return dp[l1][l2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
