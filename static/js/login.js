$(document).ready(function(){
    $("#loginFrm").submit(function(e){
        $.LoadingOverlay("show");
        e.preventDefault();
        e.stopImmediatePropagation();
        var frmDat = {
          username: $("#username").val(),
          password: $("#password").val(),
        };
        $.ajax({
          type: "POST",
          url: "/login",
          data: JSON.stringify(frmDat),
          contentType: "application/json",
          dataType: "json",
          cache: false,
          processData: false,
          success:
            function(output, status, xhr) {
                $.LoadingOverlay("hide");
                switch(output.Status) {
                  case 1:
                      window.location.href = "/";
                    break;
                  case 2:
                      OpenAlertModal(
                        "LOGIN FAILED",
                        "Username/Password not valid!");
                    break;
                }
            },
          error:
            function (xhr, textStatus, errorMessage) {
                $.LoadingOverlay("hide");
                OpenAlertModal(textStatus, xhr.responseText);
            }
        });
    });
});
