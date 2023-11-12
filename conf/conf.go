package conf

// Config 应用配置
type Config struct {
	App   *app   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

type app struct {
	Name      string `toml:"name" env:"APP_NAME"`
	Host      string `toml:"host" env:"APP_HOST"`
	Port      string `toml:"port" env:"APP_PORT"`
	Key       string `toml:"key" env:"APP_KEY"`
	EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

// Log todo
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
}

// MySQL todo
type MySQL struct {
	Host     string `toml:"host" env:"D_MYSQL_HOST"`
	Port     string `toml:"port" env:"D_MYSQL_PORT"`
	UserName string `toml:"username" env:"D_MYSQL_USERNAME"`
	Password string `toml:"password" env:"D_MYSQL_PASSWORD"`
	Database string `toml:"database" env:"D_MYSQL_DATABASE"`
	// 使用 MySQL 连接池, 需要做一些规划配置
	// 控制当前程序 MySQL 打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"D_MYSQL_MAX_OPEN_CONN"`
	// 控制 MySQL 打开的连接数
	MaxIdleConn int `toml:"max_idle_conn" env:"D_MYSQL_MAX_IDLE_CONN"`
	// 一个连接的生命周期, 这个和 MySQL Server 的配置有关, 必须小于 Server 的配置
	MaxLifeTime int `toml:"max_life_time" env:"D_MYSQL_MAX_LIFE_TIME"`
	// Idle 连接最多允许存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"D_MYSQL_MAX_idle_TIME"`
}
