package main

import (
	"fmt"
	"time"
)

func main() {

	cashier := &cashier{}

	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)

	a := time.Now()
	time.Sleep(time.Duration(1) * time.Millisecond)
	fmt.Println(time.Since(a))

	aa := time.Now()
	time.Sleep(1 * time.Millisecond)
	fmt.Println(time.Since(aa))
}
