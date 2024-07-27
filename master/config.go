package master

type Config struct {
	GRPCPort             string `json:"grpc_port"`
	Filename             string `json:"filename"`
	NumReduce            string `json:"num_reduce"`
	DeadWorkersThreshold int    `json:"dead_workers_threshold"`
}
