package cpuscheduling

type TaskState int

type Task struct {
	Name              string
	Tid               int
	Priority          int
	CpuBurst          int
	RemainingCpuBurst int
	State             TaskState
}

const (
	Runnable TaskState = iota
	Running
	Terminated
)

// func removeIndex(tasks []Task, index int) []Task {
// 	ret := []Task{}
// 	ret = append(ret, tasks[:index]...)
// 	return append(ret, tasks[index+1:]...)
// }
