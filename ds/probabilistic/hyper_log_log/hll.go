package main

import (
	"hash"
	"hash/fnv"
	"math"
)

type HyperLogLog struct {
	registers []int
	hashFunc  hash.Hash64
	m         uint32
	b         uint8
}

func NewHyperLogLog(b uint8) *HyperLogLog {
	m := uint32(1) << b
	return &HyperLogLog{
		registers: make([]int, m),
		hashFunc:  fnv.New64(),
		m:         m,
		b:         b,
	}
}

func (h *HyperLogLog) Add(item []byte) {
	h.hashFunc.Reset()
	h.hashFunc.Write(item)
	hashValue := h.hashFunc.Sum64()
	index := hashValue >> (64 - h.b)
	leadingZeros := uint8(1)
	for i := uint8(0); i < h.b; i++ {
		if (hashValue & (1 << i)) == 0 {
			leadingZeros++
		} else {
			break
		}
	}
	if h.registers[index] < int(leadingZeros) {
		h.registers[index] = int(leadingZeros)
	}
}

func (h *HyperLogLog) Count() float64 {
	sum := 0.0
	for _, v := range h.registers {
		sum += math.Pow(2, -float64(v))
	}
	estimate := h.alpha() * float64(h.m*h.m) / sum
	if estimate <= 5.0/2.0*float64(h.m) {
		v := 0
		for _, r := range h.registers {
			if r == 0 {
				v++
			}
		}
		if v != 0 {
			return float64(h.m) * math.Log(float64(h.m)/float64(v))
		}
	} else if estimate > 1.0/30.0*math.Pow(2, 32) {
		return -math.Pow(2, 32) * math.Log(1-estimate/math.Pow(2, 32))
	}
	return estimate
}

func (h *HyperLogLog) alpha() float64 {
	switch h.m {
	case 16:
		return 0.673
	case 32:
		return 0.697
	case 64:
		return 0.709
	default:
		return 0.7213 / (1 + 1.079/float64(h.m))
	}
}
