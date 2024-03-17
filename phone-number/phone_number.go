package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	builder := strings.Builder{}
	re := regexp.MustCompile(`^[+]?1?\s*[(]?[2-9]{1}[0-9]{2}[)]?(\s+|.|-)?[2-9]{1}[0-9]{2}(\s+|.|-)?[0-9]{4}\s*$`)
	if !re.MatchString(phoneNumber) {
		return "", errors.New("invalid number")
	}
	r := []rune(phoneNumber)
	for l := 0; l < len(r); l++ {
		if unicode.IsDigit(r[l]) {
			builder.WriteRune(r[l])
		}
	}

	n := builder.String()
	if len(n) < 10 || len(n) > 11 {
		return "", errors.New("invalid number")
	}
	if len(n) == 11 && n[:1] == "2" {
		return "", errors.New("invalid when 11 digits does not start with a 1")
	}
	if n[:1] == "1" {
		return n[1:], nil
	}

	return n, nil
}

func AreaCode(phoneNumber string) (string, error) {
	n, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return n[:3], nil
}

func Format(phoneNumber string) (string, error) {
	n, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", n[:3], n[3:6], n[6:]), nil
}
