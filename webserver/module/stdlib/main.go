package stdlib

import (
    "encoding/json"
    "gorm.io/gorm"
    "net/http"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/webserver/module/httpresponse"
)


func GetCookie(strToken string) http.Cookie {
    return http.Cookie {
      Name: "token",
      Value: strToken,
      Path: "/",
      MaxAge: 3600,
      HttpOnly: true,
      Secure: true,
      SameSite: http.SameSiteLaxMode,
    }
}

func GetDb(w http.ResponseWriter) *gorm.DB {
    db, err := dbo.GetDb()
    if err != nil {
        httpresponse.ErrResponseWhileDbConnect(
          w,
          err.Error())
        return nil
    }
    return db
}

func GetPayloadFromJsonBody(w http.ResponseWriter,
                            r *http.Request) map[string]interface{} {
    var out map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&out);
    if err != nil {
        httpresponse.ErrResponseForHttpBody(w, err.Error())
        return nil
    }
    return out
}

func initRedis() {

}
