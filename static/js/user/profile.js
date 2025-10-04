const PROFILE_AVATAR_URL = "/profile/avatar"

function adjustHeight() {
    let documenth = $(document).height();
    let divContentH = $("#divContent").height();
    console.log("divContent height = " + divContentH);
    if ( documenth > 250 ) {
        var padding_top = (documenth - divContentH - 68) / 2;
        $("#divInnerBody").css("padding-top", padding_top+"px")
    }
}


function onCancel() {
    $("#language").val($("#language0").val());
    $("#note").val($("#note0").val());
    var checkbox = $("#darkmode");
    if($("#darkmode0").val() != checkbox.val()) {
        if ($("#darkmode0").val() == "1") {
            checkbox.val("1");
            checkbox.prop("checked", true);
        } else {
            checkbox.val("0");
            checkbox.prop("checked", false);
        }
        checkbox.trigger("change");
    }
}


$(document).ready(function(){
    adjustHeight();
    $(window).resize(function(){
        adjustHeight();
    });
    RelocationUrl = "/dashboard";
    $.ajax({
        url: MAIN_URL,
        type: "POST",
        data: JSON.stringify({}),
        dataType: "json",
        success: function (result) {
            console.log(result);
            $("#username").html(result.username);
            $("#name").html(result.name);
            $("#email").html(result.email);
            $("#status").html(result.status);
            $("#divDarkMode").val(result.darkmode);
            $("#darkmode0").val(result.darkmode);
            $("#darkmode").val(result.darkmode);
            $("#note0").val(result.note);
            $("#note").val(result.note);
            if(result.darkmode == 1) {
                checkbox = $("#darkmode");
                checkbox.val("1");
                checkbox.prop("checked", true);
                checkbox.trigger("change");
            }
            $("#avatar").val(result.avatar);
            if(result.avatar.length === 0) {
                $("#imgAvatar").prop("src", "/static/img/icon/icons8-male-user-96.png");
            } else {
                $("#imgAvatar").prop("src", `data:image/jpg;base64,${result.avatar}`);
                $("#avatar").val("1");
            }
            $("#btnSave").prop('disabled', true);
            $("#btnCancel").prop('disabled', true);
        }
    });
    $("#darkmode").change(function () {
        if($(this).is(':checked')) {
            $("#divDarkMode > div").css("background-color", "#000000");
            $("#divDarkMode > div").removeClass("div_anim_sky");
            $("#divDarkMode > div").addClass("div_anim_night");
            $("#divDarkModeBtn").animate({marginLeft: '70px'}, 500);
            $("#divCloud").animate({marginLeft: '-300px'}, 500);
            $("#divSun").animate({opacity: '0'}, 500);
            $("#divMoon").animate({opacity: '1'}, 500);
            $("#imgStar1").animate({opacity: 1}, 1000);
            $("#imgStar2").animate({opacity: 1}, 1000);
            $("#imgStar3").animate({opacity: 1}, 1000);
            $("#darkmode").val("1");



            $("#divBody").removeClass("div-body-lightmode");
            $("#divBody").toggleClass("div-body-darkmode");
            $("#divTopRight > div:nth-child(1)").removeClass("div-top-right-1-outer-lightmode");
            $("#divTopRight > div:nth-child(1)").toggleClass("div-top-right-1-outer-darkmode");
            $("#divTopRight > div:nth-child(1) > div").removeClass("div-top-right-1-inner-lightmode");
            $("#divTopRight > div:nth-child(1) > div").toggleClass("div-top-right-1-inner-darkmode");
            $("#divTopRight > div:nth-child(2)").removeClass("div-top-right-2-outer-lightmode");
            $("#divTopRight > div:nth-child(2)").toggleClass("div-top-right-2-outer-darkmode");
            $("#divTopRight > div:nth-child(2) > div").removeClass("div-top-right-2-inner-lightmode");
            $("#divTopRight > div:nth-child(2) > div").toggleClass("div-top-right-2-inner-darkmode");
            $("#divBottom").removeClass("div-bottom-outer-lightmode");
            $("#divBottom").toggleClass("div-bottom-outer-darkmode");
            $("#divBottom > div").removeClass("div-bottom-inner-lightmode");
            $("#divBottom > div").toggleClass("div-bottom-inner-darkmode");
            $("#divTopLeft > div:nth-child(1)").removeClass("div-top-left-1-outer-lightmode");
            $("#divTopLeft > div:nth-child(1)").toggleClass("div-top-left-1-outer-darkmode");
            $("#divTopLeft > div:nth-child(2)").removeClass("div-top-left-2-inner-lightmode");
            $("#divTopLeft > div:nth-child(2)").toggleClass("div-top-left-2-inner-darkmode");
            $("#divTopLeft > div:nth-child(3)").removeClass("div-top-left-3-outer-lightmode");
            $("#divTopLeft > div:nth-child(3)").toggleClass("div-top-left-3-outer-darkmode");
            $("#divTopLeft > div:nth-child(4)").removeClass("div-top-left-4-inner-lightmode");
            $("#divTopLeft > div:nth-child(4)").toggleClass("div-top-left-4-inner-darkmode");
            $("#divTopRight > div:nth-child(3)").removeClass("div-top-right-3-outer-lightmode");
            $("#divTopRight > div:nth-child(3)").toggleClass("div-top-right-3-outer-darkmode");
            $("#divTopRight > div:nth-child(3) > div").removeClass("div-top-right-3-inner-lightmode");
            $("#divTopRight > div:nth-child(3) > div").toggleClass("div-top-right-3-inner-darkmode");
            $("#divDarkMode").removeClass("div-dark-mode-dark-outer-lightmode");
            $("#divDarkMode").toggleClass("div-dark-mode-dark-outer-darkmode");
            $("#divDarkMode > div").removeClass("div-dark-mode-dark-inner-lightmode");
            $("#divDarkMode > div").toggleClass("div-dark-mode-dark-inner-darkmode");
        } else {
            $("#divDarkMode > div").css("background-color", "#0080ff");
            $("#divDarkMode > div").removeClass("div_anim_night");
            $("#divDarkMode > div").addClass("div_anim_sky");
            $("#divDarkModeBtn").animate({marginLeft: '6px'}, 500);
            $("#divCloud").animate({marginLeft: '0'}, 500);
            $("#divMoon").animate({opacity: '0'}, 500);
            $("#divSun").animate({opacity: '1'}, 500);
            $("#imgStar1").animate({opacity: 0}, 1000);
            $("#imgStar2").animate({opacity: 0}, 1000);
            $("#imgStar3").animate({opacity: 0}, 1000);
            $('#darkmode').val("0");



            $("#divBody").removeClass("div-body-darkmode");
            $("#divBody").toggleClass("div-body-lightmode");
            $("#divTopRight > div:nth-child(1)").removeClass("div-top-right-1-outer-darkmode");
            $("#divTopRight > div:nth-child(1)").toggleClass("div-top-right-1-outer-lightmode");
            $("#divTopRight > div:nth-child(1) > div").removeClass("div-top-right-1-inner-darkmode");
            $("#divTopRight > div:nth-child(1) > div").toggleClass("div-top-right-1-inner-lightmode");
            $("#divTopRight > div:nth-child(2)").removeClass("div-top-right-2-outer-darkmode");
            $("#divTopRight > div:nth-child(2)").toggleClass("div-top-right-2-outer-lightmode");
            $("#divTopRight > div:nth-child(2) > div").removeClass("div-top-right-2-inner-darkmode");
            $("#divTopRight > div:nth-child(2) > div").toggleClass("div-top-right-2-inner-lightmode");
            $("#divBottom").removeClass("div-bottom-outer-darkmode");
            $("#divBottom").toggleClass("div-bottom-outer-lightmode");
            $("#divBottom > div").removeClass("div-bottom-inner-darkmode");
            $("#divBottom > div").toggleClass("div-bottom-inner-lightmode");
            $("#divTopLeft > div:nth-child(1)").removeClass("div-top-left-1-outer-darkmode");
            $("#divTopLeft > div:nth-child(1)").toggleClass("div-top-left-1-outer-lightmode");
            $("#divTopLeft > div:nth-child(2)").removeClass("div-top-left-2-inner-darkmode");
            $("#divTopLeft > div:nth-child(2)").toggleClass("div-top-left-2-inner-lightmode");
            $("#divTopLeft > div:nth-child(3)").removeClass("div-top-left-3-outer-darkmode");
            $("#divTopLeft > div:nth-child(3)").toggleClass("div-top-left-3-outer-lightmode");
            $("#divTopLeft > div:nth-child(4)").removeClass("div-top-left-4-inner-darkmode");
            $("#divTopLeft > div:nth-child(4)").toggleClass("div-top-left-4-inner-lightmode");
            $("#divTopRight > div:nth-child(3)").removeClass("div-top-right-3-outer-darkmode");
            $("#divTopRight > div:nth-child(3)").toggleClass("div-top-right-3-outer-lightmode");
            $("#divTopRight > div:nth-child(3) > div").removeClass("div-top-right-3-inner-darkmode");
            $("#divTopRight > div:nth-child(3) > div").toggleClass("div-top-right-3-inner-lightmode");
            $("#divDarkMode").removeClass("div-dark-mode-dark-outer-darkmode");
            $("#divDarkMode").toggleClass("div-dark-mode-dark-outer-lightmode");
            $("#divDarkMode > div").removeClass("div-dark-mode-dark-inner-darkmode");
            $("#divDarkMode > div").toggleClass("div-dark-mode-dark-inner-lightmode");

        }
    });
    $("#divDarkMode").click(function () {
        var checkbox = $("#darkmode");
        if(checkbox.prop("checked") && (checkbox.val() == "1")) {
            checkbox.prop("checked", false);
        } else if(!checkbox.prop("checked") && (checkbox.val() == "0")) {
            checkbox.prop("checked", true);
        }
        checkbox.trigger("change");
    });
    $("#btnAvatarMenuBtn").click(function () {
        if ($("#divAvatarMenuSubMenu").css("display") == "none") {
            $("#divAvatarMenuSubMenu").css("display", "block");
        } else {
            $("#divAvatarMenuSubMenu").css("display", "none");
        }
    });
    $("#btnAvatarUpd").click(function () {
        $("#divAvatarMenuSubMenu").css("display", "none");
        $("#fileAvatar").click();
    });
    $("#btnAvatarDel").click(function () {
        if($("#avatar").val().length === 0)
            return
        let frmDat = {
          location: $("#avatar").val()};
        $.ajax({
            url: PROFILE_AVATAR_URL,
            type: 'PUT',
            data: JSON.stringify(frmDat),
            contentType: "application/json",
            dataType: "json",
            cache: false,
            processData: false,
            success: function (result) {
                if (result.status = "OK") {
                    $("#imgAvatar").prop("src", "/static/img/icon/icons8-male-user-96.png");
                    $("#avatar").val("");
                }
            },
            error: function (xhr, textStatus, errorMessage) {
                //$.LoadingOverlay("hide", true);
                OpenAlertModal(textStatus,  xhr.responseText);
            }
        });
    });
    $("#fileAvatar").on("change", function (e) {
        if (this.files.length > 0) {
            const fileSize = this.files.item(0).size;
            const fileMb = fileSize / 1024 ** 2;
            if (fileMb <= 2) {
                $("#formAvatar").submit();
            } else {
                alert("kelebihan nih");
            }
        }
    });
    $("#formAvatar").submit(function (e) {
        e.preventDefault();
        var frmDat = new FormData(this);
        $.ajax({
            url: PROFILE_AVATAR_URL,
            type: 'PUT',
            data: frmDat,
            contentType: "application/json",
            success: function (result) {
                $("#divAvatarMenuSubMenu").css("display", "none");
                $("#imgAvatar").prop("src", `data:image/jpg;base64,${result.image}`);
                $("#avatar").val("1");
            },
            error: function (xhr, textStatus, errorMessage) {
                //$.LoadingOverlay("hide", true);
                OpenAlertModal(textStatus,  xhr.responseText);
            },
            cache: false,
            contentType: false,
            processData: false
        });
    });
    $("#frmInp").submit(function (e) {
        e.preventDefault();
        e.stopImmediatePropagation();
        let language = $("#language").val();
        let darkmode = parseInt($("#darkmode").val());
        let avatarLocation = $("#avatar").val();
        let note = $("#note").val();
        let frmDat = {
          language: language,
          darkmode: darkmode,
          avatarLocation: avatarLocation,
          note: note};
        $.ajax({
          type: "PUT",
          url: window.location.pathname,
          data: JSON.stringify(frmDat),
          contentType: "application/json",
          dataType: "json",
          cache: false,
          processData: false,
          success:
            function(output, status, xhr) {
                //$.LoadingOverlay("hide", true);
                if(output.Status == 1) {
                    OpenAlertModal("SUCCESS", "Profile Updated");
                }
            }
        , error:
            function (xhr, textStatus, errorMessage) {
                //$.LoadingOverlay("hide", true);
                OpenAlertModal(textStatus,  xhr.responseText);
            }
      });
    });
    CancelProcess(onCancel);
});
