package decorator

import "fmt"

type UserReport struct{}

func NewUserReport() *UserReport {
	return &UserReport{}
}

func (u *UserReport) Generate() string {
	return fmt.Sprintf(
		"User Report\n"+
			"---------------------------------\n"+
			"UserID: 1 | Name: Ali | Role: premium | Joined: 2025-01-11\n"+
			"UserID: 2 | Name: Dana | Role: basic   | Joined: 2025-02-18\n"+
			"UserID: 3 | Name: Nursultan | Role: premium | Joined: 2025-03-21\n",
	)
}