package servernports

import (
    "fmt"
    "log"
    "strconv"
    "net/http"
    "gorm.io/gorm"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/webserver/decorator"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/initredis"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/validation"
)


const Lbl = "Server & Ports"


// TYPE

type TemplateDetailPortsMonitoring struct {
    Ip string
    Name string
    ServerGroup string
    MPServerGroup []string
    Ports map[int]string
}


//dipakai di temp
func getServerNPorts(db *gorm.DB,  w http.ResponseWriter, ip string) TemplateDetailPortsMonitoring {
    temp := dbo.GetServerPorts0(db, ip)
    serverGroups, err := dbo.GetMPServerGroupsForDDL(db)
    if err != nil {
        httpresponse.ErrResponseWhenSelDb(
          w,
          "mp_server_groups",
          err.Error())
    }
    log.Println(fmt.Println(temp))
    ports := map[int]string {}
    for _, row := range temp.Ports {
        ports[row.Port] = row.Name
    }
    return TemplateDetailPortsMonitoring {
      Ip: temp.Server.Ip,
      Name: temp.Server.Name,
      ServerGroup: temp.Server.ServerGroup,
      MPServerGroup: serverGroups,
      Ports: ports}
}

func getEmails(db *gorm.DB,  w http.ResponseWriter, ip string) []string {
    serverAlertEmails, err := dbo.GetMPServerAlertEmail(db, ip)
    if err != nil {
        httpresponse.ErrResponseWhenSelDb(
          w,
          "mp_server_alert_emails",
          err.Error())
    }
    return serverAlertEmails
}

func getListForDDL(db *gorm.DB, w http.ResponseWriter) ([]string, error) {
    return dbo.GetMPServerGroupsForDDL(db)
}

func get(username string, w http.ResponseWriter) {
    html.GetTmpl1(
      username,
      Lbl,
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.Base1HtmlFilepath,
      hmonglobal.InpFrm2ColBaseHtmlFilepath,
      hmonglobal.InpFrmFilterGrpHtmlFilepath,
      hmonglobal.InpFrm2ColWithTabPagesHtmlFilepath,
      hmonglobal.AdmMonPortServerNPortsHtmlFilepath,
    )
}

func getRowsCnt(db *gorm.DB,
                grpFilter string,
                txtFilter string) (int64, error) {
    return dbo.GetCntOfMPServer(db, grpFilter, txtFilter)
}

func getRowsForTbl(db *gorm.DB,
                   grpFilter string,
                   txtFilter string,
                   page int,
                   pageLimit int) ([][]string, error) {
    return dbo.GetMPServerForTbl(db, grpFilter, txtFilter, page, pageLimit)
}

func getRow(db *gorm.DB,
            w http.ResponseWriter,
            ip string) map[string]any {
    temp := getServerNPorts(db, w, ip)
    emails := getEmails(db, w, ip)
    return map[string]interface{} {
      "ip": temp.Ip,
      "name": temp.Name,
      "servergroup": temp.ServerGroup,
      "servergroups": temp.MPServerGroup,
      "ports": temp.Ports,
      "emails": emails}
}

func post(w http.ResponseWriter, r *http.Request) {
    decorator.HttpPostForMasterDataWithTxtAndGrpFilter(
      w,
      r,
      Lbl,
      "key",
      []string{"IP", "Server", "Server-Group"},
      getRowsCnt,
      getRowsForTbl,
      getRow,
      getListForDDL)
}

func put(w http.ResponseWriter, r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    log.Println(fmt.Sprintf("Server&Ports PUT params:\n%+v\n", params))
    statDb, statDbValid := validation.StatDbParamValidation(
      w,
      params)
    if !statDbValid {
        return
    }
    ip, ipValid := validation.StrParamValidation(
      w,
      params,
      "ip",
      "IP",
      true,
      hmonglobal.RegexIp)
    if !ipValid {
        return
    }
    name, nameValid := validation.StrParamValidation(
      w,
      params,
      "name",
      "Name",
      false,
      hmonglobal.RegexId)
    if !nameValid {
        return
    }
    group, groupValid := validation.StrParamValidation(
      w,
      params,
      "group",
      "Server-Group",
      true,
      hmonglobal.RegexId)
    if !groupValid {
        return
    }
    log.Println("PORTS:\n")
    ports := map[int]string{};
    for key, val := range params["ports"].(map[string]interface{}) {
        port, err := strconv.Atoi(key)
        if err != nil {
            httpresponse.ErrResponseForHttpBody(
              w,
              "Port is not Integer")
             return
        }
        log.Println(fmt.Sprintf("port=%d value=%s type=%T\n", port, val, val))
        ports[port] = fmt.Sprintf("%s", val)
    }
    log.Println("EMAILS:\n")
    emails := []string{};
    for _, val := range params["emails"].([]interface{}) {
        emails = append(emails, val.(string))
    }
    var err error
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    if statDb == "INS" {
        err = dbo.InsServer(db, ip, name, group)
        if err != nil {
            httpresponse.ErrResponseWhenInsDb(
              w,
              "mp_servers",
              err.Error())
            return
        }
    } else if statDb == "UPD" {
        err = dbo.UpdServer(db, ip, name, group)
        if err != nil {
            httpresponse.ErrResponseWhenUpdDb(
              w,
              "mp_servers",
              err.Error())
            return
        }
    }
    dbo.UpdMPPorts(db, ip, ports)
    dbo.UpdMPServerAlertEmail(db, ip, emails)
    var success bool
    success = initredis.Init(w)
    if !success {
        return
    }
    success = initredis.SendInitSignal(w)
    if !success {
        return
    }
    httpresponse.JsonResponseForSuccessOperation(w)
}

func del(w http.ResponseWriter, r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    if params == nil {
        return
    }
    var ips []string;
    for _, x:= range params["keys"].([]interface{}) {
        ips = append(ips, x.(string))
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    var err error
    tx := db.Begin()
    err = dbo.DelMPServer(
      tx,
      ips)
    if err != nil {
        tx.Rollback()
        httpresponse.ErrResponseWhenDelDb(
          w,
          "servers",
          err.Error())
        return
    }
    err = dbo.DelMPPorts(
      tx,
      ips)
    if err != nil {
        tx.Rollback()
        httpresponse.ErrResponseWhenDelDb(
          w,
          "ports",
          err.Error())
        return
    }
    err = tx.Commit().Error
    if err != nil {
        tx.Rollback()
        httpresponse.ErrResponseWhenCommitDb(w, err.Error())
    }
    httpresponse.JsonResponseForSuccessOperation(w)
}

func Process(username string,
             w http.ResponseWriter,
             r *http.Request) {
    log.Println(fmt.Sprintf("servernports %s", r.Method))
    switch r.Method {
      case "GET":
          get(username, w)
      case "POST":
          post(w, r)
      case "PUT":
          put(w, r)
      case "DELETE":
          del(w, r)
      default:
          http.Error(w, "", http.StatusBadRequest)
    }
}
