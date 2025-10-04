package hmondbsqlite

import (
    "log"
    "fmt"
    "database/sql"
)

type Setting struct {
    Id int
    LogoutAfter1Hour sql.NullInt64
    ForceStrongPassword sql.NullInt64
    SmtpHost string
    SmtpPort int
    SenderName string
    AuthEmail string
    AuthPassword string
}

func GetSetting() Setting {
    db, err := GetDb()
    if err != nil {
        panic(fmt.Sprintf("Failed to connect SQLite database! %s", err.Error()))
    }
    var out Setting
    db.Where("id = 0").First(&out)
    return out
}

func UpdateSetting(logoutAfter1Hour int,
                   forceStrongPassword int,
                   smtphost string,
                   smtpport int,
                   sendername string,
                   authemail string,
                   authpassword string) {
    db, err := GetDb()
    if err != nil {
        panic("failed to connect database")
    }
    var settings Setting
    if err := db.Where("id = 0").First(&settings).Error; err != nil {
        panic("settings not found")
    }
    if logoutAfter1Hour == 0 {
        settings.LogoutAfter1Hour = sql.NullInt64{Int64: 0, Valid: true}
    } else {
        settings.LogoutAfter1Hour = sql.NullInt64{Int64: 1, Valid: true}
    }
    if forceStrongPassword == 0 {
        settings.ForceStrongPassword = sql.NullInt64{Int64: 0, Valid: true}
    } else {
        settings.ForceStrongPassword = sql.NullInt64{Int64: 1, Valid: true}
    }
    settings.SmtpHost = smtphost
    settings.SmtpPort = smtpport
    settings.SenderName = sendername
    settings.AuthEmail = authemail
    settings.AuthPassword = authpassword
    dbResult := db.Where("id = 0").Updates(&settings)
    if dbResult.Error != nil {
        panic(dbResult.Error.Error())
    }
    log.Println(fmt.Printf("dbResult.RowsAffected=%d\n", dbResult.RowsAffected))
}
