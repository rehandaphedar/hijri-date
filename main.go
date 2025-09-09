package main

import (
	"fmt"
	"time"

	"github.com/hablullah/go-hijri"
)

var gregorian [13]string = [13]string{
	"",
	"Muḥarram",
	"Ṣafar",
	"Rabīʿ Al Awwal",
	"Rabīʿ Ath Thānī",
	"Jumādā Al Awwal",
	"Jumādā Ath Thānī",
	"Rajab",
	"Shaʿbān",
	"Ramaḍān",
	"Shawwāl",
	"Dhul Qaʿdah",
	"Dhul Ḥijjah",
}

func main() {
	newYear := time.Now()
	hijriDate, _ := hijri.CreateHijriDate(newYear, hijri.Default)
	fmt.Printf("%02d %s %04d-%02d-%02d",
		hijriDate.Day,
		gregorian[hijriDate.Month],
		hijriDate.Year,
		hijriDate.Month,
		hijriDate.Day)
}
