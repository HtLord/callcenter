package main

import (
	"callcenter/callcenter"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//TODO: Rewrite LoadEToChannel/LoadPToChannel with Generic.
//TODO: Make TakePC's the highest rank can be config.
//TODO: Make TakePC's pause sec can be config.
//TODO: Build a visible html page and able to refill pcc.
//TODO: Build a pure console dump like.
//TODO: TC5 is not work properly. Details check TC5.Comment.

func main() {
	//read args
	al := len(os.Args)
	if al != 5 && al != 1 {
		fmt.Printf("must be run [max_rf, max_tl, max_pm, max_pc]\n")
		return
	}

	if al == 5 {
		var err error
		callcenter.MAX_FR, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("check input must be int\n")
			return
		}
		callcenter.MAX_TL, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("check input must be int\n")
			return
		}
		callcenter.MAX_PM, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Printf("check input must be int\n")
			return
		}
		callcenter.MAX_PC, err = strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Printf("check input must be int\n")
			return
		}
	}

	//Dump conf
	callcenter.DumpConfig()

	//Random seed
	rand.Seed(time.Now().UTC().UnixNano())

	//Generate PC
	callcenter.GeneratePhoneCallAutomatically()
	pcc := callcenter.LoadPCToChannel(callcenter.PCQ, callcenter.MAX_PC)

	//Generate Single E channel
	callcenter.GenerateEmployeesAutomatically()
	ec := make(chan callcenter.Employee, callcenter.MAX_FR+callcenter.MAX_TL+callcenter.MAX_PM)
	callcenter.LoadE2C(callcenter.FRQ, ec)
	callcenter.LoadE2C(callcenter.TLQ, ec)
	callcenter.LoadE2C(callcenter.PMQ, ec)

	//Dump E, PC
	callcenter.DumpAllEmployee()
	callcenter.DumpAllPhoneCall()

	//Run call center version of non-stop Single E channel
	callcenter.StoppableReceiverSingleLayer(pcc, ec)

}
