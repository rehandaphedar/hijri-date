package main

import (
	"fmt"
	"time"

	"github.com/hablullah/go-hijri"
)

var gregorian [13]string = [13]string{
	"",
	"Muharram",
	"Safar",
	"Rabi' Al-Awwal",
	"Rabi' Al-Thani",
	"Jumada Al-Awwal",
	"Jumada Al-Thani",
	"Rajab",
	"Sha'aban",
	"Ramadan",
	"Shawwal",
	"Dhu Al-Qi'dah",
	"Dhu Al-Hijjah",
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
