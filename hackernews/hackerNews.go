package hackerNews

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YoungsoonLee/bpi/utils"
)

func FetchHackerNews() {
	res, err := http.Get(utils.TopHackerNewsURL) // URL에서 데이터를 가져옴
	if err != nil {
		log.Println(err)
	}

	fmt.Println(res)
}
