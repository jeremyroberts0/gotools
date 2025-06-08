package slicetools

func Map[in, out any](inputSlice []in, mapper func(in) out) []out {
	outputSlice := []out{}
	for _, val := range inputSlice {
		outputSlice = append(outputSlice, mapper(val))
	}
	return outputSlice
}
