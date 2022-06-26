package compiler60_shared

type JsonRunnerResult struct {
	Result     string `json:"result"`
	Stdout     []byte `json:"stdout"`
	UsedTimeMs int64  `json:"used_time_ms"`
	Wstatus    uint32 `json:"wstatus"`
}
