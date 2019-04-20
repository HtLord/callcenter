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

//Enum begin
type Employee struct {
	Id       uuid.UUID
	Priority Priority
}

const (
	Fresher Priority = 1
	TL      Priority = 2
	PM      Priority = 3
)

//Enum ending

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
			FRQ = append(FRQ, Employee{uuid.New(), Fresher})
		case TL:
			TLQ = append(TLQ, Employee{uuid.New(), TL})
		case PM:
			PMQ = append(PMQ, Employee{uuid.New(), PM})
		}
	}

	return nil
}

// Generate employees by config from call center with following const:
// MAX_FR, MAX_TL, MAX_PM. Then assign to slices(FRQ, TLQ, PMQ)
func GenerateEmployeesAutomatically() {
	for i := 0; i < MAX_FR; i++ {
		FRQ = append(FRQ, Employee{uuid.New(), Fresher})
	}

	for i := 0; i < MAX_TL; i++ {
		TLQ = append(TLQ, Employee{uuid.New(), TL})
	}

	for i := 0; i < MAX_PM; i++ {
		PMQ = append(PMQ, Employee{uuid.New(), PM})
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
	if fc+tlc+pmc > MAX_PM+MAX_TL+MAX_PM {
		return errors.New("Enter more than max(" + string(MAX_PM+MAX_TL+MAX_PM) + ") number of total call center cap")
	}

	return nil
}

// Assign slice of Employee into channel of Employee for single employee channel
// e.g. freasher, technical leader, product manager will assign into single channel
func LoadE2C(vs []Employee, c chan Employee) {
	for _, v := range vs {
		c <- v
	}
}

// Assign slice of Employee into channel of Employee for return separated priority level employee channel
// e.g. freasher, technical leader, product manager will assign into its own channel
func LoadEToChannel(es []Employee, buf int) chan Employee {
	c := make(chan Employee, buf)
	for _, e := range es {
		c <- e
	}
	return c
}

// Dump E info to console
func DumpAllEmployee() {
	DumpTitle("Fresher")
	for _, v := range FRQ {
		v.dumpEmployee()
	}
	DumpTitle("Technical Leader")
	for _, v := range TLQ {
		v.dumpEmployee()
	}
	DumpTitle("Product Manager")
	for _, v := range PMQ {
		v.dumpEmployee()
	}
}

// Dump E Id and Priority to console
func (e *Employee) dumpEmployee() {
	fmt.Printf("[Id: %s, Priority: %v]\n", e.Id, e.Priority)
}

// Take a PC from occ then the employee(who call TakePC function) will pull off pcc and start solve
// the PC by random secs(0-4) and rtd(roll the dice) to decide how long solve time is. If the emp-
// loyee can solve the problem then push PC to spc, else push PC to cpc. Finally, push the employee
// back to occ. !!!CAUTION PM SOLVE PC ANYWAY!!!
//
// Details about args:
// 1. occ: A channel of employee that can be a channel of Fresher, TL, PM , or other defined role.
//			Used when employee execute phone call and is free, then push back to channel(occ).
//			TL;DR: a queue hold specific employee is free.
// 2. pcc: A phone call channel. Pop first phone call to deal, if the employee who call Occupy
// 3. spc: A phone call channel for collect solved PCs
// 4. cpc: A phone call channel as next priority channel(escalate). if current employee can not solve
//			the PC then push it to cpc for escalate.
//
// Extra:
//  	Remove push pc to cpc after Print. Because it may let next thread executed and print before current print.
func (e *Employee) MultiChanTake(occ chan<- Employee, pcc chan PhoneCall, spc chan<- PhoneCall, cpc chan<- PhoneCall) {
	pc := <-pcc
	//factor := time.Duration(rand.Intn(5))
	factor := time.Duration(0)
	time.Sleep(factor * time.Second)
	if 1 == rand.Intn(2) {
		pc.HandleBy = e.Id
		fmt.Printf("P%d: pause %s s for %s solve %s\n", e.Priority, factor, e.Id, pc.Id)
		//fmt.Printf("P%d: %s solve %s\n", e.Priority, e.Id, pc.Id)
		spc <- pc
	} else {
		if e.Priority == PM {
			fmt.Printf("P%d: pause %s s for %s solve %s\n", e.Priority, factor, e.Id, pc.Id)
			//fmt.Printf("P%d: %s solve %s\n", e.Priority, e.Id, pc.Id)
		} else {
			fmt.Printf("P%d: pause %s s for %s escalate %s\n", e.Priority, factor, e.Id, pc.Id)
			//fmt.Printf("P%d: %s escalate %s\n", e.Priority, e.Id, pc.Id)
			cpc <- pc
		}
	}
	occ <- *e
}

// Take a PC from occ then the employee(who call TakePC function) will pull off pcc. If the PC's priority
// fit the employee's then start and solve it. The PC by random secs(0-4) and rtd(roll the dice) to
// decide how long solve time is. If the employee can solve the problem then push print to console, else
// push PC to cpc. Finally, push the employee back to occ. !!!CAUTION PM SOLVE PC ANYWAY!!!
//
func (e *Employee) SingleChanTake(occ chan<- Employee, pcc chan PhoneCall) {
	pc := <-pcc

	if pc.Priority != e.Priority {
		pcc <- pc
		occ <- *e
		return
	}

	factor := time.Duration(rand.Intn(5))
	//factor := time.Duration(0)
	time.Sleep(factor * time.Second)
	if 1 == rand.Intn(2) {
		pc.HandleBy = e.Id
		fmt.Printf("P%d: pause %s s for %s solve %s\n", e.Priority, factor, e.Id, pc.Id)
		//fmt.Printf("P%d: %s solve %s\n", e.Priority, e.Id, pc.Id)

	} else {
		if e.Priority == PM {
			fmt.Printf("P%d: pause %s s for %s solve %s\n", e.Priority, factor, e.Id, pc.Id)
			//fmt.Printf("P%d: %s solve %s\n", e.Priority, e.Id, pc.Id)
		} else {
			pc.escalate()
			fmt.Printf("P%d: pause %s s for %s escalate %s\n", e.Priority, factor, e.Id, pc.Id)
			//fmt.Printf("P%d: %s escalate %s\n", e.Priority, e.Id, pc.Id)
			pcc <- pc
		}
	}

	occ <- *e
}
