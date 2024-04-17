package risk_test

import (
	"testing"

	"github.com/n00best/challenge-2000/pkg/databases"
	"github.com/n00best/challenge-2000/pkg/employees"
	"github.com/n00best/challenge-2000/pkg/risk"
)

type riskTestCase struct {
	name              string
	employee          employees.Employee
	expectedRiskValue int
}

func TestCalculateEmployeeRiskValue(t *testing.T) {
	sensitiveAccessEmployee := employees.Employee{
		Status: employees.StatusActive,
	}

	sensitiveAccessEmployee.SetAccesses([]databases.DatabaseAccess{
		{
			IsSensitiveInfo: 1,
		},
	})

	testCases := []riskTestCase{
		{
			name: "inactive employee",
			employee: employees.Employee{
				Status: employees.StatusInactive,
			},
			expectedRiskValue: 0,
		},
		{
			name: "active employee",
			employee: employees.Employee{
				Status: employees.StatusActive,
			},
			expectedRiskValue: 5,
		},
		{
			name:              "sensitive info database access",
			employee:          sensitiveAccessEmployee,
			expectedRiskValue: 55,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(subT *testing.T) {
			expectation := test
			riskValue := risk.CalculateEmployeeRiskValue(expectation.employee)

			if riskValue != expectation.expectedRiskValue {
				subT.Errorf("Test %s failed: expected %d, got %d", expectation.name, expectation.expectedRiskValue, riskValue)
			}
		})
	}
}
