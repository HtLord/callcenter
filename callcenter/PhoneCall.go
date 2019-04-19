package callcenter

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

// For convenient, we will call PhoneCall as PC.
// There will be IPC, CPC, SPC to handle PCs and will explain it
// follow by.
// IPC(Incoming phone call) Rules:
// 		1.Any phone call will be append to IPC first and wait to be solve.
// CPC(Canceled phone call) Rules:
//		1. If the phone call is coming and there is no free man append to CPC
//		2. IPC is full then it will be append to CPC ( In current CC struct it will never happened )
//		3. No one can solve the problem then remove from IPC or
// SPC(Solved phone call) Rules:
//		1. If some solved the PC which will be remove from IPC
//		2. Then append to SPC.

var PCQ []PhoneCall
var CPCQ []PhoneCall
var SPCQ []PhoneCall

type PhoneCall struct {
	Id       uuid.UUID
	Priority Priority
	HandleBy uuid.UUID
}

// Generate number of PC into IPC
func GeneratePhoneCall(num int) error {
	if len(PCQ) > MAX_PC || (len(PCQ)+num) > MAX_PC {
		return errors.New("There has no enough lines for asking number of phone call")
	}

	for i := 0; i < num; i++ {
		pcid := uuid.New()
		PCQ = append(PCQ, PhoneCall{pcid, Priority(1), uuid.Nil})
	}
	fmt.Printf("IPC(%d) is generated.\n", num)
	return nil
}

// Generate MAX IPC it will be using when
// 1. Testing and fulfill the IPC
// 2. All IPC is consumed and want to be refill the IPC
func GeneratePhoneCallAutomatically() {
	for i := 0; i < MAX_PC; i++ {
		pcid := uuid.New()
		PCQ = append(PCQ, PhoneCall{pcid, Priority(1), uuid.Nil})
	}
	fmt.Println("Max number of IPC is generated.")
}

func LoadPCToChannel(es []PhoneCall, buf int) chan PhoneCall {
	c := make(chan PhoneCall, buf)
	for _, e := range es {
		c <- e
	}
	return c
}

func DumpAllPhoneCall() {
	DumpTitle("IPC")
	for i, v := range PCQ {
		fmt.Printf("[%d, %v]\n", i+1, v)
	}
}

func DumpCanNotSolved(c PhoneCall) {
	fmt.Printf("Even PM can not solve PC[%s]\n", c)
}
