package changepassword

import (
    "fmt"
    "net/http"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/module/hashpassword"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/validation"
)

const Lbl = "Change Password"
const MARK = "webserver/controller/main/profile/changepassword"


func get(username string, w http.ResponseWriter) {
    html.GetTmpl1(
      username,
      "Change Password",
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.Base1HtmlFilepath,
      hmonglobal.InpFrm1ColHtmlFilepath,
      hmonglobal.ChangePwdHtmlFilepath,
    )
}

func put(w http.ResponseWriter, r *http.Request, username string) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    oldpassword, newpassword, passwordFound := validation.PasswordValidation(w, params)
    if !passwordFound {
        return
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    user, userFound, err := dbo.GetUser(db, username)
    if err != nil {
        httpresponse.ErrResponseForDetaiRow(
          w,
          Lbl,
          username,
          err.Error())
    }
    if !userFound {
        httpresponse.ErrResponseForDetaiRow(
          w,
          Lbl,
          username,
          fmt.Sprintf("%s not Found", user))
    }
    dbpassword := user.Password
    parampassword := hashpassword.Get(oldpassword)
    if dbpassword == parampassword {
        if err := dbo.UpdateUserPassword(db, username, hashpassword.Get(newpassword));  err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        httpresponse.JsonResponseForSuccessOperation(w)
        return
    } else {
    http.Error(w, "Wrong Password", http.StatusInternalServerError)
    }
    return
}

func Process(w http.ResponseWriter, r *http.Request, username string) {
    switch r.Method {
      case "GET":
          get(username, w)
      case "PUT":
          put(w, r, username)
      default:
          httpresponse.ErrResponseForBadRequest(w)
    }
}
