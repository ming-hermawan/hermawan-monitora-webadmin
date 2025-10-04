$(document).ready(function(){
    RelocationUrl = "/dashboard";

    var countdown = 300
    var x = setInterval(function() {
        if(countdown == 0) {
            location.reload(true);
        } else {
            countdown--;
            strCountdown = countdown.toString();
            switch(strCountdown.length) {
                case 1:
                    $("#countDown1").html("");
                    $("#countDown2").html("");
                    $("#countDown3").html(strCountdown[0]);
                break;
                case 2:
                    $("#countDown1").html("");
                    $("#countDown2").html(strCountdown[0]);
                    $("#countDown3").html(strCountdown[1]);
                break;
                case 3:
                    $("#countDown1").html(strCountdown[0]);
                    $("#countDown2").html(strCountdown[1]);
                    $("#countDown3").html(strCountdown[2]);
                break;
            }
            // 21.6 -> 7.2
            // 22 -> 8
            if($(window).width() > 425) {
                var countDown1MarginLeftBefore = 273.5;
                var countDown2MarginLeftBefore = 281.5;
                var countDown3MarginLeftBefore = 289.5;
                var topBefore = 4;
            } else {
                var countDown1MarginLeftBefore = 0;
                var countDown2MarginLeftBefore = 8;
                var countDown3MarginLeftBefore = 16;
                var topBefore = 28;
            }
            let countDown1MarginLeftAfter = countDown1MarginLeftBefore - 14;
            let countDown2MarginLeftAfter = countDown2MarginLeftBefore - 7;
            let countDown3MarginLeftAfter = countDown3MarginLeftBefore + 6;
            let topAfter = topBefore - 26;
            $("#countDown1").css({
              marginLeft: countDown1MarginLeftAfter + "px",
              marginTop: topAfter + "px",
              fontSize: "36px",
              width: "22px"});
            $("#countDown1").animate({
              marginLeft: countDown1MarginLeftBefore + "px",
              marginTop: topBefore + "px",
              fontSize: "12px",
              width: "8px"}, 300);
            $("#countDown2").css({
              marginLeft: countDown2MarginLeftAfter + "px",
              marginTop: topAfter + "px",
              fontSize: "36px",
              width: "22px"});
            $("#countDown2").animate({
              marginLeft: countDown2MarginLeftBefore + "px",
              marginTop: topBefore + "px",
              fontSize: "12px",
              width: "8px"}, 300);
            $("#countDown3").css({
              marginLeft: countDown3MarginLeftAfter + "px",
              marginTop: topAfter + "px",
              fontSize: "36px",
              width: "22px"});
            $("#countDown3").animate({
              marginLeft: countDown3MarginLeftBefore + "px",
              marginTop: topBefore + "px",
              fontSize: "12px",
              width: "8px"}, 300);
        }
    }, 1000);

    $("#oldpassword").on( "change", function() {
        $("#message").html("Input New Password");
        $("#newpassword").prop('disabled', false);
        $("#newpassword").focus();
    });
    $("#newpassword").on( "change", function() {
        $("#message").html("Input New Password Confirmation (Same as New Password)");
        if(this.value.localeCompare($("#oldpassword").val()) === 0) {
            $("#message").html("New Password must different with Old Password!");
        } else {
            $("#newpasswordconfirm").prop('disabled', false);
            $("#newpasswordconfirm").focus();
        }
    });
    $("#newpasswordconfirm").on( "change", function() {
        if(this.value.localeCompare($("#newpassword").val()) === 0) {
            $("#buttonChangePassword").prop('disabled', false);
            $("#buttonChangePassword").focus();
            $("#message").html("");
        } else {
            $("#message").html("New Password Confirmation not same as New Password!");
        }
    });
    $("#frmInp").submit(function(e){
        $.LoadingOverlay("show");
        e.preventDefault();
        e.stopImmediatePropagation();
        let oldpassword = $("#oldpassword").val();
        let newpassword = $("#newpasswordconfirm").val();
        let frmDat = {
          oldpassword: oldpassword,
          newpassword: newpassword};
        $.ajax({
          type: "PUT",
          url: MAIN_URL,
          data: JSON.stringify(frmDat),
          contentType: "application/json",
          dataType: "json",
          cache: false,
          processData: false,
          success:
            function(output, status, xhr) {
                $.LoadingOverlay("hide");
                if(output.Status == 1) {
                    OpenAlertModal("SAVED", "Password Updated");
                } else {
                    OpenAlertModal("FAILED", output.Message);
                }
            }
        , error:
            function (xhr, textStatus, errorMessage) {
                $.LoadingOverlay("hide");
                OpenAlertModal(textStatus,  xhr.responseText);
            }
        });
    });

    $("#message").html("Input Old Password first");

});
