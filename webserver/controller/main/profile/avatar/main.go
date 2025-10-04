package avatar

import (
    "net/http"
    "strings"
    "fmt"
    "log"
    "os"
    "io"
    "strconv"
    "encoding/base64"
    "path/filepath"
    "encoding/json"
    "hermawan-monitora/module/hmonenv"
    "hermawan-monitora/webserver/module/httpresponse"
    "hermawan-monitora/webserver/module/avatar"
)


// PRIVATE

func processAddAvatar(w http.ResponseWriter, r *http.Request, username string) {
    // Maximum upload of 10 MB files
    r.ParseMultipartForm(10 << 20)

    // Get handler for filename, size and headers
    file, handler, err := r.FormFile("fileAvatar")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    fileBytes, err := io.ReadAll(file)
    if err != nil {
        // Handle error
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    log.Println(fmt.Printf("Bytes:\n%v\n", fileBytes))

    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    workDirLocation, err1 := os.Getwd()
    if err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
        return
    }
    sysFileLocation := filepath.Join(
      workDirLocation,
      hmonenv.GetPicDirPath(),
      fmt.Sprintf("%s.jpg", username))
    sysDirLocation := filepath.Join(workDirLocation, hmonenv.GetPicDirPath())
    os.MkdirAll(sysDirLocation, 0755)
    log.Println(fmt.Printf("sysFileLocation = %s\n", sysFileLocation))

    dst, err3 := os.Create(sysFileLocation)
    if err3 != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer dst.Close()
    _, err = file.Seek(0, io.SeekStart)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    _, err = io.Copy(dst, file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    str := base64.StdEncoding.EncodeToString(fileBytes)
    jsonInBytes, _ := json.Marshal(map[string]string {
      "image": str})
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}


func processDelAvatar(w http.ResponseWriter, r *http.Request, username string) {
    var params map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    log.Println(fmt.Printf("DEL AVATAR params:\n%+v\n", params))
    workDirLocation, err1 := os.Getwd()
    if err1 != nil {
        http.Error(w, err1.Error(), http.StatusInternalServerError)
        return
    }
    sysFileLocation := filepath.Join(
      workDirLocation,
      hmonenv.GetPicDirPath(),
      fmt.Sprintf("%s.jpg", username))
    _, err2 := os.Stat(sysFileLocation)
    if err2 != nil {
        http.Error(w, err2.Error(), http.StatusInternalServerError)
        return
    }
    err3 := os.Remove(sysFileLocation)
    if err3 != nil {
        http.Error(w, err3.Error(), http.StatusInternalServerError)
        return
    }
    jsonInBytes, _ := json.Marshal(map[string]string {
      "status": "OK"})
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}

func get(w http.ResponseWriter, username string) {
    buffer, err := avatar.GetAvatarFromFile(username)
    if err != nil {
        log.Println(fmt.Sprintf("ERROR %s\n", err.Error()))
        buffer2, _ := avatar.GetDefaultAvatarFromFile()
        w.Header().Set("Content-Type", "image/png")
        w.Header().Set("Content-Length", strconv.Itoa(len(buffer2)))
        if _, err := w.Write(buffer2); err != nil {
            log.Println(fmt.Printf("ERROR WRITE AVATAR %s\n", err))
        }
    }
    if _, err := w.Write(buffer); err != nil {
        log.Println(fmt.Printf("ERROR WRITE AVATAR %s\n", err))
    }
    w.Header().Set("Content-Type", "image/jpeg")
    w.Header().Set("Content-Length", strconv.Itoa(len(buffer)))
}

func post(w http.ResponseWriter, username string) {
    avatarStr := ""
    avatarBytes, err2 := avatar.GetAvatarFromFile(username)
    if err2 == nil {
        avatarStr =  base64.StdEncoding.EncodeToString(avatarBytes)
    }
    result := map[string]any{
      "avatar": avatarStr}
    jsonInBytes, _ := json.Marshal(&result)
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonInBytes)
}

func put(w http.ResponseWriter, r *http.Request, username string) {
    contentType := r.Header.Get("Content-type")
    if strings.Contains(contentType, "multipart/form-data;") {
        processAddAvatar(w, r, username)
    } else {
        processDelAvatar(w, r, username)
    }
}


// PUBLIC

func Process(w http.ResponseWriter, r *http.Request, username string) {
    switch r.Method {
      case "GET":
          get(w, username)
      case "POST":
          post(w, username)
      case "PUT":
          put(w, r, username)
      default:
          httpresponse.ErrResponseForBadRequest(w)
    }
}
