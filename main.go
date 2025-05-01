package main

import (
	"fmt"
	"log"
	"net/http"
)



func AuthMidlleware(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie , err := r.Cookie("username")
		if err != nil || cookie.Value == "" {
 			http.Redirect(w, r,"/login",http.StatusFound)
		}
		next(w,r)
	}
}

func loginhandler(w http.ResponseWriter,r *http.Request)  {
	username := r.URL.Query().Get("user")
	if username == "" {
		fmt.Fprintln(w,"silahkan Masukkan ?user=namamu")
		return
	}
	cookie := http.Cookie{
		Name: "username",
		Value: username,
		Path: "/",
	  MaxAge: 3600,	
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r,"/welcome",http.StatusFound)
}

func welcomeHandler(w http.ResponseWriter,r *http.Request)  {
	cookie,_ := r.Cookie("username")
	fmt.Fprintf(w,"selamat datanmg, %s!",cookie.Value)
}



func logouthandler(w http.ResponseWriter,r *http.Request)  {
	cookie := http.Cookie{
		Name: "username",
    Value: "",
		Path: "/",
		MaxAge: -1, //hapus cookie
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "anda sudah logout")
}




func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginhandler)
  mux.HandleFunc("/welcome",welcomeHandler)
	mux.HandleFunc("/logout", logouthandler)
	log.Println("server jalan di https://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
