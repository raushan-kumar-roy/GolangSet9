package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	csvFile, err := os.Open("Cricket_Players_Stats.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	_, err = csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	newCsvFile, err := os.Create("modified_Cricket_Players_Stats.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer newCsvFile.Close()

	csvWriter := csv.NewWriter(newCsvFile)

	csvWriter.Write([]string{"Name", "Matches", "Runs"})

	for {
		row, err := csvReader.Read()
		if err != nil {
			break
		}

		name := row[0]
		matches, err := strconv.Atoi(row[2])
		if err != nil {
			log.Fatal(err)
		}
		runs, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatal(err)
		}

		matches += 10
		runs += 1000

		csvWriter.Write([]string{name, strconv.Itoa(matches), strconv.Itoa(runs)})
	}

	csvWriter.Flush()
}
