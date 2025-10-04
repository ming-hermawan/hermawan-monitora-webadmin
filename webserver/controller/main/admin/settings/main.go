package settings

import (
    "fmt"
    "encoding/json"
    "net/http"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/validation"
)

const Lbl = "Settings"

type SettingData struct {
    Name string
    SmtpHost string
    SmtpPort int
    SenderName string
    AuthEmail string
    AuthPassword string
}

func get(username string, w http.ResponseWriter) {
    html.GetTmpl1(
      username,
      Lbl,
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.Base1HtmlFilepath,
      hmonglobal.InpFrm1ColHtmlFilepath,
      hmonglobal.AdmSettingsHtmlFilepath,
    )
}

func post(w http.ResponseWriter, r *http.Request) {
    settings := dbo.GetSetting()
    var logoutAfter1Hour int64
    if settings.LogoutAfter1Hour.Valid {
        logoutAfter1Hour = settings.LogoutAfter1Hour.Int64
    } else {
        logoutAfter1Hour = 0
    }
    var forceStrongPassword int64
    if settings.ForceStrongPassword.Valid {
        forceStrongPassword = settings.ForceStrongPassword.Int64
    } else {
        forceStrongPassword = 0
    }
    result := map[string]any {
      "logoutAfter1Hour": logoutAfter1Hour,
      "forceStrongPassword": forceStrongPassword,
      "smtphost": settings.SmtpHost,
      "smtpport": settings.SmtpPort,
      "sendername": settings.SenderName,
      "authemail": settings.AuthEmail,
      "authpassword": settings.AuthPassword}
    jsonInBytes, err := json.Marshal(&result)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}

func put(w http.ResponseWriter, r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    logoutAfter1Hour, logoutAfter1HourFound := validation.IntParamValidation(
      w,
      params,
      "logoutAfter1Hour",
      "Logout After 1 Hour Status",
      true)
    if !logoutAfter1HourFound {
        return
    }
    forceStrongPassword, forceStrongPasswordFound := validation.IntParamValidation(
      w,
      params,
      "forceStrongPassword",
      "Force Strong Password Status",
      true)
    if !forceStrongPasswordFound {
        return
    }
    smtphost := fmt.Sprintf("%v", params["smtphost"])
    smtpport := int(params["smtpport"].(float64))
    sendername := fmt.Sprintf("%v", params["sendername"])
    authemail := fmt.Sprintf("%v", params["authemail"])
    authpassword := fmt.Sprintf("%v", params["authpassword"])
    dbo.UpdateSetting(logoutAfter1Hour, forceStrongPassword, smtphost, smtpport, sendername, authemail, authpassword)
    err := hmonredis.SetInt(hmonredis.LogoutAfter1Hour, logoutAfter1Hour)
    if err != nil {
        httpresponse.ErrResponseSetRedis(
          w,
          hmonredis.LogoutAfter1Hour,
          err.Error())
        return
    }
    httpresponse.JsonResponseForSuccessOperation(w)
}


func Process(username string,
             w http.ResponseWriter,
             r *http.Request) {
    switch r.Method {
        case "GET":
          get(username, w)
        case "POST":
          post(w, r)
        case "PUT":
          put(w, r)
        default:
          http.Error(w, "", http.StatusBadRequest)
    }
}
