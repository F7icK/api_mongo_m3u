package workers

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

type worker struct {
	*gocron.Scheduler
	stoppedWorker chan bool
	stopped       bool
}

type Worker interface {
	AddWorkEveryMinutes(interval uint64, jobFunc interface{}, args ...interface{}) error
	StartWorker() error
	ClearWorker()
	StopWorker() error
}

func New() Worker {
	return &worker{
		Scheduler: gocron.NewScheduler(),
		stopped:   true,
	}
}

func (w *worker) AddWorkEveryMinutes(interval uint64, jobFunc interface{}, args ...interface{}) error {
	return w.Every(interval).Minutes().Do(jobFunc, args...)
}

func (w *worker) StartWorker() error {
	if !w.stopped {
		return fmt.Errorf("workers already started")
	}

	w.stoppedWorker = w.Start()
	w.stopped = false
	return nil
}

func (w *worker) ClearWorker() {
	w.Clear()
}

func (w *worker) StopWorker() error {
	if w.stopped {
		return fmt.Errorf("workers already stopped")
	}

	w.stoppedWorker <- true
	close(w.stoppedWorker)
	w.stopped = true

	return nil
}
