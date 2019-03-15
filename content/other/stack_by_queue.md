---
title: "由栈实现的队列"
date: 2019-03-15T17:34:24+08:00
---

看到一个题目，叫做用栈来实现队列，觉得有点意思，就简单做了下实现。

{{% panel header="描述" %}}
实现一个 *内部是由栈构成* 的队列。

栈的方法有`push：入栈、pop：出栈`。

实现队列的`push：入队、poll：出队`。

{{%/ panel %}}

## 简单版本
栈的特性应该很清楚：先进后出，而队列的特性是先进先出。

意味着由栈来实现队列，在取数据的时候，需要从最底部拿出，那么最直接的想法就是把栈的数据全部都拿出来，得到想要的数据后，把栈重新装回去（因为经过拿的那一步，顺序会反了，我们需要把顺序还原回去）。

所以我们可以这样子来做：
```
// TypsScript
/**
 * 由栈实现的队列
 */
import Stack from '../lib/stack';

export default class QueueByStack<T> {
  private stack: Stack<T> = new Stack<T>();

  length(): number { return this.stack.length(); }

  private reverseStack(): void {
    const result = new Stack<T>();
    do {
      const value = this.stack.pop();
      if (!value) break;
      result.push(value);
    } while (true)
    this.stack = result;
  }

  push(t: T) {
    this.stack.push(t);
  }

  poll(): T | undefined {
    if (this.stack.length() <= 0) return undefined;
    this.reverseStack();
    const result = this.stack.pop();
    this.reverseStack();
    return result;
  }
}
```

其中`reverseStack`是实现了将栈倒过来，取完数据后，重新反转一次。

## TBD
优化

## 代码
上面的源码可以在 [{{< icon name="fa-github" size="large" >}}](https://github.com/iceyang/data_structure_and_algorithm_code/blob/master/src/queue_by_stack) 获得。
