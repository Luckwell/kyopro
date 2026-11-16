package lib

import "math"

func PowMod(a, b, mod int) int {
	if mod == 0 {
		mod = math.MaxInt
	}
	ret := 1
	power := (a%mod + mod) % mod
	if b < 0 {
		var x, y int
		_ = Extgcd(power, mod, &x, &y)
		return PowMod(x, -b, mod)
	}
	for b > 0 {
		if b&1 != 0 {
			ret *= power
			ret %= mod
		}
		power *= power
		power %= mod
		b >>= 1
	}
	return ret
}

func Extgcd(a, b int, x, y *int) int {
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}
	d := Extgcd(b, a%b, y, x)
	*y -= (a / b) * (*x)
	return d
}

func Sieve(N int) []bool {
	ret := make([]bool, N+1)
	if N < 2 {
		return ret
	}
	for i := 2; i < N+1; i++ {
		ret[i] = true
	}
	for i := 2; i*i <= N; i++ {
		if ret[i] {
			for j := i * i; j <= N; j += i {
				ret[j] = false
			}
		}
	}
	return ret
}

func Primes(N int) []int {
	isPrime := Sieve(N)
	var ret []int
	for i := 2; i <= N; i++ {
		if isPrime[i] {
			ret = append(ret, i)
		}
	}
	return ret
}
