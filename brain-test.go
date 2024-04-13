// squared error, int
// accuracy: 0.9691

package main

import (
	"brain"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// train
	file, err := os.Open("mnist/mnist_train.csv")
	if err != nil {
		log.Fatal(err)
	}
	sin := bufio.NewScanner(file)
	brain := brain.NewBrain()
	for sin.Scan() {
		line := sin.Text()
		split := strings.Split(line, ",")
		data := []int{}
		for _, v := range split {
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, int(n))
		}
		brain.Train(data[1:], data[0])
	}
	file.Close()
	fmt.Println("train data loaded. rows:", len(brain.Inputs))

	// test
	file, err = os.Open("mnist/mnist_test.csv")
	if err != nil {
		log.Fatal(err)
	}
	sin = bufio.NewScanner(file)
	nRows := 0
	nCorrect := 0
	for sin.Scan() {
		nRows++
		line := sin.Text()
		split := strings.Split(line, ",")
		data := []int{}
		for _, v := range split {
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, int(n))
		}
		// fmt.Println("data:", data)

		// predict
		best := brain.Guess(data[1:])
		if best == int(data[0]) {
			nCorrect++
		}
		fmt.Println("predicted:", best, "answer:", data[0], "accuracy:", float64(nCorrect)/float64(nRows))
	}
	file.Close()
	fmt.Println("train data loaded. rows:", nRows, "accuracy:", float64(nCorrect)/float64(nRows))
}
