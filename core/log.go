package core

import "log"

//TODO: Implement proper logging, the below is just a quick solution for now

func LogError(err interface{}) {
	if err != nil {
		log.Printf("[ERROR]: %s\n", err)
	}
}
