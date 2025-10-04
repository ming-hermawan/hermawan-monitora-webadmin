package refreshtoken

import (
    "strconv"
    "encoding/json"
    "net/http"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/module/jwt"
    "hermawan-monitora/webserver/module/stdlib"
    "hermawan-monitora/webserver/module/httpresponse"
)


func post(w http.ResponseWriter, r *http.Request) {
    logoutAfter1HourRaw, err := hmonredis.Get(hmonredis.LogoutAfter1Hour)
    if err != nil {
        httpresponse.ErrResponseGetRedis(
          w,
	  hmonredis.LogoutAfter1Hour,
	  err.Error())
        return
    }
    logoutAfter1Hour, _ := strconv.Atoi(logoutAfter1HourRaw)
    if logoutAfter1Hour == 1 {
        httpresponse.ErrResponseForInvalidToken(
          w,
          "Not allowed to refresh token, 'Logout After 1 Hour' is true.")
        return
    }
    oldcookie, err := r.Cookie("token")
    if err != nil {
        httpresponse.ErrResponseForInvalidToken(w, err.Error())
        return
    }
    username, err := jwt.GetUsernameFromToken(oldcookie.Value)
    if err != nil {
        httpresponse.ErrResponseForInvalidToken(w, err.Error())
	return
    }
    token := jwt.GetJwtToken(w, username)
    if token == "" {
        return
    }
    newcookie := stdlib.GetCookie(token);
    http.SetCookie(w, &newcookie)

    result := map[string]string {
      "status": "SUCCESS"}

    jsonInBytes, err := json.Marshal(&result)
    if err != nil {
        httpresponse.ErrResponseForInvalidToken(w, err.Error())
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}

func Process(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case "POST":
        post(w, r)

      default:
        http.Error(w, "", http.StatusBadRequest)
    }
}
