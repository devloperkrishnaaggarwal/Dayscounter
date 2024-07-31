package main

import (
	"fmt"
	"time"
)

func takeAge() time.Time {
	fmt.Println("Enter your Date of Birth: YYYY/MM/DD ")
	var age string
	fmt.Scan(&age)
	date, err := time.Parse("2006/01/02", age)
	fmt.Println("Error occurred in parsing date:", err)

	return date
}

func nowTime() time.Time {
	return time.Now()
}

func extraTime(age string, now time.Time) time.Time {

	date, err := time.Parse("2006/01/02", age)
	if err != nil {
		fmt.Println("Error occurred in parsing date:", err)
	}

	duration := date.Sub(now)

	newDate := now.Add(duration)
	return newDate
}

func specificAge(age string) time.Time {
	date, err := time.Parse("2006/01/02", age)
	if err != nil {
		fmt.Println("Error occurred in parsing date:", err)
	}
	return date
}

func DataPresentation(a time.Time, b time.Time, c1 string, c2 string) string {
	myTime := b.Sub(a)
	myTimeDays := int(myTime.Hours() / 24)

	finalOutput := c1 + fmt.Sprintf("%d", myTimeDays) + c2
	return finalOutput
}

func main() {
	fmt.Println("Enter your Date of Birth: YYYY/MM/DD ")
	var age string
	fmt.Scan(&age)
	fmt.Println(DataPresentation(specificAge(age), nowTime(), "You are ", " days old."))
	fmt.Println(DataPresentation(nowTime(), specificAge("2100/01/01"), "You have to wait ", " days to see the next 22nd century"))
	fmt.Println(DataPresentation(nowTime(), specificAge("2061/07/28"), "You have to wait ", " days to see Halley's Comet"))

}
