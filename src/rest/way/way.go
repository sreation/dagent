package way

import (
    "net/http"
    "common"
    "command"
)

func NginxStart(w http.ResponseWriter, r *http.Request) {
    result := command.CmdNginxStart()
    if result != "" {
        common.IOWrite(w, http.StatusOK, 0, "service: nginx , nginx is started", result)
    }else {
        common.IOWrite(w, http.StatusOK, 1, "service: nginx , nginx start failed", result)
    }
}

func NginxStop(w http.ResponseWriter, r *http.Request) {
    result := command.CmdNginxStop()
    if result != "" {
        common.IOWrite(w, http.StatusOK, 0, "service: nginx , nginx is stopped", result)
    }else {
        common.IOWrite(w, http.StatusOK, 1, "service: nginx , nginx stop failed", result)
    }

}

func NginxTest(w http.ResponseWriter, r *http.Request) {
    result := command.CmdNginxTest()
    if result != "" {
        common.IOWrite(w, http.StatusOK, 0, "service: nginx , nginx is test ok", result)
    }else {
        common.IOWrite(w, http.StatusOK, 1, "service: nginx , nginx test failed", result)
    }
}

func NginxCheck(w http.ResponseWriter, r *http.Request) {
    result := command.CmdNginxCheck()
    if result != "" {
        common.IOWrite(w, http.StatusOK, 0, "service: nginx , nginx is check ok", result)
    }else {
        common.IOWrite(w, http.StatusOK, 1, "service: nginx , nginx check failed", result)
    }
}


func NginxReload(w http.ResponseWriter, r *http.Request) {
    result := command.CmdNginxReload()
    if result != "" {
        common.IOWrite(w, http.StatusOK, 0, "service: nginx , nginx is reloaded", result)
    }else {
        common.IOWrite(w, http.StatusOK, 1, "service: nginx , nginx reload failed", result)
    }
}

func NginxShowConf(w http.ResponseWriter, r *http.Request) {
    bo, msg, result := command.ShowNginxConf()
    if bo != false {
        common.IOWrite(w, http.StatusOK, 0, "succ", result)
    }else {
        common.IOWrite(w, http.StatusOK, 1, msg, result)
    }
}

func NginxReadConf(w http.ResponseWriter, r *http.Request) {
    confname := r.PostFormValue("confname")
    if confname == "" {
        common.IOWrite(w, http.StatusOK, 1, "params is error", "")
        return
    }
    bo, msg, result := command.ReadNginxConf(confname)
    if bo != false {
        common.IOWrite(w, http.StatusOK, 0, "succ", result)
    }else {
        common.IOWrite(w, http.StatusOK, 1, msg, result)
    }
}

func NginxWriteConf(w http.ResponseWriter, r *http.Request) {
    confname := r.PostFormValue("confname")
    content := r.PostFormValue("content")
    if (confname == "" || content == "") {
        common.IOWrite(w, http.StatusOK, 1, "params is error", "")
        return
    }

    bo, msg := command.WriteNginxConf(confname, content)
    if bo != false {
        common.IOWrite(w, http.StatusOK, 0, "succ", "")
    }else {
        common.IOWrite(w, http.StatusOK, 1, msg, "")
    }
}