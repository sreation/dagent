package common

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "encoding/xml"
    "os"
    "strings"
)

const XMLPath = "/etc/dagent/config.xml"
//const XMLPath = "/Users/dingdan/dinggo/go_event/conf/config.xml"

var CfgXML = ConfigStruct{}

type IOStruct struct {
    Code int        `json:"code"`
    Message string  `json:"message"`
    Data string     `json:"data"`
}

func JsonEncodeIO (c int, m string, da string) (data string) {
    /*
    int 0 - succ 1 - fail
    */
    d, errs := json.Marshal(IOStruct{Code: c, Message: m, Data: da})
    if errs != nil {
        fmt.Println(errs.Error())
	}else {
        data = string(d)
    }
    return
}

func IOWrite (w http.ResponseWriter, httpCode int, code int, message string, data string) {
    w.WriteHeader(httpCode)
    io.WriteString(w, JsonEncodeIO(code, message, data))
}

//xml parse
type ConfigStruct struct {
    XMLName xml.Name    `xml:"config"`
    Server ServerStruct `xml:"server"`
    Nginx NginxStruct   `xml:"nginx"`
}

type ServerStruct struct {
	Auth string       `xml:"auth"`
    LogPath string    `xml:"logPath"`
    RemoteHost string `xml:"remoteHost"`
}

type NginxStruct struct {
    ConfdDir string `xml:"confdDir"`
    CmdStart string `xml:"cmdStart"`
    CmdStop string `xml:"cmdStop"`
    CmdReload string `xml:"cmdReload"`
    CmdTest string `xml:"cmdTest"`
    CmdCheck string `xml:"cmdCheck"`
}



func ConfigParse() (result ConfigStruct, b bool){
    content, err := ioutil.ReadFile(XMLPath)
    result = ConfigStruct{}
    b = true

    if err != nil {
        fmt.Println(err.Error())
        b = false
    }

	err = xml.Unmarshal(content, &result)
	if err != nil {
		fmt.Println(err.Error())
        b = false
	}
    result.Nginx.ConfdDir = strings.TrimRight(result.Nginx.ConfdDir, "/")
    return result, b
}

func JsonDecode(s string) (bool, string) {
    f := ""
    err := json.Unmarshal([]byte(s), &f)

    if err != nil {
        return false, err.Error()
    } else {
        return true, f
    }
}

func CommonInit() {
    r, b:= ConfigParse()
    if b == true {
        CfgXML = r
    }else {
        os.Exit(1)
    }
}
