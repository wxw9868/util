package delayqueue

import (
	"fmt"
	"testing"
	"time"
)

func TestDelayMessage_Start(t *testing.T) {
	// 创建延迟消息
	dm := NewDelayMessage()
	// 添加任务
	_ = dm.AddTask(time.Now().Add(time.Second*10), "test1", func(args ...interface{}) {
		fmt.Println(args...)
	}, []interface{}{1, 2, 3})
	_ = dm.AddTask(time.Now().Add(time.Second*10), "test2", func(args ...interface{}) {
		fmt.Println(args...)
	}, []interface{}{4, 5, 6})
	_ = dm.AddTask(time.Now().Add(time.Second*10), "test3", func(args ...interface{}) {
		fmt.Println(args...)
	}, []interface{}{"hello", "world", "test"})
	_ = dm.AddTask(time.Now().Add(time.Second*10), "test4", func(args ...interface{}) {
		sum := 0
		for arg := range args {
			sum += arg
		}
		fmt.Println("sum: ", sum)
	}, []interface{}{1, 2, 3})

	// 40秒后关闭
	time.AfterFunc(time.Second*40, func() {
		dm.Close()
	})
	dm.Start()
}
