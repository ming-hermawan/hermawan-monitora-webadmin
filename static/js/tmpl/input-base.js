function EnableButton() {
    $('#btnSave').prop('disabled', false);
    $('#btnCancel').prop('disabled', false);
}


function DisableButton() {
    $('#btnSave').prop('disabled', true);
    $('#btnCancel').prop('disabled', true);
}


function CancelProcess(onCancel) {
    $("#btnCancel").click(function() {
        onCancel();
        DisableButton();
    });
}


$(document).ready(function() {
    DisableButton();
    $(".h-inp").change(function() {
        EnableButton();
    });
    $(".h-inp-chkbox").change(function() {
        EnableButton();
    });
    $(".h-inp-textarea").change(function() {
        EnableButton();
    });
    $(".h-btn-inp").click(function() {
        EnableButton();
    });
});
