package cpuscheduling

import (
	"log"
	"time"
)

type CPU struct {
	currentlyExecutingTask *Task
	// quantum int
}

func newCPU() *CPU {
	return &CPU{currentlyExecutingTask: nil}
}

func (cpu *CPU) isIdle() bool {
	return cpu.currentlyExecutingTask == nil ||
		cpu.currentlyExecutingTask.RemainingCpuBurst == 0
}

// Run a task for the specified time slice.
func (cpu *CPU) runTask(task *Task, slice int) {
	if cpu.currentlyExecutingTask != nil {
		cpu.currentlyExecutingTask.State = Runnable
	}
	cpu.currentlyExecutingTask = task
	log.Printf(
		"started running task named %s with id %d for %d time units.\n",
		task.Name, task.Tid, slice,
	)

	cpu.execute()
}

func (cpu *CPU) execute() {
	for {
		time.Sleep(time.Second * 1)
		if cpu.isIdle() {
			return
		}

		log.Printf(
			"running one CPU cycle for task with id %d - remaining CPU burst: %d\n",
			cpu.currentlyExecutingTask.Tid, cpu.currentlyExecutingTask.RemainingCpuBurst,
		)
		cpu.currentlyExecutingTask.RemainingCpuBurst -= 1
		if cpu.currentlyExecutingTask.RemainingCpuBurst == 0 {
			log.Printf("finished task with id %d\n", cpu.currentlyExecutingTask.Tid)
		}
	}
}
