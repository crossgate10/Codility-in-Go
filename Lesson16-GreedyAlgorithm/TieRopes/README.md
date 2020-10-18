
There are N ropes numbered from 0 to N − 1, whose lengths are given in an array A, lying on the floor in a line. For each I (0 ≤ I < N), the length of rope I on the line is A[I].

We say that two ropes I and I + 1 are adjacent. Two adjacent ropes can be tied together with a knot, and the length of the tied rope is the sum of lengths of both ropes. The resulting new rope can then be tied again.

For a given integer K, the goal is to tie the ropes in such a way that the number of ropes whose length is greater than or equal to K is maximal.

For example, consider K = 4 and array A such that:

    A[0] = 1
    A[1] = 2
    A[2] = 3
    A[3] = 4
    A[4] = 1
    A[5] = 1
    A[6] = 3
The ropes are shown in the figure below.

![avatar][figure1]

We can tie:

rope 1 with rope 2 to produce a rope of length A[1] + A[2] = 5;
rope 4 with rope 5 with rope 6 to produce a rope of length A[4] + A[5] + A[6] = 5.
After that, there will be three ropes whose lengths are greater than or equal to K = 4. It is not possible to produce four such ropes.

Write a function:

func Solution(K int, A []int) int

that, given an integer K and a non-empty array A of N integers, returns the maximum number of ropes of length greater than or equal to K that can be created.

For example, given K = 4 and array A such that:

    A[0] = 1
    A[1] = 2
    A[2] = 3
    A[3] = 4
    A[4] = 1
    A[5] = 1
    A[6] = 3
the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
K is an integer within the range [1..1,000,000,000];
each element of array A is an integer within the range [1..1,000,000,000].

#### 心得
```
題意：從A中找出最多可以產生幾條長度大於K的繩子
```


[figure1]:data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAloAAAAOCAQAAABd0nqpAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAgY0hSTQAAeiYAAICEAAD6AAAAgOgAAHUwAADqYAAAOpgAABdwnLpRPAAAAAJiS0dEAP+Hj8y/AAAACXBIWXMAAABIAAAASABGyWs+AAACNklEQVR42u2bu43bQBCG//9wgGNWYGBb2AKcsAXC2YXrEnQlyCVIueHgWmDiAlTCCXAFzA5wNA7Ms+8Ag8QOtY+R9mPADbTAPzsP7ksU2IKdTKU1XMIK4BrsaDT+R9r45lrR4gCPUcbSwzCr2eGMDpM8lVayyYoBsG/FBvs7HHCsJabUVmTOjLoycUVp0vi+X/uBPLGmGcFJRoB7mE73P860bsUGAo6lJWwnd2ZUlomLSoGU8b1atOpCRoAeBr42y7BHf6slix6n0hoaaUkb33elzYuFHs7CFHkZGfEVfWkVhejt+6+xTNr4NjbTooe7jp0gmVhaQilO7OFhZKnT0JEyvleLFgM8QCcV7ELQY48TPTr5UlrLJjsCJnQ4l9ZRBhnpMNgvWrkzo6ZMXFWaNL5XTw8bKaDD1K48NK6VtPF9BwAM3LFTSOu5p9OYxD2z7ujUplTOl3YpO+4YVD2V3reD1ov5x9SO0mUuH9/vcliAZwgEAieIeXCYGyGyX5gbh7h++seOUrWFbm48R/dUet/Oo/Vi/jG1o7SAF9/k8D17vM5AvvFHROH7gIe5NUTeuhnm9wNf8OvS9di0Uj2f5rdj1Nmc2vuW+Dy/Bz5Gff9fZy6O3/Ezot/Hv2Ma8BilVBtv1+/9dzls7spDo9G4cdrysCalagtvYIGgHhsziy47Sgt48U0OUwAwoMMxfuuMPXocJfpokw4h77+o7ChVW9ghYNIch2u9bwetF/OPqR2l+fmXw78BB2dQVEv4nZsAAAAldEVYdGRhdGU6Y3JlYXRlADIwMTYtMTEtMTVUMTc6NTE6MjcrMDA6MDC+mbGlAAAAJXRFWHRkYXRlOm1vZGlmeQAyMDE2LTExLTE1VDE3OjUxOjI3KzAwOjAwz8QJGQAAAB50RVh0cHM6SGlSZXNCb3VuZGluZ0JveAA2MDJ4MTQtMy0y07+P0AAAABx0RVh0cHM6TGV2ZWwAQWRvYmUtMy4wIEVQU0YtMy4wCptwu+MAAAAZdEVYdHBzOlNwb3RDb2xvci0wAGZvbnQgQ01SMTDtW08rAAAAAElFTkSuQmCC