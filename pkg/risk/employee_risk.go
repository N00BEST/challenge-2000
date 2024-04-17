package risk

import (
	"github.com/n00best/challenge-2000/pkg/databases"
	"github.com/n00best/challenge-2000/pkg/employees"
)

func CalculateEmployeeRiskValue(employee employees.Employee) int {
	if !employee.IsActive() {
		return 0
	}

	total := 5

	total += employee.CountSensitiveAccesses() * 50 / databases.GetTotalSensitiveDatabases()

	return total
}
