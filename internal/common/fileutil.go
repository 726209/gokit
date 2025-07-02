package common

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// WithTempDir 创建临时目录，执行用户提供的操作函数，结束后自动删除该目录。
func WithTempDir(do func(tempDir string) error) error {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "temp-op-*")
	if err != nil {
		return fmt.Errorf("创建临时目录失败: %w", err)
	}
	// 确保退出时删除临时目录
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	// 执行操作
	return do(tempDir)
}

// DefaultLogPath 外部调用（推荐使用这个）自动判断操作系统并返回默认日志路径
func DefaultLogPath(app string) string {
	return DefaultLogPathWithName(app, "")
}

// DefaultLogPathWithName 返回日志完整路径：<log_dir>/<filename>，若 filename 为空则自动生成。
func DefaultLogPathWithName(name string, filename string) string {
	if filename == "" {
		// 使用时间戳和随机数生成唯一日志名
		timestamp := time.Now().Format("20060102-150405")
		randSuffix := rand.Intn(1000)
		filename = fmt.Sprintf("app-%s-%03d.log", timestamp, randSuffix)
	}

	var dir string
	if runtime.GOOS == "darwin" {
		homeDir, _ := os.UserHomeDir()
		dir = filepath.Join(homeDir, "Library", "Logs", name)
	} else {
		cacheDir, _ := os.UserCacheDir()
		dir = filepath.Join(cacheDir, name)
	}

	_ = os.MkdirAll(dir, 0755) // 创建目录
	return filepath.Join(dir, filename)
}

func DefaultDBPathWithName(name string, dsn string) string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = filepath.Join("./_DATA", dsn)
	} else {
		dir = filepath.Join(dir, name, "_DATA", dsn)
	}
	_ = os.MkdirAll(filepath.Dir(dir), os.ModePerm)
	return dir
}

func DefaultDBPath(dsn string) string {
	return DefaultDBPathWithName("_APP_NAME_", dsn)
}

func DefaultConfigPathWithName(name string, filename string) string {
	dir, err := os.UserConfigDir()
	if err != nil {
		dir = filepath.Join("./_CONFIG", filename)
	} else {
		dir = filepath.Join(dir, name, "_CONFIG", filename)
	}
	_ = os.MkdirAll(filepath.Dir(dir), os.ModePerm)
	return dir
}

func DefaultConfigPath(filename string) string {
	return DefaultConfigPathWithName("_APP_NAME_", filename)
}

func DefaultDownloadPathWithName(name string, filename string) string {
	dir, _ := os.UserConfigDir()
	dir = filepath.Join(dir, name, "./_DOWNLOAD", filename)
	_ = os.MkdirAll(filepath.Dir(dir), os.ModePerm)
	return dir
}

func DefaultDownloadPath(filename string) string {
	return DefaultDownloadPathWithName("_APP_NAME_", filename)
}
