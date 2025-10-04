package webserver

import (
    "strconv"
    "net/http"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/global/menu"
    "hermawan-monitora/webserver/controller/main/auth/login"
    "hermawan-monitora/webserver/controller/main/admin/settings"
)


func settingsProcess(w http.ResponseWriter, r *http.Request) {
    username, _ := authProcess(r)
    if username == "" {
        login.Get(w)
        return
    }
    admMonPortSettingsRaw, _ := hmonredis.Get(
      hmonredis.GetUsrMenu(
        username,
        menu.AdmMonPortSettings))
    admMonPortSettings, _ := strconv.Atoi(admMonPortSettingsRaw)
    if admMonPortSettings != 1 {
        login.Get(w)
        return
    }
    settings.Process(username, w, r)
}
