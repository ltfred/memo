package types

type MemoStatus uint8

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
