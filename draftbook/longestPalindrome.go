package draftbook

import "fmt"

func main4() {
	ans := longestPalindrome("babad")
	fmt.Println(ans)
}

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	var ans string = ""
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 2 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return ans
}
