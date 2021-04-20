1. 下面的代码是做什么的？用Python/C/C++/Java或其他你擅长的语言重新实现它。注意，在重新实现时不可照搬伪代码的逻辑，而需要在对样例代码理解的基础上使用更简洁高效的方式。

```
 1: def Boolean no_name(a : String, b : String)
 2:     if a.length != b.length
 3:         return false
 4:
 5:     for(x : Integer = 0; x < b.length; x++)
 6:         if a[0] == b[x]
 7:             return no_name(utilityFunction(a, 0), utilityFunction(b, x))
 8:         end
 9:     end
10:
11:     return b.length == 0
12: end
13:
14: def String utilityFunction(s : String, j : Integer)
15:     ret = new char[s.length - 1]
16:     int d = 0
17:     for (k : Integer = 0; k < s.length; k++)
18:         if (k == j)
19:             d = 1
20:         else
21:             ret[k - d] = s[k]
22:         end
23:     end
24:     return new String(ret)
25: end
```

2. 为你在上一题中实现的代码编写单元测试。

3. 实现[Python 3的max函数](https://docs.python.org/3/library/functions.html#max)。需要实现 Python 3 标准库的 max 完整的 API，包括 key, default 等 keyword only 参数的支持。考察点不是基础算法，而是工程上去写用于生产的代码需要关注的点：API 的设计，文档，代码逻辑的清晰等。

4. 在以前的开发经历中，你遇到过最大的挑战是什么？有何收获？

5. 在你过去共事过的同事/同学里，有没有哪个人是你最想感谢的？为什么？
