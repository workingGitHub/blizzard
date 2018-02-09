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
	file_save_path = "./static/img/"
)

/********************[数据表]*******************************/
type ImgTable struct {
	Id       int64
	Filename string `xorm:"varchar(100)  notnull unique "`
	Name     string `xorm:"varchar(40)  notnull unique "`
}

var g_orm *xorm.Engine

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
		t, err := template.ParseFiles("view/addimg.html")
		if err != nil {
			fmt.Print("template.ParseFiles error:", err, "\n")
		}

		err = t.Execute(w, "")
		if err != nil {
			fmt.Print("t.Execute error:", err, "\n")
		}
	} else if r.Method == "POST" {
		// 获取信息
		file, file_head, file_err := r.FormFile("uploaded")
		if file_err != nil {
			fmt.Fprintln(w, "file upload fail:%s", file_err)
		}
		file_save := file_save_path + file_head.Filename

		//打开 已只读,文件不存在创建 方式打开  要存放的路径资源
		f, f_err := os.OpenFile(file_save, os.O_WRONLY|os.O_CREATE, 0666)
		if f_err != nil {
			fmt.Fprintln(w, "file open fail:%s", f_err)
		}
		//文件 copy
		_, copy_err := io.Copy(f, file)
		if copy_err != nil {
			fmt.Fprintln(w, "file copy fail:%s", copy_err)
		}
		//关闭对应打开的文件
		defer f.Close()
		defer file.Close()

		fmt.Fprintln(w, file_save)
		fmt.Println("test2", file_head.Filename, file_save, file_err)
		r.ParseMultipartForm(1024)
		name := r.MultipartForm.Value["name"][0]
		//	数据库存库
		img := ImgTable{Filename: file_head.Filename, Name: name}
		fmt.Println("test3", img)
		i, err := g_orm.Insert(img)
		fmt.Println("test4", i, err)
		if err != nil {
			fmt.Fprintln(w, "存库失败", err)
		} else {
			fmt.Fprintln(w, "存库成功", img)
		}

	} else {
		w.WriteHeader(500)
		fmt.Fprintln(w, "不支持这种调用方式!")
	}
}
func showimg(w http.ResponseWriter, r *http.Request) {
	var imgs []ImgTable
	err := g_orm.Find(&imgs)
	if err != nil {
		fmt.Fprintln(w, "数据库查询失败!")
	} else {
		t, _ := template.ParseFiles("view/showimg.html")
		fmt.Println("test2", imgs)
		err = t.Execute(w, imgs)

		if err != nil {
			fmt.Fprintln(w, "模板查询失败!")
		}

	}

}
func main() {
	//创建数据库
	orm, err := xorm.NewEngine("sqlite3", "./blizzard.db")
	if err != nil {
		fmt.Println("orm failed to initialized: %v", err)
	}
	err = orm.Sync2(new(ImgTable))
	if err != nil {
		fmt.Println("orm Sync2 failed: %v", err)
	}
	g_orm = orm

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	server := http.Server{
		Addr: "127.0.0.1:6789",
	}

	http.Handle("/admin/addimg", log(addimg))
	http.Handle("/admin/showimg", log(showimg))

	server.ListenAndServe()

}
