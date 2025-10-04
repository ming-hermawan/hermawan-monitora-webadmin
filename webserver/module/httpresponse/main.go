package httpresponse

import (
    "fmt"
    "log"
    "encoding/json"
    "net/http"
    "hermawan-monitora/hmonglobal/lang"
    "hermawan-monitora/module/hmonlog"
)


// PRIVATE

func errResponse(w http.ResponseWriter,
                 status int,
                 title string,
                 msg string) {
    id, err := hmonlog.WriteLog(status, title, msg)
    if err != nil {
        log.Println(
          fmt.Printf(
            "%s\n",
            err.Error()))
    }
    http.Error(
      w,
      fmt.Sprintf(
        "%s<br>%s",
        lang.InternalServerWebError,
        id),
      http.StatusInternalServerError)
}


// PUBLIC

// 1. Error Response

// 1.a. Bad Request

func ErrResponseForBadRequest(w http.ResponseWriter) {
    http.Error(
      w,
      lang.UnknownWebError,
      http.StatusBadRequest)
}

// 1.b. DB Error Response

func ErrResponseDb(w http.ResponseWriter,
                   title string,
                   msg string) {
    errResponse(
      w,
      hmonlog.DbLog,
      title,
      msg)
}

func ErrResponseWhileDbConnect(w http.ResponseWriter,
                               msg string) {
    ErrResponseDb(
      w,
      lang.DbConnErrTitle,
      msg)
}

func ErrResponseForMasterDataRowsCount(w http.ResponseWriter,
                                       name string,
                                       msg string) {
    ErrResponseDb(
      w,
      lang.MasterDataRowsCountErrTitle(name),
      msg)
}

func ErrResponseForMasterDataRows(w http.ResponseWriter,
                                  name string,
                                  msg string) {
    ErrResponseDb(
      w,
      lang.MasterRowsErrTitle(name),
      msg)
}

func ErrResponseForDetaiRow(w http.ResponseWriter,
                            tblName string,
                            key string,
                            msg string) {
    ErrResponseDb(
      w,
      lang.DetailRowErrTitle(tblName, key),
      msg)
}

func ErrResponseWhenSelDb(w http.ResponseWriter,
                          tblName string,
                          msg string) {
    ErrResponseDb(
      w,
      lang.SelDbErrTitle(tblName),
      msg)
}

func ErrResponseWhenInsDb(w http.ResponseWriter,
                          tblName string,
                          msg string) {
    ErrResponseDb(
      w,
      lang.InsDbErrTitle(tblName),
      msg)
}

func ErrResponseWhenUpdDb(w http.ResponseWriter,
                          tblName string,
                          msg string) {
    ErrResponseDb(
      w,
      lang.UpdDbErrTitle(tblName),
      msg)
}

func ErrResponseWhenDelDb(w http.ResponseWriter,
                          tblName string,
                          msg string) {
    ErrResponseDb(
      w,
      lang.DelDbErrTitle(tblName),
      msg)
}

func ErrResponseWhenCommitDb(w http.ResponseWriter,
                             msg string) {
    ErrResponseDb(w,
      lang.CommitErrTitle,
      msg)
}


// 1.c. Redis Error Response

func ErrResponseRedis(w http.ResponseWriter,
                      title string,
                      msg string) {
    errResponse(
      w,
      hmonlog.RedisLog,
      title,
      msg)
}

func ErrResponseSetRedis(w http.ResponseWriter,
                         key string,
                         msg string) {
    errResponse(
      w,
      hmonlog.RedisLog,
      lang.SetRedisErrTitle(key),
      msg)
}

func ErrResponseGetRedis(w http.ResponseWriter,
                         key string,
                         msg string) {
    errResponse(
      w,
      hmonlog.RedisLog,
      lang.GetRedisErrTitle(key),
      msg)
}

// 1.c. Web Server Error Response

func ErrResponseWs(w http.ResponseWriter,
                   title string,
                   msg string) {
    errResponse(
      w,
      hmonlog.WebserverLog,
      title,
      msg)
}

func ErrResponseForHttpBody(w http.ResponseWriter,
                            msg string) {
    ErrResponseWs(
      w,
      lang.HttpBodyErrTitle,
      msg)
}

func ErrResponseForInvalidToken(w http.ResponseWriter,
                                msg string) {
    ErrResponseWs(
      w,
      lang.TokenErrTitle,
      msg)
}

func ErrResponseWhenCnvToJson(w http.ResponseWriter,
                              msg string) {
    ErrResponseWs(
      w,
      lang.CnvJsonErrTitle,
      msg)
}

// 2. Json Response

func JsonResponseForSuccessOperation(w http.ResponseWriter) {
    jsonInBytes, _ := json.Marshal(map[string]any {
      "Status": 1})
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}

func JsonResponseForUnauthorizedLogin(w http.ResponseWriter, msg string) {
    id, _ := hmonlog.WriteLog(hmonlog.WebserverLog, "Login", msg)
    jsonInBytes, _ := json.Marshal(map[string]any {
      "Status": 2,
      "Code": id})
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}

func JsonResponseForCannotLoginBecauseWrongPassword3Times(w http.ResponseWriter,
                                                          username string) {
    id, _ := hmonlog.WriteLog(
      hmonlog.WebserverLog,
      lang.LockedAccountTitle,
      lang.CannotLoginBecauseWrongPassword3Times(username))
    jsonInBytes, _ := json.Marshal(map[string]any {
      "Status": 3,
      "Code": id})
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}

func JsonResponseForDDL(w http.ResponseWriter, payload []string) {
    jsonInBytes, err := json.Marshal(&payload)
    if err != nil {
        ErrResponseWhenCnvToJson(w, err.Error())
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}

func JsonResponseForMasterDataList(w http.ResponseWriter,
                                   headers []string,
                                   rowsCount int64,
                                   pageCount int,
                                   rows [][]string,
                                   groupList [] string) {
    result := map[string]any {}
    result["headers"] = headers
    result["rowsCount"] = rowsCount
    result["pageCount"] = pageCount
    result["rows"] = rows
    if groupList != nil {
        result["groupList"] = groupList
    }
    jsonInBytes, err := json.Marshal(&result)
    if err != nil {
        ErrResponseWs(
          w,
          "JSON Convert Error",
          err.Error())
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}
