package calc

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func Add[T Number](a, b T) T {
	return a + b
}

func Sub[T Number](a, b T) T {
	return a - b
}

func Mult[T Number](a, b T) T {
	return a * b
}

func Div[T Number](a, b T) T {
	return a / b
}
