package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"reflect"
	"runtime"

	"github.com/go-xorm/xorm"

	_ "github.com/mattn/go-sqlite3"
)

const (
	//图片文件保存文件夹
	file_save_path = "./img/"
)

/********************[数据表]*******************************/
type ImgTable struct {
	Id       int64
	filename string `xorm:"varchar(25) notnull unique "`
	name     string `xorm:"varchar(25) notnull unique "`
	imgpath  string `xorm:"varchar(25)  "`
}

var orm *xorm.Engine

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handle Function called -" + name)
		h(w, r)
	}
}
func addimg(w http.ResponseWriter, r *http.Request) {
	//输出对应的 请求方式
	fmt.Println(r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("view/addimg.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		// 获取信息
		file, file_head, file_err := r.FormFile("filename")
		r.ParseMultipartForm(32 << 20)
		name := r.MultipartForm.Value["name"]
		if file_err != nil {
			fmt.Fprintf(w, "file upload fail:%s", file_err)
		}
		file_save := file_save_path + file_head.Filename
		//打开 已只读,文件不存在创建 方式打开  要存放的路径资源
		f, f_err := os.OpenFile(file_save, os.O_WRONLY|os.O_CREATE, 0666)
		if f_err != nil {
			fmt.Fprintf(w, "file open fail:%s", f_err)
		}
		//文件 copy
		_, copy_err := io.Copy(f, file)
		if copy_err != nil {
			fmt.Fprintf(w, "file copy fail:%s", copy_err)
		}
		//关闭对应打开的文件
		defer f.Close()
		defer file.Close()

		fmt.Fprintf(w, file_save)

		// TODO: 后续需要吧密码做加密处理
		img := ImgTable{filename: file_head.Filename, name: name[0], imgpath: file_save}
		_, err := orm.Insert(img)
		if err != nil {
			fmt.Fprintf(w, "存库失败", copy_err)
		}
	} else {
		w.WriteHeader(500)
		fmt.Fprintln(w, "不支持这种调用方式!")
	}
}
func showimg(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/showimg.html")
	t.Execute(w, nil)
}
func main() {
	//创建数据库
	var err error
	orm, err = xorm.NewEngine("sqlite3", "./blizzard.db")
	if err != nil {
		fmt.Println("orm failed to initialized: %v", err)
	}
	err = orm.Sync2(new(ImgTable))

	server := http.Server{
		Addr: "127.0.0.1:6789",
	}

	http.Handle("/admin/addimg", log(addimg))
	http.Handle("/admin/showimg", log(showimg))

	server.ListenAndServe()

}
