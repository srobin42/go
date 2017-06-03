//Package pascal computes Pascals Triangle to the nth level.
package pascal

const testVersion = 1

//Triangle computes and returns Pascals Triangle, output int array.
func Triangle(i int) [][]int {
	tri := [][]int{}

	//value of row one always one
	tri = append(tri, []int{1})
	//build the rest of the triangle, if applicable
	for j := 0; j < i-1; j++ {
		newRow := []int{}
		prev := 0
		for _, val := range tri[j] {
			newRow = append(newRow, val+prev)
			prev = val
		}
		newRow = append(newRow, 1)
		tri = append(tri, newRow)
	}
	return tri
}
