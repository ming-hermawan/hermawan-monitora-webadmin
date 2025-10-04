package logout

import (
    "net/http"
    "hermawan-monitora/hmonglobal"
    "hermawan-monitora/webserver/module/html"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/stdlib"
)

func get(w http.ResponseWriter, r *http.Request) {
    cookie := stdlib.GetCookie("");
    http.SetCookie(w, &cookie)
    html.GetPage("Logout", w, hmonglobal.LogoutHtmlFilepath)
}

func Process(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case "GET":
          get(w, r)
      default:
          httpresponse.ErrResponseForBadRequest(w)
    }
}
