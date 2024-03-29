package main

import "fmt"

func main() {
	/*
		В этом фрагменте кода создается новый экземпляр TDigest с параметром компрессии 100.
		Параметр компрессии определяет, насколько "тонкие" будут кластеры в T-Digest.
		Чем больше компрессия, тем больше вероятность того, что новые центроиды будут создаваться в отдельных кластерах.
		В данном случае мы используем значение 100, что означает, что каждый центроид будет отдельным кластером.
	*/
	td := NewTDigest(100)

	// Затем мы добавляем несколько значений в T-Digest с весом 1 каждое.
	// Это означает, что каждое добавленное значение будет считаться отдельным центроидом.
	td.Add(1.0, 1)
	td.Add(2.0, 1)
	td.Add(3.0, 1)
	td.Add(4.0, 1)
	td.Add(5.0, 1)

	/*
		После добавления значений, мы вычисляем квантили для медианы (50-го процентиля), 90-го и 95-го процентилей.
		Функция Quantile вычисляет значение, которое находится на заданном процентиле распределения.
		Например, медиана - это центр распределения, то есть 50-й процентиль. 90-й процентиль - это значение, ниже которого находится 90% всех значений.
		Аналогично, 95-й процентиль - это значение, ниже которого находится 95% всех значений.
	*/
	fmt.Println("Median:", td.Quantile(0.5))
	fmt.Println("90th Percentile:", td.Quantile(0.9))
	fmt.Println("95th Percentile:", td.Quantile(0.95))
}
