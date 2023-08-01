package heap_rune

type runeSlice []rune

func (r runeSlice) Len() int           { return len(r) }
func (r runeSlice) Less(i, j int) bool { return r[i] > r[j] }
func (r runeSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func (r *runeSlice) Push(x interface{}) {
	*r = append(*r, x.(rune))
}

func (r *runeSlice) Pop() interface{} {
	x := (*r)[len(*r)-1]
	*r = (*r)[:len(*r)-1]
	return x
}
