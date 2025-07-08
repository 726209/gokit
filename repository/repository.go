package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/726209/gokit"
	"github.com/726209/gokit/logger"
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*

// SimulateGenData 模拟生成数据
func SimulateGenData(repo *db.Repository, len int) {
	if len < 1 {
		len = 10000
	}
	gofakeit.Seed(0)
	pterm.Info.Println("Username:", gofakeit.Username())
	var users []User
	for i := 0; i < len; i++ {
		user := User{
			Name:     gofakeit.Name(),
			Username: gofakeit.Username(),
			Email:    gofakeit.Email(),
			Phone:    gofakeit.Phone(),
			Address:  gofakeit.Address().Address,
			Bio:      GenerateBio(10), // 生成10个词的个人签名
		}
		users = append(users, user)
	}
	repo.CreateBatch(users, 100)
}
*/

// Repository wraps gorm.DB and provides utility helpers.
type Repository struct {
	*gorm.DB // 匿名嵌入，没有指定字段名，自动将嵌入类型的类型名作为字段名，通过嵌入字段直接访问 A 的字段
	//db *gorm.DB // 小写，外部无法访问
}

// InitRepository 从环境变量中读取 DSN 并创建数据库连接
func InitRepository() (*Repository, error) {
	// 读取 .env 环境变量
	_ = godotenv.Load()

	dsn := os.Getenv("DSN")
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("DSN 不能为空，请在环境变量中设置，或者调用`CrateRepository(dsn string)`方法。")
	}

	return CrateRepository(dsn)
}

// CrateRepository 根据 dsn 创建 MySQL 或 SQLite 连接
func CrateRepository(dsn string) (repo *Repository, err error) {
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("DSN 不能为空，请提供数据库连接字符串。")
	}
	var dialect gorm.Dialector
	if isSQLite(dsn) {
		dialect = sqlite.Open(dsn)
		dbpath, _ := filepath.Abs(dsn) // SQLite 文件绝对路径
		logger.Infof("使用 SQLite(dsn:%s) 数据库: %s", dsn, dbpath)
	} else {
		dialect = mysql.Open(dsn)
		logger.Infof("使用 MySQL 数据库: %s", dsn)
	}
	level := logger.GormLogLevel()
	db, err := gorm.Open(dialect, &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		Logger: &GormLogger{
			LogLevel:      level,
			SlowThreshold: 200 * time.Millisecond,
		},

		//Logger: logger.New(
		//	log.New(os.Stdout, "\r\n[GORM] ", log.LstdFlags), // 日志输出目标
		//	logger.Config{
		//		SlowThreshold:             time.Second, // 慢查询阈值
		//		LogLevel:                  logger.Info, // 日志级别：Silent, Error, Warn, Info
		//		IgnoreRecordNotFoundError: true,        // 忽略 ErrRecordNotFound 错误
		//		Colorful:                  true,        // 彩色输出
		//	},
		//),
	})
	if err != nil {
		return nil, fmt.Errorf("连接数据库(dsn:%s) 失败，Error = %w", dsn, err)
	}
	db.Logger.Info(nil, "创建（打开）数据库(dsn:%s) 成功！", dsn)
	pterm.Info.Printf("创建（打开）数据库(dsn:%s) 成功！", dsn)
	return New(db), nil
}

func isSQLite(dsn string) bool {
	dsn = strings.ToLower(dsn)
	return strings.HasSuffix(dsn, ".db") ||
		strings.HasSuffix(dsn, ".sqlite") ||
		strings.HasSuffix(dsn, ".sqlite3") ||
		strings.TrimSpace(dsn) == ":memory:"
}

// New creates a new DB wrapper.
func New(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// WithContext binds the context to DB operations.
func (repo *Repository) WithContext(ctx context.Context) *Repository {
	return &Repository{DB: repo.DB.WithContext(ctx)}
}

func (repo *Repository) WithUnscoped() *Repository {
	return &Repository{DB: repo.DB.Unscoped()}
}

// Paginate applies offset and limit.
func (repo *Repository) Paginate(offset, limit int) *Repository {
	return &Repository{DB: repo.DB.Offset(offset).Limit(limit)}
}

// Transaction wraps GORM transaction handling.
func (repo *Repository) Transaction(ctx context.Context, fn func(tx *Repository) error) error {
	return repo.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(&Repository{DB: tx})
	})
}

// Count counts the total number of records for a model.
func (repo *Repository) Count(model interface{}) (int64, error) {
	var count int64
	err := repo.DB.Model(model).Count(&count).Error
	return count, err
}

