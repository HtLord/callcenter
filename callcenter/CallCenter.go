package callcenter

// Config Call Center
// Controll
// 1. the number of Employee(MAX_FR, MAX_TL, MAX_PM, MAX_TCC)
// 2. the number of IPC(MAX_IPC, more details in ./PhoneCall)

const (
	MAX_FR  int = 10
	MAX_TL  int = 1
	MAX_PM  int = 1
	MAX_TCC int = 5
	MAX_IPC int = 10
)

type CallCenter struct {
	Employees []Employee
	PC        PhoneCall
}
