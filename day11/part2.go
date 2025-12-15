package main

func Part2(lines []string) int {
	// inputs, _ := ParseCables(lines)
	_, outputs := ParseCables(lines)
	// fmt.Println(inputs["you"])
	// fmt.Println(outputs["you"])
	paths := Paths2ToOut(outputs, "svr", false, false)
	return paths
}

type paths2State struct {
	node   string
	hitDAC bool
	hitFFT bool
}

// Paths2ToOut counts paths from node to "out" that pass through both "dac" and "fft".
// For large graphs with redundant structure, this is heavily memoized by (node, hitDAC, hitFFT).
func Paths2ToOut(outputs map[string][]string, node string, hitDAC, hitFFT bool) int {
	cache := make(map[paths2State]int)
	return paths2ToOutMemo(outputs, node, hitDAC, hitFFT, cache)
}

func paths2ToOutMemo(
	outputs map[string][]string,
	node string,
	hitDAC, hitFFT bool,
	cache map[paths2State]int,
) int {
	st := paths2State{node: node, hitDAC: hitDAC, hitFFT: hitFFT}
	if v, ok := cache[st]; ok {
		return v
	}

	if node == "out" {
		if hitDAC && hitFFT {
			cache[st] = 1
			return 1
		}
		cache[st] = 0
		return 0
	}

	yesDAC := hitDAC || node == "dac"
	yesFFT := hitFFT || node == "fft"
	pathsCount := 0
	for _, o := range outputs[node] {
		pathsCount += paths2ToOutMemo(outputs, o, yesDAC, yesFFT, cache)
	}
	cache[st] = pathsCount
	return pathsCount
}
