package hmondbsqlite

import (
  "fmt"
  "gorm.io/gorm"
)


// DB Struc

type UserGroup struct {
  UserGroup string `gorm:"primaryKey"`
  SortNumber int
}

type UserGroupMenu struct {
  UserGroup string `gorm:"primaryKey"`
  MenuCode string `gorm:"primaryKey"`
}


// Func

func GetUserGroupsForDDL(db *gorm.DB) ([]string, error) {
    var userGroups []UserGroup
    dbResult := db.Find(&userGroups)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    out := []string{}
    for k := range userGroups {
        out = append(out, userGroups[k].UserGroup)
    }
    return out, nil
}

func GetCntOfUserGroupsForTbl(db *gorm.DB,
                              txtFilter string) (int64, error) {
    var out int64
    var dbResult *gorm.DB
    if (txtFilter != "") && (txtFilter != "null") {
        sql := "user_group LIKE ?";
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        dbResult = db.Model(&UserGroup{}).Where(
          sql,
          txtLikeFilter).Count(&out)
    } else {
        dbResult = db.Model(&UserGroup{}).Count(&out)
    }
    if dbResult.Error != nil {
        return -1, dbResult.Error
    }
    return out, nil
}

func GetUserGroupsForTbl(db *gorm.DB,
                         txtFilter string,
                         pageNumber int,
                         pageLimit int) ([][]string, error) {
    var userGroups []UserGroup
    var dbResult *gorm.DB
    offset := (pageNumber - 1) * pageLimit
    if (txtFilter != "") && (txtFilter != "null") {
        sql := "user_group LIKE ?"
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        dbResult = db.Where(
          sql,
          txtLikeFilter).Limit(pageLimit).Offset(offset).Find(&userGroups)
    } else {
        dbResult = db.Limit(pageLimit).Offset(offset).Find(&userGroups)
    }
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    out := [][]string{}
    for k := range userGroups {
        temp := []string{userGroups[k].UserGroup}
        out = append(out, temp)
    }
    return out, nil
}

func GetUserGroup(db *gorm.DB,
                  usergroup string) (UserGroup, bool, error) {
    var out UserGroup
    dbResult := db.Where("user_group = ?", usergroup).First(&out)
    return out,(dbResult.RowsAffected == 1), dbResult.Error
}

func GetUserGroupMenus(db *gorm.DB,
                       userGroup string) ([]string, error) {
    var userGroupMenus []UserGroupMenu
    dbResult := db.Where("user_group = ?", userGroup).Find(&userGroupMenus)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    out := []string{}
    for k := range userGroupMenus {
        out = append(out, userGroupMenus[k].MenuCode)
    }
    return out, nil
}

func InsUserGroup(db *gorm.DB,
                  usergroup string,
                  sortnumber int) error {
    var userGroup UserGroup
    userGroup.UserGroup = usergroup
    userGroup.SortNumber = sortnumber
    result := db.Create(userGroup)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func UpdUserGroup(db *gorm.DB,
                  usergroup string,
                  sortnumber int) error {
    var userGroup UserGroup
    if err := db.Where("user_group = ?", usergroup).First(&userGroup).Error; err != nil {
        return err
    }
    userGroup.SortNumber = sortnumber
    result := db.Where("user_group = ?", usergroup).Save(&userGroup)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func DelUserGroups(db *gorm.DB,
                   usergroups []string) (bool, error) {
    var err error
    var isUserGroupUsedInServers bool
    isUserGroupUsedInServers, err = IsUserGroupUsedInUsers(db, usergroups)
    if err != nil {
        return false, err
    }
    if isUserGroupUsedInServers {
        return true, nil
    }
    var userGroups []UserGroup
    err = db.Where(
      "user_group in (?)",
      userGroups).Find(&userGroups).Error;
    if err != nil {
        return false, err
    }
    dbResult := db.Where(
      "user_group in (?)",
      usergroups).Delete(&userGroups)
    if dbResult.Error != nil {
        return false, dbResult.Error
    }
    return false, nil
}

func DelUserGroupMenus(db *gorm.DB, usergroups []string) error {
    var userGroupMenus []UserGroupMenu
    if err := db.Where(
      "user_group in (?)",
      userGroupMenus).Find(&userGroupMenus).Error; err != nil {
        return err
    }
    dbResult := db.Where("user_group in (?)", userGroupMenus).Delete(&userGroupMenus)
    if dbResult.Error != nil {
        return dbResult.Error
    }
    return nil
}

func UpdUserGroupMenus(db *gorm.DB, usergroup string, menus []string) error {
    dbResult := db.Where("user_group = ?", usergroup).Delete(&UserGroupMenu{})
    if dbResult.Error != nil {
        return dbResult.Error
    }
    for _, menu := range menus {
        if menu != "" {
            userGroupMenus := &UserGroupMenu {UserGroup: usergroup, MenuCode: menu}
            dbResult := db.Create(userGroupMenus)
            if dbResult.Error != nil {
                return dbResult.Error
            }
        }
    }
    return nil
}
