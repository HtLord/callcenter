package callcenter

import (
	"fmt"
	"math/rand"
	"time"
)

// Config Call Center
// Controll
// 1. the number of Employee(MAX_FR, MAX_TL, MAX_PM, MAX_TCC)
// 2. the number of IPC(MAX_IPC, more details in ./PhoneCall)

const (
	MAX_FR int = 10
	MAX_TL int = 1
	MAX_PM int = 1
	MAX_PC int = 15
)

type CallCenter struct {
	Employees []Employee
	PC        PhoneCall
}

func Sender() {

}

func Receiver3Layer(pcc chan PhoneCall, es ...chan Employee) {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(es) != 3 {
		fmt.Println("Receiver must have 3 layer mean while 3 different channels.")
		return
	}

	var destructedEs []chan Employee
	for _, e := range destructedEs {
		destructedEs = append(destructedEs, e)
	}

	spc := make(chan PhoneCall, MAX_PC)
	tlpcc := make(chan PhoneCall, MAX_PC)
	pmpcc := make(chan PhoneCall, MAX_PC)
	cpc := make(chan PhoneCall, MAX_PC)

	DumpTitle("Result")
	for {
		select {
		case fr := <-destructedEs[0]:
			go fr.TakePC(destructedEs[0], pcc, spc, tlpcc)
		case tl := <-destructedEs[1]:
			go tl.TakePC(destructedEs[1], tlpcc, spc, pmpcc)
		case pm := <-destructedEs[2]:
			go pm.TakePC(destructedEs[2], pmpcc, spc, cpc)
		default:
		}
	}
}
