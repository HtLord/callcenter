package callcenter

import (
	"fmt"
	"math/rand"
	"time"
)

// Config Call Center
// 1. the number of Employee(MAX_FR, MAX_TL, MAX_PM)
// 2. the number of IPC(MAX_PC, more details in ./PhoneCall)
const (
	MAX_FR int = 3
	MAX_TL int = 1
	MAX_PM int = 1
	MAX_PC int = 10
)

// Just define but not use yet
type CallCenter struct {
	Employees []Employee
	PC        PhoneCall
}

// A non-stop consumer. Generate pcc which single chan and employees which
// single chan. Then start consume PC from pcc and test taking phone call
// process in multi-thread.
func Receiver3Layer(pcc chan PhoneCall, ec ...chan Employee) {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(ec) != 3 {
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
			go fr.MultiChanTake(destructedEs[0], pcc, spc, tlpcc)
		case tl := <-destructedEs[1]:
			go tl.MultiChanTake(destructedEs[1], tlpcc, spc, pmpcc)
		case pm := <-destructedEs[2]:
			go pm.MultiChanTake(destructedEs[2], pmpcc, spc, cpc)
		default:
		}
	}
}

// A non-stop consumer. Generate pcc which single chan and employees which
// single chan. Then start consume PC from pcc and test taking phone call
// process in multi-thread.
// TL;DR: Consume PCs(mean while pcc) it will always consume PCs from pcc.
func NonStopReceiverSingleLayer(pcc chan PhoneCall, ec chan Employee) {
	rand.Seed(time.Now().UTC().UnixNano())

	DumpTitle("Result")
	for {
		select {
		case e := <-ec:
			go e.SingleChanTake(ec, pcc)
		default:

		}
	}
}

// A stoppable consumer. Generate pcc which single chan and employees which
// single chan. Then start consume PC from pcc and test taking phone call
// process in multi-thread.
// TL;DR: Consume PCs(mean while pcc) while all PCs are solve return.
func StoppableReceiverSingleLayer(pcc chan PhoneCall, ec chan Employee) {
	rand.Seed(time.Now().UTC().UnixNano())

	c := 0
	DumpTitle("Result")
	for {
		select {
		case e := <-ec:
			go func() {
				r := e.SingleChanTakeR(ec, pcc)
				if r {
					c++
				}
			}()
		default:

		}
		if c == MAX_PC {
			break
		}
	}
}
