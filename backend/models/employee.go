package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type Employee struct {
	ID                string         `json:"id_employee" db:"id_employee"`
	LastName          string         `json:"empl_surname" db:"empl_surname"`
	FirstName         string         `json:"empl_name" db:"empl_name"`
	MiddleName        sql.NullString `json:"empl_patronymic,omitempty" db:"empl_patronymic"`
	Role              Role           `json:"empl_role" db:"empl_role"`
	Salary            float64        `json:"salary" db:"salary"`
	BirthDate         time.Time      `json:"date_of_birth" db:"date_of_birth"`
	StartDate         time.Time      `json:"date_of_start" db:"date_of_start"`
	PhoneNumber       string         `json:"phone_number" db:"phone_number"`
	City              string         `json:"city" db:"city"`
	Street            string         `json:"street" db:"street"`
	ZipCode           string         `json:"zip_code" db:"zip_code"`
	Username          string         `json:"username" db:"username"`
	Password          string         `json:"password" db:"password"`
	IsPasswordDefault bool           `json:"is_password_default" db:"is_password_default"`
}

func (employee *Employee) HashAndSavePassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	employee.Password = string(bytes)

	return nil
}

func (employee *Employee) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (employee *Employee) VerifyCorrectness() error {
	stringParameters := []string{employee.FirstName, employee.LastName, employee.PhoneNumber, employee.City, employee.Street, employee.ZipCode, employee.Username, employee.Password}
	for _, stringParameter := range stringParameters {
		if stringParameter == "" {
			return errors.New("cannot have empty strings")
		}
	}

	if employee.Role != Manager && employee.Role != Cashier {
		return errors.New("incorrect employee role")
	}

	if !employee.BirthDate.Before(time.Now().AddDate(-18, 0, 0)) {
		return errors.New("employee too young")
	}

	if !employee.StartDate.Before(time.Now()) {
		return errors.New("cannot add an employee who hasn't yet started working")
	}

	if employee.StartDate.IsZero() {
		employee.StartDate = time.Now()
	}

	if !strings.HasPrefix(employee.PhoneNumber, "+") {
		return errors.New("phone number must start with a '+'")
	}

	if employee.Salary < 0 {
		return errors.New("salary cannot be less than 0")
	}

	return nil
}
