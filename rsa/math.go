package rsa

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}

	return gcd(y, x%y)
}

func modInverse(x, y int) int {
	i := 1
	for i < y {
		if ((x%y)*(i%y))%y == 1 {
			return i
		}
		i++
	}
	return -1
}

func NthPrime(n int) int {
	i := 1
	num := 2

	for i != n {
		num++
		if isPrime(num) {
			i++
		}
	}

	return num
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
