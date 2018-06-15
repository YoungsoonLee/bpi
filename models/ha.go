package models

// Manage models.
// if you want to extend, you just add some models.

// Result ...
// save result for stdout or csv
type Result struct {
	By    string `json:"by"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Story ...
// save story infomation from HA
type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}
