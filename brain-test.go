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
	nRowsTrain := 0
	for sin.Scan() {
		nRowsTrain++
		line := sin.Text()
		split := strings.Split(line, ",")
		data := []int{}
		for _, v := range split {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, n)
		}
		brain.Train(data[1:], data[0])
	}
	file.Close()
	fmt.Println("train data loaded. rows:", nRowsTrain)

	// test
	file, err = os.Open("mnist/mnist_test.csv")
	if err != nil {
		log.Fatal(err)
	}
	sin = bufio.NewScanner(file)
	nRowsTest := 0
	nCorrect := 0
	for sin.Scan() {
		nRowsTest++
		line := sin.Text()
		split := strings.Split(line, ",")
		data := []int{}
		for _, v := range split {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, n)
		}
		// fmt.Println("data:", data)

		// predict
		best := brain.Guess(data[1:])
		if best == data[0] {
			nCorrect++
		}
		fmt.Println("predicted:", best, "answer:", data[0], "accuracy:", float64(nCorrect)/float64(nRowsTest))
	}
	file.Close()
}
