package main

import (
	"fmt"
	"time"
)

//Epoch
//Starting Point: 00:00:00 UTC on January 1, 1970 (not counting leap seconds)
//Epoch Time Units: Seconds Milliseconds
//Epoch Time Values:Positive Values (time after 1970 1 1),Negative Values((time before 1970 1 1))
//Unix Time Functions:time.Now() time.Unix() time.Since()
//Epoch Applications: Database Storage, System Timestamps,Cross-Platform Compatibility(跨平台兼容)

func main() {
	// 00:00:00 UTC on Jan 1, 1970
	now:=time.Now()
	unixTime:=now.Unix()
	fmt.Println("current unix time:",unixTime) //1758796933

	//func time.Unix(sec int64, nsec int64) time.Time
  t:=time.Unix(unixTime,0) 
	fmt.Println(t) //2025-09-25 10:45:14 +0000 UTC
	fmt.Println("Time:", t.Format("2006-01-02")) //Time: 2025-09-25
}
