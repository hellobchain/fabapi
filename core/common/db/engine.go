package db

import (
	"github.com/wsw365904/wswlog/wlogging"
	"io/ioutil"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	yaml "gopkg.in/yaml.v2"
)

//var logger = log.GetLogger("xorm")
var logger = wlogging.MustGetLoggerWithoutName()

type Xorm struct {
	Config *MysqlConfig `yaml:"xorm"`
}

type MysqlConfig struct {
	Drivename string `yaml:"drivename"`
	Database  string `yaml:"database"`
	Ip        string `yaml:"ip"`
	Port      string `yaml:"port"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	LogLevel  string `yaml:"loglevel"`
	Showsql   bool   `yaml:"showsql"`
	Maxidle   int    `yaml:"maxidle"`
	Maxopen   int    `yaml:"maxopen"`
}

func newXorm() *Xorm {
	return &Xorm{
		Config: &MysqlConfig{},
	}
}
func loadConfig(file string) *MysqlConfig {
	cfg, err := ioutil.ReadFile(file)
	if err != nil {
		logger.Error(err.Error())
	}
	var xorm = newXorm()
	err = yaml.Unmarshal(cfg, xorm)
	if err != nil {
		logger.Error(err.Error())
	}
	return xorm.Config
}

func GetEngine(configFile string) *xorm.Engine {
	config := loadConfig(configFile)
	//conn string
	return XormEngineInit(config)
}

func XormEngineInit(config *MysqlConfig) *xorm.Engine {
	//conn string
	conn := config.User + ":" + config.Password + "@tcp(" + config.Ip + ":" + config.Port + ")/" + config.Database + "?charset=utf8"
	engine, err := xorm.NewEngine(config.Drivename, conn)
	if err != nil {
		logger.Error(err.Error())
	}
	// 打印sql
	xormLogger := &OrmLogger{
		logger: logger,
		level:  OrmLevel(config.LogLevel),
	}
	engine.SetLogger(xormLogger)
	engine.ShowSQL(config.Showsql)
	engine.SetMaxIdleConns(config.Maxidle)
	engine.SetMaxOpenConns(config.Maxopen)
	//连接生存时间半个小时
	engine.SetConnMaxLifetime(1800 * time.Second)
	return engine
}

func OrmLevel(level string) core.LogLevel {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return core.LOG_DEBUG
	case "INFO":
		return core.LOG_INFO
	case "WARNING":
		return core.LOG_WARNING
	case "ERROR":
		return core.LOG_ERR
	case "OFF":
		return core.LOG_OFF
	case "UNKNOWN":
		return core.LOG_UNKNOWN
	default:
		logger.Warning("config orm level unknown, current level set info")
		return core.LOG_INFO
	}
}
