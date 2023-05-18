package parser

type Memo struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	Content  string `json:"content"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
	CreateAt int64  `json:"createAt"`
}
