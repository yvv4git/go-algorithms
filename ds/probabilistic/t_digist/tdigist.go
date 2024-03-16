package main

import (
	"sort"
)

type Centroid struct {
	Mean   float64
	Weight int
}

type TDigest struct {
	Centroids   []Centroid
	Size        int
	Compression float64
}

func NewTDigest(compression float64) *TDigest {
	return &TDigest{
		Centroids:   make([]Centroid, 0),
		Size:        0,
		Compression: compression,
	}
}

func (td *TDigest) Add(x float64, w int) {
	td.Size += w
	added := false

	// Проверяем, может ли новый центроид быть добавлен в существующий кластер
	for i := range td.Centroids {
		if td.Centroids[i].Mean == x {
			td.Centroids[i].Weight += w
			added = true
			break
		}
	}

	// Если новый центроид не был добавлен, пробуем добавить его в существующий кластер
	if !added {
		for i := range td.Centroids {
			if td.Centroids[i].Weight+w <= int(td.Compression) {
				td.Centroids[i].Mean = (td.Centroids[i].Mean*float64(td.Centroids[i].Weight) + x*float64(w)) / float64(td.Centroids[i].Weight+w)
				td.Centroids[i].Weight += w
				added = true
				break
			}
		}
	}

	// Если новый центроид не был добавлен, создаем новый кластер
	if !added {
		td.Centroids = append(td.Centroids, Centroid{Mean: x, Weight: w})
	}

	// Сортируем центроиды по возрастанию значений
	sort.Slice(td.Centroids, func(i, j int) bool {
		return td.Centroids[i].Mean < td.Centroids[j].Mean
	})
}

func (td *TDigest) Quantile(q float64) float64 {
	targetWeight := q * float64(td.Size)
	currentWeight := 0.0

	for _, centroid := range td.Centroids {
		currentWeight += float64(centroid.Weight)
		if currentWeight >= targetWeight {
			return centroid.Mean
		}
	}

	return 0.0
}
