package types

type MemoPriority uint8

var MemoPriorityValues []string

func init() {
	MemoPriorityValues = []string{"Important", "Generally"}
}

const (
	MEMO_PRIORITY__UNKNOWN MemoPriority = iota
	MEMO_PRIORITY__IMPORTANT
	MEMO_PRIORITY__GENERALLY
)

func (priority MemoPriority) String() string {
	switch priority {
	case MEMO_PRIORITY__IMPORTANT:
		return "Important"
	case MEMO_PRIORITY__GENERALLY:
		return "Generally"
	default:
		return "Unknown"
	}
}

func ParseMemoPriorityFromString(s string) MemoPriority {
	switch s {
	case "Important":
		return MEMO_PRIORITY__IMPORTANT
	case "Generally":
		return MEMO_PRIORITY__GENERALLY
	default:
		return MEMO_PRIORITY__UNKNOWN
	}
}
