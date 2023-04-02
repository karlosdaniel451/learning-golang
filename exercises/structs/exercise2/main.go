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

	fmt.Printf("employees: %v\n", employees)

	fmt.Printf("employe with the highest salary: %v\n", GetEmployeesWithHighestSalary(employees))
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

func GetEmployeesWithHighestSalary(employees []Employee) []Employee {
	// When there is no employee at all.
	if len(employees) == 0 {
		return []Employee{}
	}

	// When there is only one employee.
	if len(employees) == 1 {
		return employees[0:]
	}

	employeesWithHighestSalary := []Employee{}

	highestSalary := GetHighestSalaryInCentsOfEmployees(employees)

	for _, employee := range employees {
		if employee.MonthlySalaryInCents == highestSalary {
			employeesWithHighestSalary = append(employeesWithHighestSalary, employee)
		}
	}

	return employeesWithHighestSalary
}

func GetHighestSalaryInCentsOfEmployees(employees []Employee) int {
	highestSalary := employees[0].MonthlySalaryInCents

	for _, employee := range employees[1:] {
		if employee.MonthlySalaryInCents > highestSalary {
			highestSalary = employee.MonthlySalaryInCents
		}
	}

	return highestSalary
}
