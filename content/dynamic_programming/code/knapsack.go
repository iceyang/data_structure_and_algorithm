package main

/**
 * 假设：现在有一个容量为 w 的背包，有 n 件物品，它们的重量为 weights=[w1, w2, ..., w(n)]，它们的价值为values=[v1, v2, ..., v(n)]。
 * 问：背包能装入的物品的最大价值是多少？
 */

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func knapsack(w int, weights []int, values []int) int {
	n := len(weights)
	res := make([][]int, n+1)
	for i := range res {
		res[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		weight := weights[i-1]
		value := values[i-1]
		for j := 1; j <= w; j++ {
			if j < weight {
				res[i][j] = res[i-1][j]
				continue
			}
			res[i][j] = max(res[i-1][j], res[i-1][j-weight]+value)
		}
	}

	return res[n][w]
}

/**
 * 根据动态规划的特性，我们可以将二维空间降至一维。
 * 第i件物品是否放入背包，只依赖于第i-1件物品时容量为j-w的结果，即：res[j-weights[i]]
 */
func knapsack2(w int, weights []int, values []int) int {
	n := len(weights)
	res := make([]int, w+1)

	for i := 1; i <= n; i++ {
		weight := weights[i-1]
		value := values[i-1]
		// 此时的res[j-weight]相当于二维时的res[i-1][j-weight]
		for j := w; j >= weight; j-- {
			if res[j-weight]+value > res[j] {
				res[j] = res[j-weight] + value
			}
		}
		// 逆序计算是为了避免重复计算的情况
		// 以w=8, weights=[2,3,4,5]，values=[3,4,5,6]为例，当i为1时，如果采用正序计算一轮的结果如下：
		// [0 0 3 0 0 0 0 0 0]
		// [0 0 3 3 0 0 0 0 0]
		// [0 0 3 3 6 0 0 0 0]
		// [0 0 3 3 6 6 0 0 0]
		// [0 0 3 3 6 6 9 0 0]
		// [0 0 3 3 6 6 9 9 0]
		// [0 0 3 3 6 6 9 9 12]
		// 可以思考下为什么会这样
	}

	return res[w]
}
