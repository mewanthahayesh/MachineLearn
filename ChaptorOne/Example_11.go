package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./CData/labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	var observed []int
	var predicted []int
	line := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if line == 1 {
			line++
			continue
		}
		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
			continue
		}
		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
			continue
		}

		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++

	}
	var truePosNeg int
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}
	accuracy := float64(truePosNeg) / float64(len(observed))
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}
