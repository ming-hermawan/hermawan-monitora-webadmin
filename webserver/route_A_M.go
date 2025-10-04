package webserver

import (
    "strconv"
    "net/http"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/global/menu"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/controller/main/auth/login"
    "hermawan-monitora/webserver/controller/main/admin/monitoring/port/servergroup"
    "hermawan-monitora/webserver/controller/main/admin/monitoring/port/servernports"
    "hermawan-monitora/webserver/controller/main/admin/monitoring/port/uploadcsv"
    "hermawan-monitora/webserver/controller/main/admin/monitoring/port/email"
    ddlServerGroup "hermawan-monitora/webserver/controller/ddl/admin/monitoring/port/servergroup"
)


func adminServerGroupProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    admMonPortServerGrpRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmMonPortServerGrp))
    admMonPortServerGrp, _ := strconv.Atoi(admMonPortServerGrpRaw)
    if admMonPortServerGrp != 1 {
        login.Get(w)
        return
    }
    servergroup.Process(username, w, r)
}

func adminServerNPortsProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    admMonPortServerNPortsRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmMonPortServerNPorts))
    admMonPortServerNPorts, _ := strconv.Atoi(admMonPortServerNPortsRaw)
    if admMonPortServerNPorts != 1 {
        login.Get(w)
        return
    }
    servernports.Process(username, w, r)
}

func uploadCSVProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    admMonPortUploadCsvRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmMonPortUploadCsv))
    admMonPortUploadCsv, _ := strconv.Atoi(admMonPortUploadCsvRaw)
    if admMonPortUploadCsv != 1 {
        login.Get(w)
        return
    }
    uploadcsv.Process(username, w, r)
}

func admMonEmailProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    admMonPortAlertEmailRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmMonPortAlertEmail))
    admMonPortAlertEmail, _ := strconv.Atoi(admMonPortAlertEmailRaw)
    if admMonPortAlertEmail != 1 {
        login.Get(w)
        return
    }
    email.Process(username, w, r)
}

func ddlServerGroupProcess(w http.ResponseWriter, r *http.Request) {
    _, err := authProcess(r)
    if err != nil {
        httpresponse.ErrResponseForInvalidToken(w, err.Error())
        return
    }
    ddlServerGroup.PostProcess(w, r)
}
