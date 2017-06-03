//Package diffsquares finds the difference between the square of the sum and the sum of the squares of the first N natural numbers, ouput int.
package diffsquares

const testVersion = 1

//Difference returns the difference between square of sums and sum of squares of input, output int.
func Difference(i int) int {
	var sum, squ int

	sum = SumOfSquares(i)
	squ = SquareOfSums(i)
	return squ - sum

}

//SumOfSquares calculates the sum of squares of input, output int.
func SumOfSquares(i int) int {
	var sum int

	for i > 0 {
		sum += i * i
		i--
	}
	return sum
}

//SquareOfSums calculates square of sum of input, output int.
func SquareOfSums(i int) int {
	var square int

	for i > 0 {
		square += i
		i--
	}
	return square * square
}
