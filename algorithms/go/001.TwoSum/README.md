# 题目
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

> 给定 nums = [2, 7, 11, 15], target = 9
> 
> 因为 nums[0] + nums[1] = 2 + 7 = 9
> 所以返回 [0, 1]

# 原思路
切片两次循环，判断索引

# 优化思路
a + b = target 可转换成 a = target - b
利用map结构，构造一个map[a] = index
如果值a存在于map，则返回index值，一次循环即可完成。

# 新知识点
* [如何判断两个切片相等](https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices)

> reflect.DeepEqual()
> 
> DeepEqual is a recursive relaxation of Go's == operator.
> 
> DeepEqual reports whether x and y are “deeply equal,” > defined as follows. Two values of identical type are > deeply equal if one of the following cases applies. Values > of distinct types are never deeply equal.
> 
> Array values are deeply equal when their corresponding > elements are deeply equal.
> 
> Struct values are deeply equal if their corresponding > fields, both exported and unexported, are deeply equal.
> 
> Func values are deeply equal if both are nil; otherwise > they are not deeply equal.
> 
> Interface values are deeply equal if they hold deeply > equal concrete values.
> 
> Map values are deeply equal if they are the same map > object or if they have the same length and their > corresponding keys (matched using Go equality) map to > deeply equal values.
> 
> Pointer values are deeply equal if they are equal using > Go's == operator or if they point to deeply equal values.
> 
> Slice values are deeply equal when all of the following > are true: they are both nil or both non-nil, they have the > same length, and either they point to the same initial > entry of the same underlying array (that is, &x[0] == &y[0]> ) or their corresponding elements (up to length) are > deeply equal. Note that a non-nil empty slice and a nil > slice (for example, []byte{} and []byte(nil)) are not > deeply equal.
> 
> Other values - numbers, bools, strings, and channels - are > deeply equal if they are equal using Go's == operator.

    reflect.DeepEqual(a, b)

