package types

type MemoPriority uint8

var MemoPriorityValues []string

func init() {
	MemoPriorityValues = []string{"High", "Low"}
}

const (
	MemoPriorityUnknown MemoPriority = iota
	MemoPriorityHigh
	MemoPriorityLow
)

func (priority MemoPriority) String() string {
	switch priority {
	case MemoPriorityHigh:
		return "High"
	case MemoPriorityLow:
		return "Low"
	default:
		return "Unknown"
	}
}

func ParseMemoPriorityFromString(s string) MemoPriority {
	switch s {
	case "High":
		return MemoPriorityHigh
	case "Low":
		return MemoPriorityLow
	default:
		return MemoPriorityUnknown
	}
}
