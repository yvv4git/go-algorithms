package main

import (
	"hash"
	"hash/fnv"
	"math"
)

// CountMinSketch структура данных
type CountMinSketch struct {
	matrix [][]uint
	hashes []hash.Hash64
}

// NewCountMinSketch создает новый экземпляр CountMinSketch
func NewCountMinSketch(width, depth int) *CountMinSketch {
	matrix := make([][]uint, depth)
	for i := range matrix {
		matrix[i] = make([]uint, width)
	}
	hashes := make([]hash.Hash64, depth)
	for i := range hashes {
		hashes[i] = fnv.New64a()
	}
	return &CountMinSketch{
		matrix: matrix,
		hashes: hashes,
	}
}

// Add добавляет элемент в CountMinSketch
func (c *CountMinSketch) Add(item string) {
	for i, h := range c.hashes {
		h.Reset()
		h.Write([]byte(item))
		index := h.Sum64() % uint64(len(c.matrix[i]))
		c.matrix[i][index]++
	}
}

// Count возвращает приблизительное количество элемента в CountMinSketch
func (c *CountMinSketch) Count(item string) uint {
	minCount := uint(math.MaxUint64)
	for i, h := range c.hashes {
		h.Reset()
		h.Write([]byte(item))
		index := h.Sum64() % uint64(len(c.matrix[i]))
		if c.matrix[i][index] < minCount {
			minCount = c.matrix[i][index]
		}
	}
	return minCount
}
