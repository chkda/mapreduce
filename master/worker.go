package master

import (
	"sync"

	pbw "github.com/chkda/mapreduce/rpc/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type WorkerStatus int

const (
	WORKER_ALIVE WorkerStatus = iota
	WORKER_DEAD
)

type Worker struct {
	Uuid   string
	mu     sync.RWMutex
	conn   *grpc.ClientConn
	IP     string
	State  WorkerStatus
	Client pbw.WorkerClient
}

func NewWorker(uuid string, ip string) (*Worker, error) {
	worker := &Worker{
		Uuid:  uuid,
		IP:    ip,
		State: WORKER_ALIVE,
	}
	conn, err := grpc.Dial(worker.IP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pbw.NewWorkerClient(conn)
	worker.Client = client
	return worker, nil
}

func (w *Worker) GetUuid() string {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.Uuid
}

func (w *Worker) GetIP() string {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.IP
}

func (w *Worker) GetState() WorkerStatus {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.State
}

func (w *Worker) SetState(state WorkerStatus) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.State = state
}

func (w *Worker) GetClient() pbw.WorkerClient {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.Client
}

func (w *Worker) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.conn != nil {
		return w.conn.Close()
	}
	return nil
}

func (w *Worker) UpdateClient() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.conn != nil {
		w.conn.Close()
	}

	conn, err := grpc.Dial(w.IP, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	w.Client = pbw.NewWorkerClient(conn)
	w.conn = conn
	return nil
}
