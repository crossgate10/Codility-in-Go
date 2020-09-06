package SieveOfEratosthenes

// 可以篩出小於 num 的質數
func Sieve(num int) []int {
	prime := make([]bool, num+1)

	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true
			}
		}
	}
	return getPrime(num, prime)
}

func getPrime(num int, prime []bool) []int {
	result := []int{}
	for i := 2; i <= num; i++ {
		if prime[i] == false {
			result = append(result, i)
		}
	}
	return result
}