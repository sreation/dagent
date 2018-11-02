package common

import (
    "log"
    "os"
    "fmt"
    "path"
)

var (
    Warning *log.Logger
    Info    *log.Logger
    Error   *log.Logger
)

func LogInit() {
    logPath := CfgXML.Server.LogPath
    logDir := path.Dir(logPath)
    _, err := os.Stat(logDir)

    if err != nil {
        if os.IsNotExist(err) {
            err2 := os.MkdirAll(logDir, os.ModePerm)
            if err2 != nil {
                fmt.Println(err2.Error())
                os.Exit(1)
            }
        }else {
            fmt.Println(err.Error())
            os.Exit(1)
        }

    }

    file, err := os.OpenFile(logPath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, os.ModePerm)
    if err!= nil && os.IsNotExist(err) {
        file2, err2 := os.Create(logPath)
        defer file.Close()
        if err2 != nil {
            fmt.Println(err2.Error())
            os.Exit(1)
        }
        file = file2
    }

    Info = log.New(file, "[Info]", log.LstdFlags)
    Warning = log.New(file, "[Warning]", log.LstdFlags)
    Error = log.New(file, "[Error]", log.LstdFlags)
}
