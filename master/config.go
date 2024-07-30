package master

type Config struct {
	GRPCPort             string
	Filename             string
	NumReduce            int
	DeadWorkersThreshold int
}
