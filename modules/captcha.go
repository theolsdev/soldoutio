package modules

import (
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type CaptchaSession struct {
	UUID   string       `json:"uuid"`
	Client *http.Client `json:"client"`
	Key    string       `json:"key"`
}

func (s *CaptchaSession) InitSession(key string) {
	u, _ := uuid.NewV4()
	s.UUID = u.String()
	s.Key = key
	s.Client = &http.Client{
		Transport: &http.Transport{},
		Timeout:   5 * time.Second,
	}
}
