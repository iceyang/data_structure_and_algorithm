---
title: "经典题目：01背包问题"
date: 2019-07-04T21:50:00+08:00
---

{{% panel header="题目描述" %}}

假设：现在有一个容量为 W 的背包，有 n 件物品，它们的重量为 [w1, w2, ..., w(n)]，它们的价值为 [v1, v2, ..., v(n)]。

问：背包能装入的物品的最大价值是多少？
 
{{%/ panel %}}

{{% panel theme="danger" header="题目分析" %}}

对于大问题，我们先寻找有没有办法拆解成小问题。

背包中的第i件物品，它是否放入背包，取决于背包中前 i-1 件物品的放置情况。

我们假定 V(i, j) 代表当背包容量为 j 时，前 i 件物品放入背包的最大价值。

那么对于第 i 件物品是否放入容量为 j 的背包中，我们有下面的推导过程：

1. 如果第 i 件物品的重量大于背包容量，那么最大价值等于前 i-1 件物品放入背包的最大价值，即 V(i, j) = V(i-1, j)
2. 如果第 i 件物品可以放入背包，那么
    * 当该物品放入背包时，背包需要腾出w(i)的空间，所以价值为 V(i-1, j-w(i)) + v(i)
    * 当该物品不放入背包时，价值为 V(i-1, j)
    * 最终最大价值 V(i, j) = max{ V(i-1, j), V(i-1, j-w(i)) + v(i) }

所以，我们得到最终的公式为：

```
1) 当j <w(i)时，V(i,j)=V(i-1,j)
2) 当i>=w(i)时，V(i,j)=max{V(i-1,j),V(i-1,j-w(i))+v(i)}
```

从上面的分析可以得出，这是符合动态规划特性的，属于多阶段决策问题：

* 重复子问题：子问题存在多次计算的情况；
* 无后效性：当前问题可以由子问题推导出来，子问题状态一旦确定，便不会更改；
* 最优子结构：每个解都是当前情况的最优解，问题的子问题都是当时状态下的最优选择。
 
{{%/ panel %}}

{{% panel theme="info" header="动态规划解题代码" %}}

我们定义了一个函数，输入为：

* w 为背包容量
* weights 为所有物品的重量
* values 为所有物品的价值

输出结果为背包所能装入的最大价值。

我们采用一个二维数组保存结果，第一个维度代表着放入前i个物品，第二个维度是指放入第i个物品时的背包容量j，那么 res[i][j] 代表的是当背包容量为j时，放入前i个物品的最大价值。

第二层循环内，res[i][j] = max(res[i-1][j], res[i-1][j-weight]+value) 是关键所在，对应着我们上面的分析。

```Go
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
```
 
{{%/ panel %}}

{{% panel theme="info" header="进一步优化" %}}


根据动态规划的特性，我们可以将二维空间降至一维。
第 i 件物品是否放入背包，只依赖于第 i-1 件物品时容量为j-w 的结果，即：res[j-weights[i]]

```Go
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
```
 
{{%/ panel %}}

到此，我们01背包问题已经采用动态规划算法进行了解答，若有什么疑问或者错误的地方，还望指出改正。

代码可以在 进行查看，疑问可以在同个项目进行issue提交。
[{{< icon name="fa-github" size="large" >}}](https://github.com/iceyang/data_structure_and_algorithm/blob/master/content/dynamic_programming/code)

感谢您的宝贵时间。
