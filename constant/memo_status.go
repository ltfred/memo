package constant

var (
	MemoStatusValues    = []string{MemoStatusUndo, MemoStatusDoing}
	MemoStatusValuesMap = map[string]int{MemoStatusUndo: 1, MemoStatusDoing: 2}
)

const (
	MemoStatusUndo  string = "undo"
	MemoStatusDoing string = "doing"
)
