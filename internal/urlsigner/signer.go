package urlsigner

import (
	"fmt"
	"strings"
	"time"

	goalone "github.com/bwmarrin/go-alone"
)

type Signer struct {
	Secret []byte
}

func (s *Signer) GenerateTokenFromString(data string) string {
	var urlString string

	crypt := goalone.New(s.Secret, goalone.Timestamp)
	if strings.Contains(data, "?") {
		urlString = fmt.Sprintf("%s&hash=", data)
	} else {
		urlString = fmt.Sprintf("%s?hash=", data)
	}

	tokenBytes := crypt.Sign([]byte(urlString))
	token := string(tokenBytes)
	return token
}

func (s *Signer) VerifyToken(token string) bool {
	crypt := goalone.New(s.Secret, goalone.Timestamp)
	_, err := crypt.Unsign([]byte(token))

	if err != nil {
		fmt.Sprintln(err)
		return false
	}

	return true
}

func (s *Signer) Expired(token string, minutesUnitExpire int) bool {
	crypt := goalone.New(s.Secret, goalone.Timestamp)
	ts := crypt.Parse([]byte(token))

	return time.Since(ts.Timestamp) > time.Duration(minutesUnitExpire)*time.Minute
}
