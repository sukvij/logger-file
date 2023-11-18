package main

import (
	"vijju/log"
	"vijju/user"
)

func main() {

	log := log.NewAgreeGateLogger()
	err := user.AddUser(log)
	if err != nil {
		log.Error(err)
	}
}
