package cpuscheduling

import "time"

func schedule(taskManager TaskManager) {
	for {
		nextTaskToBeRun, ok := taskManager.pickNextTask()
		if !ok {
			break
		}
		// Only select another process/thread/task to run when there is no one
		// currently running in the CPU
		for {
			isCpuIdle := taskManager.GetCpu().isIdle()
			if isCpuIdle {
				break
			}
			time.Sleep(time.Second * 1)
		}
		previouslyExecutingTask := taskManager.GetCpu().currentlyExecutingTask
		if previouslyExecutingTask != nil {
			previouslyExecutingTask.State = Terminated
		}
		nextTaskToBeRun.State = Running
		taskManager.GetCpu().runTask(nextTaskToBeRun, nextTaskToBeRun.CpuBurst)
	}
}
