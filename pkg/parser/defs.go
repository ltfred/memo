package parser

import "github.com/ltfred/memo/types"

type Memo struct {
	Name     string             `json:"name"`
	Date     string             `json:"date"`
	Content  string             `json:"content"`
	Priority types.MemoPriority `json:"priority"`
}
