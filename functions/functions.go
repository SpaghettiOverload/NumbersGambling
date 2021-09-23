package functions

import (
	"fmt"
	"math/rand"
	"time"
)

func CleanLine() {
	fmt.Print("\r")
}

func NewRandomNum(r int) int {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(r)
	fmt.Print("Generating new number")
	Loading()
	CleanLine()
	return randomNum
}

func SpareRandomNum(r int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(r)
}

func Wait() {
	time.Sleep(500 * time.Millisecond)
}

func WaitLonger() {
	time.Sleep(1500 * time.Millisecond)
}

func Loading(){
	for i := 0; i < 2; i++ {
		fmt.Print(".")
		Wait()
	}
	fmt.Print("DONE!")
	Wait()
	CleanLine()
}
