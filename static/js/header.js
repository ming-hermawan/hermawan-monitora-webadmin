function on_mouse_over_profile_menu() {
    $("#divOuterLayer").show();
    $("#divHoverMenuConfig").hide();
    $("#divHoverMenuMonitoring").hide();
    $("#divHoverProfileMenu").show();
}
function on_mouse_over_config_menu() {
    $("#divOuterLayer").show();
    $("#divHoverProfileMenu").hide();
    $("#divHoverMenuMonitoring").hide();
    $("#divHoverMenuConfig").show();
}
function on_mouse_over_monitoring_menu() {
    $("#divOuterLayer").show();
    $("#divHoverProfileMenu").hide();
    $("#divHoverMenuConfig").hide();
    $("#divHoverMenuMonitoring").show();
}



function on_mouse_leave_menu() {
    $("#divHoverMenuConfig").hide();
    $("#divHoverMenuMonitoring").hide();
    $("#divHoverProfileMenu").hide();
    $("#divOuterLayer").hide();
}
