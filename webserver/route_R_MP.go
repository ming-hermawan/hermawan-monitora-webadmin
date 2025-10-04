package webserver

import (
    "strconv"
    "net/http"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/global/menu"
    "hermawan-monitora/webserver/controller/main/auth/login"
    "hermawan-monitora/webserver/controller/main/report/monitoring/port"
    "hermawan-monitora/webserver/controller/main/report/monitoring/port/report"
)


func reportMonPortProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    reportMonPortRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.ReportMonPort))
    reportMonPort, _ := strconv.Atoi(reportMonPortRaw)
    if reportMonPort != 1 {
        login.Get(w)
        return
    }
    port.Process(w, r, username)
}

func reportMonPortReportProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    reportMonPortRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.ReportMonPort))
    reportMonPort, _ := strconv.Atoi(reportMonPortRaw)
    if reportMonPort != 1 {
        login.Get(w)
        return
    }
    report.Process(w, r, username)
}
