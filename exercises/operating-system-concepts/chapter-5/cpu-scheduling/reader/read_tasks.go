package reader

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/karlosdaniel451/cpu-scheduling/cpu_scheduling"
)

func ReadTasksFromCSV(filepath string) ([]cpuscheduling.Task, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err // When an error when opening the file ocurrs
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err // When an error when reading the CSV file ocurrs
	}

	tasks, err := parseTasksFromString(data)
	if err != nil {
		return nil, err // When a parsing error occurs
	}

	return tasks, nil
}

func parseTasksFromString(tasksToBeParsed [][]string) ([]cpuscheduling.Task, error) {
	tasks := []cpuscheduling.Task{}

	for i, line := range tasksToBeParsed[1:] {
		log.Printf("parsing line %d\n", i)
		var task cpuscheduling.Task
		for j, field := range line {
			field := strings.TrimSpace(field) // Remove leading and trailing white spaces
			switch j {
			case 0:
				task.Name = field
			case 1:
				priority, err := strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				task.Priority = priority
			case 2:
				cpuBurst, err := strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				task.CpuBurst = cpuBurst
			}
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
