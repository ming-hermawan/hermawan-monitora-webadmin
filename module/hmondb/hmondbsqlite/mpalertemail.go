package hmondbsqlite

import (
  "fmt"
  "gorm.io/gorm"
)

type MPAlertEmail struct {
    Email string `gorm:"primaryKey"`
}

func GetMPAlertEmails(db *gorm.DB) ([]MPAlertEmail, error) {
    var alertEmails []MPAlertEmail
    dbResult := db.Find(&alertEmails)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    return alertEmails, nil
}

func GetCntOfMPAlertEmailForTbl(db *gorm.DB,
                                txtFilter string) (int64, error) {
    var out int64
    var dbResult *gorm.DB
    if (txtFilter != "") && (txtFilter != "null") {
        sql := "email LIKE ?"
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        dbResult = db.Model(&MPAlertEmail{}).Where(
          sql,
          txtLikeFilter).Count(&out)
    } else {
        dbResult = db.Model(&MPAlertEmail{}).Count(&out)
    }
    if dbResult.Error != nil {
        return -1, dbResult.Error
    }
    return out, nil
}

func GetMPAlertEmailForTbl(db *gorm.DB,
                           txtFilter string,
                           pageNumber int,
                           pageLimit int) ([][]string, error) {
    var alertEmails []MPAlertEmail
    var dbResult *gorm.DB
    offset := (pageNumber - 1) * pageLimit
    if (txtFilter != "") && (txtFilter != "null") {
        sql := "email LIKE ?"
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        dbResult = db.Where(
          sql,
          txtLikeFilter).Limit(pageLimit).Offset(offset).Find(&alertEmails)
    } else {
        dbResult = db.Limit(pageLimit).Limit(pageLimit).Offset(offset).Find(&alertEmails)
    }
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    out := [][]string{}
    for k := range alertEmails {
        temp := []string{alertEmails[k].Email}
        out = append(out, temp)
    }
    return out, nil
}

func GetMPAlertEmail(db *gorm.DB,
                     email string) (MPAlertEmail, bool, error) {
    var out MPAlertEmail
    dbResult := db.Where("email = ?", email).First(&out)
    return out,(dbResult.RowsAffected == 1), dbResult.Error
}

func InsMPAlertEmail(db *gorm.DB,
                     email string) error {
    var alertEmail MPAlertEmail
    alertEmail.Email = email
    result := db.Create(alertEmail)
    if result.Error != nil {
        return result.Error
    }
    return nil
}


func DelMPAlertEmails(db *gorm.DB, emails []string) error {
    var alertEmails []MPAlertEmail
    if err := db.Where(
      "email in (?)",
      emails).Find(&alertEmails).Error; err != nil {
        return err
    }
    result := db.Where(
      "email in (?)",
      emails).Delete(&alertEmails)
    if result.Error != nil {
        return result.Error
    }
    return nil
}
