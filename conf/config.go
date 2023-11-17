package conf

// 全局 Config 对象，内存中的配置对象
// 需要在配置加载时被初始化
//		LoadConfigFromToml
//		LoadConfigFromEnv
// 为了被程序在运行时恶意修改，设置为私有变量
// 通过函数获取配置
var config *Config

func C() *Config {
	return config
}

// 默认的 Config 配置
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMySQL(),
	}
}

// Config 应用配置
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

// 默认的 App 配置
func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8050",
	}
}

type App struct {
	Name      string `toml:"name" env:"APP_NAME"`
	Host      string `toml:"host" env:"APP_HOST"`
	Port      string `toml:"port" env:"APP_PORT"`
	Key       string `toml:"key" env:"APP_KEY"`
	EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

// 默认的 Log 配置
func NewDefaultLog() *Log {
	return &Log{
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}

type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

// 默认的 MySQL 配置
func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "demo",
		Password:    "123456",
		Database:    "demo",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
		MaxLifeTime: 60 * 60,
		MaxIdleTime: 60 * 60,
	}
}

type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	// 使用 MySQL 连接池, 需要做一些规划配置
	// 控制当前程序 MySQL 打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 控制 MySQL 打开的连接数
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 一个连接的生命周期, 这个和 MySQL Server 的配置有关, 必须小于 Server 的配置
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// Idle 连接最多允许存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`
}
