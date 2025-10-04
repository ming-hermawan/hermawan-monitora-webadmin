package profile

import (
    "fmt"
    "log"
    "encoding/base64"
    "encoding/json"
    "net/http"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/webserver/module/avatar"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/validation"
)

const Lbl = "Profile"


func get(username string, w http.ResponseWriter) {
    html.GetTmpl1(
      username,
      "Profile",
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.Base1HtmlFilepath,
      hmonglobal.ProfileHtmlFilepath,
    )
}

func post(w http.ResponseWriter, r *http.Request, username string) {
    language := "en"
    darkmode := 0
    note := ""
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
    name := user.Name
    email := user.Email
    var status string
    if user.IsAdmin == 1 {
        status = "Admin"
    } else {
        status = "User"
    }
    userSetting, dbResult := dbo.GetUserSetting(db, username)
    if dbResult.RowsAffected == 1 {
        language = userSetting.Language
        darkmode = userSetting.Darkmode
        note = userSetting.Note
    }
    avatarStr := ""
    avatarBytes, err3 := avatar.GetAvatarFromFile(username)
    if err3 == nil {
        avatarStr =  base64.StdEncoding.EncodeToString(avatarBytes)
    }
    result := map[string]any{
      "username": username,
      "name": name,
      "email": email,
      "status": status,
      "language": language,
      "darkmode": darkmode,
      "avatar": avatarStr,
      "note": note}
    jsonInBytes, _ := json.Marshal(&result)
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}

func put(w http.ResponseWriter, r *http.Request, username string) {
    var params map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    log.Println(fmt.Printf("PUT params:\n%+v\n", params))
    language, languageFound := validation.StrParamValidation(
      w,
      params,
      "language",
      "Language",
      true,
      hmonglobal.RegexPwd)
    if !languageFound {
        return
    }
    darkMode, darkModeFound := validation.IntParamValidation(
      w,
      params,
      "darkmode",
      "Dark-Mode",
      true)
    if !darkModeFound {
        return
    }
    note, noteFound := validation.StrParamValidation(
      w,
      params,
      "note",
      "Note",
      true,
      "")
    if !noteFound {
        return
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    dbo.SavSetting(db, username, language, darkMode, note)
    httpresponse.JsonResponseForSuccessOperation(w)
}

func Process(w http.ResponseWriter, r *http.Request, username string) {
    switch r.Method {
      case "GET":
          get(username, w)
      case "POST":
          post(w, r, username)
      case "PUT":
          put(w, r, username)
      default:
          httpresponse.ErrResponseForBadRequest(w)
    }
}
