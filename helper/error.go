package helper

import "log"

func ErrorIfNotNil(err error){
	if err != nil {
		log.Printf(err.Error())
	}
}
