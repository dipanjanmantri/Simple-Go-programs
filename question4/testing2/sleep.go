package sleep

import "time"

func sleeping(sec int) {
	select {
		case <- time.After(time.Duration(sec) * time.Second):
			return
	}
}
