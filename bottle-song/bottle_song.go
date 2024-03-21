package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	translation := map[int]string{0: "no", 1: "one",
		2: "two", 3: "three", 4: "four",
		5: "five", 6: "six", 7: "seven",
		8: "eight", 9: "nine", 10: "ten"}
	l := []string{}

	for i := startBottles; i > startBottles-takeDown; i-- {
		for range []int{1, 2} {
			if i != 1 {
				l = append(l, fmt.Sprintf("%s green bottles hanging on the wall,", strings.Title(translation[i])))
			} else {
				l = append(l, fmt.Sprintf("%s green bottle hanging on the wall,", strings.Title(translation[i])))

			}
		}
		l = append(l, fmt.Sprintf("And if one green bottle should accidentally fall,"))

		if i-1 != 1 {
			l = append(l, fmt.Sprintf("There'll be %s green bottles hanging on the wall.", translation[i-1]))
		} else {
			l = append(l, fmt.Sprintf("There'll be %s green bottle hanging on the wall.", translation[i-1]))
		}

		if i-1 > startBottles-takeDown {
			l = append(l, "")
		}

	}
	return l
}
