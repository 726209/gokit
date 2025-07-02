package gokit

import (
	"github.com/726209/gokit/internal/common"
)

// ******************** basic 类【BEGIN】 ********************
// ********************* basic 类【END】 *********************

// ******************** time 类【BEGIN】 ********************
var ElapsedTimeSince = common.ElapsedTimeSince

// TimeCost @brief：耗时统计函数
var TimeCost = common.TimeCost

// ********************* time 类【END】 *********************

// ******************** string 类【BEGIN】 ********************

// Capitalize 可选汇总入口（封装常用函数）
var Capitalize = common.Capitalize
var PrettyStruct = common.PrettyStruct
var PrettyJSONString = common.PrettyJSONString
var JSONString = common.JSONString

// ********************* string 类【END】 *********************

// ******************** file 类【BEGIN】 ********************

// WithTempDir 创建临时目录，执行用户提供的操作函数，结束后自动删除该目录。
var WithTempDir = common.WithTempDir

// DefaultLogPath 外部调用（推荐使用这个）自动判断操作系统并返回默认日志路径
var DefaultLogPath = common.DefaultLogPath

// DefaultLogPathWithName 返回日志完整路径：<log_dir>/<filename>，若 filename 为空则自动生成。
var DefaultLogPathWithName = common.DefaultLogPathWithName
var DefaultDBPathWithName = common.DefaultDBPathWithName
var DefaultDBPath = common.DefaultDBPath
var DefaultConfigPathWithName = common.DefaultConfigPathWithName
var DefaultConfigPath = common.DefaultConfigPath
var DefaultDownloadPathWithName = common.DefaultDownloadPathWithName
var DefaultDownloadPath = common.DefaultDownloadPath

// ********************* file 类【END】 *********************
