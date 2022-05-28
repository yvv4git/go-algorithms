package main

func sum(x, y, field int) int {
	return (x+y) % field
}

func inversion(x, field int) int {
	return x % field
}