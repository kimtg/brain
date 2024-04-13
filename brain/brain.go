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

type Brain struct {
	Inputs  [][]int
	Outputs []int
}

func NewBrain() Brain {
	var b Brain
	b.Inputs = [][]int{}
	b.Outputs = []int{}
	return b
}

func (b *Brain) Train(input []int, output int) {
	b.Inputs = append(b.Inputs, input)
	b.Outputs = append(b.Outputs, output)
}

func (b *Brain) Guess(input []int) int {
	best := 0
	minError := -1
	for i, d := range b.Inputs {
		error := sumSqErr(input, d)
		if minError < 0 || error < minError {
			minError = error
			best = int(b.Outputs[i])
		}
	}
	return best
}
