package initredis

import (
    "encoding/json"
    "net/http"
    "hermawan-monitora/webserver/module/httpresponse"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/module/hmonredis"
)

func Init(w http.ResponseWriter) bool {
    db := stdlib.GetDb(w)
    // Emails
    mpAlertEmails, err := dbo.GetMPAlertEmails(db)
    if err != nil {
        httpresponse.ErrResponseWhenSelDb(
          w,
          "mp_alert_emails",
          err.Error())
        return false
    }
    serverEmails, err := dbo.GetServerEmails(db, mpAlertEmails)
    if err != nil {
        httpresponse.ErrResponseWhenSelDb(
          w,
          "mp_server_alert_emails",
          err.Error())
        return false
    }
    for key, val := range serverEmails {
        redisKey := hmonredis.GetServerMails(key)
        hmonredis.SetStr(redisKey, val)
        if err != nil {
            httpresponse.ErrResponseSetRedis(
              w,
              redisKey,
              err.Error())
            return false
        }
    }
    // Server & Ports
    mpServerNPorts, err := dbo.GetMPServerNPortsForRedis(db)
    if err != nil {
        httpresponse.ErrResponseWhenSelDb(
          w,
          "mp_servers/mp_ports",
          err.Error())
        return false
    }
    for _, val := range mpServerNPorts {
        result := map[string]string {
          "serverName": val.ServerName,
          "serviceName": val.ServiceName}
        jsonInBytes, err := json.Marshal(&result)
        if err != nil {
            httpresponse.ErrResponseWhenCnvToJson(
              w,
	      err.Error())
            return false
        }
        redisKey := hmonredis.GetServerNPortKey(val.Ip, val.Port)
        hmonredis.SetRaw(redisKey, jsonInBytes)
        if err != nil {
            httpresponse.ErrResponseSetRedis(
              w,
              redisKey,
              err.Error())
            return false
        }
    }
    return true
}

func SendInitSignal(w http.ResponseWriter) bool {
    err := hmonredis.SetStr(hmonredis.PortScanStatus, "INIT")
    if err != nil {
        httpresponse.ErrResponseSetRedis(
          w,
          hmonredis.PortScanStatus,
          err.Error())
        return false
    }
    return true
}
