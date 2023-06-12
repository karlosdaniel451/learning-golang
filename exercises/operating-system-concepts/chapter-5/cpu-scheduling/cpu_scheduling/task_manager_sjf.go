package cpuscheduling

import (
	"github.com/ugurcsen/gods-generic/queues/priorityqueue"
	"github.com/ugurcsen/gods-generic/utils"
)

type TaskManagerShortestJobFirst struct {
	cpu           *CPU
	runnableTasks priorityqueue.Queue[*Task]
	nextTaskID    int
}

// Comparator in order to sort values of `Task` in descendent order according
// to the value of the field`Priority`.
func taskPriorityComparator(task1, task2 *Task) int {
	return utils.NumberComparator(task1.CpuBurst, task2.CpuBurst)
}

func NewTaskManagerShortestJobFirst() *TaskManagerShortestJobFirst {
	return &TaskManagerShortestJobFirst{
		nextTaskID: 1,
		// runnableTasks: *linkedlistqueue.New[*Task](),
		runnableTasks: *priorityqueue.NewWith(taskPriorityComparator),
		cpu:           newCPU(),
		// scheduler:     RoundRobinScheduler{},
	}
}

func (taskManager *TaskManagerShortestJobFirst) GetRunnableTasks() []*Task {
	return taskManager.runnableTasks.Values()
}

func (taskManager *TaskManagerShortestJobFirst) GetNextTaskID() int {
	defer func() {
		taskManager.nextTaskID++
	}()

	return taskManager.nextTaskID
}

func (taskManager *TaskManagerShortestJobFirst) CreateTask(
	name string,
	priority int,
	cpuBurst int,
) *Task {
	task := &Task{
		Name:              name,
		Tid:               taskManager.GetNextTaskID(),
		Priority:          priority,
		CpuBurst:          cpuBurst,
		RemainingCpuBurst: cpuBurst,
		State:             Runnable,
	}

	taskManager.runnableTasks.Enqueue(task)

	return task
}

func (taskManager *TaskManagerShortestJobFirst) GetCpu() *CPU {
	return taskManager.cpu
}

func (taskManager *TaskManagerShortestJobFirst) Schedule() {
	schedule(taskManager)
}

func (taskManager *TaskManagerShortestJobFirst) pickNextTask() (task *Task, ok bool) {
	// log.Printf("number of runnable tasks: %d", taskManager.runnableTasks.Size())
	return taskManager.runnableTasks.Dequeue()
}
