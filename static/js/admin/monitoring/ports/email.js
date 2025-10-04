function onSelect(result) {
    $("#email").val(result.email);
    $("#email").prop('disabled', true);;
}


function getFrmDat() {
    email = $("#email").val();
    return {
      "email": email};
}


function onNew() {
    $("#email").val("");
    $("#email").prop('disabled', false);;
}


function submitValidation() {
    let email = $("#email").val().trim();
    if(email === "") {
        $("#divInputMessage").html("Email can't be empty!");
        return false;
    }
}

$(document).ready(function() {
    $("#divInputBtn").css("margin-left", "calc(100% - 80px)");
    $("#divInputBtn").css("width", "80px");
    $("#btnCancel").css("display", "none");
    InitTbl(0, onSelect);
    NewProcess(onNew);
    SubmitProcess(submitValidation, getFrmDat);
});
