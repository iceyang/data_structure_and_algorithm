---
title: "理论知识与问题特征"
date: 2019-06-09T17:19:00+08:00
---

直接看理论枯燥无味，且印象不深，我们先从一些简单的例子开始，从而引出对应的知识点。

## 斐波拉契数列（Fibonacci）

关于斐波拉契数列，大家应该很熟悉了，它的定义公示如下：

``` shell
Fibonacci(n) = Fibonacci(n-1) + Fibonacci(n-2)

其中：Fibonacci(0) = 1，Fibonacci(1) = 1
```

作为递归的经典入门例子，我们很快能想到它的递归解法：

``` Go
// 语言是Golang

func fibonacci(n int) int {
        if n == 0 || n == 1 {
                return 1
        }
        return fibonacci(n-1) + fibonacci(n-2)
}
```
