package work

import (
	"cloudauto/models"
	"cloudauto/utils"
	"flag"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"path/filepath"
)

var (
	defaultConf = "conf/app.conf"
	confPath    = flag.String("conf", "", "please set autocloud conf path")
	version     = flag.Bool("version", false, "autocould version")
	Version     = "v1.0"
	//RootDir     = ""
)

func init() {
	initFalg()
	initConfig()
	initMySQL()
}

func initFalg() {
	flag.Parse()
	if *version == true {
		fmt.Printf(Version)
		os.Exit(0)
	}
}

func initConfig() {
	RootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	confFile := *confPath
	if *confPath == "" {
		confFile = filepath.Join(RootDir, defaultConf)
	}
	ok, _ := utils.NewFile().PathIsExists(confFile)
	if ok == false {
		log.Println("conf file" + confFile + " not exist.")
		os.Exit(1)
	}

	//init config file
	beego.LoadAppConfig("ini", confFile)

	//init name
	beego.AppConfig.Set("sys.name", "mm-wiki")
	beego.BConfig.AppName = "sys.name"
	beego.BConfig.ServerName = "sys.name"

	//set static path
	beego.SetStaticPath("/static", filepath.Join(RootDir, "static"))

	//set views path
	beego.SetViewsPath(filepath.Join(RootDir, "views"))

	//init logs
	logConfigs, err := beego.AppConfig.GetSection("log")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	for adapter, config := range logConfigs {
		logs.SetLogger(adapter, config)
	}
	logs.SetLogFuncCall(true)
}
func initMySQL() {
	host, _ := beego.AppConfig.String("mysql::host")
	user, _ := beego.AppConfig.String("mysql::user")
	pass, _ := beego.AppConfig.String("mysql::pass")
	port, _ := beego.AppConfig.String("mysql::port")
	maxIdle, _ := beego.AppConfig.Int("mysql::conn_max_idle")
	maxConn, _ := beego.AppConfig.Int("mysql::conn_max_connection")
	database, _ := beego.AppConfig.String("mysql::db_name")
	dsn := fmt.Sprintf(user + ":" + pass + "@" + "tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(dsn)
	db, err := gorm.Open("mysql", dsn)
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxConn)
	fmt.Println(maxConn, maxIdle)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	models.G = db
	logs.Info("mysql init success.")
}
