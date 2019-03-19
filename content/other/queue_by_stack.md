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

## 优化版
上面简单版本的实现，有一个问题，就是每次出队，会有两次整个栈的倒装操作。

我们可以用两个栈来保存队列中的内容，假设为S1和S2。

S1负责入队操作，所有新push的数据，直接进入到S1中。

出队由S2负责，当S2中不存在元素时，将S1的内容pop到S2中，然后再从S2 pop出去。

为什么可以这样做？

由于栈的先进后出特性，当S2没有元素时，S1的内容直接装到S2，S2的元素顺序经过S1的倒装，就变成了先进先出的效果。

而如果S2有元素，为什么出队就不需要理S1了？因为S2的元素肯定是比S1的元素先到达队列的，所以可以直接忽视S1，直到S2没元素了，才去S1中取。

下面是简单的代码实现：

```
/**
 * 优化版本，保存两个栈，一个用于入队(S1)，一个用于出队(S2)。
 * 入队时：直接进栈S1
 * 出队时：假如S2为空，则将S1倒入S2中，S2出栈；如果S2不为空，则S2直接出栈
 */
export class QueueByStack2<T> implements Queue<T>{
  private s1: Stack<T> = new Stack<T>();
  private s2: Stack<T> = new Stack<T>();

  length(): number { return this.s1.length() + this.s2.length(); }

  push(t: T) {
    this.s1.push(t);
  }

  poll(): T | undefined {
    if (this.length() <= 0) return undefined;
    if (this.s2.length() > 0) return this.s2.pop();
    while (this.s1.length() > 0) {
      this.s2.push(<T>this.s1.pop());
    }
    return this.s2.pop();
  }
}
```

## 代码
上面的源码可以在 [{{< icon name="fa-github" size="large" >}}](https://github.com/iceyang/data_structure_and_algorithm_code/blob/master/src/queue_by_stack) 获得。
