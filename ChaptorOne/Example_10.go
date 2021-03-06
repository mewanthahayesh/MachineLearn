package main

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/gonum/stat"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./CData/continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)

	var observed []float64
	var predicted []float64

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

		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed,unexpected type \n", line)
			continue
		}
		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed,unexpected type \n", line)
			continue
		}
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)

	}
	var mAE float64
	var mSE float64
	for idx, oVal := range observed {
		mAE += math.Abs(oVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(oVal-predicted[idx], 2) / float64(len(observed))
	}
	var mean = stat.Mean(observed, nil)
	var rsq = stat.RSquaredFrom(observed, predicted, nil)

	fmt.Printf("Mean of Observed %0.2f\n", mean)
	fmt.Printf("\n nMAE=%0.2f\n", mAE)
	fmt.Printf("\n nMSE=%0.2f\n", mSE)
	fmt.Printf("\n RSquared=%0.2f\n", rsq)

}
