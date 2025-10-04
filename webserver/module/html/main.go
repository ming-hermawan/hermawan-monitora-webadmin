package html

import (
    "strconv"
    "html/template"
    "net/http"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/global"
    "hermawan-monitora/webserver/global/menu"
    "hermawan-monitora/webserver/module/httpresponse"
)
func GetPage(name string,
             w http.ResponseWriter,
             htmlFilepath string) {
    tmpl, err := template.ParseFiles(htmlFilepath)
    if err != nil {
        httpresponse.ErrResponseWs(w, "template.ParseFiles Error", err.Error())
    }
    data := global.IndexData{
      Name: name}
    err = tmpl.Execute(w, data)
    if err != nil {
        httpresponse.ErrResponseWs(w, "ExecuteTemplate Error", err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func GetTmpl0(name string,
              w http.ResponseWriter,
              t ...string) {
    tmpl, err := template.ParseFiles(t...)
    if err != nil {
        httpresponse.ErrResponseWs(w, "template.ParseFiles Error", err.Error())
    }
    data := global.IndexData{
      Name: name}
    err = tmpl.ExecuteTemplate(w, "base0", data)
    if err != nil {
        httpresponse.ErrResponseWs(w, "ExecuteTemplate Error", err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func GetTmpl1(username string,
              name string,
              w http.ResponseWriter, t ...string) {
    admUsrMgtUsrGrpRaw, err := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmUsrMgtUsrGrp))
    var admUsrMgtUsrGrp int
    if(err == nil) {
        admUsrMgtUsrGrp, _ = strconv.Atoi(admUsrMgtUsrGrpRaw)
    } else {
        admUsrMgtUsrGrp = 0
    }
    admUsrMgtUsrRaw, err := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmUsrMgtUsr))
    var admUsrMgtUsr int
    if(err == nil) {
        admUsrMgtUsr, _ = strconv.Atoi(admUsrMgtUsrRaw)
    } else {
        admUsrMgtUsr = 0
    }
    admMonPortSettingsRaw, err := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmUsrMgtUsr))
    var admMonPortSettings int
    if(err == nil) {
        admMonPortSettings, _ = strconv.Atoi(admMonPortSettingsRaw)
    } else {
        admMonPortSettings = 0
    }
    admMonPortServerGrpRaw, err := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmMonPortServerGrp))
    var admMonPortServerGrp int
    if(err == nil) {
        admMonPortServerGrp, _ = strconv.Atoi(admMonPortServerGrpRaw)
    } else {
        admMonPortServerGrp = 0
    }
    admMonPortServerNPortsRaw, err := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmMonPortServerNPorts))
    var admMonPortServerNPorts int
    if(err == nil) {
        admMonPortServerNPorts, _ = strconv.Atoi(admMonPortServerNPortsRaw)
    } else {
        admMonPortServerNPorts = 0
    }
    admMonPortAlertEmailRaw, err := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmMonPortAlertEmail))
    var admMonPortAlertEmail int
    if(err == nil) {
        admMonPortAlertEmail, _ = strconv.Atoi(admMonPortAlertEmailRaw)
    } else {
        admMonPortAlertEmail = 0
    }
    admMonPortUploadCsvRaw, err := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmMonPortUploadCsv))
    var admMonPortUploadCsv int
    if(err == nil) {
        admMonPortUploadCsv, _ = strconv.Atoi(admMonPortUploadCsvRaw)
    } else {
        admMonPortUploadCsv = 0
    }
    monPortsRaw, err := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.MonPort))
    var monPorts int
    if(err == nil) {
        monPorts, _ = strconv.Atoi(monPortsRaw)
    } else {
        monPorts = 0
    }
    reportMonPortRaw, err := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.ReportMonPort))
    var reportMonPort int
    if(err == nil) {
        reportMonPort, _ = strconv.Atoi(reportMonPortRaw)
    } else {
        reportMonPort = 0
    }
    tmpl, err := template.ParseFiles(t...)
    if err != nil {
        httpresponse.ErrResponseWs(w, "template.ParseFiles Error", err.Error())
    }
    data := global.IndexData{
      Name: name,
      AdmUsrMgtUsrGrp: admUsrMgtUsrGrp,
      AdmUsrMgtUsr: admUsrMgtUsr,
      AdmMonPortSettings: admMonPortSettings,
      AdmMonPortServerGrp: admMonPortServerGrp,
      AdmMonPortServerNPorts: admMonPortServerNPorts,
      AdmMonPortUploadCsv: admMonPortUploadCsv,
      AdmMonPortAlertEmail: admMonPortAlertEmail,
      MonPort: monPorts,
      ReportMonPort: reportMonPort}
    err = tmpl.ExecuteTemplate(w, "base0", data)
    if err != nil {
        httpresponse.ErrResponseWs(w, "ExecuteTemplate Error", err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
