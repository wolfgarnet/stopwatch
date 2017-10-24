package util

import (
	"fmt"
	"time"
)

type lap struct {
	name     string
	duration time.Duration
}

type StopWatch struct {
	started, stopped time.Time
	lastLap          time.Time
	laps             []lap
}

func NewStopWatch() *StopWatch {
	return &StopWatch{}
}

func (sw *StopWatch) Start() *StopWatch {
	sw.started = time.Now()
	sw.lastLap = sw.started
	return sw
}

func (sw *StopWatch) NewLap(name string) {
	now := time.Now()
	duration := now.Sub(sw.lastLap)
	sw.lastLap = now
	sw.laps = append(sw.laps, lap{name, duration})
}

func (sw *StopWatch) Stop() *StopWatch {
	sw.stopped = time.Now()
	return sw
}

func (sw *StopWatch) Duration() time.Duration {
	return sw.stopped.Sub(sw.started)
}

func (sw *StopWatch) Laps() {
	d := sw.Duration()
	fmt.Printf("Laps: Total=%v\n", d)
	for i, lap := range sw.laps {
		fmt.Printf(" [%v] %-15v: %-12v (%v%%)\n", i, lap.name, lap.duration, float32(int((float32(lap.duration.Nanoseconds())/float32(d.Nanoseconds()))*10000))/100)
	}
}
