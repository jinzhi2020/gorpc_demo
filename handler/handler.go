package handler

import (
	"encoding/json"
	"fmt"
	"time"
)

const TimeServiceName = "handler/TimeService"

type TimeService struct{}

func (s *TimeService) Now(request string, reply *string) error {
	t := map[string]string{
		"Now": time.Now().Format("2006-01-02 15:04:05"),
	}
	jsonStr, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Json encode failed!")
	}
	*reply = string(jsonStr)
	return nil
}