// Exists returns true if the given model exists with condition.
func (repo *Repository) Exists(model interface{}, query string, args ...interface{}) (bool, error) {
	var count int64
	err := repo.DB.Model(model).Where(query, args...).Count(&count).Error
	return count > 0, err
}

// UpdateBatch updates a list of models in a transaction.
func (repo *Repository) UpdateBatch(ctx context.Context, models []interface{}) error {
	return repo.Transaction(ctx, func(tx *Repository) error {
		for _, model := range models {
			if err := tx.DB.Save(model).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// LogSQL wraps an operation and logs its duration.
func (repo *Repository) LogSQL(label string, fn func(*Repository) error) error {
	start := time.Now()
	err := fn(repo)
	duration := time.Since(start)
	log.Printf("[dbx] %s took %s", label, duration)
	return err
}

// CreateTable : Migrate the schema
func (repo *Repository) CreateTable(dst ...interface{}) error {
	pterm.Info.Println("自动迁移表结构(Schema)")
	pterm.Info.Printf("%T", dst[0])
	pterm.Debug.Println("模型结构:", reflect.TypeOf(dst[0]).String())
	err := repo.DB.AutoMigrate(dst...)
	if err != nil {
		return fmt.Errorf("迁移表结构 schema 失败，Error = %w", err)
	}
	pterm.Info.Println("创建（打开）数据表(products) 成功！", gokit.JSONString(repo.printTableName(dst...)))
	return nil
}

func (repo *Repository) printTableName(dst ...interface{}) map[string]any {
	tables := make(map[string]any)
	for _, model := range dst {
		stmt := &gorm.Statement{DB: repo.DB}
		if err := stmt.Parse(model); err != nil {
			fmt.Printf("解析失败: %v", err)
			continue
		}
		tables[fmt.Sprintf("模型 %T ==", model)] = stmt.Schema.Table
	}
	return tables
}

func (repo *Repository) Create(value interface{}) {
	pterm.Info.Printf("插入数据:(%s)", gokit.JSONString(value))
	result := repo.DB.Create(value)
	if result.Error != nil {
		panic(fmt.Sprintf("插入数据失败，Error:%v", result.Error))
	}
	pterm.Info.Printfln("插入数据成功(%d row)", result.RowsAffected)
}

func (repo *Repository) CreateBatch(value interface{}, batchSize int) {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Slice {
		panic("BatchCreate 仅支持 slice 类型")
	}
	length := v.Len()
	logger.Infof("批量插入 %d 条数据 (batchSize=%d)", length, batchSize)

	tx := repo.DB.CreateInBatches(value, batchSize)
	if tx.Error != nil {
		panic(fmt.Sprintf("批量插入数据失败，Error: %v", tx.Error))
	}
	logger.Infof("批量插入完成，共 %d 条记录，实际插入 %d 条", length, tx.RowsAffected)
}

func (repo *Repository) Read(dest interface{}, queryFunc func(*Repository, interface{}, ...interface{}) *gorm.DB, conditions ...interface{}) {
	tx := queryFunc(repo, dest, conditions...)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			logger.Warnf("未找到记录: %v", tx.Error)
			return
		}
		panic(fmt.Sprintf("查询数据失败，Error:%v", tx.Error))
	}
	logger.Infof("查询成功，共 %d 行", tx.RowsAffected)
}

type PageResult[T any] struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`       // 总记录数
	TotalPages int   `json:"total_pages"` // 总页数
	Count      int   `json:"count"`       // 当前页记录数
	List       []T   `json:"list"`        // 当前页数据
}

// Paginating 分页查询
//
// 示例调用：
// var db *gorm.DB // your initialized DB
// res, err := Paginating[User](db.Where("age > ?", 18), 0, 10)
// if err != nil {
// log.Fatal(err)
// }
// fmt.Printf("共 %d 条记录，当前页 %d 条\n", res.Total, res.Count)
func Paginating[T any](db *gorm.DB, page, pageSize int) (PageResult[T], error) {
	var (
		total int64
		list  []T
	)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 1000 {
		pageSize = 10
	}

	// 查询总数
	if err := db.Model(new(T)).Count(&total).Error; err != nil {
		return PageResult[T]{}, err
	}
	// 查询当前页数据
	if err := db.
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&list).Error; err != nil {
		return PageResult[T]{}, err
	}

	// 构造结果
	return PageResult[T]{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
		Count:      len(list),
		List:       list,
	}, nil
}
