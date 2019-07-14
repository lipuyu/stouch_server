package utils

import (
	"github.com/google/uuid"
	"strings"
)

func GetUUID() string {
	uuid1, _ := uuid.NewUUID()
	salt := uuid1.String()
	return strings.Replace(salt, "-", "", -1)
}