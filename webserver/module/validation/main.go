package validation

import (
    "fmt"
    "regexp"
    "time"
    "net/http"
    "net/mail"
    "hermawan-monitora/hmonglobal"
    "hermawan-monitora/webserver/module/httpresponse"
    dbo "hermawan-monitora/module/hmondb/hmondbsqlite"
)

var (
    regexUpperCase *regexp.Regexp
    regexLowerCase *regexp.Regexp
    regexNumber *regexp.Regexp
    regexPwdSpecialChar *regexp.Regexp
    regexDate *regexp.Regexp
)


// PUBLIC

// const

const errorValidationTitle = "Parameter Validations Error!"

// func

func StatDbParamValidation(w http.ResponseWriter,
                           params map[string]interface{}) (string, bool) {
    out, found := params["statDb"].(string)
    if !found {
        msg := fmt.Sprintf("statDb not found!")
        httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
        return "", false
    }
    if((out != "INS") && (out != "UPD")) {
        msg := fmt.Sprintf("statDb value is not 'INS' or 'UPD'!")
        httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
        return "", false
    }
    return out, true
}

func StrParamValidation(w http.ResponseWriter,
                        params map[string]interface{},
                        key string,
                        lbl string,
                        mandatory bool,
                        regexPattern string) (string, bool) {
    out, found := params[key].(string)
    if !found && mandatory {
        msg := fmt.Sprintf("%s not found!", lbl)
        httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
        return "", false
    }
    if regexPattern != "" {
        regex, _ := regexp.Compile(regexPattern)
        isMatch := regex.MatchString(out)
        if !isMatch {
            msg := fmt.Sprintf("%s value is not valid!", lbl)
            httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
            return "", false
        }
    }
    return out, true
}

func DateParamValidation(w http.ResponseWriter,
                         val string,
                         lbl string,
                         mandatory bool) (time.Time, bool) {
    isMatch := regexDate.MatchString(val)
    if !isMatch {
        msg := fmt.Sprintf("%s is not Date!", lbl)
        httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
        return time.Time{}, false
    }
    strTime := fmt.Sprintf("%s 00:00:00", val)
    out, err := time.Parse(hmonglobal.DateLayout, strTime)
    if err != nil {
        httpresponse.ErrResponseWs(w, "Error Parsing Date!", err.Error())
        return time.Time{}, false
    }
    return out, true
}

func IntParamValidation(w http.ResponseWriter,
                        params map[string]interface{},
                        key string,
                        lbl string,
                        mandatory bool) (int, bool) {
    out, found := params[key].(float64)
    if !found {
        msg := fmt.Sprintf("%s not found!", lbl)
        httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
        return -1, false
    }
    return int(out), true
}

func EmailParamvalidation(w http.ResponseWriter,
                          params map[string]interface{},
                          key string,
                          lbl string,
                          mandatory bool) (string, bool) {
    out, found := params[key].(string)
    if !found {
        if mandatory {
            msg := fmt.Sprintf("%s not found!", lbl)
            httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
            return "", false
        }
        return "", true
    }
    _, err := mail.ParseAddress(out)
    if err != nil {
        msg := fmt.Sprintf("%s value is not email!", lbl)
        httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
        return "", false
    }
    return out, true
}


func PasswordValidation(w http.ResponseWriter,
                        params map[string]interface{}) (string, string, bool) {
    oldpassword, oldpasswordFound := StrParamValidation(
      w,
      params,
      "oldpassword",
      "Old Password",
      true,
      hmonglobal.RegexPwd)
    if !oldpasswordFound {
        return "", "", false
    }
    newpassword, newpasswordFound := StrParamValidation(
      w,
      params,
      "newpassword",
      "New Password",
      true,
      hmonglobal.RegexPwd)
    if !newpasswordFound {
        return "", "", false
    }
    if oldpassword == newpassword {
        http.Error(
          w,
          "Old password must different with new password!",
          http.StatusInternalServerError)
        return "", "", false
    }
    settings := dbo.GetSetting()
    var forceStrongPassword int64
    if settings.ForceStrongPassword.Valid {
        forceStrongPassword = settings.ForceStrongPassword.Int64
    } else {
        forceStrongPassword = 0
    }
    if forceStrongPassword == 1 {
        // Upper Case
        isMatch1 := regexUpperCase.MatchString(newpassword)
        if !isMatch1 {
            msg := "Password doesn't have an uppercase alphabet."
            httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
            return "", "", false
        }
        // Lower Case
        isMatch2 := regexLowerCase.MatchString(newpassword)
        if !isMatch2 {
            msg := "Password doesn't have a lowercase character."
            httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
            return "", "", false
        }
        // Number
        isMatch3 := regexNumber.MatchString(newpassword)
        if !isMatch3 {
            msg := "Password doesn't have a number."
            httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
            return "", "", false
        }
        // Special Characters
        isMatch4 := regexPwdSpecialChar.MatchString(newpassword)
        if !isMatch4 {
            msg := "Password doesn't have a number."
            httpresponse.ErrResponseWs(w, errorValidationTitle, msg)
            return "", "", false
        }
    }
    return oldpassword, newpassword, true
}

func init() {
    regexUpperCase, _ = regexp.Compile(hmonglobal.RegexUpperCase)
    regexLowerCase, _ = regexp.Compile(hmonglobal.RegexLowerCase)
    regexNumber, _ = regexp.Compile(hmonglobal.RegexNumber)
    regexPwdSpecialChar, _ = regexp.Compile(hmonglobal.RegexSpecialCharForPwd)
    regexDate, _ = regexp.Compile(hmonglobal.RegexDate)
}
