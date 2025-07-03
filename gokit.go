package gokit

import (
	"github.com/726209/gokit/internal/common"
	"time"
)

// ******************** basic 类【BEGIN】 ********************
var (
	// Capitalize 可选汇总入口（封装常用函数）
	Capitalize       = common.Capitalize
	PrettyStruct     = common.PrettyStruct
	PrettyJSONString = common.PrettyJSONString
	JSONString       = common.JSONString
	SnakeToCamel     = common.SnakeToCamel
	ToSnakeCase      = common.ToSnakeCase
)

// ********************* basic 类【END】 *********************

// ******************** string 类【BEGIN】 ********************

// ********************* string 类【END】 *********************

// ******************** time 类【BEGIN】 ********************

var Time = struct {
	ElapsedTimeSince func(start time.Time) string
	TimeCost         func() func()
}{
	ElapsedTimeSince: common.ElapsedTimeSince,
	TimeCost:         common.TimeCost,
}

// ********************* time 类【END】 *********************

// ******************** collection【BEGIN】 ********************
var (
	// Map 映射函数：将 []T 映射为 []R
	Map = common.Map

	// Filter 过滤函数：返回满足条件的 []T 子集
	Filter = common.Filter

	// Reduce 累加函数：将 []T 聚合成一个 R 值
	Reduce = common.Reduce

	// Find 查找函数：返回首个满足条件的元素及是否找到
	Find = common.Find
)

// ********************* collection【END】 *********************

// ******************** file 类【BEGIN】 ********************

var (
	Exists = common.Exists
	// WithTempDir 创建临时目录，执行用户提供的操作函数，结束后自动删除该目录。
	WithTempDir = common.WithTempDir

	// DefaultLogPath 外部调用（推荐使用这个）自动判断操作系统并返回默认日志路径
	DefaultLogPath = common.DefaultLogPath

	// DefaultLogPathWithName 返回日志完整路径：<log_dir>/<filename>，若 filename 为空则自动生成。
	DefaultLogPathWithName      = common.DefaultLogPathWithName
	DefaultDBPathWithName       = common.DefaultDBPathWithName
	DefaultDBPath               = common.DefaultDBPath
	DefaultConfigPathWithName   = common.DefaultConfigPathWithName
	DefaultConfigPath           = common.DefaultConfigPath
	DefaultDownloadPathWithName = common.DefaultDownloadPathWithName
	DefaultDownloadPath         = common.DefaultDownloadPath
)

// ********************* file 类【END】 *********************
