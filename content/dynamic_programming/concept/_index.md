---
title: "理论知识与问题特征"
date: 2019-06-09T17:19:00+08:00
---

直接看理论枯燥无味，且印象不深，我们先从一些简单的例子开始，从而引出对应的知识点。

## 斐波拉契数列

关于斐波拉契数列，大家应该很熟悉了，它的定义公示如下：

``` shell
Fibonacci(n) = Fibonacci(n-1) + Fibonacci(n-2)

其中：Fibonacci(0) = 1，Fibonacci(1) = 1
```

作为递归的经典入门例子，我们很快能想到它的递归解法：

``` Go
func fibonacci(n int) int {
        if n == 0 || n == 1 {
                return 1
        }
        return fibonacci(n-1) + fibonacci(n-2)
}
```

假设我们现在根据上面的递归代码要求出F(5)的值，按照递归铺开，可以得到下面这个计算过程：

![根据递归求F(5)](concept_1.png)

可以明显看到，里面有某些值是重复计算的，比如：F(2)就计算了3次。

我们可以对代码做一下优化：

``` Go
func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}
```

采用自底向上的方法，将过程中的计算结果保存起来，那么问题就变成了求解红色框框部分：

![自底向上求F(5)](concept_2.png)

```
到了这里，我们引入了动态规划的其中一个特性：
*重复子问题*
```
