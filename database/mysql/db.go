package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

/*
* MysqlConnectiPool
* 数据库连接操作库
* 基于gorm封装开发
 */
type MysqlConnectiPool struct {
}

var instance *MysqlConnectiPool
var once sync.Once

var db *gorm.DB
var err_db error

func GetInstance() *MysqlConnectiPool {
	once.Do(func() {
		instance = &MysqlConnectiPool{}
	})
	return instance
}

/*
* @fuc 初始化数据库连接(可在mail()适当位置调用)
 */
func (m *MysqlConnectiPool) InitDataPool(dsn string) (issucc bool) {
	db, err_db = gorm.Open("mysql",  dsn)
	fmt.Println(err_db)
	if err_db != nil {
		panic(err_db)
		return false
	}
	db.DB().SetMaxOpenConns(200)
	db.DB().SetMaxIdleConns(100)
	db.SingularTable(true)
	db.LogMode(true)
	//关闭数据库，db会被多个goroutine共享，可以不调用
	return true
}

/*
* @fuc  对外获取数据库连接对象db
 */
func (m *MysqlConnectiPool) NewDB() (dbCon *gorm.DB) {
	return db
}