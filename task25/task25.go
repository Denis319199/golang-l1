package task25

import (
	"fmt"
	"syscall"
	"time"
)

func Sleep(duration time.Duration) {
	deadline := time.Now().Add(duration)

	// Active, this thread takes CPU time
	for deadline.After(time.Now()) {
	}
}

// SleepSys - implementation for the Windows OS only
func SleepSys(milli uint32) {
	handle, err := syscall.GetCurrentProcess()
	if err != nil {
		return
	}

	// Passive, this thread doesn't take CPU time
	_, err = syscall.WaitForSingleObject(handle, milli)
	if err != nil {
		return
	}
}

func Task25() {
	p1 := time.Now()
	Sleep(100 * time.Millisecond)
	p2 := time.Now()
	fmt.Println(p2.Sub(p1))

	p1 = time.Now()
	SleepSys(100)
	p2 = time.Now()
	fmt.Println(p2.Sub(p1))
}
