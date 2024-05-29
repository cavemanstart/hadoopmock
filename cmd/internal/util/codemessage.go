package util

type CodeMessage struct {
	Code int
	Msg  string
}

var (
	Success = CodeMessage{200, "success"}
	MockErr = CodeMessage{501, "mock error"}
	DbErr   = CodeMessage{502, "mongo error"}
)
