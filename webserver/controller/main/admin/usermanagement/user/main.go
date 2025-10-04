package user

import (
    "fmt"
    "gorm.io/gorm"
    "net/http"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/module/hashpassword"
    "hermawan-monitora/webserver/decorator"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/validation"
    ddlUserGroup "hermawan-monitora/webserver/controller/ddl/admin/usermanagement/usergroup"
)


// CONST

const Lbl = "Users"


// PRIVATE

func getListForDDL(db *gorm.DB,
                   w http.ResponseWriter) ([]string, error) {
    return dbo.GetUserGroupsForDDL(db)
}

func getRowsCnt(db *gorm.DB,
                grpFilter string,
                txtFilter string) (int64, error) {
    return dbo.GetCntOfUserForTbl(
      db,
      grpFilter,
      txtFilter)
}

func getRowsForTbl(db *gorm.DB,
                     grpFilter string,
                     txtFilter string,
                     page int,
                     pageLimit int) ([][]string, error) {
    return dbo.GetUserForTbl(
      db,
      grpFilter,
      txtFilter,
      page,
      pageLimit)
}

func getRowForDetail(db *gorm.DB,
                     w http.ResponseWriter,
                     username string) map[string]any {
    tblName := "users"
    user, userFound, err := dbo.GetUser(
      db,
      username)
    if err != nil {
        httpresponse.ErrResponseForDetaiRow(
          w,
          tblName,
          username,
          err.Error())
    }
    if !userFound {
        httpresponse.ErrResponseForDetaiRow(
          w,
          tblName,
          username,
          fmt.Sprintf("%s not Found", user))
    }
    usergroupList := ddlUserGroup.Get(db, w)
    if usergroupList == nil {
        return nil
    }
    return map[string]any {
      "username": user.Username,
      "name": user.Name,
      "isAdmin": user.IsAdmin,
      "email": user.Email,
      "usergroup": user.UserGroup,
      "usergroupList": usergroupList}
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
      hmonglobal.InpFrmFilterGrpHtmlFilepath,
      hmonglobal.InpFrm2ColNormalHtmlFilepath,
      hmonglobal.AdmUsrMgtHtmlFilepath,
    )
}


func post(w http.ResponseWriter,
          r *http.Request) {
    decorator.HttpPostForMasterDataWithTxtAndGrpFilter(
      w,
      r,
      Lbl,
      "key",
      []string{"Username", "Name", "E-Mail"},
      getRowsCnt,
      getRowsForTbl,
      getRowForDetail,
      getListForDDL)
}

func put(w http.ResponseWriter,
         r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    if params == nil {
        return
    }
    statDb, statDbValid := validation.StatDbParamValidation(
      w,
      params)
    if !statDbValid {
        return
    }
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
    password, passwordFound := validation.StrParamValidation(
      w,
      params,
      "password",
      "Password",
      false,
      "")
    var newpassword string
    if passwordFound {
        newpassword = hashpassword.Get(password)
    } else {
        newpassword = ""
    }
    name, nameFound := validation.StrParamValidation(
      w,
      params,
      "name",
      "Name",
      true,
      "")
    if !nameFound {
        return
    }
    isAdmin, isAdminFound := validation.IntParamValidation(
      w,
      params,
      "isAdmin",
      "Admin",
      true)
    if !isAdminFound {
        return
    }
    email, emailFound := validation.EmailParamvalidation(
      w,
      params,
      "email",
      "Email",
      false)
    if !emailFound {
        return
    }
    group, groupFound := validation.StrParamValidation(
      w,
      params,
      "group",
      "User-Group",
      true,
      hmonglobal.RegexId)
    if !groupFound {
        return
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    if statDb == "INS" {
        err := dbo.InsUser(
          db,
          username,
          newpassword,
          name,
          isAdmin,
          email,
          group)
        if err != nil {
            httpresponse.ErrResponseWhenInsDb(
              w,
              "server_groups",
              err.Error())
            return
        }
    } else if statDb == "UPD" {
        err := dbo.UpdUser(
          db,
          username,
          newpassword,
          name,
          isAdmin,
          email,
          group)
        if err != nil {
            httpresponse.ErrResponseWhenUpdDb(
              w,
              "server_groups",
              err.Error())
            return
        }
    }
    httpresponse.JsonResponseForSuccessOperation(w)
}


func del(w http.ResponseWriter,
         r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    if params == nil {
        return
    }
    var users []string;
    for _, x:= range params["keys"].([]interface{}) {
        users = append(users, x.(string))
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    err := dbo.DelUser(
      db,
      users)
    if err != nil {
        httpresponse.ErrResponseWhenDelDb(
          w,
          "users",
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
