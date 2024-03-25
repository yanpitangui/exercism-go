package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	re := regexp.MustCompile(`^What is -?\d+( (multiplied by|divided by|plus|minus) -?\d+)*\?$`)
	if !re.MatchString(question) {
		return 0, false
	}

	segments := strings.Split(question[:len(question)-1], " ")
	n, _ := strconv.Atoi(segments[2])
	for i := 3; i < len(segments); {
		switch segments[i] {
		case "minus":
			minus, _ := strconv.Atoi(segments[i+1])
			n = n - minus
			i += 2
			break

		case "plus":
			plus, _ := strconv.Atoi(segments[i+1])
			n = n + plus
			i += 2
			break

		case "multiplied":
			multiplied, _ := strconv.Atoi(segments[i+2])
			n = n * multiplied
			i += 3
			break
		case "divided":
			divided, _ := strconv.Atoi(segments[i+2])
			n = n / divided
			i += 3
			break
		}
	}
	return n, true

}
