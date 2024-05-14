package helper

import (
	"math"
	"time"
)

func ValidateString(str string, min int, max int) bool {
	if len(str) < min || len(str) > max {
		return false
	}
	return true
}

func ValidateNip(nip int64, role int) bool {
	gender := getGenderCodeFromNip(nip)
	year :=  getYearFromNip(nip)
	month := getMonthFromNip(nip)

	if countDigits(nip) != 13 {
		return false
	}
	if getRoleCodeFromNip(nip) != role{
		return false
	}
	if !checkGenderCode(gender) {
		return false
	}
	if year < 2000 || year > getCurrentYear() {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	
	return true
}

func countDigits(num int64) int {
	if num == 0 {
		return 1
	}

	count := 0
	if num < 0 {
		num = -num
	}

	for num != 0 {
		num /= 10
		count++
	}

	return count
}

func checkGenderCode(num int) bool {
	return num == 1 || num == 2
}

func getRoleCodeFromNip(nip int64) int {
	divisor := int64(math.Pow(10, 10))
	code := nip/divisor
	return int(code)
}

func getGenderCodeFromNip(nip int64) int {
	divisor := int64(math.Pow(10, 9))
	reducedNip := nip/divisor
	code := reducedNip % 10
	return int(code)
}

func getCurrentYear() int {
	return time.Now().Year()
}

func getYearFromNip(nip int64) int {
	divisor := int64(math.Pow(10, 5))
	reducedNip := nip/divisor
	divisor = int64(math.Pow(10, 4))
	code := reducedNip % divisor
	return int(code)
}

func getMonthFromNip(nip int64) int {
	divisor := int64(math.Pow(10, 3))
	reducedNip := nip/divisor
	divisor = int64(math.Pow(10, 2))
	code := reducedNip % divisor
	return int(code)
}