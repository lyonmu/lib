package common

import (
	"github.com/sony/sonyflake"
	"strconv"
	"sync"
)

var idInstance *sonyflake.Sonyflake

var once = new(sync.Once)

func init() {
	once.Do(func() {
		settings := sonyflake.Settings{}
		idInstance = sonyflake.NewSonyflake(settings) // 创建snowflake实例
	})
}

// ID 生成全局唯一ID
func ID() string {
	v, _ := idInstance.NextID()
	return strconv.FormatUint(v, 10)
}
