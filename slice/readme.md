我描述一下它所做的事情。首先，我用内建函数make声明了一个[]int类型的变量s1。我传给 make函数的第二个参数是5，从而指明了该切片的长度。我用几乎同样的方式声明了切片s2， 只不过多传入了一个参数8以指明该切片的容量。
现在，具体的问题是:切片s1和s2的容量都是多少? 这道题的典型回答:切片s1和s2的容量分别是5和8。
问题解析
解析一下这道题。s1的容量为什么是5呢?因为我在声明s1的时候把它的长度设置成了5。 当我们用make函数初始化切片时，如果不指明其容量，那么它就会和长度一致。如果在初始化
时指明了容量，那么切片的实际容量也就是它了。这也正是s2的容量是8的原因。 我们顺便通过s2再来明确下长度、容量以及它们的关系。我在初始化s2代表的切片时同时指定
了它的长度和容量。
在这种情况下，它的容量实际上代表了它的底层数组的长度，这里是8。注意，切片的底层数组 等同于我们前面讲到的数组，其长度不可变。
现在你需要跟着我一起想象:有一个窗口，你可以通过这个窗口看到一个数组，但是不一定能看
到该数组中的所有元素，有时候只能看到连续的一部分元素。
现在，这个数组就是切片s2的底层数组，而这个窗口就是切片s2本身。s2的长度实际上指明的 就是这个窗口的宽度，决定了你透过s2可以看到其底层数组中的哪几个连续的元素。
由于s2的长度是5，所以你可以看到其底层数组中的第 1 个元素到第 5 个元素，对应的底层数 组的索引范围是 [0, 4]。
}
// 示例 1。
s1 := make([]int, 5)
fmt.Printf("The length of s1: %d\n", len(s1)) fmt.Printf("The capacity of s1: %d\n", cap(s1)) fmt.Printf("The value of s1: %d\n", s1)
s2 := make([]int, 5, 8)
fmt.Printf("The length of s2: %d\n", len(s2)) fmt.Printf("The capacity of s2: %d\n", cap(s2)) fmt.Printf("The value of s2: %d\n", s2)
https://time.geekbang.org/column/article/14106
3/12

切片代表的窗口也会被划分成一个一个的小格子，就像我们家里的窗户那样。每个小格子都对应
着其底层数组中的某一个元素。
我们继续拿s2为例，这个窗口最左边的那个小格子对应的正好是其底层数组中的第一个元素， 即索引为0的那个元素。因此可以说，s2中的索引从0到4所指向的元素恰恰就是其底层数组中索 引从0到4代表的那 5 个元素。
请记住，当我们用make函数或切片值字面量(比如[]int{1, 2, 3})初始化一个切片时，该 窗口最左边的那个小格子总是会对应其底层数组中的第 1 个元素。