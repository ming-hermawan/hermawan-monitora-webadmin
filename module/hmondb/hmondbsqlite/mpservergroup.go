package hmondbsqlite

import (
  "errors"
  "fmt"
  "gorm.io/gorm"
)


// DB Struc

type MPServerGroup struct {
    ServerGroup string `gorm:"primaryKey"`
    SortNumber int
}


// Func

func IsMPServerGroupExists(db *gorm.DB,
                         serverGroup string) (bool, error) {
    var cnt int64
    sql := "server_group = ?"
    dbResult := db.Model(&MPServerGroup{}).Where(
      sql,
      serverGroup).Count(&cnt)
    if dbResult.Error != nil {
        return false, dbResult.Error
    }
    if (cnt == 0) {
        return false, nil
    } else {
        return true, nil
    }
}

func GetMPServerGroupsForDDL(db *gorm.DB) ([]string, error) {
    var serverGroups []MPServerGroup
    dbResult := db.Find(&serverGroups)
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    out := []string{}
    for k := range serverGroups {
        out = append(out, serverGroups[k].ServerGroup)
    }
    return out, nil
}

func GetCntOfMPServerGroupForTbl(db *gorm.DB,
                                 txtFilter string) (int64, error) {
    var out int64
    var dbResult *gorm.DB
    if (txtFilter != "") && (txtFilter != "null") {
        sql := "server_group LIKE ?"
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        dbResult = db.Model(&MPServerGroup{}).Where(
          sql,
          txtLikeFilter).Count(&out)
    } else {
        dbResult = db.Model(&MPServerGroup{}).Count(&out)
    }
    if dbResult.Error != nil {
        return -1, dbResult.Error
    }
    return out, nil
}

func GetMPServerGroupForTbl(db *gorm.DB,
                            txtFilter string,
                            pageNumber int,
                            pageLimit int) ([][]string, error) {
    var serverGroups []MPServerGroup
    var dbResult *gorm.DB
    offset := (pageNumber - 1) * pageLimit
    if (txtFilter != "") && (txtFilter != "null") {
        sql := "server_group LIKE ?"
        txtLikeFilter := fmt.Sprintf("%%%s%%", txtFilter)
        dbResult = db.Where(
          sql,
          txtLikeFilter).Limit(pageLimit).Offset(offset).Find(&serverGroups)
    } else {
        dbResult = db.Limit(pageLimit).Limit(pageLimit).Offset(offset).Find(&serverGroups)
    }
    if dbResult.Error != nil {
        return nil, dbResult.Error
    }
    out := [][]string{}
    for k := range serverGroups {
        temp := []string{serverGroups[k].ServerGroup}
        out = append(out, temp)
    }
    return out, nil
}

func GetMPServerGroup(db *gorm.DB,
                      servergroup string) (MPServerGroup, bool, error) {
    var out MPServerGroup
    dbResult := db.Where("server_group = ?", servergroup).First(&out)
    return out,(dbResult.RowsAffected == 1), dbResult.Error
}

func InsMPServerGroup(db *gorm.DB,
                      servergroup string,
                      sortnumber int) error {
    var serverGroup MPServerGroup
    serverGroup.ServerGroup = servergroup
    serverGroup.SortNumber = sortnumber
    result := db.Create(serverGroup)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func BulkInsMPServerGroup(db *gorm.DB, serverGroups []MPServerGroup) error {
    result := db.Create(serverGroups)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected >= 0 {
        return nil
    } else {
        return errors.New("No Rows Affected")
    }
}

func UpdMPServerGroup(db *gorm.DB,
                      servergroup string,
                      sortnumber int) error {
    var serverGroup MPServerGroup
    if err := db.Where("server_group = ?", servergroup).Find(&serverGroup).Error; err != nil {
        return err
    }
    serverGroup.SortNumber = sortnumber
    result := db.Where("server_group = ?", servergroup).Save(&serverGroup)
    if result.Error != nil {
        return result.Error
    }
    return nil
}


func DelMPServerGroups(db *gorm.DB, servergroups []string) error {
    isServerGroupUsedInServers, err := IsServerGroupUsedInServers(servergroups)
    if err != nil {
        return err
    }
    if isServerGroupUsedInServers {
        return errors.New("Can't be deleted")
    }
    var serverGroups []MPServerGroup
    if err := db.Where(
      "server_group in (?)",
      servergroups).Find(&serverGroups).Error; err != nil {
        return err
    }
    result := db.Where(
      "server_group in (?)",
      servergroups).Delete(&serverGroups)
    if result.Error != nil {
        return result.Error
    }
    return nil
}
