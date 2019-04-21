package test

import (
	"callcenter/callcenter"
	"math/rand"
	"time"
)

var pcc chan callcenter.PhoneCall
var frc chan callcenter.Employee
var tlc chan callcenter.Employee
var pmc chan callcenter.Employee
var ec chan callcenter.Employee

// Test case 1: Generate employees(Es) automatically and dump it
func TC1() {
	callcenter.GenerateEmployeesAutomatically()
	callcenter.DumpAllEmployee()
}

// Test case 2: Generate phone calls(PCs) automatically and dump it
func TC2() {
	callcenter.GeneratePhoneCallAutomatically()
	callcenter.DumpAllPhoneCall()
}

// Test case 3: Load Es and return result as a buffered channel
func TC3() {
	frc = callcenter.LoadEToChannel(callcenter.FRQ, callcenter.MAX_FR)
	tlc = callcenter.LoadEToChannel(callcenter.TLQ, callcenter.MAX_TL)
	pmc = callcenter.LoadEToChannel(callcenter.PMQ, callcenter.MAX_PM)
}

// Test case 4: Load PCs and return result as a buffered channel
func TC4() {
	pcc = callcenter.LoadPCToChannel(callcenter.PCQ, callcenter.MAX_PC)
}

// Test case 5: A non-stop consumer. Generate pcc which single chan and
// employees which separated chan. Then start consume PC from
// pcc and test taking phone call process in multi-thread.
//
// Comment: It will run and then holt. But not all pcc are solved.
func TC5() {
	rand.Seed(time.Now().UTC().UnixNano())

	callcenter.GeneratePhoneCallAutomatically()
	pcc = callcenter.LoadPCToChannel(callcenter.PCQ, callcenter.MAX_PC)

	callcenter.GenerateEmployeesAutomatically()
	ec := make(chan callcenter.Employee, callcenter.MAX_FR+callcenter.MAX_TL+callcenter.MAX_PM)
	callcenter.LoadE2C(callcenter.FRQ, ec)
	callcenter.LoadE2C(callcenter.TLQ, ec)
	callcenter.LoadE2C(callcenter.PMQ, ec)

	spc := make(chan callcenter.PhoneCall, callcenter.MAX_PC)
	tlpcc := make(chan callcenter.PhoneCall, callcenter.MAX_PC)
	pmpcc := make(chan callcenter.PhoneCall, callcenter.MAX_PC)
	cpc := make(chan callcenter.PhoneCall, callcenter.MAX_PC)

	callcenter.DumpTitle("Result")

	//var wg sync.WaitGroup
	//wg.Add(callcenter.MAX_PC)

	for {
		select {
		case fr := <-frc:
			go fr.MultiChanTake(frc, pcc, spc, tlpcc)
			//go fr.MultiChanTakeAndSync(frc, pcc, spc, tlpcc, &wg)
		case tl := <-tlc:
			go tl.MultiChanTake(frc, tlpcc, spc, pmpcc)
			//go tl.MultiChanTakeAndSync(tlc, tlpcc, spc, pmpcc, &wg)
		case pm := <-pmc:
			go pm.MultiChanTake(frc, pmpcc, spc, cpc)
		//go pm.MultiChanTakeAndSync(pmc, pmpcc, spc, cpc, &wg)
		case <-cpc:
		default:
		}
	}

}

// Test case 6: A prototype version of non-stop consumer. Generate pcc which
// single chan and employees which single chan. Then start consume PC from
// pcc and test taking phone call process in multi-thread.
func TC6() {
	//Random seed
	rand.Seed(time.Now().UTC().UnixNano())

	//Generate PC
	callcenter.GeneratePhoneCallAutomatically()
	pcc = callcenter.LoadPCToChannel(callcenter.PCQ, callcenter.MAX_PC)

	//Generate Single E channel
	callcenter.GenerateEmployeesAutomatically()
	ec = make(chan callcenter.Employee, callcenter.MAX_FR+callcenter.MAX_TL+callcenter.MAX_PM)
	callcenter.LoadE2C(callcenter.FRQ, ec)
	callcenter.LoadE2C(callcenter.TLQ, ec)
	callcenter.LoadE2C(callcenter.PMQ, ec)

	//Run prototype version of non-stop Single E channel
	for {
		select {
		case e := <-ec:
			go e.SingleChanTake(ec, pcc)
		default:

		}
	}
}

// Test case 7: A non-stop consumer. But using the method from call center
func TC7() {
	//Random seed
	rand.Seed(time.Now().UTC().UnixNano())

	//Generate PC
	callcenter.GeneratePhoneCallAutomatically()
	pcc = callcenter.LoadPCToChannel(callcenter.PCQ, callcenter.MAX_PC)

	//Generate Single E channel
	callcenter.GenerateEmployeesAutomatically()
	ec = make(chan callcenter.Employee, callcenter.MAX_FR+callcenter.MAX_TL+callcenter.MAX_PM)
	callcenter.LoadE2C(callcenter.FRQ, ec)
	callcenter.LoadE2C(callcenter.TLQ, ec)
	callcenter.LoadE2C(callcenter.PMQ, ec)

	//Run call center version of non-stop Single E channel
	callcenter.NonStopReceiverSingleLayer(pcc, ec)
}

// Test case 8: A stoppable consumer. using the method from call center
func TC8() {
	//Random seed
	rand.Seed(time.Now().UTC().UnixNano())

	//Generate PC
	callcenter.GeneratePhoneCallAutomatically()
	pcc = callcenter.LoadPCToChannel(callcenter.PCQ, callcenter.MAX_PC)

	//Generate Single E channel
	callcenter.GenerateEmployeesAutomatically()
	ec = make(chan callcenter.Employee, callcenter.MAX_FR+callcenter.MAX_TL+callcenter.MAX_PM)
	callcenter.LoadE2C(callcenter.FRQ, ec)
	callcenter.LoadE2C(callcenter.TLQ, ec)
	callcenter.LoadE2C(callcenter.PMQ, ec)

	//Dump E, PC
	callcenter.DumpAllEmployee()
	callcenter.DumpAllPhoneCall()

	//Run call center version of non-stop Single E channel
	callcenter.StoppableReceiverSingleLayer(pcc, ec)
}
