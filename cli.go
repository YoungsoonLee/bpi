package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/YoungsoonLee/bpi/getnews"
)

// CLI responsible for processing command line arguments.
type CLI struct{}

// this is usage print function
// it is print except input ./bpi or ./bpi -csv
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  -csv Create excel csv ")
}

// Run parses command line arguments and processes commands.
// if there is no any argument, just run the program for stdout.
// if argument is "-csv", run program for CSV file create.
// In the main process function, made by simple pipeline.
// So, first make a CSV file, and then pass the file to main process function for pipeline.
// for extend, need to category for call main process.
// for example, if yuo want to extend r/golang, call main process with category like getnews.GetTopStory("RA", nil).
func (cli *CLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println(">> start get hackernews to stdout ...")
		getnews.GetTopStory("HA", nil)
		fmt.Println(">> end get hackernews to stdout ...")
		os.Exit(1)
	} else {
		switch os.Args[1] {
		case "-csv":
			fmt.Println(">> start get hackernews to bpi_csv_file.csv ...")

			// make a file
			file, err := os.OpenFile("bpi_csv_file.csv", os.O_CREATE|os.O_WRONLY, 0777)
			// close file
			defer file.Close()

			if err != nil {
				log.Println(err)
				os.Exit(1)
			}

			//make first header
			headers := []string{"No", "Title", "By", "Date", "Url"}
			csvWriter := csv.NewWriter(file)
			csvWriter.Write(headers)
			csvWriter.Flush()

			// call main process function with file
			getnews.GetTopStory("HA", file)

			fmt.Println(">> end get hackernews to bpi_csv_file.csv ...")
			os.Exit(1)
		default:
			cli.printUsage()
			os.Exit(1)
		}
	}
}
