package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

///Users/chenping/go/src/Advancedcombat/videopost
var dbSessions = map[string]string{} // session ID, user ID

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./resources/*.html"))
}

func main() {
	fmt.Println("启动服务")
	//资源文件目录
	///Users/chenping/go/src/Advancedcombat/videopost/resources
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./resources"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/post", vPost)
	http.HandleFunc("/uuid", genUUID)
	http.HandleFunc("/qrcode", qrCode)
	http.ListenAndServe(":8080", nil)
	defer fmt.Println("服务结束")
}

func index(rw http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		ck = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(rw, ck)
		//fmt.Fprintln(rw, "index.html")
	}
	tpl.ExecuteTemplate(rw, "index.html", nil)
	// fmt.Println("请求Session", ck)
	// fmt.Fprintf(rw, fmt.Sprintf("{\"message\":\"请求成功\",\"code\":200,\"session\":\"%s\"}", ck.Value))
}

//请求合成参数
func vPost(rw http.ResponseWriter, r *http.Request) {
	//解析请求参数
	r.ParseForm()
	//获取请求参数
	http.NotFound(rw, r)
}

//用于生成uuid
func genUUID(rw http.ResponseWriter, r *http.Request) {
	uid := uuid.NewV4()
	var sessions = map[string]string{}
	sessions["uuid"] = uid.String()
	rsp := &Respond{
		Message: "请求成功",
		Code:    200,
		Data:    sessions,
	}
	rjson, err := json.Marshal(rsp)
	if err != nil {
		fmt.Println("JOSN转换异常")
		fmt.Fprintln(rw, "uuid 生成失败")
		return
	}
	fmt.Fprintln(rw, string(rjson))
}

//生成二维码图片
func qrCode(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}
