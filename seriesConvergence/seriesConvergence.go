package seriesconvergence

func NextElement(n int) int {
	if (n % 2) == 0 {
		return int(n / 2)
	}
	return (n * 3) + 1
}

func ChainLength(mem map[int]int, n int) int {
	if n == 1 {
		return 1
	}
	if val, inMap := mem[n]; inMap {
		return val
	} else {
		nxtElem := NextElement(n)
		subLen := ChainLength(mem, nxtElem)
		len := subLen + 1
		mem[n] = len
		return len
	}
}

func KeyOfMaxValueInMap(mp map[int]int) (int, int) {
	maxKey := -1
	maxElement := -1
	for key, element := range mp {
		if element > maxElement {
			maxKey = key
			maxElement = element

		}
	}
	return maxKey, maxElement
}

func LongestChainUnder(num int) (int, int) {
	mem := make(map[int]int)
	for i := 1; i <= num; i++ {
		ChainLength(mem, i)
	}
	maxKey, maxElement := KeyOfMaxValueInMap(mem)
	return maxKey, maxElement
}
