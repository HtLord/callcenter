package callcenter

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

var FRQ []Employee
var TLQ []Employee
var PMQ []Employee

//Enum
type Employee struct {
	Id       uuid.UUID
	IsFree   bool
	Priority Priority
}

const (
	Fresher Priority = 1
	TL      Priority = 2
	PM      Priority = 3
)

// TL(Team Leader), PM(Product Manager)
type Priority int

// Generate and return employees as a array by formula.
// The formula will be looks like [ 1, 2, 3, 1, 1]
// and stand for 3 Freasher, 1 TL, and 1 PM in the call center
func GenerateEmployeesByFormula(formula []Priority) error {

	// Make sure the call center will get MAX_TL, MAX_PM, MAX_FR
	err := ValidateFormula(formula)
	if err != nil {
		return err
	}

	for _, f := range formula {
		switch f {
		case Fresher:
			FRQ = append(FRQ, Employee{uuid.New(), true, Fresher})
		case TL:
			TLQ = append(TLQ, Employee{uuid.New(), true, TL})
		case PM:
			PMQ = append(PMQ, Employee{uuid.New(), true, PM})
		}
	}

	return nil
}

func GenerateEmployeesAutomatically() {
	for i := 0; i < MAX_FR; i++ {
		FRQ = append(FRQ, Employee{uuid.New(), true, Fresher})
	}

	for i := 0; i < MAX_TL; i++ {
		TLQ = append(TLQ, Employee{uuid.New(), true, TL})
	}

	for i := 0; i < MAX_PM; i++ {
		PMQ = append(PMQ, Employee{uuid.New(), true, PM})
	}
	fmt.Println("Max number of freasher, team leader, and product manager is generated.")

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

func LoadEToChannel(es []Employee, buf int) chan Employee {
	c := make(chan Employee, buf)
	for _, e := range es {
		c <- e
	}
	return c
}

func DumpAllEmployee() {
	TitleDump("Fresher")
	for _, v := range FRQ {
		v.dumpEmployee()
	}
	TitleDump("Technical Leader")
	for _, v := range TLQ {
		v.dumpEmployee()
	}
	TitleDump("Product Manager")
	for _, v := range PMQ {
		v.dumpEmployee()
	}
}

func (e *Employee) dumpEmployee() {
	fmt.Printf("[Id: %s, Priority: %v]\n", e.Id, e.Priority)
}

func (e *Employee) Occupy(occ chan<- Employee, pcc chan PhoneCall, spc chan<- PhoneCall, cpc chan<- PhoneCall) {

	pc := <-pcc
	factor := time.Duration(rand.Intn(5))
	//factor := time.Duration(0)
	time.Sleep(factor * time.Second)
	if 1 == rand.Intn(2) {
		pc.HandleBy = e.Id
		spc <- pc
		fmt.Printf("P%d: pause %s s for %s solve %s\n", e.Priority, factor, e.Id, pc.Id)
		//fmt.Printf("P%d: %s solve %s\n", e.Priority, e.Id, pc.Id)
	} else {
		if e.Priority == PM {
			fmt.Printf("P%d: pause %s s for %s solve %s\n", e.Priority, factor, e.Id, pc.Id)
			//fmt.Printf("P%d: %s solve %s\n", e.Priority, e.Id, pc.Id)
		} else {
			cpc <- pc
			fmt.Printf("P%d: pause %s s for %s escalate %s\n", e.Priority, factor, e.Id, pc.Id)
			//fmt.Printf("P%d: %s escalate %s\n", e.Priority, e.Id, pc.Id)
		}
	}

	occ <- *e
}
