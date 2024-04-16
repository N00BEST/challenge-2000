package risk

import "github.com/n00best/challenge-2000/pkg/employees"

func CalculateEmployeeRiskValue(employee employees.Employee) int {
	if !employee.IsActive() {
		return 0
	}

	return 100
}
