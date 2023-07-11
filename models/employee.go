package models


type Employee struct {
	Id          uint64  `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
	FirstName   string  `gorm:"column:first_name;" json:"first_name"`
	LastName    string  `gorm:"column:last_name;" json:"last_name"`
	DateOfBirth string  `gorm:"column:date_of_birth;type:date;" json:"date_of_birth"`
	Email       string  `gorm:"column:email;" json:"email"`
	PhoneNumber string  `gorm:"column:phone_number;" json:"phone_number"`
	Address     string  `gorm:"column:address;" json:"address"`
	City        string  `gorm:"column:city;" json:"city"`
	State       string  `gorm:"column:state;" json:"state"`
	Country     string  `gorm:"column:country;" json:"country"`
	PostalCode  string  `gorm:"column:postal_code;" json:"postal_code"`
	Position    string  `gorm:"column:position;" json:"position"`
	Department  string  `gorm:"column:department;" json:"department"`
	HireDate    *string `gorm:"column:hire_date;" json:"hire_date,omitempty"`
	Salary      float32 `gorm:"column:salary;" json:"salary"`
}

type EmployeeOutput struct {
	Completed bool
}
