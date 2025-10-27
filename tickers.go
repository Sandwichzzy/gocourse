package main

import (
	"fmt"
	"time"
)

// -ticker 常用于需要定期执行的任务，比如定期收集数据、定期发送心跳、定期检查状态等。
// -与 timer 的区别
// timer 只在指定的时间后触发一次，而 ticker 会周期性地触发。
// Timer - 单次
// timer := time.NewTimer(2 * time.Second)
// <-timer.C // 等待2秒，只触发一次
// Ticker - 重复
// ticker := time.NewTicker(2 * time.Second)
// for range ticker.C { // 每2秒触发一次
//     fmt.Println("tick")
// }
// timer 可以通过 Reset 重新设置，而 ticker 可以通过 Reset 来改变其周期
// （但注意，ticker 的 Reset 方法在使用上有一些注意事项，通常建议在停止后重置，并且确保通道中的事件已经被取出）。

// ========== ticker along with timer
func main() {
	// 创建一个每秒触发一次的 ticker
	ticker := time.NewTicker(1 * time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t:= <-ticker.C:
			fmt.Println("Tick at:", t)
		case <-stop:
			fmt.Println("Stopping ticker")
			return
		}
	}

}






//  ========= SCHEDULING LOGGING, 定期任务执行 PERIODIC TASKS, POLLING FOR UPDATES
// func periodicTask(){
// 	fmt.Println("Performing periodic task at",time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select{
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// func main() {

// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()
// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick at:", tick)
// 	// }
// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println(i)
// 	}

// 	// for tick := range ticker.C {
// 	// 	i *= 2
// 	// 	fmt.Println(tick)
// 	// }
// }
