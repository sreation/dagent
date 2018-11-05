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

func CmdNginxCheck() string {
    return Cmd(common.CfgXML.Nginx.CmdCheck)
}

func ShowNginxConf() (bool, string, string) {
    return common.ShowFile(common.CfgXML.Nginx.ConfdDir)
}

func ReadNginxConf(confName string) (bool, string, string) {
    return common.ReadFile(common.CfgXML.Nginx.ConfdDir + "/" + confName)
}

func WriteNginxConf(confName string, content string) (bool, string) {
    return common.WriteFile(common.CfgXML.Nginx.ConfdDir + "/" + confName, content)
}