package getnews

import (
	"fmt"
	"testing"
)

// TestProcess
// test main pipeline proces.
func TestProcess(t *testing.T) {
	// GIVEN
	input := make(chan string)
	defer close(input)

	done := make(chan bool)
	defer close(done)

	go func() {
		input <- "https://hacker-news.firebaseio.com/v0/item/8863.json"
		done <- true
	}()

	// WHEN
	output := process("HA", input)
	<-done // blocks until the input write routine is finished

	// THEN
	want := "dhouston" //by
	got := <-output    // blocks until the output has contents

	assertEqual(t, got.By, want)
}

// TestGetItemDetail
// test a top story detail information with itemid from HA
func TestGetItemDetail(t *testing.T) {
	// GIVEN
	r := getItemDetail("HA", "https://hacker-news.firebaseio.com/v0/item/8863.json")

	// TEHN
	want := "dhouston" // by
	assertEqual(t, r.By, want)

}

// TestGetTopStoriesItemID
// test 20 top stories from HA
func TestGetTopStoriesItemID(t *testing.T) {
	// GIVEN
	i, _ := getTopStoriesItemID("HA")

	// THEN
	want := "20" // count
	got := fmt.Sprintf("%d", len(i))
	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Wrong, got '%s' want '%s'", got, want)
	}
}
