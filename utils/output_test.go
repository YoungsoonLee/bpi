package utils

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"

	"github.com/YoungsoonLee/bpi/models"
)

// StdCsv
// test StdCsv
func TestStdCsv(t *testing.T) {
	// GIVEN
	testFileName := "bpi_csv_test.csv"
	// make a file
	file, _ := os.OpenFile(testFileName, os.O_CREATE|os.O_WRONLY, 0777)
	r := models.Result{Title: "test title", By: "test by", Time: 11111111, URL: "test url"}

	StdCsv(file, 1, r)

	// THEN
	csvFile, _ := os.Open(testFileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, _ := reader.Read()

	got := line[1] //title
	want := "test title"
	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Wrong, got '%s' want '%s'", got, want)
	}
}
