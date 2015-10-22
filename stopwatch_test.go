package stopwatch

import (
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	sw := NewStopWatch()
	sw.Start("test1")

	time.Sleep(100 * time.Millisecond)

	err := sw.Stop("test1")
	if err != nil {
		println(err)
	}
	sw.Print()
}