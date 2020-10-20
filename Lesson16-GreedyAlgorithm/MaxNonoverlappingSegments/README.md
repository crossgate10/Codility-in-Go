Located on a line are N segments, numbered from 0 to N − 1, whose positions are given in arrays A and B. For each I (0 ≤ I < N) the position of segment I is from A[I] to B[I] (inclusive). The segments are sorted by their ends, which means that B[K] ≤ B[K + 1] for K such that 0 ≤ K < N − 1.

Two segments I and J, such that I ≠ J, are overlapping if they share at least one common point. In other words, A[I] ≤ A[J] ≤ B[I] or A[J] ≤ A[I] ≤ B[J].

We say that the set of segments is non-overlapping if it contains no two overlapping segments. The goal is to find the size of a non-overlapping set containing the maximal number of segments.

For example, consider arrays A, B such that:

    A[0] = 1    B[0] = 5
    A[1] = 3    B[1] = 6
    A[2] = 7    B[2] = 8
    A[3] = 9    B[3] = 9
    A[4] = 9    B[4] = 10
The segments are shown in the figure below.

![avatar][figure1]

The size of a non-overlapping set containing a maximal number of segments is 3. For example, possible sets are {0, 2, 3}, {0, 2, 4}, {1, 2, 3} or {1, 2, 4}. There is no non-overlapping set with four segments.

Write a function:

func Solution(A []int, B []int) int

that, given two arrays A and B consisting of N integers, returns the size of a non-overlapping set containing a maximal number of segments.

For example, given arrays A, B shown above, the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..30,000];
each element of arrays A, B is an integer within the range [0..1,000,000,000];
A[I] ≤ B[I], for each I (0 ≤ I < N);
B[K] ≤ B[K + 1], for each K (0 ≤ K < N − 1).

#### 心得
貪婪演算法（Greedy algorithm）與動態規劃（Dynamic Programming）差異：
||貪婪| 動態規劃|
|-|-|-|
|有無遞迴|無|有使用遞迴|
|解|局部最佳解|全域最佳解|
|記憶體使用|因為不需修改前次紀錄，記憶體更為高效|需要存先前計算結果|
|效率|通常較快|通常較慢|


[figure1]:data:image/*;base64,iVBORw0KGgoAAAANSUhEUgAAAQYAAAAkCAQAAADBL1w+AAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAgY0hSTQAAeiYAAICEAAD6AAAAgOgAAHUwAADqYAAAOpgAABdwnLpRPAAAAAJiS0dEAP+Hj8y/AAAACXBIWXMAAABIAAAASABGyWs+AAACiUlEQVR42u2csY7TQBBA35yQoN0OCYrT8gmmoLvG9wkRBVcb/iBI/ECu4AMS0Z0oIA0VDW7oKEh7okmE+ADc0KKhOMu5Khevgydx5qXZIp6ZHY13ZjezEcUZChJAqw7PDycYZERGqaW1HYbzh0Cl82QJwwkGkByONxhqH0z0deqz96yNd3aH5OQkrwtwYj0BZ3doySV5+vMeDIOiS/k4qDQhBRlI1Jm1JWbzrwisOkgYUgF57Eik8q2lsxNMagbJZCofZWQ9+UFZ+kKu5Ze87STDYmWQ3wQAnurCQP0wLf1bv9gX+j5VhkEBKUXtYPgkP/rX34JHjaUFL62N2YRcNWv8Gw4pGFifEc75bKB/ex5yVY/2fF3gHRc7sFQNPkxRlO8EC+0DtfQaRfnDaboMo92ERMK+Z+GDs/SMx+n1AvjW0rmFH0c7DR4MToMHg9PgweA0eDA4DR4MToNZMEjoLuNw2c/ZWx06jYhUrI6zfVXGrLr1MSfoDEyZbfa31coQ9VJnZEbarVnovPfZF9zZAWYVDDf9OAdxzLt7tATJ6HFVlGwbX3sBaYJkxF5TZL6NNqtguCmgjjRNSEbss14AFpKTETd/yao7eiVjquNME5IxYSEZQXtrmNFSIiPuaJY1+9Wy6yVRZ/ecAEguE4ldRbVDOzV1t0WiTKTDXaMEjYk+lSBjKZI0FjJOOb+4ZWnTzaMU1t06/60LqKgH0940JvqUWA+WrTUu60FMt1QILOty7icf+ntzeuU5pwBUPOljPZJ0n57xrB6dt9ltSM6XeviNry303ecVDwAo9dy3ls4aTxOeJpo0odT3+mfa4crmviORot9/dUn1qQQKqpTLw1IQmLVPg2tL/wFaD5rWAmhwbQAAACV0RVh0ZGF0ZTpjcmVhdGUAMjAxNi0xMS0xNVQxNzo1MDowNSswMDowMIThzM8AAAAldEVYdGRhdGU6bW9kaWZ5ADIwMTYtMTEtMTVUMTc6NTA6MDUrMDA6MDD1vHRzAAAAH3RFWHRwczpIaVJlc0JvdW5kaW5nQm94ADI2MngzNisyNS0zZNE+tAAAABx0RVh0cHM6TGV2ZWwAQWRvYmUtMy4wIEVQU0YtMy4wCptwu+MAAAAZdEVYdHBzOlNwb3RDb2xvci0wAGZvbnQgQ01SMTDtW08rAAAAAElFTkSuQmCC

