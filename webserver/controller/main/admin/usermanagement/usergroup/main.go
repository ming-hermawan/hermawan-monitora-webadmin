package usergroup

import (
    "fmt"
    "net/http"
    "gorm.io/gorm"
    "hermawan-monitora/hmonglobal"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/webserver/decorator"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/validation"
)


// CONST

const Lbl = "User-Groups"


// PRIVATE

func getRowsCnt(db *gorm.DB,
                txtFilter string) (int64, error) {
    return dbo.GetCntOfUserGroupsForTbl(
      db,
      txtFilter)
}

func getRowsForTbl(db *gorm.DB,
                   txtFilter string,
                   page int,
                   pageLimit int) ([][]string, error) {
    return dbo.GetUserGroupsForTbl(
      db,
      txtFilter,
      page,
      pageLimit)
}

func getRowForDetail(db *gorm.DB,
                     w http.ResponseWriter,
                     usergroup string) map[string]any {
    var userGroup dbo.UserGroup
    var userGroupFound bool
    var userGroupMenus []string
    var err error
    tblName := "user_groups"
    userGroup, userGroupFound, err = dbo.GetUserGroup(
     db,
     usergroup)
    if err != nil {
        httpresponse.ErrResponseForDetaiRow(
          w,
          tblName,
          usergroup,
          err.Error())
    }
    if !userGroupFound {
        httpresponse.ErrResponseForDetaiRow(
          w,
          tblName,
          usergroup,
          fmt.Sprintf("%s not Found", usergroup))
    }
    userGroupMenus, err = dbo.GetUserGroupMenus(db, usergroup)
    if err != nil {
        httpresponse.ErrResponseForDetaiRow(
          w,
          "user_group_menus",
          usergroup,
          err.Error())
    }
    return map[string]any {
      "usergroup": userGroup.UserGroup,
      "sortnumber": userGroup.SortNumber,
      "menus": userGroupMenus}
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
      hmonglobal.InpFrm2ColWithTabPagesHtmlFilepath,
      hmonglobal.AdmUsrMgtUsrGrpHtmlFilepath,
    )
}

func post(w http.ResponseWriter, r *http.Request) {
    decorator.HttpPostForMasterDataWithTxtFilter(
      w,
      r,
      Lbl,
      "key",
      []string {"User-Group"},
      getRowsCnt,
      getRowsForTbl,
      getRowForDetail)
}

func put(w http.ResponseWriter, r *http.Request) {
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
    usergroup, usergroupFound := validation.StrParamValidation(
      w,
      params,
      "usergroup",
      "User-Group",
      true,
      hmonglobal.RegexId)
    if !usergroupFound {
        return
    }
    sortnumber, sortnumberFound := validation.IntParamValidation(
      w,
      params,
      "sortnumber",
      "Sort-Number",
      true)
    if !sortnumberFound {
        return
    }
    rawmenus, rawmenusFound := params["menus"].([]interface{})
    if !rawmenusFound {
        http.Error(w, "Menus not found!", http.StatusInternalServerError)
        return
    }
    menus := make([]string, len(rawmenus))
    for i, v := range rawmenus {
        menus[i] = v.(string)
    }
    var err error
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    tx := db.Begin()
    if statDb == "INS" {
        err = dbo.InsUserGroup(tx, usergroup, sortnumber)
        if err != nil {
            tx.Rollback()
            httpresponse.ErrResponseWhenInsDb(
              w,
              "user_groups",
              err.Error())
            return
        }
    } else if statDb == "UPD" {
        err = dbo.UpdUserGroup(tx, usergroup, sortnumber)
        if err != nil {
            tx.Rollback()
            httpresponse.ErrResponseWhenUpdDb(
              w,
              "user_groups",
              err.Error())
            return
        }
    }
    err = dbo.UpdUserGroupMenus(tx, usergroup, menus)
    if err != nil {
        tx.Rollback()
        httpresponse.ErrResponseWhenUpdDb(w,
          "user_group_menus",
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

func del(w http.ResponseWriter, r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    if params == nil {
        return
    }
    var usergroups []string;
    for _, x:= range params["keys"].([]interface{}) {
        usergroups = append(usergroups, x.(string))
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    var err error
    tx := db.Begin()
    err = dbo.DelUserGroups(db, usergroups)
    if err != nil {
        tx.Rollback()
        httpresponse.ErrResponseWhenDelDb(
          w,
          "user_groups",
          err.Error())
        return
    }
    err = dbo.DelUserGroupMenus(db, usergroups)
    if err != nil {
        tx.Rollback()
        httpresponse.ErrResponseWhenDelDb(
          w,
          "user_group_menus",
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
          httpresponse.ErrResponseForBadRequest(w)
    }
}
