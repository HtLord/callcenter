package main

import (
	"callcenter/callcenter"
	"math/rand"
	"time"
)

var pcc chan callcenter.PhoneCall
var frc chan callcenter.Employee
var tlc chan callcenter.Employee
var pmc chan callcenter.Employee

func main() {
	tc1_GEAD()
	tc2_GPCAD()
	tc3_LETC()
	tc4_LPCTC()
	tc5_MT()
}

// Test case 1: Generate employees(Es) automatically and dump it
func tc1_GEAD() {
	callcenter.GenerateEmployeesAutomatically()
	callcenter.DumpAllEmployee()
}

// Test case 2: Generate phone calls(PCs) automatically and dump it
func tc2_GPCAD() {
	callcenter.GeneratePhoneCallAutomatically()
	callcenter.DumpAllPhoneCall()
}

// Test case 3: Load Es and return result as a buffered channel
func tc3_LETC() {
	frc = callcenter.LoadEToChannel(callcenter.FRQ, callcenter.MAX_FR)
	tlc = callcenter.LoadEToChannel(callcenter.TLQ, callcenter.MAX_TL)
	pmc = callcenter.LoadEToChannel(callcenter.PMQ, callcenter.MAX_PM)
}

// Test case 4: Load PCs and return result as a buffered channel
func tc4_LPCTC() {
	pcc = callcenter.LoadPCToChannel(callcenter.PCQ, callcenter.MAX_PC)
}

func tc5_MT() {
	rand.Seed(time.Now().UTC().UnixNano())

	spc := make(chan callcenter.PhoneCall, callcenter.MAX_PC)
	tlpcc := make(chan callcenter.PhoneCall, callcenter.MAX_PC)
	pmpcc := make(chan callcenter.PhoneCall, callcenter.MAX_PC)
	cpc := make(chan callcenter.PhoneCall, callcenter.MAX_PC)

	callcenter.DumpTitle("Result")

	for {
		select {
		case fr := <-frc:
			go fr.TakePC(frc, pcc, spc, tlpcc)
		case tl := <-tlc:
			go tl.TakePC(tlc, tlpcc, spc, pmpcc)
		case pm := <-pmc:
			go pm.TakePC(pmc, pmpcc, spc, cpc)
		default:
		}
	}
}
