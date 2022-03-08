package fibonacci

// Fibonacci - hold two fibonacci integer
type Fibonacci struct {
	x1, x2 int
}

// NewFibonacci - simple factory for create instance Fibonacci
func NewFibonacci() *Fibonacci {
	return &Fibonacci{-1, 1}
}

// Next - used for getting new fibonacci digitals
func (f *Fibonacci) Next() int {
	f.x1, f.x2 = f.x2, f.x1+f.x2
	return f.x2
}

// fibonacciByStruct - used for find fibonacci digital by num
func fibonacciByStruct(n int) int {
	fibo := NewFibonacci()

	result := 0
	for i := 0; i <= n; i++ {
		result = fibo.Next()
	}

	return result
}
