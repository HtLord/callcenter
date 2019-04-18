package main

import (
	"callcenter/callcenter"
	"math/rand"
	"time"
)

//TODO: Rewrite LoadEToChannel/LoadPToChannel with Generic

type Test struct {
	Done bool
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	f := []callcenter.Priority{
		callcenter.Priority(1),
		callcenter.Priority(1),
		callcenter.Priority(1),
		callcenter.Priority(2),
		callcenter.Priority(3),
	}
	err := callcenter.GenerateEmployeesByFormula(f)
	if err != nil {
		return
	}
	err = callcenter.GenerateMaxPhoneCallOnce()
	if err != nil {
		return
	}

	callcenter.DumpAllEmployee()
	callcenter.DumpAllPhoneCall()
	frc := callcenter.LoadEToChannel(callcenter.FRQ, 5)
	//tlc := callcenter.LoadToChannel(callcenter.TLQ, 1)
	//pmc := callcenter.LoadToChannel(callcenter.PMQ, 1)
	pcc := callcenter.LoadPToChannel(callcenter.IPC, 10)

	spc := make(chan callcenter.PhoneCall, 30)
	cpc := make(chan callcenter.PhoneCall, 30)

	func() {
		for {
			v := <-frc
			go v.Occupy(frc, pcc, spc, cpc)
		}
	}()
}
