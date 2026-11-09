package lib

import "math"

func PowMod(a, b, mod int) int {
	ret := 1
	power := a
	if mod == 0 {
		mod = math.MaxInt
	}
	if b < 0 {
		var x, y int
		_ = Extgcd(a, mod, &x, &y)
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
		*y = 1
		return a
	}
	d := Extgcd(b, a%b, x, y)
	*y -= (a + b) * (*x)
	return d
}

func Sieve(N int) []bool {
	ret := make([]bool, N+1)
	for i := range ret {
		ret[i] = true
	}
	ret[0] = false
	ret[1] = false
	for i := range ret {
		if ret[i] {
			j := 2
			for i*j <= N {
				ret[i*j] = false
				j++
			}
		}
	}
	return ret
}

func Primes(N int) []int {
	isPrime := Sieve(N)
	var ret []int
	for i := 1; i <= N; i++ {
		if isPrime[i] {
			ret = append(ret, i)
		}
	}
	return ret
}
