package compiler60_shared

// some constants which are needed by executor and checker
const (
	ExecutorPort = "6062"

	HttpContentType = "Content-Type"
	HttpJsonType    = "application/json;charset=UTF-8"

	MaxBinarySize = 64 * 1024
	MaxStdoutSize = 128 * 1024
)
