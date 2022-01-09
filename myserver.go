package main
import
(
    "fmt"
    "net/http"
)
//参考https://www.topgoer.com/%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B/http%E7%BC%96%E7%A8%8B.html
func myHandle(w http.ResponseWriter, r *http.Request){
    //w: 写给客户端的数据内容
    w.Write([]byte("This is a http web server."))

    //r：从客户端读到的内容
    fmt.Println("Header:",r.Header)
    fmt.Println("URL:",r.URL)
    fmt.Println("Method:",r.Method)
    fmt.Println("Host:",r.Host)
    fmt.Println("RemoteAddr:",r.RemoteAddr)
    fmt.Println("Body:",r.Body)
    //fmt.Println("VERSION:",VERSION)

    //将request中的header写入response header
    //将当前系统环境变量VERSION写入response header
    //w.Header:=r.Header
    fmt.Println(w.Header)
    //sw.Header().Set(r.Header)
    //ssw.Header().Set("VERSION",VERSION)
    //w.WriteHeader(500)


}

func main () {
    fmt.Println("Hello wentingada!   \n\n         - from Golang .")
    //注册回调函数，该函数在客户端访问服务器时，自动被调用
    http.HandleFunc("/", myHandle)

    //绑定服务器监听地址
    http.ListenAndServe("127.0.0.1:8080", nil)
}
