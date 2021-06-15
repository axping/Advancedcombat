package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

var Config string = `
#公告设置
[app]
	db="SQLite"
[server]
	address=":5900"
	root="public"
`
var (
	rootPath string = "public"
	address  string = ":5900"
	dbType   string = "SQLite"
)

//以下用于返回结果响应体想改内容
//Respond 响应体
type R struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

//创建一个返回结果
func NewR(message string, code int, data interface{}) *R {
	return &R{Message: message, Code: code, Data: data}
}

//
type D map[string]interface{}

//将结果返回未JSON数据
func (r R) JSON() string {
	result, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(result)
}

//响应结果结束

var MyDB *sql.DB

func main() {
	//如果是第一次启动需要控制生成配置文件
	checkConf()
	//读取配置文件
	loadConf()
	//初始化数据库
	loadData()
	//按照启动服务
	server()
}

//检查配置文件
func checkConf() {
	_, err := os.Open("config.toml")
	if err != nil {
		log.Println("Success to loading config!")
	}
	//如果是文件不存在错误，则创建配置文件
	if os.IsNotExist(err) {
		f, err := os.Create("config.toml")
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("The config file was generated successfully！Please restart this program")
			f.Write([]byte(Config))
			os.Exit(0)
		}
		defer f.Close()
	}
}

//加载配置文件
func loadConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")    // 配置文件路径
	err := viper.ReadInConfig() // 读取配置文件信息
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	dbType = viper.GetString("app.db")
	address = viper.GetString("server.address")
	rootPath = viper.GetString("server.root")
	fmt.Printf("数据库类型[%s],服务监听地址[%s],静态文件目录[%s]", dbType, address, rootPath)
}

//初始化数据库
func loadData() {
	dbfile := "deql.db"
	_, err := os.Open(dbfile)
	//检查Db文件是否存在
	if os.IsNotExist(err) {
		MyDB, err = sql.Open("sqlite3", dbfile)
		//表不存在
		if err != nil {
			log.Println("异常一", err.Error())
			checkErr(err)
		}
		//初始化创建表
		initTable()
		return
	}
	MyDB, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Println("异常三", err.Error())
		checkErr(err)
	}
}

//初始化创建表
func initTable() {
	//创建用户表
	user := `CREATE TABLE usertab(
		_id integer not null primary key,
		name VARCHAR(32) not null,
		session VARCHAR(32),
	
		uid text,
		password text not null,
		phone text
		)`
	_, err := MyDB.Exec(user)
	if err != nil {
		log.Println("create userinfo table happen error", err.Error())
		os.Remove("deql.db")
		os.Exit(0)
	}
	//创建定时任务Crontab
	cron := `CREATE TABLE db_crontab(
		_id integer not null primary key,
		uid text, 
		cron text not null
	)
	`
	_, err = MyDB.Exec(cron)
	if err != nil {
		log.Println("create corn table happen error", err.Error())
		os.Remove("deql.db")
		os.Exit(0)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func server() {
	http.HandleFunc("/", index)
	//URL 定义
	//用户注册:/user/register
	//用户登入:/user/login
	//用户属性:/user/info
	//用户授权:/user/licence

	//发布文章属性

	//生成uuid
	http.HandleFunc("/uuid/generate", genUuid)
	//授权信息生成器
	http.HandleFunc("/licence", licence)
	// //用于生成短网址链接
	// http.HandleFunc("/url", shortUrl)
	//启动服务监听地址
	http.ListenAndServe(address, nil)
}

//页面访问求请求首页
func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "请求成功!")
	} else {
		http.NotFound(w, r)
	}
}

//按照特定条件获取一个许可证
func licence(rw http.ResponseWriter, r *http.Request) {
	//appkey appsecret backsecret

}

//用户生成UUID
func genUuid(rw http.ResponseWriter, r *http.Request) {
	uid := uuid.NewV4()
	var sessions = map[string]string{}
	sessions["uuid"] = uid.String()
	//创建一个返回结果类型
	rsp := NewR("请求成功", 200, D{"uuid": uid.String()})
	//直接将结果返回
	fmt.Fprintln(rw, rsp.JSON())
}
