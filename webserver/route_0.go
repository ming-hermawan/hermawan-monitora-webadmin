package webserver

import (
    "errors"
    "net/http"
    "hermawan-monitora/webserver/module/jwt"
    "hermawan-monitora/webserver/controller/main/dashboard"
    "hermawan-monitora/webserver/controller/main/auth/login"
    "hermawan-monitora/webserver/controller/main/profile/avatar"
)


func authProcess(r *http.Request) (string, error) {
    cookie, err := r.Cookie("token")
    if err != nil {
        return "", err
    }
    if cookie.Value == "" {
        return "", errors.New("Cookie is Empty")
    }
    username, _ := jwt.GetUsernameFromToken(cookie.Value)
    return username, nil
}


func dashboardProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    dashboard.Process(username, w, r)
}

func avatarProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    avatar.Process(w, r, username)
}
