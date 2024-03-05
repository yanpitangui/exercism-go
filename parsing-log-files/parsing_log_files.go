package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	if re.MatchString(text) {
		return true
	}

	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<([~*=-])*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)

	count := 0
	for _, val := range lines {
		if re.MatchString(val) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line(\d)+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	li := []string{}
	re := regexp.MustCompile(`User\s+([A-Za-z0-9]+)`)

	for _, val := range lines {

		sl := re.FindStringSubmatch(val)
		if len(sl) <= 1 {
			li = append(li, val)
		} else {
			li = append(li, fmt.Sprintf("[USR] %s %s", sl[1], val))
		}
	}
	return li

}
