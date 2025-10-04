package email

import (
    "fmt"
    "net/http"
    "gorm.io/gorm"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/webserver/decorator"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/validation"
)


// CONST

const Lbl = "Email-Alert"


// PRIVATE

func getRowsCnt(db *gorm.DB,
                txtFilter string) (int64, error) {
    return dbo.GetCntOfMPAlertEmailForTbl(
      db,
      txtFilter)
}

func getRowsForTbl(db *gorm.DB,
                   txtFilter string,
                   page int,
                   pageLimit int) ([][]string, error) {
    return dbo.GetMPAlertEmailForTbl(
      db,
      txtFilter,
      page,
      pageLimit)
}

func getRowForDetail(db *gorm.DB,
                     w http.ResponseWriter,
                     alertemail string) map[string]any {
    tblName := "mp_alert_emails"
    alertEmail, alertEmailFound, err := dbo.GetMPAlertEmail(
      db,
      alertemail)
    if err != nil {
        httpresponse.ErrResponseForDetaiRow(
          w,
          tblName,
          alertemail,
          err.Error())
    }
    if !alertEmailFound {
        httpresponse.ErrResponseForDetaiRow(
          w,
          tblName,
          alertemail,
          fmt.Sprintf("%s not Found", alertemail))
    }
    return map[string]any {
      "email": alertEmail.Email}
}

func get(username string,
         w http.ResponseWriter) {
    html.GetTmpl1(
      username,
      Lbl,
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.Base1HtmlFilepath,
      hmonglobal.InpFrm2ColBaseHtmlFilepath,
      hmonglobal.InpFrmFilterStdHtmlFilepath,
      hmonglobal.InpFrm2ColNormalHtmlFilepath,
      hmonglobal.AdmMonPortEmailHtmlFilepath,
    )
}

func post(w http.ResponseWriter, r *http.Request) {
    decorator.HttpPostForMasterDataWithTxtFilter(
      w,
      r,
      Lbl,
      "key",
      []string{"Email"},
      getRowsCnt,
      getRowsForTbl,
      getRowForDetail)
}

func put(w http.ResponseWriter, r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    if params == nil {
        return
    }
    email, emailFound := validation.EmailParamvalidation(
      w,
      params,
      "email",
      "Email",
      true)
    if !emailFound {
        return
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    err := dbo.InsMPAlertEmail(
      db,
      email)
    if err != nil {
        httpresponse.ErrResponseWhenInsDb(
          w,
          "email",
          err.Error())
        return
    }
    httpresponse.JsonResponseForSuccessOperation(w)
}

func del(w http.ResponseWriter, r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    if params == nil {
        return
    }
    var emails []string;
    for _, x:= range params["keys"].([]interface{}) {
        emails = append(emails, x.(string))
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    err := dbo.DelMPAlertEmails(
      db,
      emails)
    if err != nil {
        httpresponse.ErrResponseWhenDelDb(
          w,
          "emails",
          err.Error())
        return
    }
    httpresponse.JsonResponseForSuccessOperation(w)
}


// PUBLIC

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
      case "DELETE":
          del(w, r)
      default:
          http.Error(w, "", http.StatusBadRequest)
    }
}
