package utils

import (
	"github.com/google/uuid"
	"strconv"
)

func GetUUIDString() string {
	u := uuid.New()
	return u.String()
}

func GetUUIDNumber() string {
	u := uuid.New()
	return strconv.Itoa(int(u.ID()))
}
