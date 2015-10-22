package stopwatch

import (
	"time"
	"fmt"
)

type task struct {
	name string
	start, end time.Time
}

func (t task) duration() time.Duration {
	return t.end.Sub(t.start)
}

type StopWatch struct  {
	initial time.Time
	tasks map[string]*task
}

func NewStopWatch() StopWatch {
	sw := StopWatch{time.Now(), make(map[string]*task)}
	return sw
}

func (sw StopWatch) Start(name string) {
	task := &task{name, time.Now(), time.Unix(0,0)}
	sw.tasks[name] = task
}

func (sw StopWatch) Stop(name string) error {
	task, ok := sw.tasks[name]
	if !ok {
		return fmt.Errorf("Task with name \"%v\" does not exist", name)
	}

	task.end = time.Now()
	return nil
}

func (sw StopWatch) Print() {
	for k, v := range sw.tasks {
		fmt.Printf("%v: %v\n", k, v.duration())
	}
}

