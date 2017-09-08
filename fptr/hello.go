package fptr

type arithmeticOperation func(int, int)int

func SayHello(suffix string) string{

	return "Hello," + suffix
}

func Calculate(ao arithmeticOperation, x, y int) int {
	return 	ao(x, y)
}

func Plus(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Multiply(a,b int) int {
	return a * b
}
