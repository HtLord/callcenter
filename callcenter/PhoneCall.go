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

var IPC []PhoneCall
var CPC []PhoneCall
var SPC []PhoneCall

type PhoneCall struct {
	Id       uuid.UUID
	IsIdle   bool
	HandleBy uuid.UUID
}

// Generate number of PC into IPC
func GeneratePhoneCall(num int) error {
	if len(IPC) > MAX_IPC || (len(IPC)+num) > MAX_IPC {
		return errors.New("There has no enough lines for asking number of phone call")
	}

	for i := 0; i < num; i++ {
		pcid := uuid.New()
		IPC = append(IPC, PhoneCall{pcid, true, uuid.Nil})
	}
	fmt.Printf("IPC(%d) is generated.\n", num)
	return nil
}

// Generate MAX IPC it will be using when
// 1. Testing and fulfill the IPC
// 2. All IPC is consumed and want to be refill the IPC
func GenerateMaxPhoneCallOnce() error {
	if len(IPC) > MAX_IPC {
		return errors.New("There has no enough lines for asking number of phone call")
	}

	for i := 0; i < MAX_IPC; i++ {
		pcid := uuid.New()
		IPC = append(IPC, PhoneCall{pcid, true, uuid.Nil})
	}
	fmt.Println("Max number of IPC is generated.")
	return nil
}

func LoadPToChannel(es []PhoneCall, buf int) chan PhoneCall {
	c := make(chan PhoneCall, buf)
	for _, e := range es {
		c <- e
	}
	return c
}

func DumpAllPhoneCall() {
	fmt.Printf("IPC\n")
	for i, v := range IPC {
		fmt.Printf("[%d, %v]\n", i, v)
	}
	fmt.Printf("CPC[%s]\n", len(CPC))
	fmt.Printf("SPC[%s]\n", len(SPC))
}

func Solved(pc PhoneCall) {
}

// Search next avaliable rank to handle PC
func Escalate(pc PhoneCall) {

}

// Remove IPC to CPC
func Cancel() {

}