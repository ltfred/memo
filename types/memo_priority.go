package types

type MemoPriority uint8

var MemoPriorityValues []string

func init() {
	MemoPriorityValues = []string{"Important", "Generally"}
}

const (
	MemoPriorityUnknown MemoPriority = iota
	MemoPriorityImportant
	MemoPriorityGenerally
)

func (priority MemoPriority) String() string {
	switch priority {
	case MemoPriorityImportant:
		return "Important"
	case MemoPriorityGenerally:
		return "Generally"
	default:
		return "Unknown"
	}
}

func ParseMemoPriorityFromString(s string) MemoPriority {
	switch s {
	case "Important":
		return MemoPriorityImportant
	case "Generally":
		return MemoPriorityGenerally
	default:
		return MemoPriorityUnknown
	}
}
