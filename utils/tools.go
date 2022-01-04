package utils

import (
	"github.com/gofrs/uuid"
	"strings"
)

func GenerateUUid() string {
	u, _ := uuid.NewV4()
	u1 := uuid.Must(u, nil)
	result := u1.String()
	result = strings.ReplaceAll(result, "-", "")
	return result
}
