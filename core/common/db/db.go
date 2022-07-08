package db

import (
	"sync"
	"time"

	"fabapi/core/common/e"
	"fabapi/pkg/utils"

	"github.com/go-xorm/xorm"
)

var db Database

type Database struct {
	db    *xorm.Engine
	clock sync.Mutex

	config *MysqlConfig
	closed bool
	retryC int
}

func NewDatabase(db *xorm.Engine, config *MysqlConfig, retryC int) *Database {
	return &Database{
		db:     db,
		config: config,
		retryC: retryC,
	}
}

type Session struct {
	*xorm.Session
}

func DatabaseInit(config *MysqlConfig) {
	db.config = config
	db.retryC = 3
	err := db.init()
	if err != nil {
		panic(err)
	}
}
func GetDefaultDatabse() *Database {
	return &db
}

func (d *Database) init() error {
	d.db = getDBEngine(d.config)
	return d.db.Ping()
}

func getDBEngine(config *MysqlConfig) *xorm.Engine {
	logger.Debug("数据库引擎初始化")
	engine := XormEngineInit(config)
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logger.Error("设置时区失败:", err)
	} else {
		engine.TZLocation = location
		engine.DatabaseTZ = location
	}
	return engine
}

func (d *Database) Ping() *xorm.Engine {
	err := d.db.Ping()
	if err != nil {
		logger.Warning("数据库链接异常", err)
		err := d.retry()
		if err != nil {
			logger.Error("")
			panic(err)
		}
	}
	return d.db
}

func (d *Database) retry() error {
	d.clock.Lock()
	defer d.clock.Unlock()
	var err error
	for i := 0; i < d.retryC; i++ {
		if err = d.init(); err == nil {
			break
		}
		logger.Warning("数据库链接异常", err)
	}
	return err

}

func (d *Database) SessionHandle(action func(session *xorm.Session) error) error {
	d.Ping()
	session := d.db.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		logger.Errorf("session begin error : %+v", err)
		return utils.NewError(e.ERROR_DB_SESSION_BEGIN_FAIL, err)
	}

	err := action(session)
	if err != nil {
		if err := session.Rollback(); err != nil {
			logger.Warningf("session rollback error : %+v", err)
		}
		//error 在业务层封装了， 不需要再包装
		return err
	}

	if err := session.Commit(); err != nil {
		logger.Errorf("session commit error : %+v", err)
		if err := session.Rollback(); err != nil {
			logger.Warningf("session rollback error : %+v", err)
		}
		return utils.NewError(e.ERROR_DB_SESSION_COMMIT_FAIL, err)
	}

	return nil
}
