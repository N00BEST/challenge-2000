package employees

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
}

func (e Employee) IsActive() bool {
	return e.Status == StatusActive
}
