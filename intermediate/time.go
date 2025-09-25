package intermediate

import (
	"fmt"
	"time"
)


func main() {

	//current local time
	fmt.Println(time.Now()) //2025-09-25 09:12:30.836946924 +0000 UTC m=+0.000013042

	//sepcific time
	specificTime := time.Date(2025,time.September, 25, 17,0,0,0,time.UTC)
	fmt.Println(specificTime) //2025-09-25 17:00:00 +0000 UTC

	//parse time 
	//func time.Parse(layout string, value string) (time.Time, error)
	parsedTime, _ :=time.Parse("2006-01-02","2020-05-01") // Mon Jan 2 15:04:05  2006  模版 1 2 3 4 5 6
	parsedTime1, _ :=time.Parse("06-01-02","20-05-01") 
	parsedTime2, _ :=time.Parse("06-1-2","20-5-1") 
	parsedTime3, _ :=time.Parse("06-1-2 15-04","20-5-1 18-03") 
	fmt.Println(parsedTime) //2020-05-01 00:00:00 +0000 UTC
	fmt.Println(parsedTime1) //2020-05-01 00:00:00 +0000 UTC
	fmt.Println(parsedTime2) //2020-05-01 00:00:00 +0000 UTC
	fmt.Println(parsedTime3) //2020-05-01 18:03:00 +0000 UTC

	//Formatting time
	t:=time.Now()
	//func (t time.Time) Format(layout string) string
	fmt.Println("Formatteed time:",t.Format("Monday 06-01-02 04-15")) //Formatteed time:Thursday 25-09-25 41-09

	oneDayLater :=t.Add(time.Hour*24)
	fmt.Println(oneDayLater) //2025-09-26 09:47:50.947353778 +0000 UTC m=+86400.000109335
	fmt.Println(oneDayLater.Weekday()) //Friday

	fmt.Println("Rounded Time:",t.Round(time.Hour)) //Rounded Time: 2025-09-25 10:00:00 +0000 UTC

	// loc,_:=time.LoadLocation("Asia/Kolkata")
	// t = time.Date(2025,time.September,25,17,56,40,00,time.UTC)

	// tLocal:=t.In(loc)

	// //perform rounding
	// roundedTime :=t.Round(time.Hour)
	// roundedTimeLocal :=roundedTime.In(loc)
	// fmt.Println("Original Time (UTC):",t) //Original Time (UTC): 2025-09-25 17:56:40 +0000 UTC
	// fmt.Println("Original Time (Local):",tLocal) //Original Time (Local): 2025-09-25 23:26:40 +0530 IST
	// fmt.Println("rounded Time (UTC):",roundedTime) //rounded Time (UTC): 2025-09-25 18:00:00 +0000 UTC
	// fmt.Println("rounded Time  (Local):",roundedTimeLocal) //rounded Time  (Local): 2025-09-25 23:30:00 +0530 IST

	fmt.Println("truncated Time",t.Truncate(time.Hour)) //round down 向下取整
	loc,_:=time.LoadLocation("America/New_York")

	//convert time to location
	tInNY:=time.Now().In(loc)
	fmt.Println("NewYork Time",tInNY) //NewYork Time 2025-09-25 06:32:35.868220122 -0400 EDT


	t1 := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)
	duration := t2.Sub(t1)
	fmt.Println("Duration:", duration)

	// Compare times
	fmt.Println("t2 is after t1?", t2.After(t1))
}