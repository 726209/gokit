package common

import (
	"fmt"
	"time"
)

func ElapsedTimeSince(beginTime time.Time) string {
	elapsedTime := time.Since(beginTime)
	nanoseconds := time.Now().UnixNano() - beginTime.UnixNano() // ns（int64）
	var duration string
	if elapsedTime >= time.Hour {
		duration = fmt.Sprintf("%.2f(小时, %d ns)", elapsedTime.Hours(), nanoseconds)
	} else if elapsedTime >= time.Minute {
		duration = fmt.Sprintf("%.2f(分钟, %d ns)", elapsedTime.Minutes(), nanoseconds)
	} else if elapsedTime >= time.Second {
		duration = fmt.Sprintf("%.2f(秒, %d ns)", elapsedTime.Seconds(), nanoseconds)
	} else if elapsedTime >= time.Millisecond {
		duration = fmt.Sprintf("%v(毫秒, %d ns)", elapsedTime, nanoseconds)
	} else if elapsedTime >= time.Microsecond {
		duration = fmt.Sprintf("%v(微秒, %d ns)", elapsedTime, nanoseconds)
	} else if elapsedTime >= time.Nanosecond {
		duration = fmt.Sprintf("%v(纳秒, %d ns)", elapsedTime, nanoseconds)
	} else {
		duration = fmt.Sprintf("%v(unkonw)", elapsedTime)
	}
	return duration
}

/**
func fib(f int) int {
	if f <= 2 {
		return f
	}
	return fib(f-1) + fib(f-2)
}

func fibN() {
	defer timeCost()() //注意，是对 timeCost()返回的函数进行调用，因此需要加两对小括号
	num := 50
	fmt.Printf("fib(%d) = %v\n", num, fib(num))
}
*/

// TimeCost @brief：耗时统计函数
func TimeCost() func() {
	beginTime := time.Now() // 开始时间
	return func() {
		elapsedTime := time.Since(beginTime)
		nanoseconds := time.Now().UnixNano() - beginTime.UnixNano() // ns（int64）
		var duration string
		if elapsedTime >= time.Hour {
			duration = fmt.Sprintf("%.2f(小时, %d ns)", elapsedTime.Hours(), nanoseconds)
		} else if elapsedTime >= time.Minute {
			duration = fmt.Sprintf("%.2f(分钟, %d ns)", elapsedTime.Minutes(), nanoseconds)
		} else if elapsedTime >= time.Second {
			duration = fmt.Sprintf("%.2f(秒, %d ns)", elapsedTime.Seconds(), nanoseconds)
		} else if elapsedTime >= time.Millisecond {
			duration = fmt.Sprintf("%v(毫秒, %d ns)", elapsedTime, nanoseconds)
		} else if elapsedTime >= time.Microsecond {
			duration = fmt.Sprintf("%v(微秒, %d ns)", elapsedTime, nanoseconds)
		} else if elapsedTime >= time.Nanosecond {
			duration = fmt.Sprintf("%v(纳秒, %d ns)", elapsedTime, nanoseconds)
		} else {
			duration = fmt.Sprintf("%v(unkonw)", elapsedTime)
		}
		fmt.Printf("执行耗时：%s \n", duration)
	}
}
