package diffsquares

func SquareOfSum(n int) int {
	return ((n*n + n)/2)*((n*n + n)/2)
}

func SumOfSquares(n int) int {
	return ((2*n + 1)*n*(n + 1))/6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
