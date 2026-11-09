package main

import (
	"kyopro/lib"
	"os"
)

func main() {
	fastin := lib.NewScanner(os.Stdin)
	fastout := lib.NewWriter(os.Stdout)
	defer fastout.Flush()

	T := fastin.Int()
	for range T {
		A := string(fastin.Bytes())
		B := string(fastin.Bytes())
		zarr := lib.ZAlgorithm(B + A + A)
		ans := false
		for i := range A {
			if zarr[len(A)+i] >= len(A) {
				fastout.PrintInt(i)
				fastout.PrintByte('\n')
				ans = true
				break
			}
		}
		if ans {
			continue
		}
		fastout.PrintInt(-1)
		fastout.PrintByte('\n')
	}
}
