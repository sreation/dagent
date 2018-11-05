package main
import (
    //"io"
    "net/http"
    //"strings"
    "time"
    "rest"
    "common"
    //"fmt"
    //"io"
    "strings"
)


type myServer struct {}

var (
    server = &http.Server{
        Addr:           ":9949",
        Handler:        &myServer{},
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        //MaxHeaderBytes: 1 << 20,
    }
)

func (*myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    remoteaddr := strings.Split(r.RemoteAddr, ":")[0]
    if (common.CfgXML.Server.RemoteHost != "0.0.0.0") && (common.CfgXML.Server.RemoteHost != remoteaddr) {
        common.IOWrite(w, http.StatusUnauthorized, 1, "remoteaddr not allow", "")
        return
    }
    myUrl := r.URL.String()
    var h, ok = rest.RouterMap[myUrl]
    if ok {
        if r.Method != h.Method {
            common.IOWrite(w, http.StatusMethodNotAllowed, 1, "method not allow", "")
            //w.WriteHeader(http.StatusMethodNotAllowed)
            //io.WriteString(w, common.JsonEncode(1, "method not allow"))
            return
        }
        h.Run(w, r)
        if h.IsLog == true {
            common.Info.Println(remoteaddr + " url: " + myUrl)
        }
    }
    //if h, ok := rest.RouterMap[r.URL.String()]; ok {
    //    fmt.Println(h)
    //    fmt.Println(ok)
    //    h(w, r)
    //}
    //io.WriteString(w, "URL"+r.URL.String())
}

func main() {
    //init configxml
    common.CommonInit()
    //init log
    common.LogInit()
    common.Info.Println("******agent is started******")
    server.ListenAndServe()
}