function onSelect(result) {
    $("#servergroup").val(result.servergroup);
    $("#sortnumber").val(result.sortnumber);
    $("#servergroup0").val(result.servergroup);
    $("#sortnumber0").val(result.sortnumber);
    $("#servergroup").prop('disabled', true);
}


function getFrmDat() {
    servergroup = $("#servergroup").val();
    sortnumber = parseInt($("#sortnumber").val());
    return {
      "servergroup": servergroup,
      "sortnumber": sortnumber};
}


function onCancel() {
    $("#servergroup").val($("#servergroup0").val());
    $("#sortnumber").val($("#sortnumber0").val());
}


function onNew() {
    $("#servergroup").val("");
    $("#sortnumber").val(null);
    $("#servergroup0").val("");
    $("#sortnumber0").val(null);
    $("#servergroup").prop('disabled', false);;
}


function submitValidation() {
    let servergroup = $("#servergroup").val().trim();
    if(servergroup === "") {
        $("#divInputMessage").html("Server-Group Name can't be empty!");
        return false;
    }
    let sortnumber = $("#sortnumber").val().trim();
    if( (sortnumber === "") || (isNaN(sortnumber)) ) {
        $("#divInputMessage").html("Sort Number should be number.!");
        return false;
    }
    return true;
}

$(document).ready(function() {
    InitTbl(0, onSelect);
    NewProcess(onNew);
    CancelProcess(onCancel);
    SubmitProcess(submitValidation, getFrmDat);
});
