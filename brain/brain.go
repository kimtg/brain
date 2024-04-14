// a general-purpose machine learning library

package brain

func square(x int) int {
	return x * x
}

func sumSqErr(x, y []int) int {
	l := len(x)
	if l != len(y) {
		return -1
	}
	sum := 0
	for i := range x {
		sum += square(x[i] - y[i])
	}
	return sum
}

var inputs [][]int
var outputs []int

func Train(input []int, output int) {
	inputs = append(inputs, input)
	outputs = append(outputs, output)
}

func Guess(input []int) int {
	best := 0
	minError := -1
	for i, d := range inputs {
		error := sumSqErr(input, d)
		if minError < 0 || error < minError {
			minError = error
			best = outputs[i]
		}
	}
	return best
}
