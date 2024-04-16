package employees_test

import (
	"testing"

	"github.com/n00best/challenge-2000/pkg/employees"
)

func TestIsActive(t *testing.T) {
	activeEmployee := employees.Employee{
		Status: employees.StatusActive,
	}

	inactiveEmployee := employees.Employee{
		Status: employees.StatusInactive,
	}

	if activeEmployee.IsActive() != true {
		t.Error("The active employee is not marked as one")
	}

	if inactiveEmployee.IsActive() == true {
		t.Error("The inactive employee is marked as active")
	}
}
