package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Syntax error: Must specify CSV file\n Example: meshcalc input.csv\n")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	r := csv.NewReader(f)
	points, err := r.ReadAll()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	if len(points) < 2 {
		log.Fatalf("Invalid number of measurements: must be 2 or greater (got %d)\n", len(points))
	}
	if len(points[0]) != 2 {
		log.Fatalf("Invalid measurement record: must be 2 numbers representing plan-view distance from previous point in metres, and measured height in metres. (got %d numbers)\n", len(points[1]))
	}
	area := 0.0
	spans := len(points) - 1
	for s := 0; s < spans; s++ {
		thisHeight, err := strconv.ParseFloat(points[s][1], 64)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		nextHeight, err := strconv.ParseFloat(points[s+1][1], 64)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		width, err := strconv.ParseFloat(points[s+1][0], 64)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		area += ((thisHeight + nextHeight) / 2.0) * width
	}
	fmt.Printf("Area: %.2fm²\n", area)
	os.Exit(0)
}
