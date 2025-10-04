package login

import (
    "encoding/json"
    "hermawan-monitora/module/hmonredis"
)


// PRIVATE

func setUserFailedLoginInfo(username string,
                            timeUnix int64,
                            attemp int) {
    val := map[string]any {
      "time": timeUnix,
      "attemp": attemp}
    jsonInBytes, _ := json.Marshal(&val)
    redisKey := hmonredis.GetFailedLoginKey(username)
    hmonredis.SetRawWithExpired(
      redisKey,
      jsonInBytes)
}

func getUserFailedLoginInfo(username string) (int64, int) {
    key, err := hmonredis.Get(hmonredis.GetFailedLoginKey(username))
    if err != nil {
        return 0, 0
    }
    if key == "" {
        return 0, 0
    }
    var reading map[string]interface{}
    if err := json.Unmarshal([]byte(key), &reading); err != nil {
        return 0, 0
    }
    timeUnix := int64(reading["time"].(float64))
    attemp := int(reading["attemp"].(float64))
    return timeUnix, attemp
}

func delUserFailedLoginInfo(username string) {
    hmonredis.Del(hmonredis.GetFailedLoginKey(username))
}
