package utilstest

import (
	"log"
)


func UtilLog(a string) (b string){
	log.Printf("THIS IS A LOG")
	string2 := a + "From Utils Package"
	return string2
	
}
