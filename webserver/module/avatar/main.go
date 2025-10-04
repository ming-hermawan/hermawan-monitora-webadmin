package avatar

import (
    "os"
    "fmt"
    "path/filepath"
    "hermawan-monitora/hmonglobal"
    "hermawan-monitora/module/hmonenv"
)


func GetAvatarFromFile(username string) ([]byte, error) {
    var err error
    var data []byte
    sysFileLocation := filepath.Join(hmonenv.GetPicDirPath(),
                                     fmt.Sprintf("%s.jpg", username))
    _, err = os.Stat(sysFileLocation)
    if err != nil {
        return nil, err
    }
    data, err = os.ReadFile(sysFileLocation)
    return data, err
}

func GetDefaultAvatarFromFile() ([]byte, error) {
    var err error
    var data []byte
    if err != nil {
        return nil, err
    }
    sysFileLocation := filepath.Join(hmonglobal.StaticDirPath,
				     "/img/icon/icons8-male-user-96.png")
    _, err = os.Stat(sysFileLocation)
    if err != nil {
        return nil, err
    }
    data, err = os.ReadFile(sysFileLocation)
    return data, err
}
