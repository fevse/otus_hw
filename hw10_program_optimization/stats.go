package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	gojson "github.com/goccy/go-json"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain), nil
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	br := bufio.NewReader(r)
	user := &User{}
	var i int
	var line []byte
	for {
		line, _, err = br.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return
		}
		if err = gojson.Unmarshal(line, &user); err != nil {
			return
		}
		result[i] = *user
		i++
	}
	return result, nil
}

func countDomains(u users, domain string) DomainStat {
	result := make(DomainStat)
	domain = "." + domain
	for _, user := range u {
		if strings.HasSuffix(user.Email, domain) {
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]++
		}
	}
	return result
}
