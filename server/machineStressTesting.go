package server

import (
	"crypto/rand"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/mackerelio/go-osstat/memory"
	"github.com/patrickmn/go-cache"
)

// RunCPULoad run CPU load in specify cores count and percentage
func RunCPULoad(coresCount int, timeSeconds int, percentage int) {
	runtime.GOMAXPROCS(coresCount)

	// second     ,s  * 1
	// millisecond,ms * 1000
	// microsecond,μs * 1000 * 1000
	// nanosecond ,ns * 1000 * 1000 * 1000

	// every loop : run + sleep = 1 unit

	// 1 unit = 100 ms may be the best
	unitHundresOfMicrosecond := 1000
	runMicrosecond := unitHundresOfMicrosecond * percentage
	sleepMicrosecond := unitHundresOfMicrosecond*100 - runMicrosecond
	for i := 0; i < coresCount; i++ {
		go func() {
			runtime.LockOSThread()
			// endless loop
			for {
				begin := time.Now()
				for {
					// run 100%
					if time.Now().Sub(begin) > time.Duration(runMicrosecond)*time.Microsecond {
						break
					}
				}
				// sleep
				time.Sleep(time.Duration(sleepMicrosecond) * time.Microsecond)
			}
		}()
	}
	// how long
	time.Sleep(time.Duration(timeSeconds) * time.Second)
}

func RunMem(timeSeconds int, percentage int) {
	memory, _ := memory.Get()

	dividedNumber := int(100 / percentage)
	byteSize := int(memory.Total) / dividedNumber

	// random value in []byte
	token := make([]byte, byteSize)
	rand.Read(token)

	// cache value in memory
	c := cache.New(time.Duration(timeSeconds)*time.Second, time.Duration(timeSeconds)*time.Second)
	c.SetDefault("stress-memory", token)
}

type MachineStressTestingRequest struct {
	Memory int `json:"memory"`
	CPU    int `json:"cpu"`
}

func (s *echoServer) MachineStressTesting(c echo.Context) error {
	body := new(MachineStressTestingRequest)

	err := c.Bind(body)
	if err != nil {
		return err
	}

	if body.CPU > 0 {
		go RunCPULoad(runtime.NumCPU(), 20, body.CPU)
	}

	if body.Memory > 0 {
		go RunMem(20, body.Memory)
	}

	return nil
}
