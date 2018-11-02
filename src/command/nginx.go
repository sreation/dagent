package command

import (
    "common"
    "os/exec"
    "bytes"
)

func Cmd(cmdStr string) string {
    cmd := exec.Command("sh", "-c", cmdStr)

    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err := cmd.Run()

    if err != nil {
        common.Error.Println("command:"+cmdStr+", fail. output:" + stderr.String())
        return ""
    }
    return string(out.String())
}

/*
nginx commands
*/

func CmdNginxStart() string {
    return Cmd(common.CfgXML.Nginx.CmdStart)
}

func CmdNginxStop() string {
    return Cmd(common.CfgXML.Nginx.CmdStop)
}

func CmdNginxReload() string {
    return Cmd(common.CfgXML.Nginx.CmdReload)
}

func CmdNginxTest() string {
    return Cmd(common.CfgXML.Nginx.CmdTest)
}

func ShowNginxConf() (string, bool){
    return common.ShowFile(common.CfgXML.Nginx.ConfdDir)
}

func ReadNginxConf(confName string) (string, bool){
    return common.ReadFile(common.CfgXML.Nginx.ConfdDir + "/" + confName)
}