package main

// clockAfter - определяет, сколько будет времени на часах, через заданное количество часов.
func clockAfter(nowTime, hourseAfter int) int {
	return (nowTime + hourseAfter) % 12
}
