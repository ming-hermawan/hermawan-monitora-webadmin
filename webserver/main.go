package webserver

import (
    "fmt"
    "log"
    "sync"
    "net/http"
    socketio "github.com/googollee/go-socket.io"
    "hermawan-monitora/hmonglobal"
    "hermawan-monitora/hmonglobal/lang"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
    "hermawan-monitora/module/hmonenv"
    "hermawan-monitora/module/hmonredis"
    "hermawan-monitora/webserver/controller/main/cookieswarning"
    "hermawan-monitora/webserver/controller/main/auth/login"
    "hermawan-monitora/webserver/controller/main/auth/refreshtoken"
    "hermawan-monitora/webserver/controller/main/monitoring/port"
    "hermawan-monitora/webserver/controller/main/auth/logout"
    "hermawan-monitora/webserver/global/resetsignal"
)

var ServerPort map[string]interface{}

func checkResetSignal(s socketio.Conn) {
  for {
    if resetsignal.Get() {
        log.Println("RESET SIGNAL")
        message1 := fmt.Sprintf("{\"status\":\"RESET\"}")
        s.Emit("reply", message1)
        resetsignal.Set(false)
    }
  }
}

func Run() {
    server := socketio.NewServer(nil)

    var wg sync.WaitGroup
    wgN := 0

    settings := dbo.GetSetting()
    var logoutAfter1Hour int
    if settings.LogoutAfter1Hour.Valid {
        logoutAfter1Hour = int(settings.LogoutAfter1Hour.Int64)
    } else {
        logoutAfter1Hour = 0
    }
    err := hmonredis.SetInt(hmonredis.LogoutAfter1Hour, logoutAfter1Hour)
    if err != nil {
        panic(
          lang.RedisSetErr(
            hmonredis.LogoutAfter1Hour,
            err.Error()))
    }

    server.OnConnect("/", func(s socketio.Conn) error {
        s.SetContext("")
        log.Println(fmt.Println("SOCKET connected:", s.ID()))

        ServerPort = port.GetServerPort("", "", "")
        log.Println(fmt.Printf("ServerPort:\n%+v\n", ServerPort))
        for kServer, rowServerPort := range ServerPort {
            temp1 := rowServerPort.(map[string]interface{})["ports"].(map[int]map[string]string)
            for kPort, rowPort := range temp1 {
                log.Println(fmt.Printf("rowPort:\n%+v\n---\n", rowPort))
                wg.Add(wgN)
                wgN += 1
                go port.RedisSubscribe(s, kServer, kPort, rowPort["uniqueId"])
            }
        }

        go checkResetSignal(s)
        return nil
    })

    server.OnError("/", func(s socketio.Conn, e error) {
        log.Println(fmt.Println("meet error:", e))
    })

    server.OnDisconnect("/", func(s socketio.Conn, reason string) {
        // Add the Remove session id. Fixed the connection & mem leak
        // server.Remove(s.ID())
        log.Println(fmt.Println("closed", reason))
    })

    wg.Add(wgN)
    go server.Serve()
  	defer server.Close()
    wg.Wait()

    http.Handle("/socket.io/", server)
    fs := http.FileServer(http.Dir(hmonglobal.StaticDirPath))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", dashboardProcess)

    http.HandleFunc(hmonglobal.MonPortUrl, portsMonitoringProcess)
    http.HandleFunc(hmonglobal.ChangePwdUrl, changePasswordProcess)
    http.HandleFunc(hmonglobal.ProfileUrl, profileProcess)
    http.HandleFunc(hmonglobal.AdmUsrMgtUsrUrl, adminUserProcess)
    http.HandleFunc(hmonglobal.AdmUsrMgtUsrGrpUrl, adminUserGroupProcess)

    http.HandleFunc(hmonglobal.LoginUrl, login.Process)
    http.HandleFunc(hmonglobal.RefreshTokenUrl, refreshtoken.Process)

    http.HandleFunc(hmonglobal.AdmMonPortServerGrpUrl, adminServerGroupProcess)
    http.HandleFunc(hmonglobal.AdmMonPortServerNPortsUrl, adminServerNPortsProcess)
    http.HandleFunc(hmonglobal.AdmMonPortUploadCsvUrl, uploadCSVProcess)
    http.HandleFunc(hmonglobal.AdmMonPortEmailUrl, admMonEmailProcess)

    http.HandleFunc(hmonglobal.DdlAdmUsrMgtUsrGrpUrl, ddlUserGroupProcess)
    http.HandleFunc(hmonglobal.DdlAdmMonPortServerGrpUrl, ddlServerGroupProcess)


    http.HandleFunc(hmonglobal.TmpGetServerPorts, port.ProcessGetServerPort)

    http.HandleFunc(hmonglobal.AdmSettingUrl, settingsProcess)

    http.HandleFunc(hmonglobal.MonPortReport, reportMonPortProcess)
    http.HandleFunc(hmonglobal.MonPortReportReport, reportMonPortReportProcess)

    http.HandleFunc(hmonglobal.ProfileAvatarUrl, avatarProcess)
    http.HandleFunc(hmonglobal.LogoutUrl, logout.Process)

    http.HandleFunc(hmonglobal.CookiesWarningUrl, cookieswarning.Process)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", hmonenv.GetPort()), nil))
}
