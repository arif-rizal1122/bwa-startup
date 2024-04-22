package helper

import "log"



func IfError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}