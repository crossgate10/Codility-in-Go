You are given two non-empty arrays A and B consisting of N integers. These arrays represent N planks. More precisely, A[K] is the start and B[K] the end of the K−th plank.

Next, you are given a non-empty array C consisting of M integers. This array represents M nails. More precisely, C[I] is the position where you can hammer in the I−th nail.

We say that a plank (A[K], B[K]) is nailed if there exists a nail C[I] such that A[K] ≤ C[I] ≤ B[K].

The goal is to find the minimum number of nails that must be used until all the planks are nailed. In other words, you should find a value J such that all planks will be nailed after using only the first J nails. More precisely, for every plank (A[K], B[K]) such that 0 ≤ K < N, there should exist a nail C[I] such that I < J and A[K] ≤ C[I] ≤ B[K].

For example, given arrays A, B such that:

    A[0] = 1    B[0] = 4
    A[1] = 4    B[1] = 5
    A[2] = 5    B[2] = 9
    A[3] = 8    B[3] = 10
four planks are represented: [1, 4], [4, 5], [5, 9] and [8, 10].

Given array C such that:

    C[0] = 4
    C[1] = 6
    C[2] = 7
    C[3] = 10
    C[4] = 2
if we use the following nails:

0, then planks [1, 4] and [4, 5] will both be nailed.
0, 1, then planks [1, 4], [4, 5] and [5, 9] will be nailed.
0, 1, 2, then planks [1, 4], [4, 5] and [5, 9] will be nailed.
0, 1, 2, 3, then all the planks will be nailed.
Thus, four is the minimum number of nails that, used sequentially, allow all the planks to be nailed.

Write a function:

func Solution(A []int, B []int, C []int) int

that, given two non-empty arrays A and B consisting of N integers and a non-empty array C consisting of M integers, returns the minimum number of nails that, used sequentially, allow all the planks to be nailed.

If it is not possible to nail all the planks, the function should return −1.

For example, given arrays A, B, C such that:

    A[0] = 1    B[0] = 4
    A[1] = 4    B[1] = 5
    A[2] = 5    B[2] = 9
    A[3] = 8    B[3] = 10

    C[0] = 4
    C[1] = 6
    C[2] = 7
    C[3] = 10
    C[4] = 2
the function should return 4, as explained above.

Write an efficient algorithm for the following assumptions:

N and M are integers within the range [1..30,000];
each element of arrays A, B, C is an integer within the range [1..2*M];
A[K] ≤ B[K].

#### 心得
```
1. 二分搜尋+前綴合
N块木板，可以看作N线段，给定两个长度为N的正整数数组A[],B[]，[A[k],B[k]]表示木板（线段）的起点和终点，A[k] <= B[k]。有M个钉子，它们分别在长度为M的正整数数组里。钉子I可以固定住木板K，当且今当A[K]<=C[I]<=B[K]。问按顺序使用钉子，至少使用前多少个钉子可以固定住所有木板？无解返回-1。
数据范围： 木板数N和M的范围［1..30000]， A B C数组元素范围为[1..2 * M]
要求复杂度： 时间O((N+M)*log(M)) ， 空间O(M)
分析： 一个显然的并且符合要求的算法是二分答案，问题是如何判断，显然我们不能循环木板和钉子。但是我们可以计算从开头到当前位置一共有多少个钉子，这是前缀和的思想。计算前缀和需要O(M)，判断需要O(N)，二分是O(logM)，所以整好是要求的时间复杂度。空间上需要存前缀和O(M)。

2. 單調隊列(Monotone queue) 與 雙向隊列
更快的算法，如果我们建立一个长度为2 * M的数组，每个位置表示该位置上钉子的最小编号（可能同一个位置有多个钉子，取编号最小的），没有钉子的位置值为无穷大。那么固定第i块木板的最小编号钉子，相当于[A[i],B[i]]区间的最小值。但是我们这个题实际上是求这些最小值的max，首先如果一个木板A的覆盖区间完全包含另外一个木板B,则实际上我们只考虑木板B即可。因为固定木板B同时能固定木板A，并且我们一定要固定木板B，即使A覆盖区间有更小的值，也无法改变最终取最大值的结果。
于是，我们可以建立一个数组plank[x]表示右端点为x的木板的最大左边界，没有木板的话，认为边界是0。我们从左到右遍历木板右边界，假设这之前（更左）的木板已经被固定了，已经固定的区间范围是[left,right) (右开区间），然后对当前这个木板，如果显然它的右边界更大（我们遍历右边界是按当增的顺序），如果该木板start <= left,则它已经被前面固定住了，不影响结果。否则，要求[start,end]之间的最大值。这个问题有点像滑动窗口最大值的问题。本质在于：我们不断查询最大值，每次查询的时候窗口的左边界和右有边界是单调递增的，于是我们可以动态更新窗口维护最大值。这个经典问题可以用单调队列实现，这也是把单调队列发挥到了极致。
结论： 查询窗口最大值的时候，如果窗口向右滑动的过程中，查询时左边界和右边界都是单增的，则可以使用单调队列解决。
时间复杂度 O(N + M)达到了线性

ref: http://uml.idehub.cn/article/282
```
