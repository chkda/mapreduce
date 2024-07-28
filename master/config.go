package master

type Config struct {
	GRPCPort             string `json:"grpc_port"`
	Filename             string `json:"filename"`
	NumReduce            int    `json:"num_reduce"`
	DeadWorkersThreshold int    `json:"dead_workers_threshold"`
}
