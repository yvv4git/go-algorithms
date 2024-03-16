package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	target         = "To be or not to be."
	charSet        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ.! "
	populationSize = 100
	mutationRate   = 0.01
	generations    = 10000
)

type Individual struct {
	DNA     []byte
	Fitness float64
}

func createIndividual() Individual {
	dna := make([]byte, len(target))
	for i := range dna {
		dna[i] = charSet[rand.Intn(len(charSet))]
	}
	return Individual{DNA: dna, Fitness: 0}
}

func (ind Individual) calculateFitness() float64 {
	score := 0
	for i := range ind.DNA {
		if ind.DNA[i] == target[i] {
			score++
		}
	}
	ind.Fitness = float64(score)
	return ind.Fitness
}

func crossover(parent1, parent2 Individual) Individual {
	child := Individual{DNA: make([]byte, len(target))}
	split := rand.Intn(len(target))
	for i := range child.DNA {
		if i > split {
			child.DNA[i] = parent1.DNA[i]
		} else {
			child.DNA[i] = parent2.DNA[i]
		}
	}
	return child
}

func mutate(ind Individual) Individual {
	for i := range ind.DNA {
		if rand.Float64() < mutationRate {
			ind.DNA[i] = charSet[rand.Intn(len(charSet))]
		}
	}
	return ind
}

func main() {
	/*
		Простой пример генетического алгоритма на Go, который будет пытаться найти строку, соответствующую заданной целевой строке.
		Это будет простой пример, и в реальных задачах оптимизации используются более сложные алгоритмы и структуры данных.

		В этом примере мы используем символьный набор для генерации новых решений, где каждый символ представляет собой возможный ген.
		Мы используем метод элитзма, где лучшие индивидуумы автоматически переходят в следующее поколение.
	*/
	rand.Seed(time.Now().UnixNano())

	population := make([]Individual, populationSize)
	for i := range population {
		population[i] = createIndividual()
	}

	for gen := 0; gen < generations; gen++ {
		for i := range population {
			population[i].calculateFitness()
		}

		sort.Slice(population, func(i, j int) bool {
			return population[i].Fitness > population[j].Fitness
		})

		if population[0].Fitness == float64(len(target)) {
			break
		}

		newPopulation := make([]Individual, populationSize)
		copy(newPopulation, population[:2]) // Elitism

		for i := 2; i < populationSize; i++ {
			parent1 := population[rand.Intn(5)]
			parent2 := population[rand.Intn(5)]
			child := crossover(parent1, parent2)
			child = mutate(child)
			newPopulation[i] = child
		}

		population = newPopulation
	}

	fmt.Println("Best solution:", string(population[0].DNA))
}
