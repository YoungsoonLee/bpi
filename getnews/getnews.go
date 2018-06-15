package getnews

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"

	"github.com/YoungsoonLee/bpi/models"
	"github.com/YoungsoonLee/bpi/utils"
)

// GetTopStory ...
// Function that has main logic with goroutine and channels like pipeline.
// category input param is that call URL.
//    for example, the category is "HA", call HA, the category is "R", call R. of course, first need to add function for "R".
// file input param is that the file to write to CSV.
// For pipeline, communicate with urls channel and done channel.
// urls channal is the top story url with itemid and done channel is for stop goroutine.
func GetTopStory(category string, file io.Writer) {
	// Get top stoires item ids
	list, err := getTopStoriesItemID(category)
	if err != nil {
		log.Panic(err)
	}

	// Channel for reguest get top story detail
	urls := make(chan string)
	defer close(urls)

	// Channel for sending stop to goroutine
	done := make(chan struct{})

	output := process(category, urls)

	// send url to channal to start
	go func() {
		// Request jobs with just having same top stories counts 20
		for i := 0; i < len(list); i++ {
			// make top story url with itemid
			istr := strconv.FormatInt(list[i], 10) + ".json"
			// send data to url channel
			urls <- utils.ITEMHAURL + istr
		}
		// close channel and all worker gorutine stop
		close(done)
	}()

	count := 0
	// receive data from channel
	// write to stdout or csv like pipeline
	for r := range output {

		if file != nil {
			// csv
			utils.StdCsv(file, count+1, r)
		} else {
			// stdout
			utils.StdOut(count+1, r)
			fmt.Println()
		}
		// check counts 20, and then stop pipeline
		count++
		if count == utils.TOPCOUNTS {
			break
		}
	}
}

// Process ...
// simple pipeline function.
// get the url channel and then send worker function for getting informaition.
// get the return from workwe, return result channal for print or create.
func process(category string, urls <-chan string) <-chan models.Result {
	var wg sync.WaitGroup
	// just use same top stories count for setting go routine
	wg.Add(utils.TOPCOUNTS)

	output := make(chan models.Result)

	// make a go routine
	// just make same top stories counts 20
	for i := 1; i < utils.TOPCOUNTS; i++ {
		go func() {
			for url := range urls {
				output <- woker(category, url)
			}
			wg.Done()
		}()
	}

	go func() {
		// Waiting until done goroutine
		wg.Wait()
		// Close result channel if wait done
		close(output)
	}()

	return output
}

// woker ...
// operation to get top story with itemID
func woker(category string, url string) models.Result {
	r := getItemDetail(category, url)
	return r
}

// getTopStories ...
// Get top 500 stories ItemID from HA.
// If you want to extend, add the function in switch with category.
func getTopStoriesItemID(category string) ([]int64, error) {
	ret := make([]int64, 0, 20)

	// If need to extend, use category
	switch category {
	case "HA":
		// get the response body []byte from HA
		responseData, err := utils.GetResponse(utils.TOPHAURL)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// unmarshall from response body
		var decoded []interface{}
		if err := json.Unmarshal(responseData, &decoded); err != nil {
			log.Println(err)
			return nil, err
		}

		// get 20 stories
		decoded = decoded[:utils.TOPCOUNTS] // get 20 stories
		for _, d := range decoded {
			ret = append(ret, int64(d.(float64)))
		}
	}
	return ret, nil
}

// getItemDetail
// get top story from HA with itemID
// If you want to extend, add the function in switch with category.
func getItemDetail(category string, url string) models.Result {
	switch category {
	case "HA":
		// get top story with itemid from HA
		responseData, err := utils.GetResponse(url)
		if err != nil {
			log.Println(err)
		}
		// marshalling
		var st models.Story
		if err := json.Unmarshal(responseData, &st); err != nil {
			log.Println(err)
		}

		// send data to result channel
		return models.Result{
			By:    st.By,
			Time:  st.Time,
			Title: st.Title,
			URL:   st.URL,
		}
	}

	return models.Result{}
}
