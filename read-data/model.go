package read_data

type Record struct {
	Date      string  `json:"date"`
	Timestamp int64   `json:"timestamp"`
	Money     float64 `json:"money"`
	Category  string  `json:"category"`
	Type      string  `json:"type"`
	Desc      string  `json:"desc"`
}
