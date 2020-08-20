package integer_to_roman

import (
	"math"
	"strings"
)

func intToRoman(num int) string {
	var i int
	var roman strings.Builder
	pow10 := 1
	a := make([]int, 0)
	m := map[int]rune{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}

	for num > 0 {
		switch num % 10 {
		case 1, 2, 3, 6, 7, 8:
			a = append(a, pow10)
			num--
			continue
		case 4:
			a = append(a, pow10*5, pow10)
			num -= 4
			continue
		case 9:
			a = append(a, pow10*10, pow10)
			num -= 9
			continue
		case 5:
			a = append(a, pow10*5)
			num -= 5
			continue
		default: // 0
			num /= 10
			i++
			pow10 = int(math.Pow10(i))
		}
	}

	for i = len(a) - 1; i >= 0; i-- {
		roman.WriteRune(m[a[i]])
	}
	return roman.String()
}
