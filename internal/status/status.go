package status

type TaskStatus int

const (
	IDLE TaskStatus = iota
	INPROGRESS
	COMPLETE
	FAILED
)
