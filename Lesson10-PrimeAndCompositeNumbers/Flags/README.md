#### 心得
```
題目限制：
1. 0 < P < N − 1
2. A[P − 1] < A[P] > A[P + 1]
3. if you take K flags, 
then the distance between any two flags should be greater than or equal to K

心得：
1. 先找出peak的數目，再計算第一座peak與最後一座peak的平方根，此即為理論上的最大旗子數量
2. 依照求出的最大旗子數量理論值反向依序，遍歷peaks[0]-[-1]，當flag_count大於等於flag後回傳

Codility 官方解答 https://codility.com/media/train/solution-flags.pdf
```
