package webserver

import (
    "net/http"
    "hermawan-monitora/webserver/controller/main/auth/login"
    "hermawan-monitora/webserver/controller/main/profile/changepassword"
    "hermawan-monitora/webserver/controller/main/profile/profile"
)


func changePasswordProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    changepassword.Process(w, r, username)
}

func profileProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    profile.Process(w, r, username)
}
