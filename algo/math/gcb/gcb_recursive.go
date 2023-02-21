package gcb

func gcb(m, n int) int {
	/*
		GCB - наибольший общий делитель находим через рекурсивный алгоритм.
	*/
	if m < n {
		return gcb(n, m) // Шаг рекурсии(функция продолжает вызывать саму себя)
	}

	if m%n == 0 {
		return n // Базовый случай(дальнейший вызов самой себя должен прекратится и осуществляется выход)
	}

	return gcb(n, m%n)
}
