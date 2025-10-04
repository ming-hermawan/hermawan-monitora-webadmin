package login

import (
    "fmt"
    "time"
    "net/http"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/module/hashpassword"
    "hermawan-monitora/webserver/global/menu"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/jwt"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/validation"
)


// PRIVATE

func post(w http.ResponseWriter, r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    username, usernameFound := validation.StrParamValidation(
      w,
      params,
      "username",
      "Username",
      true,
      hmonglobal.RegexId)
    if !usernameFound {
        return
    }
    lastLogin, loginAttempt := getUserFailedLoginInfo(username)
    if lastLogin > 0 {
        timeLimit := time.Unix(lastLogin, 0).Add(time.Minute * 15).Unix()
        if (time.Now().Unix() < timeLimit) && (loginAttempt >= 3) {
            httpresponse.JsonResponseForCannotLoginBecauseWrongPassword3Times(w, username)
            return
        }
    }
    password, passwordFound := validation.StrParamValidation(
      w,
      params,
      "password",
      "Password",
      true,
      hmonglobal.RegexPwd)
    if !passwordFound {
        return
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    user, userFound, err := dbo.GetUser(db, username)
    if err != nil {
        httpresponse.ErrResponseWhenSelDb(
          w,
          "users",
          err.Error())
        return
    }
    if !userFound {
        httpresponse.JsonResponseForUnauthorizedLogin(
          w,
          fmt.Sprintf("User %s not Found.", username))
        return
    }
    if user.Password != hashpassword.Get(password) {
        setUserFailedLoginInfo(username, time.Now().Unix(), (loginAttempt + 1))
        httpresponse.JsonResponseForUnauthorizedLogin(
          w,
          fmt.Sprintf("Password for %s not match.", username))
        return
    }
    menuAccess := make(map[string]int)
    if user.IsAdmin == 1 {
        for _, x := range menu.Menus {
            menuAccess[x] = 1
        }
    } else if user.IsAdmin == 0 {
        for _, x := range menu.Menus {
            menuAccess[x] = 0
        }
        userGroup, userGroupFound, _ := dbo.GetUserGroup(db, user.UserGroup)
        if !userGroupFound {
            httpresponse.JsonResponseForUnauthorizedLogin(
              w,
              fmt.Sprintf("User-Group %s not Found.", username))
            return
        }
        userGroupMenus, _ := dbo.GetUserGroupMenus(db, userGroup.UserGroup)
        for _, x := range userGroupMenus {
            menuAccess[x] = 1
        }
    }
    for key, val := range menuAccess {
        redisKey := hmonredis.GetUsrMenu(username, key)
        err := hmonredis.SetInt(
          redisKey,
          val)
        if err != nil {
            httpresponse.ErrResponseSetRedis(
              w,
              redisKey,
              err.Error())
            return
        }
    }
    delUserFailedLoginInfo(username)
    token := jwt.GetJwtToken(w, username)
    if token == "" {
        return
    }
    cookie := stdlib.GetCookie(token);
    http.SetCookie(w, &cookie)
    httpresponse.JsonResponseForSuccessOperation(w)
}


// PUBLIC

const MARK = "webserver/controller/main/auth/login"


func Get(w http.ResponseWriter) {
    html.GetTmpl0(
      "Login",
      w,
      hmonglobal.Base0HtmlFilepath,
      hmonglobal.LoginHtmlFilepath,
    )
}

func Process(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case "GET":
          Get(w)
      case "POST":
          post(w, r)
      default:
          httpresponse.ErrResponseForBadRequest(w)
    }
}
