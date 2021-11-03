package errors

type ErrorChannel struct {
	Type   string
	Errors error
}

const (
	FS_ERROR       int = 1001
	REDIS_ERROR    int = 1002
	POPULATE_ERROR int = 1003
	VALIDATE_ERROR int = 1004
)
