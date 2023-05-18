package constant

var (
	MemoPriorityValues    = []string{MemoPriorityLow, MemoPriorityHigh}
	MemoPriorityValuesMap = map[string]int{MemoPriorityLow: 1, MemoPriorityHigh: 2}
)

const (
	MemoPriorityLow  string = "low"
	MemoPriorityHigh string = "high"
)
