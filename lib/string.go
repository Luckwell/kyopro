package lib

type RLE struct {
	Char  byte
	Count int
}

func ZAlgorithm(s string) []int {
	zArray := make([]int, len(s))
	i, j := 1, 0
	for i < len(s) {
		if j+zArray[j] <= i {
			zArray[i] = 0
		} else {
			zArray[i] = min(j+zArray[j]-i, zArray[i-j])
		}
		for i+zArray[i] < len(s) && s[zArray[i]] == s[i+zArray[i]] {
			zArray[i]++
		}
		if j+zArray[j] <= i+zArray[i] {
			j = i
		}
		i++

	}
	return zArray
}

func Runlength(s string) []RLE {
	var ret []RLE
	for i := range s {
		if i == 0 {
			ret = append(ret, RLE{s[i], 1})
		} else {
			if ret[len(ret)-1].Char == s[i] {
				ret[len(ret)-1].Count++
			} else {
				ret = append(ret, RLE{s[i], 1})
			}
		}
	}
	return ret
}
