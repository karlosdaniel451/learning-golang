package cpuscheduling

type TaskManager interface {
	GetRunnableTasks() []*Task
	GetNextTaskID() int
	CreateTask(name string, priority int, cpuBurst int) *Task
	Schedule(cpu *CPU)
	pickNextTask(cpu *CPU) *Task
}
