package webserver

import (
    "strconv"
    "net/http"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/global/menu"
    "hermawan-monitora/webserver/controller/main/auth/login"
    "hermawan-monitora/webserver/controller/main/monitoring/port"
)


func portsMonitoringProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    monPortsRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.MonPort))
    monPorts, _ := strconv.Atoi(monPortsRaw)
    if monPorts != 1 {
        login.Get(w)
        return
    }
    port.GetPortsMonitoringPage(username, w, r)
}
