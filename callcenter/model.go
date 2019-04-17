package callcenter

import (
	"errors"
)

type CallCenter struct {
	Employees []Employee
	PC        PhoneCall
}

type Employee struct {
	Priority Priority
}

// TL(Team Leader), PM(Product Manager)
type Priority int

const (
	Fresher Priority = 1
	TL      Priority = 2
	PM      Priority = 3
	MAX_FR  int      = 10
	MAX_TL  int      = 1
	MAX_PM  int      = 1
	MAX_TCC int      = 10
	MAX_PC  int      = 10
)

// Generate and return employees as a array by formula.
// The formula will be looks like [ 1, 2, 3, 1, 1]
// and stand for 3 Freasher, 1 TL, and 1 PM in the call center
func GenerateEmployeesByFormula(formula []Priority) ([]Employee, error) {
	var es []Employee

	// Make sure the call center will get max 1 TL, 1 PM, n Freasher
	err := ValidateFormula(formula)
	if err != nil {
		return nil, err
	}

	for _, f := range formula {
		switch f {
		case Fresher:
			es = append(es, Employee{Fresher})
		case TL:
			es = append(es, Employee{TL})
		case PM:
			es = append(es, Employee{PM})
		}
	}

	return es, nil
}

// Make sure formula will match the follow by rules:
// 1. The max number of Freasher
// 2. The max number of TL
// 3. The max number of PM
// 4. The max number of people of call center
func ValidateFormula(formula []Priority) error {
	fc := 0
	tlc := 0
	pmc := 0

	for _, f := range formula {
		switch f {
		case Fresher:
			fc++
		case TL:
			tlc++
		case PM:
			pmc++
		}
	}

	if tlc > MAX_TL {
		return errors.New("Enter more than max(" + string(MAX_TL) + ") TL")
	}
	if tlc > MAX_TL {
		return errors.New("Enter more than max(" + string(MAX_TL) + ") TL")
	}
	if pmc > MAX_PM {
		return errors.New("Enter more than max(" + string(MAX_PM) + ") PM")
	}
	if fc+tlc+pmc > MAX_TCC {
		return errors.New("Enter more than max(" + string(MAX_TCC) + ") number of total call center cap")
	}

	return nil
}
