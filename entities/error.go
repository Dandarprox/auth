package entities

import (
	"fmt"
)

type ErrorEntity struct {
	Message   string `json:"message"`
	ErrorCode int16  `json:"errorCode"`
	ip        string
	_         any // Force attribute names to be used
}

func (e *ErrorEntity) LogError() {
	fmt.Printf("[IP] %s", e.ip)
}

func NewErrorEntiy(message, ip string, errorCode int16) *ErrorEntity {
	return &ErrorEntity{
		Message:   message,
		ip:        ip,
		ErrorCode: errorCode,
	}
}
