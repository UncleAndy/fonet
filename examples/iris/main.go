package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/UncleAndy/fonet"
	"github.com/gocarina/gocsv"
)

type IrisCase struct {
	SepalLength float64 `csv:"sepal_length"`
	SepalWidth  float64 `csv:"sepal_width"`
	PetalLength float64 `csv:"petal_length"`
	PetalWidth  float64 `csv:"petal_width"`
	Setosa      float64 `csv:"setosa"`
	Virginica   float64 `csv:"virginica"`
	Versicolor  float64 `csv:"versicolor"`
}

func main() {
	n, err := fonet.NewNetwork([]int{4, 5, 5, 3}, fonet.Sigmond, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	samples := makeSamples("train.csv")
	log.Println("Training started!")
	n.Train(samples, 10000, 1.111, false)
	log.Println("Training finished!")
	tests := makeSamples("test.csv")
	for _, t := range tests {
		fmt.Printf("Predicted: %v ->", n.Predict(t[0]))
		for _, p := range n.Predict(t[0]) {
			fmt.Printf(" %v", math.Round(p))
		}
		fmt.Printf(", Expected: %v\n", t[1])
	}
}

func makeSamples(path string) [][][]float64 {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var cases []IrisCase
	err = gocsv.Unmarshal(f, &cases)
	var out [][][]float64
	for _, c := range cases {
		out = append(out, [][]float64{
			[]float64{
				c.SepalLength,
				c.SepalWidth,
				c.PetalLength,
				c.PetalWidth,
			},
			[]float64{
				c.Setosa,
				c.Virginica,
				c.Versicolor,
			},
		})
	}
	return out
}
