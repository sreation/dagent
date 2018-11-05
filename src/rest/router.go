package rest

import (
    "net/http"
    "rest/way"
)

type wayType func(http.ResponseWriter, *http.Request)

type RouterStruct struct{
    Method string
    IsLog bool
    Run wayType
}

//var RouterMap = make(map[string]wayType)
var RouterMap = make(map[string]RouterStruct)

func init() {
    //RouterMap["/nginx/start"] = way.NginxStart
    RouterMap["/nginx/start"] = RouterStruct{Method: "GET", IsLog:true, Run: way.NginxStart}
    RouterMap["/nginx/stop"] = RouterStruct{Method: "GET", IsLog:true, Run: way.NginxStop}
    RouterMap["/nginx/test"] = RouterStruct{Method: "GET", IsLog:false, Run: way.NginxTest}
    RouterMap["/nginx/reload"] = RouterStruct{Method: "GET", IsLog:true, Run: way.NginxReload}
    RouterMap["/nginx/check"] = RouterStruct{Method: "GET", IsLog:false, Run: way.NginxCheck}
    RouterMap["/nginx/confd"] = RouterStruct{Method: "GET", IsLog:false, Run: way.NginxShowConf}
    RouterMap["/nginx/confdread"] = RouterStruct{Method: "POST", IsLog:false, Run: way.NginxReadConf}
    RouterMap["/nginx/confdfile"] = RouterStruct{Method: "POST", IsLog:true, Run: way.NginxWriteConf}
}
