package hmondbsqlite

import (
  "fmt"
  "gorm.io/gorm"
)


// DB Struc

type User struct {
  Username string `gorm:"primaryKey"`
  Password string
  Name string
  IsAdmin int
  Email string
  UserGroup string
}

type UserSetting struct {
  Username string `gorm:"primaryKey"`
  Language string
  Darkmode int
  Note string
}


// Func

func IsUserGroupUsedInUsers(db *gorm.DB,
                            usergroups []string) (bool, error) {
    var cnt int64
    sql := "user_group IN (?)"
    dbResult := db.Model(&User{}).Where(
      sql,
      usergroups).Count(&cnt)
    if dbResult.Error != nil {
        return false, dbResult.Error
    }
    if (cnt == 0) {
        return false, nil
    } else {
        return true, nil
    }
}

func GetCntOfUserForTbl(db *gorm.DB,
                        grpFilter string,
                        txtFilter string) (int64, error) {
    var out int64
    var dbResult *gorm.DB
    if (txtFilter != "") {
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        if (grpFilter != "") && (grpFilter != "null") {
            sql := "(user_group = ?) AND ((username LIKE ?) OR (name LIKE ?) OR (email LIKE ?))"
            dbResult = db.Model(&[]User{}).Where(
              sql,
              grpFilter,
              txtLikeFilter,
              txtLikeFilter,
              txtLikeFilter).Count(&out)
        } else {
            sql := "(username LIKE ?) OR (name LIKE ?) OR (email LIKE ?)"
            dbResult = db.Model(&[]User{}).Where(
              sql,
              txtLikeFilter,
              txtLikeFilter,
              txtLikeFilter).Count(&out)
        }
    } else if (grpFilter != "") {
        sql := "user_group = ?"
        dbResult = db.Model(&[]User{}).Where(
          sql,
          grpFilter).Count(&out)
    } else {
        dbResult = db.Model(&[]User{}).Count(&out)
    }
    if dbResult.Error != nil {
        return -1, dbResult.Error
    }
    return out, nil
}

func GetUserForTbl(db *gorm.DB,
                   grpFilter string,
                   txtFilter string,
                   pageNumber int,
                   pageLimit int) ([][]string, error) {
    var users []User
    var dbResult *gorm.DB
    offset := (pageNumber - 1) * pageLimit
    if (txtFilter != "") {
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        if (grpFilter != "") && (grpFilter != "null") {
            sql := "(user_group = ?) AND ((username LIKE ?) OR (name LIKE ?) OR (email LIKE ?))"
            dbResult = db.Debug().Where(
              sql,
              grpFilter,
              txtLikeFilter,
              txtLikeFilter,
              txtLikeFilter).Limit(pageLimit).Offset(offset).Find(&users)
        } else {
            sql := "(username LIKE ?) OR (name LIKE ?) OR (email LIKE ?)"
            txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
            dbResult = db.Debug().Where(
              sql,
              txtLikeFilter,
              txtLikeFilter,
              txtLikeFilter).Limit(pageLimit).Offset(offset).Find(&users)
        }
    } else if (grpFilter != "") {
        sql := "user_group = ?"
        dbResult = db.Where(
          sql,
          grpFilter).Limit(pageLimit).Offset(offset).Find(&users)
    } else {
        dbResult = db.Limit(pageLimit).Offset(offset).Find(&users)
    }
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    out := [][]string{}
    for k := range users {
        temp := []string{users[k].Username, users[k].Name, users[k].Email}
        out = append(out, temp)
    }
    return out, nil
}

func GetUser(db *gorm.DB,
             username string) (User, bool, error)  {
    var out User
    dbResult := db.Where("username = ?", username).First(&out)
    return out,(dbResult.RowsAffected == 1), dbResult.Error
}

func InsUser(db *gorm.DB,
             username string,
             password string,
             name string,
             is_admin int,
             email string,
             group string) error {
    var user User
    user.Username = username
    if password != "" {
        user.Password = password
    }
    user.Name = name
    user.IsAdmin = is_admin
    user.Email = email
    user.UserGroup = group
    result := db.Create(user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func UpdUser(db *gorm.DB,
             username string,
             password string,
             name string,
             is_admin int,
             email string,
             group string) error {
    var user User
    if err := db.Where("username = ?", username).Find(&user).Error; err != nil {
        return err
    }
    if password != "" {
        user.Password = password
    }
    user.Name = name
    user.IsAdmin = is_admin
    user.Email = email
    user.UserGroup = group
    result := db.Where("username = ?", username).Save(&user)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func DelUser(db *gorm.DB,
             usernames []string) error {
    var users []User
    if err := db.Where("username in (?)", usernames).Find(&users).Error; err != nil {
        return err
    }
    result := db.Where("username in (?)", usernames).Delete(&users)
    if result.Error != nil {
        return result.Error
    }
    return nil
//    db.Find(&users, paramusers).Delete(&users)
}



func UpdateUserPassword(db *gorm.DB, username string, password string) error {
    var user User
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return err
    }
    user.Password = password
    db.Where("username = ?", username).Save(&user)
    return nil
}

func GetUserSetting(db *gorm.DB, username string) (UserSetting, *gorm.DB) {
    var out UserSetting
    result := db.Where("username = ?", username).First(&out)
    return out, result
}

func SavSetting(db *gorm.DB,
                username string,
		language string,
		darkmode int,
                note string) {
    var userSetting UserSetting
    updateStatus := true
    if err := db.Where("username = ?", username).First(&userSetting).Error; err != nil {
        updateStatus = false
    }
    userSetting.Language = language
    userSetting.Darkmode = darkmode
    userSetting.Note = note
    if updateStatus {
        db.Where("username = ?", username).Save(&userSetting)
    } else {
        userSetting.Username = username
        db.Create(userSetting)
    }
}

/*
func SavUser(db *gorm.DB,
             username string,
             password string,
             name string,
             is_admin int,
             email string,
             group string) {
    var user User
    updateStatus := true
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        updateStatus = false
    }
    if password != "" {
        user.Password = password
    }
    user.Name = name
    user.IsAdmin = is_admin
    user.Email = email
    user.UserGroup = group
    if updateStatus {
        db.Where("username = ?", username).Save(&user)
    } else {
        user.Username = username
        db.Create(user)
    }
}*/
