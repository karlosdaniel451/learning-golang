package cpuscheduling

import (
	"log"
	"sync"
	"time"
)

type CPU struct {
	currentlyExecutingTask *Task
	// quantum int
	mu sync.Mutex
}

func newCPU() *CPU {
	return &CPU{currentlyExecutingTask: nil}
}

func (cpu *CPU) isIdle() bool {
	cpu.mu.Lock()
	defer cpu.mu.Unlock()
	return cpu.currentlyExecutingTask == nil ||
		cpu.currentlyExecutingTask.RemainingCpuBurst == 0
}

// Run a task for the specified time slice.
func (cpu *CPU) runTask(task *Task, slice int) {
	cpu.mu.Lock()
	defer cpu.mu.Unlock()

	if cpu.currentlyExecutingTask != nil {
		cpu.currentlyExecutingTask.State = Runnable
	}
	cpu.currentlyExecutingTask = task
	log.Printf(
		"started running task named %s with id %d for %d time units.\n",
		task.Name, task.Tid, slice,
	)

	go cpu.execute()
}

func (cpu *CPU) execute() {
	for {
		// time.Sleep(time.Duration(cpu.currentlyExecutingTask.remainingCpuBurst) * time.Second)
		time.Sleep(time.Second * 1)
		if cpu.isIdle() {
			return
		}

		cpu.mu.Lock()
		log.Printf(
			"running one CPU cycle for task with id %d - remaining CPU burst: %d\n",
			cpu.currentlyExecutingTask.Tid, cpu.currentlyExecutingTask.RemainingCpuBurst,
		)
		cpu.currentlyExecutingTask.RemainingCpuBurst -= 1
		if cpu.currentlyExecutingTask.RemainingCpuBurst == 0 {
			log.Printf("finished task with id %d\n", cpu.currentlyExecutingTask.Tid)
		}
		cpu.mu.Unlock()
	}
}

// func (cpu *CPU) preempt() (preemptedTask *Task) {
// 	if cpu.currentlyExecutingTask != nil {
// 		return nil
// 	}

// 	preemptedTask = cpu.currentlyExecutingTask
// 	cpu.currentlyExecutingTask = nil

// 	return preemptedTask
// }
