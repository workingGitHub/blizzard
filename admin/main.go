package main

import  {
	"fmt"
	"net/http"
}

func showimg(w http.ResponseWriter, r *http.Request){
	fm.Fprintf(w,"showimg")
}

func addimg(w http.ResponseWriter, r *http.Request){
	fm.Fprintf(w,"addimg")
}

func main(){
	service := http.Server{
		Addr: "127.0.0.1:6789"
	}	
	
	http.Handle("/admin/addimg",addimg)
	http.Handle("/admin/showimg",showimg)
}