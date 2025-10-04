package decorator

import (
    "fmt"
    "math"
    "encoding/json"
    "gorm.io/gorm"
    "net/http"
    "hermawan-monitora/module/hmonenv"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/stdlib"
)


// PRIVATE

// 1. Type

type getRowsForDDLFunc func(*gorm.DB) ([]string, error)
type getRowForDetailFunc func(*gorm.DB,
                              http.ResponseWriter,
                              string) map[string]any
type getRowsCntWithTxtFilterFunc func(*gorm.DB,
                                      string) (int64, error)
type getRowsCntWithTxtAndGrpFilterFunc func(*gorm.DB,
                                            string,
                                            string) (int64, error)
type getRowsForTblWithTextFilterFunc func(*gorm.DB,
                                          string,
                                          int,
                                          int) ([][]string, error)
type getRowsForTblWithTextAndGrpFilterFunc func(*gorm.DB,
                                                string,
                                                string,
                                                int,
                                                int) ([][]string, error)
type getListForDDLFunc func(*gorm.DB,
                            http.ResponseWriter) ([]string, error)

// 2. Function

func getHttpPostParamsForMasterData(w http.ResponseWriter,
                                    r *http.Request,
                                    keyName string,
                                    getRowForDetail getRowForDetailFunc) map[string]interface{} {
    params := stdlib.GetPayloadFromJsonBody(w, r)
    if params == nil {
        return nil
    }
    key, keyFound := params[keyName].(string)
    if keyFound {
        db := stdlib.GetDb(w)
        if db == nil {
            return nil
        }
        result := getRowForDetail(db, w, key)
        if result == nil {
            return nil
        }
        jsonInBytes, err := json.Marshal(&result)
        if err != nil {
            httpresponse.ErrResponseWhenCnvToJson(
              w,
              err.Error())
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonInBytes)
        return nil
    } else {
        return params
    }
}

func getPageParam(params map[string]interface{}) int {
    float64page, found := params["page"].(float64)
    out := 1
    if (found) {
        out = int(float64page);
    }
    return out
}

func getTxtFilterParam(params map[string]interface{}) string {
    out, found := params["txtFilter"].(string)
    if !found {
        out = "";
    }
    return out
}

func getGrpFilterParam(params map[string]interface{}) string {
    out, found := params["grpFilter"].(string)
    if (!found) || (out == "Semua")  {
        out = "";
    }
    return out
}

func getPageCount(rowsCount int64, pageLimit int) int {
    out := int(math.Ceil(float64(rowsCount) / float64(pageLimit))) + 1
    return out
}


// PUBLIC

func GetDDL(db *gorm.DB,
            w http.ResponseWriter,
            name string,
            getRowsForDDL getRowsForDDLFunc) []string {
    out, err := getRowsForDDL(db)
    if err != nil {
        httpresponse.ErrResponseDb(
          w,
          fmt.Sprintf(
            "Get %s List for DDL Error",
            name),
          err.Error())
        return nil
    }
    return out
}

func HttpPostToGetDDL(w http.ResponseWriter,
                      name string,
                      getRowsForDDL getRowsForDDLFunc) {
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    out := GetDDL(db, w, name, getRowsForDDL)
    if out == nil {
        return
    }
    httpresponse.JsonResponseForDDL(w, out)
}

func HttpPostForMasterDataWithTxtFilter(w http.ResponseWriter,
                                        r *http.Request,
                                        name string,
                                        keyName string,
                                        headers []string,
                                        getRowsCnt getRowsCntWithTxtFilterFunc,
                                        getRowsForTbl getRowsForTblWithTextFilterFunc,
                                        getRowForDetail getRowForDetailFunc) {
    params := getHttpPostParamsForMasterData(
      w,
      r,
      keyName,
      getRowForDetail)
    if params == nil {
        return
    }
    pageLimit := hmonenv.GetPageLimit()
    page := getPageParam(params)
    txtFilter := getTxtFilterParam(params)
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    rowsCount, err := getRowsCnt(db, txtFilter)
    if err != nil {
        httpresponse.ErrResponseForMasterDataRowsCount(
          w,
          name,
          err.Error())
    }
    pageCount := getPageCount(rowsCount, pageLimit)
    rows, err := getRowsForTbl(db, txtFilter, page, pageLimit)
    if err != nil {
        httpresponse.ErrResponseForMasterDataRows(
          w,
          name,
          err.Error())
    }
    httpresponse.JsonResponseForMasterDataList(
      w,
      headers,
      rowsCount,
      pageCount,
      rows,
      nil)
}

func HttpPostForMasterDataWithTxtAndGrpFilter(w http.ResponseWriter,
                                              r *http.Request,
                                              name string,
                                              keyName string,
                                              headers []string,
                                              getRowsCnt getRowsCntWithTxtAndGrpFilterFunc,
                                              getRowsForTbl getRowsForTblWithTextAndGrpFilterFunc,
                                              getRowForDetail getRowForDetailFunc,
                                              getListForDDL getListForDDLFunc) {
    params := getHttpPostParamsForMasterData(
      w,
      r,
      keyName,
      getRowForDetail)
    if params == nil {
        return
    }
    pageLimit := hmonenv.GetPageLimit()
    page := getPageParam(params)
    txtFilter := getTxtFilterParam(params)
    grpFilter := getGrpFilterParam(params)
    db := stdlib.GetDb(w)
    if db == nil {
        return
    }
    rowsCount, err := getRowsCnt(db, grpFilter, txtFilter)
    if err != nil {
        httpresponse.ErrResponseForMasterDataRowsCount(
          w,
          name,
          err.Error())
    }
    pageCount := getPageCount(rowsCount, pageLimit)
    rows, err := getRowsForTbl(
      db,
      grpFilter,
      txtFilter,
      page,
      pageLimit)
    if err != nil {
        httpresponse.ErrResponseForMasterDataRows(
          w,
          name,
          err.Error())
    }
    groupList, err := getListForDDL(db, w)
    if err != nil {
        httpresponse.ErrResponseDb(
          w,
          "Get Group List for DDL Error",
          err.Error())
    }
    httpresponse.JsonResponseForMasterDataList(
      w,
      headers,
      rowsCount,
      pageCount,
      rows,
      groupList)
}
