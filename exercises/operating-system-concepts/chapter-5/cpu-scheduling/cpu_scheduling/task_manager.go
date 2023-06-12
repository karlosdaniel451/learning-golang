package cpuscheduling

type TaskManager interface {
	GetRunnableTasks() []*Task
	GetNextTaskID() int
	CreateTask(name string, priority int, cpuBurst int) *Task
	GetCpu() *CPU
	Schedule()
	pickNextTask() (task *Task, ok bool)

}
