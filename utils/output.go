package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"

	"github.com/YoungsoonLee/bpi/models"
)

// StdOut ...
// print top 20 stories to stdout.
// print No, Title, By, date, url
func StdOut(i int, r models.Result) {
	fmt.Printf("No: %d, Title: %s, By: %s, Date: %s, Url: %s", i, r.Title, r.By, time.Unix(r.Time, 0), r.URL)
}

// StdCsv ...
// print top 20 stories to csv file.
// print No, Title, By, date, url
func StdCsv(file io.Writer, i int, r models.Result) {
	values := []string{fmt.Sprintf("%d", i), r.Title, r.By, fmt.Sprintf("%s", time.Unix(r.Time, 0)), r.URL}
	csvWriter := csv.NewWriter(file)
	csvWriter.Write(values)
	csvWriter.Flush()
}
