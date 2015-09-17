package sleep

import "time"
import "testing"

func TestingSleep(t1 *testing.T) {
	start := time.Now()
	sleeping(25)
	end := time.Now()

	if end.Sub(start) < 25 {
		t1.Error("Sleep exited before 10s")
	}
}
