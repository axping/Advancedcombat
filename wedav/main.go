package main

//sesstions:登入信息session表格
//session_id
//login_name

func main() {
	r := WServe()
	//设置静态文件目录
	r.Static("/static/", "./static/*")
	r.Run(":4200")
}
