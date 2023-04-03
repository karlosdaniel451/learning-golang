package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	Name                 string
	MonthlySalaryInCents int
	DepartmentName       string
}

func main() {
	var employees = make([]Employee, 0, 100)

	for {
		newEmployee, err := readEmployeeFromStdin()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}

		if newEmployee == nil {
			break
		}

		employees = append(employees, *newEmployee)
	}

	fmt.Printf("\nemployees: \n%v\n", employees)

	highestPaidEmployeeByDepartment := getHighestPaidEmployeeByDepartmentName(employees)

	fmt.Printf("\nnumber of departments: %d\n", len(highestPaidEmployeeByDepartment))
	

	for departmentName, employee := range highestPaidEmployeeByDepartment {
		fmt.Printf("%s: %#v\n", departmentName, employee)
	}
}

func readEmployeeFromStdin() (*Employee, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(`Employee's name or "\q" to quit: `)
	name, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	name = strings.TrimSpace(name)

	if name == `\q` {
		return nil, nil
	}

	fmt.Printf(`Employee's salary IN CENTS or "\q" to quit: `)
	stringMonthlySalaryInCents, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	stringMonthlySalaryInCents = strings.TrimSpace(stringMonthlySalaryInCents)

	if stringMonthlySalaryInCents == `\q` {
		return nil, nil
	}

	intMonthlySalaryInCents, err := strconv.Atoi(stringMonthlySalaryInCents)

	if err != nil {
		return nil, err
	}

	fmt.Printf(`Employee's department name or "\q" to quit: `)
	departmentName, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	departmentName = strings.TrimSpace(departmentName)

	if departmentName == `\q` {
		return nil, nil
	}

	return &Employee{
		Name:                 name,
		MonthlySalaryInCents: intMonthlySalaryInCents,
		DepartmentName:       departmentName,
	}, nil
}


// Return a map[string]Employee where the keys are the department names of the passed
// employees and the value are the Employee with the highest salary in the corresponding
// department.
func getHighestPaidEmployeeByDepartmentName(employees []Employee) map[string]Employee {
	highestPaidEmployeeByDepartment := make(map[string]Employee)

	for _, employee := range employees {
		if employee.MonthlySalaryInCents > highestPaidEmployeeByDepartment[employee.DepartmentName].MonthlySalaryInCents {
			highestPaidEmployeeByDepartment[employee.DepartmentName] = employee
		}
	}

	return highestPaidEmployeeByDepartment
}
