package testdata

import (
	"log"
	"time"
)

func doesNotSleep() {
}

func sleeps() {
	time.Sleep(1 * time.Nanosecond) //want "no-sleep-lint"

	const renownedConstant = "ConstantField"
	justAnotherVariable := "i am a variable!"

	log := log2{}
	log.WithField("Field1", "Value1")            //want "no-log-withField-lint"
	log.WithField(justAnotherVariable, "Value1") //want "no-log-withField-lint"
	log.WithField(renownedConstant, "Value1")
}

type log2 struct {
	field string
}

func (l log2) WithField(field string, value interface{}) {
	log.Println("->Field:", field, " ->Value", value)
}
