package way

import (
    "net/http"
    "common"
    "command"
)

func NginxStart(w http.ResponseWriter, r *http.Request) {
    result := command.CmdNginxStart()
    if result != "" {
        common.IOWrite(w, http.StatusOK, 0, "service: nginx , nginx is started", "")
    }else {
        common.IOWrite(w, http.StatusOK, 1, "service: nginx , nginx start failed", "")
    }
}

func NginxStop(w http.ResponseWriter, r *http.Request) {
    result := command.CmdNginxStop()
    if result != "" {
        common.IOWrite(w, http.StatusOK, 0, "service: nginx , nginx is stopped", "")
    }else {
        common.IOWrite(w, http.StatusOK, 1, "service: nginx , nginx stop failed", "")
    }

}

func NginxTest(w http.ResponseWriter, r *http.Request) {
    result := command.CmdNginxTest()
    if result != "" {
        common.IOWrite(w, http.StatusOK, 0, "service: nginx , nginx is test ok", "")
    }else {
        common.IOWrite(w, http.StatusOK, 1, "service: nginx , nginx test failed", "")
    }
}

func NginxReload(w http.ResponseWriter, r *http.Request) {
    result := command.CmdNginxReload()
    if result != "" {
        common.IOWrite(w, http.StatusOK, 0, "service: nginx , nginx is reloaded", "")
    }else {
        common.IOWrite(w, http.StatusOK, 1, "service: nginx , nginx reload failed", "")
    }
}

func NginxShowConf(w http.ResponseWriter, r *http.Request) {
    result,bo := command.ShowNginxConf()
    if bo != false {
        common.IOWrite(w, http.StatusOK, 0, "", result)
    }else {
        common.IOWrite(w, http.StatusOK, 1, "command: NginxShowConf failed", "")
    }
}

func NginxReadConf(w http.ResponseWriter, r *http.Request) {
    confname := r.PostFormValue("confname")
    result, bo := command.ReadNginxConf(confname)
    if bo != false {
        common.IOWrite(w, http.StatusOK, 0, "", result)
    }else {
        common.IOWrite(w, http.StatusOK, 1, "command: NginxReadConf failed", "")
    }
}