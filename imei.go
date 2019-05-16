package util

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// generate IMEI
func Generate() string {
	var pos, sum, final_digit, t, len_offset, leng int
	str := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	leng = 15
	imei := ""

	rbi := []string{"01", "10", "30", "33", "35", "44", "45", "49", "50", "51", "52", "53", "54", "86", "91", "98", "99"}
	rn := math.Floor(rand.Float64() * float64(len(rbi)))
	arr := strings.Split(rbi[int(rn)], "")
	str[0], _ = strconv.Atoi(arr[0])
	str[1], _ = strconv.Atoi(arr[1])
	pos = 2

	for {
		if pos < leng-1 {
			str[pos] = (int)(rand.Intn(10) % 10)
		} else {
			break
		}
		pos++
	}

	len_offset = (leng + 1) % 2
	for pos = 0; pos < leng-1; pos++ {
		if (pos+len_offset)%2 != 0 {
			t = str[pos] * 2
			if t > 9 {
				t -= 9
			}

			sum += t
		} else {
			sum += str[pos]
		}
	}

	final_digit = (10 - (sum % 10)) % 10
	str[leng-1] = final_digit
	for _, d := range str {
		imei += fmt.Sprint(d)
	}
	return imei
}

// validate IMEL
func Validate(imei string) bool {
	source := strings.Split(imei, "")
	checksum := 0
	double := false
	for i := len(source) - 1; i > -1; i-- {
		t, _ := strconv.ParseInt(source[i], 10, 8)
		n := int(t)

		if double {
			n = n * 2
		}
		double = !double
		if n >= 10 {
			n = n - 9
		}
		checksum += n
	}
	return checksum%10 == 0
}
