package main

func hammingWeight(num uint32) int {
	count := uint32(0)

	for ; num > uint32(0); num >>= 1 {
		count += num & 1
	}

	return int(count)
}

/* 
Т.е. каждый раз мы сдвигаем наше число вправо.
num=11[00001011]  count=1
num=5[00000101]  count=2
num=2[00000010]  count=2
num=1[00000001]  count=3
*/