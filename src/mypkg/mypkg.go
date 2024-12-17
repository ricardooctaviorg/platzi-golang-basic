package mypkh

import "fmt"

type CarPublic struct {
	Name  string
	Brand string
	Year  int
	color string
	Motor string
}

type carPrivate struct {
	name string
	year int
}

func PrintMessage() {
	fmt.Println("Mensage X ")
}

func printMessage(text string) {
	fmt.Println(text)
}

type Pc struct {
	Ram  int
	Disk int
	Cpu  string
}

func (myPc Pc) Ping() {
	fmt.Println(myPc.Ram, "PONG")
}

func (myPc *Pc) DuplicateRam() {
	myPc.Ram = myPc.Ram * 2
}

func (myPc Pc) String() string {
	return fmt.Sprintf("RAM: %d DISK: %d CPU:%s", myPc.Ram, myPc.Disk, myPc.Cpu)
}
