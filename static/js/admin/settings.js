function onCancel(e) {
    if ($("#forceStrongPassword0").val() == "1") {
        $("#forceStrongPassword").prop("checked", true);
        $("#forceStrongPassword").val("1");
    } else {
        $("#forceStrongPassword").prop('checked', false);
        $("#forceStrongPassword").val("0");
    }
    if ($("#logoutAfter1Hour0").val() == "1") {
        $("#logoutAfter1Hour").prop("checked", true);
        $("#logoutAfter1Hour").val("1");
    } else {
        $("#logoutAfter1Hour").prop('checked', false);
        $("#logoutAfter1Hour").val("0");
    }
    $("#smtphost").val($("#smtphost0").val());
    $("#smptport").val($("#smptport0").val());
    $("#sendername").val($("#sendername0").val());
    $("#authemail").val($("#authemail0").val());
    $("#authpassword").val($("#authpassword0").val());
}


$(document).ready(function(){
    RelocationUrl = "/dashboard";
    $(".h-inp").change(function() {
        $('#btnSave').prop('disabled', false);
    });
    $.ajax({
        url: MAIN_URL,
        type: "POST",
        data: JSON.stringify({}),
        dataType: "json",
        success: function (result) {
            $("#smtphost").val(result.smtphost);
            $("#smtpport").val(result.smtpport);
            $("#sendername").val(result.sendername);
            $("#authemail").val(result.authemail);
            $("#authpassword").val(result.authpassword);
            $("#smtphost0").val(result.smtphost);
            $("#smtpport0").val(result.smtpport);
            $("#sendername0").val(result.sendername);
            $("#authemail0").val(result.authemail);
            $("#authpassword0").val(result.authpassword);
            switch(result.forceStrongPassword) {
                case 0:
                    $("#forceStrongPassword").prop('checked', false);
                break;
                case 1:
                    $("#forceStrongPassword").prop('checked', true);
                break;
            }
            $("#forceStrongPassword0").val(result.forceStrongPassword);
            switch(result.logoutAfter1Hour) {
                case 0:
                    $("#logoutAfter1Hour").prop('checked', false);
                break;
                case 1:
                    $("#logoutAfter1Hour").prop('checked', true);
                break;
            }
            $("#logoutAfter1Hour0").val(result.logoutAfter1Hour);
        }
    });
    $("#frmInp").submit(function(e){
        e.preventDefault();
        smtphost = $("#smtphost").val();
        smtpport = parseInt($("#smtpport").val());
        sendername = $("#sendername").val();
        authemail = $("#authemail").val();
        authpassword = $("#authpassword").val();
        var forceStrongPassword = 0;
        if($("#forceStrongPassword").is(":checked")) {
            forceStrongPassword = 1;
        }
        var logoutAfter1Hour = 0;
        if($("#logoutAfter1Hour").is(":checked")) {
            logoutAfter1Hour = 1;
        }
        var frmDat = {
          "forceStrongPassword": forceStrongPassword,
          "logoutAfter1Hour": logoutAfter1Hour,
          "smtphost": smtphost,
          "smtpport": smtpport,
          "sendername": sendername,
          "authemail": authemail,
          "authpassword": authpassword};
        $.ajax({
          url: "/admin/setting",
          type: "PUT",
          data: JSON.stringify(frmDat),
          dataType: "json",
          success: function(output, status, xhr) {
            if(output.Status == 1) {
                  OpenAlertModal("SUCCESS", "Settings Updated");
            } else {
                 window.alert(result);
            }
          }
        });
    });
    CancelProcess(onCancel);
});
