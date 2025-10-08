package lang

import (
    "fmt"
    "strings"
)

const InternalServerWebError = "Internal Server Error"
const UnknownWebError = "Unknown"
const DbConnErrTitle = "Database Connection Error"
const CommitErrTitle = "Commit Error"
const HttpBodyErrTitle = "HTTP Body Error"
const TokenErrTitle = "Token Error"
const CnvJsonErrTitle = "Get JSON Body Error"
const LockedAccountTitle = "Locked Account"

func Help() string {
    return `usage: hermawan-monitora-webadmin [-h] [--env ENV]

Hermawan-Monitora is a monitoring application, and the main feature is for monitoring services' health.
This is the web admin service.
For more description and help please check:
https://github.com/ming-hermawan/hermawan-monitora-manual/blob/master/hemawan-monitora-manual.pdf

optional arguments:
  -h, --help  show this help message and exit
  -env        env file path
`
}

func RedisSetErr(lbl string, errMsg string) string {
    return fmt.Sprintf(
      "Can't Set %s in Redis, %s",
      lbl,
      errMsg)
}

func MasterDataRowsCountErrTitle(lbl string) string {
    return fmt.Sprintf(
      "Select Count of Master Rows %s Error",
      lbl)
}

func MasterRowsErrTitle(lbl string) string {
    return fmt.Sprintf(
      "Select Master Rows %s Error",
      lbl)
}

func DetailRowErrTitle(key string, tblName string) string {
    return fmt.Sprintf(
      "Get %s in %s Error",
      key,
      tblName)
}

func SelDbErrTitle(lbl string) string {
    return fmt.Sprintf(
      "Select %s Error",
      lbl)
}

func InsDbErrTitle(lbl string) string {
    return fmt.Sprintf(
      "Insert into %s Error",
      lbl)
}

func UpdDbErrTitle(lbl string) string {
    return fmt.Sprintf(
      "Update %s Error",
      lbl)
}

func DelDbErrTitle(lbl string) string {
    return fmt.Sprintf(
      "Delete %s Error",
      lbl)
}

func GetRedisErrTitle(lbl string) string {
    return fmt.Sprintf(
      "Error read %s from Redis!",
      lbl)
}

func UsedCannotBeDeleted(groups []string) string {
    return fmt.Sprintf(
      "%s is used, so it can't be deleted!",
      strings.Join(groups, ","))
}

func SetRedisErrTitle(lbl string) string {
    return fmt.Sprintf(
      "Error set %s in Redis!",
      lbl)
}

func CannotLoginBecauseWrongPassword3Times(lbl string) string {
    return fmt.Sprintf(
      "%s can't Login because wrong password 3 times.",
      lbl)
}
