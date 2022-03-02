package constantarg

import (
	"log"
)

func constantArgumentAnalyzerData() {

	const renownedConstant = "ConstantField"
	justAnotherVariable := "i am a variable!"

	log := log2{}
	log.WithField("Field1", "Value1")            //want "log-withField"
	log.WithField(justAnotherVariable, "Value1") //want "log-withField"
	log.WithField(renownedConstant, "Value1")

}

type log2 struct {
	field string
}

func (l log2) WithField(field string, value interface{}) {
	log.Println("->Field:", field, " ->Value", value)
}
