package utils

import (
	"testing"
)

// TestGetResponse
// test GetResponse with url
func TestGetResponse(t *testing.T) {
	// GIVEN
	_, err := GetResponse("https://hacker-news.firebaseio.com/v0/item/8863.json")

	// THEN
	if err != nil {
		t.Error("Wrong ", err)
	}
}
