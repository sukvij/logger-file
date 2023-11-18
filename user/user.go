package user

import (
	"errors"
	"vijju/log"
)

func AddUser(log *log.AgreeGateLoager) error {
	log.Error("problem")
	return errors.New("user does not exist")
}
