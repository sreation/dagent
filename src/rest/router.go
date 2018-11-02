package rest

import (
    "net/http"
    "rest/way"
)

type wayType func(http.ResponseWriter, *http.Request)

type RouterStruct struct{
    Method string
    Run wayType
}

//var RouterMap = make(map[string]wayType)
var RouterMap = make(map[string]RouterStruct)

func init() {
    //RouterMap["/nginx/start"] = way.NginxStart
    RouterMap["/nginx/start"] = RouterStruct{Method: "GET", Run: way.NginxStart}
    RouterMap["/nginx/stop"] = RouterStruct{Method: "GET", Run: way.NginxStop}
    RouterMap["/nginx/test"] = RouterStruct{Method: "GET", Run: way.NginxTest}
    RouterMap["/nginx/reload"] = RouterStruct{Method: "GET", Run: way.NginxReload}
    RouterMap["/nginx/confd"] = RouterStruct{Method: "GET", Run: way.NginxShowConf}
    RouterMap["/nginx/confdfile"] = RouterStruct{Method: "POST", Run: way.NginxReadConf}
}
