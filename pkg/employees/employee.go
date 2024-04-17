package employees

import "github.com/n00best/challenge-2000/pkg/databases"

type Employee struct {
	ID     uint   `json:"id"`
	Status Status `json:"status"`

	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`

	DivisionCode   uint   `json:"div_code"`
	Division       string `json:"division"`
	DepartmentCode uint   `json:"department_code"`
	Department     string `json:"department"`

	CountryID uint   `json:"country_id"`
	Country   string `json:"country"`

	DateIn  string `json:"date_in"`
	DateOut string `json:"date_out"`

	accesses []databases.DatabaseAccess
}

func (e Employee) IsActive() bool {
	return e.Status == StatusActive
}

func (e Employee) CountSensitiveAccesses() int {
	return databases.CountSensitiveDatabases(e.accesses)
}

func (e *Employee) SetAccesses(accesses []databases.DatabaseAccess) {
	e.accesses = accesses
}

func (e Employee) GetAccesses() []databases.DatabaseAccess {
	return e.accesses
}
