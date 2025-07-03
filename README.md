
# gokit

🚀 A lightweight and extensible Go utility library providing enhanced **database abstraction**, **configurable logging**, and **developer-friendly features** out of the box.

---

## ✨ Features

### 🛠️ Database Repository Helper (GORM-based)
- Simplified **DB initialization** (MySQL / SQLite auto-detection)
- Context-aware `Repository` with chainable helpers
- Built-in helpers:
    - `Paginate()`, `Transaction()`, `UpdateBatch()`, `Exists()`, `Count()`
    - Auto table creation: `CreateTable()`
- Generic pagination: `Paginating[T]()`
- Clean separation from business logic

### 📒 Logging System
- Unified logging with support for:
    - **Console output** (standard or pretty via [pterm](https://github.com/pterm/pterm))
    - **Rotating file logging** (via [zap + lumberjack](https://pkg.go.dev/go.uber.org/zap))
- Configurable via `Config` struct:
    - `Level`: debug, info, warn, error
    - `PrettyPrint`: enable colorful and structured pterm logs
    - `OutputPath`: specify log file path for persistent logging
- Logs written to **console and file simultaneously**
- Supports log level conversion across:
    - `zapcore.Level`
    - `pterm.LogLevel`
    - `gormlogger.LogLevel`

---

## 📦 Installation

```bash
go get github.com/726209/gokit
```

---

## 🚀 Quick Start

### Initialize Logger

```go
import "github.com/726209/gokit/logger"

logger.InitLogger(logger.Config{
    Level:       logger.InfoLevel,
    PrettyPrint: true,
    OutputPath:  "logs/app.log",
})
```

### Setup Database

```go
import "github.com/726209/gokit/repository"

repo, err := repository.CreateRepository(os.Getenv("DSN"))
if err != nil {
    panic(err)
}
repo.CreateTable(&User{}) // Automigrate schema
```

### Paginate Example

```go
type User struct {
    ID   uint
    Name string
}

result, err := repo.Paginating[User](1, 10)
if err != nil {
    panic(err)
}
fmt.Println(result.Data)
```

---

## 🧪 Testing

```bash
go test ./...
```

---

## 📁 Project Structure

```
gokit/
├── repository/  // Repository & GORM wrappers
├── logger/      // zap + pterm unified logging
├── examples/    // Usage examples
├── utils/       // Helper functions
└── ...
```

---

## 🛡 License

MIT License. ©️ 2025 [726209](https://github.com/726209)
