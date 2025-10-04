package servergroup

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

const Lbl = "Server-Group"


// PRIVATE

func getRowsCnt(db *gorm.DB,
                txtFilter string) (int64, error) {
    return dbo.GetCntOfMPServerGroupForTbl(
      db,
      txtFilter)
}

func getRowsForTbl(db *gorm.DB,
                   txtFilter string,
                   page int,
                   pageLimit int) ([][]string, error) {
    return dbo.GetMPServerGroupForTbl(
      db,
      txtFilter,
      page,
      pageLimit)
}

func getRowForDetail(db *gorm.DB,
                     w http.ResponseWriter,
                     servergroup string) map[string]any {
    tblName := "mp_server_groups"
    serverGroup, serverGroupFound, err := dbo.GetMPServerGroup(
      db,
      servergroup)
    if err != nil {
        httpresponse.ErrResponseForDetaiRow(
          w,
          tblName,
          servergroup,
          err.Error())
    }
    if !serverGroupFound {
        httpresponse.ErrResponseForDetaiRow(
          w,
          tblName,
          servergroup,
          fmt.Sprintf("%s not Found", servergroup))
    }
    return map[string]any {
      "servergroup": serverGroup.ServerGroup,
      "sortnumber": serverGroup.SortNumber}
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
      hmonglobal.AdmMonPortServerGrpHtmlFilepath,
    )
}

func post(w http.ResponseWriter, r *http.Request) {
    decorator.HttpPostForMasterDataWithTxtFilter(
      w,
      r,
      Lbl,
      "key",
      []string{"Server-Group"},
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
    servergroup, servergroupFound := validation.StrParamValidation(
      w,
      params,
      "servergroup",
      "Server-Group",
      true,
      hmonglobal.RegexId)
    if !servergroupFound {
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
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    if statDb == "INS" {
        err := dbo.InsMPServerGroup(
          db,
          servergroup,
          sortnumber)
        if err != nil {
            httpresponse.ErrResponseWhenInsDb(
              w,
              "server_groups",
              err.Error())
            return
        }
    } else if statDb == "UPD" {
        err := dbo.UpdMPServerGroup(
          db,
          servergroup,
          sortnumber)
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

func del(w http.ResponseWriter, r *http.Request) {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    if params == nil {
        return
    }
    var servergroups []string;
    for _, x:= range params["keys"].([]interface{}) {
        servergroups = append(servergroups, x.(string))
    }
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    err := dbo.DelMPServerGroups(
      db,
      servergroups)
    if err != nil {
        httpresponse.ErrResponseWhenDelDb(
          w,
          "server_groups",
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
