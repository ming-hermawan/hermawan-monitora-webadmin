function init_menu_checkbox(name) {
    $("#"+name).change(function () {
        $("."+name).prop("checked", this.checked);
        $("."+name).val("checked", this.checked);
    });
}


function onSelect(result) {
    $("#usergroup").val(result.usergroup);
    $("#usergroup").prop('disabled', true);
    $("#usergroup0").val(result.usergroup);
    $("#sortnumber").val(result.sortnumber);
    $("#sortnumber0").val(result.sortnumber);
    $(".h-inp-checkbox").prop("checked", false);
    $(".menu0").val(false);
    for(menu of result.menus) {
        $("#menu-"+menu).prop("checked", true);
        $("#menu0-"+menu).val(true);
    }
}


function getFrmDat() {
    usergroup = $("#usergroup").val();
    sortnumber = parseInt($("#sortnumber").val());
    const menus = []
    $(".h-inp-checkbox").each(function( index ) {
        if($(this).prop("checked")) {
            const menu_code = $(this).data("menu");
            if(menu_code !== "") {
                menus.push(menu_code);
            }
        }
    });
    return {
      "usergroup": usergroup,
      "sortnumber": sortnumber,
      "menus": menus};
}


function onCancel() {
    $("#usergroup").val($("#usergroup0").val());
    $("#sortnumber").val($("#sortnumber0").val());
    $("#menu-A").prop("checked", ($("#menu0-A").val() === 'true'));
    $("#menu-A-UM").prop("checked", ($("#menu0-A-UM").val() === 'true'));
    $("#menu-A-UM-usergroup").prop("checked", ($("#menu0-A-UM-usergroup").val() === 'true'));
    $("#menu-A-UM-user").prop("checked", ($("#menu0-A-UM-user").val() === 'true'));
    $("#menu-A-M").prop("checked", ($("#menu0-A-M").val() === 'true'));
    $("#menu-A-M-ports").prop("checked", ($("#menu0-A-M-ports").val() === 'true'));
    $("#menu-A-M-ports-settings").prop("checked", ($("#menu0-A-M-ports-settings").val() === 'true'));
    $("#menu-A-M-ports-servergroup").prop("checked", ($("#menu0-A-M-ports-servergroup").val() === 'true'));
    $("#menu-A-M-ports-serversNports").prop("checked", ($("#menu0-A-M-ports-serversNports").val() === 'true'));
    $("#menu-A-M-P-email").prop("checked", ($("#menu0-A-M-P-email").val() === 'true'));
    $("#menu-A-M-P-csv").prop("checked", ($("#menu0-A-M-P-csv").val() === 'true'));
    $("#menu-M").prop("checked", ($("#menu0-M").val() === 'true'));
    $("#menu-M-ports").prop("checked", ($("#menu0-M-ports").val() === 'true'));
    $("#menu-R").prop("checked", ($("#menu0-R").val() === 'true'));
    $("#menu-R-M").prop("checked", ($("#menu0-R-M").val() === 'true'));
    $("#menu-R-M-ports").prop("checked", ($("#menu0-R-M-ports").val() === 'true'));
}


function onNew() {
    $("#method").val("POST");
    $(".h-inp-text").val("");
    $(".h-inp-number").val("");
    $(".h-inp-checkbox").prop("checked", false);
    $("#divInputPageBody > div").css("display", "none");
    $("#divPageMain").css("display", "block");
    $("#divInput").css("display", "block");
    $("#usergroup").prop("disabled", false);
}


function submitValidation() {
    let usergroup = $("#usergroup").val().trim();
    if(usergroup === "") {
        $("#divInputMessage").html("User-Group Name can't be empty!");
        return false;
    }
    let sortnumber = $("#sortnumber").val().trim();
    if( (sortnumber === "") || (isNaN(sortnumber)) ) {
        $("#divInputMessage").html("Sort Number should be number.!");
        return false;
    }
    return true;
}


$(document).ready(function(){
    InitTbl(0, onSelect);
    NewProcess(onNew);
    CancelProcess(onCancel);
    SubmitProcess(submitValidation, getFrmDat);

    init_menu_checkbox("menu-A");
    init_menu_checkbox("menu-A-UM");
    init_menu_checkbox("menu-A-M");
    init_menu_checkbox("menu-A-M-ports");
    init_menu_checkbox("menu-M");
    init_menu_checkbox("menu-R");
    init_menu_checkbox("menu-R-M");

    $(".h-inp-checkbox").each(function( index ) {
        $(this).change(function () {
            console.log($(this).get(0).id +" " + $(this).data("menu"));
        });
    });
});
