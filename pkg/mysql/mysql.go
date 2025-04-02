package mysql

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
	"wrblog-api-go/config"
	"wrblog-api-go/pkg/mylog"
)

func Db(name ...string) *gorm.DB {
	if len(name) > 0 {
		return dbs[name[0]]
	} else {
		return dbs["master"]
	}
}
func Cx(name ...string) *gorm.DB {
	if len(name) > 0 {
		return dbs[name[0]].Begin()
	} else {
		return dbs["master"].Begin()
	}
}

var dbs = make(map[string]*gorm.DB)

// database-mysql 配置映射
type databaseConf struct {
	Database struct {
		Mysql struct {
			SourceName string        `yaml:"source-name"`
			Master     *sourceInfo   `yaml:"master"`
			Salves     []*sourceInfo `yaml:"salves"`
		} `yaml:"mysql"`
	} `yaml:"database"`
}

type sourceInfo struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db-name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Timeout  string `yaml:"timeout"`
}

func init() {
	//初始化mysql配置信息
	analysisMysql().initMysqlDbs()
	fmt.Println("Mysql db initialization connection successful......")
}

func analysisMysql() (mysqlConf *databaseConf) {
	// 打开 YAML 文件
	file, err := os.ReadFile(config.ConfPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse mysql configuration:%s", err.Error()))
	}
	err = yaml.Unmarshal(file, &mysqlConf)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse mysql configuration:%s", err.Error()))
	}
	return
}

func (mysqlConf *databaseConf) initMysqlDbs() {
	master := mysqlConf.Database.Mysql.Master
	masterDb, err := master.getMysqlDb()
	if err != nil {
		panic(fmt.Sprintf("Mysql Connection Master initialization failed:%s", err.Error()))
	}
	dbs["master"] = masterDb
	dbs[master.DbName] = masterDb
	fmt.Println(fmt.Sprintf("Mysql Connection initialization successful: Master ---> %s", master.DbName))
	salves := mysqlConf.Database.Mysql.Salves
	for index, salve := range salves {
		salveDb, errSalve := salve.getMysqlDb()
		if errSalve != nil {
			panic(fmt.Sprintf("Mysql Connection Salve initialization failed:%s", err.Error()))
		}
		dbs[fmt.Sprintf("salve-%d", index)] = salveDb
		dbs[salve.DbName] = salveDb
		fmt.Println(fmt.Sprintf("Mysql Connection initialization successful: Salve-%d ---> %s", index, salve.DbName))
	}
}

// 初始化链接
func (source *sourceInfo) getMysqlDb() (db *gorm.DB, err error) {
	//拼接数据库配置
	MysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s", source.Username, source.Password, source.Host, source.Port, source.DbName, source.Charset, source.Timeout)
	// 打开连接失败
	db, err = gorm.Open(mysql.Open(MysqlDSN), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: logger.New(mylog.Logger, logger.Config{
			SlowThreshold: 100 * time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,            // 日志级别
			Colorful:      true,                   // 彩色打印
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tp_",
			SingularTable: true,
		},
	})
	return db, err
}
