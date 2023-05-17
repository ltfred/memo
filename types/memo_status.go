package types

type MemoStatus uint8

var MemoStatusValues []string

func init() {
	MemoStatusValues = []string{"Undo", "Doing"}
}

const (
	MemoStatusUnknown MemoStatus = iota
	MemoStatusUndo
	MemoStatusDoing
)

func (status MemoStatus) String() string {
	switch status {
	case MemoStatusUndo:
		return "Undo"
	case MemoStatusDoing:
		return "Doing"
	default:
		return "Unknown"
	}
}

func ParseMemoStatusFromString(s string) MemoStatus {
	switch s {
	case "Undo":
		return MemoStatusUndo
	case "Doing":
		return MemoStatusDoing
	default:
		return MemoStatusUnknown
	}
}
