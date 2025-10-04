var sec = 0;

function show_profile_wide_menu() {
    $("#divOuterLayer1").show();
    $("#divHoverMenuConfig").hide();
    $("#divHoverMenuMonitoring").hide();
    $("#divHoverHelpMenu").hide();
    $("#divHoverProfileMenu").show();
    $("#divHoverReportMenu").hide();
}
function show_config_menu() {
    $("#divOuterLayer1").show();
    $("#divHoverProfileMenu").hide();
    $("#divHoverMenuMonitoring").hide();
    $("#divHoverHelpMenu").hide();
    $("#divHoverMenuConfig").show();
    $("#divHoverReportMenu").hide();
}
function show_monitoring_menu() {
    $("#divOuterLayer1").show();
    $("#divHoverProfileMenu").hide();
    $("#divHoverMenuConfig").hide();
    $("#divHoverHelpMenu").hide();
    $("#divHoverMenuMonitoring").show();
    $("#divHoverReportMenu").hide();
}
function show_report_menu() {
    $("#divOuterLayer1").show();
    $("#divHoverProfileMenu").hide();
    $("#divHoverMenuConfig").hide();
    $("#divHoverMenuMonitoring").hide();
    $("#divHoverHelpMenu").hide();
    $("#divHoverReportMenu").show();
}
function show_help_menu() {
    $("#divOuterLayer1").show();
    $("#divHoverProfileMenu").hide();
    $("#divHoverMenuConfig").hide();
    $("#divHoverMenuMonitoring").hide();
    $("#divHoverHelpMenu").show();
    $("#divHoverReportMenu").hide();
}
function hide_menu() {
    $("#divHoverMenuConfig").hide();
    $("#divHoverMenuMonitoring").hide();
    $("#divHoverProfileMenu").hide();
    $("#divHoverHelpMenu").hide();
    $("#divOuterLayer1").hide();
    $("#divHoverReportMenu").hide();
}



function click_profile_menu() {
    var profile_menu_visible;
    if(profile_menu_visible) {
        profile_menu_visible = false;
        hide_menu();
    } else {
        profile_menu_visible = true;
        show_profile_menu();
    }
}



function myTimer() {
    const d = new Date();
    var minutes = d.getMinutes();
    if((minutes == 15) || (minutes == 45)) {
        var frmDat = {};
        $.ajax({
          url: "/refreshtoken",
          type: "POST",
          data: JSON.stringify(frmDat),
          dataType: "json",
          success: function (result) {
              console.log(result);
          }
        });
    }
    document.getElementById("div-clock").innerHTML = d.toLocaleTimeString();
}



$(document).ready(function(){
    $("#wideMenuProfile").mouseover(function() {
        show_profile_wide_menu();
    });
    $("#wideMenuProfile").click(function() {
        show_profile_wide_menu();
    });
    $("#narrowMenuProfile").click(function() {
        if($("#narrowMenuAreaProfile").css("display") == "none") {
            // Profile
            $("#narrowMenuProfile").css("border-bottom", "0");
            $("#narrowMenuAreaProfile").css("display", "block");
            // Admin
            $("#narrowMenuAreaAdmin").css("display", "none");
            $("#narrowMenuAdmin").css("border-bottom", "1px solid #ffffff");
            // Monitoring
            $("#narrowMenuAreaMonitoring").css("display", "none");
            $("#narrowMenuMonitoring").css("border-bottom", "1px solid #ffffff");
            // Report
            $("#narrowMenuAreaReport").css("display", "none");
            $("#narrowMenuReport").css("border-bottom", "1px solid #ffffff");
            // About
            $("#narrowMenuAreaAbout").css("display", "none");
            $("#narrowMenuAbout").css("border-bottom", "1px solid #ffffff");
        } else if ($("#narrowMenuAreaProfile").css("display") == "block") {
            $("#narrowMenuProfile").css("border-bottom", "1px solid #ffffff");
            $("#narrowMenuAreaProfile").css("display", "none");
        }
    });
    $("#narrowMenuAdmin").click(function() {
        if($("#narrowMenuAreaAdmin").css("display") == "none") {
            // Profile
            $("#narrowMenuAreaProfile").css("display", "none");
            $("#narrowMenuProfile").css("border-bottom", "1px solid #ffffff");
            // Admin
            $("#narrowMenuAdmin").css("border-bottom", "0");
            $("#narrowMenuAreaAdmin").css("display", "block");
            // Monitoring
            $("#narrowMenuAreaMonitoring").css("display", "none");
            $("#narrowMenuMonitoring").css("border-bottom", "1px solid #ffffff");
            // Report
            $("#narrowMenuAreaReport").css("display", "none");
            $("#narrowMenuReport").css("border-bottom", "1px solid #ffffff");
            // About
            $("#narrowMenuAreaAbout").css("display", "none");
            $("#narrowMenuAbout").css("border-bottom", "1px solid #ffffff");
        } else if ($("#narrowMenuAreaAdmin").css("display") == "block") {
            $("#narrowMenuAdmin").css("border-bottom", "1px solid #ffffff");
            $("#narrowMenuAreaAdmin").css("display", "none");
        }
    });
    $("#narrowMenuMonitoring").click(function() {
        if($("#narrowMenuAreaMonitoring").css("display") == "none") {
            // Profile
            $("#narrowMenuAreaProfile").css("display", "none");
            $("#narrowMenuProfile").css("border-bottom", "1px solid #ffffff");
            // Admin
            $("#narrowMenuAreaAdmin").css("display", "none");
            $("#narrowMenuAdmin").css("border-bottom", "1px solid #ffffff");
            // Monitoring
            $("#narrowMenuMonitoring").css("border-bottom", "0");
            $("#narrowMenuAreaMonitoring").css("display", "block");
            // Report
            $("#narrowMenuAreaReport").css("display", "none");
            $("#narrowMenuReport").css("border-bottom", "1px solid #ffffff");
            // About
            $("#narrowMenuAbout").css("border-bottom", "none");
            $("#narrowMenuAreaAbout").css("display", "1px solid #ffffff");
        } else if ($("#narrowMenuAreaMonitoring").css("display") == "block") {
            $("#narrowMenuMonitoring").css("border-bottom", "1px solid #ffffff");
            $("#narrowMenuAreaMonitoring").css("display", "none");
        }
    });
    $("#narrowMenuReport").click(function() {
        if($("#narrowMenuAreaReport").css("display") == "none") {
            // Profile
            $("#narrowMenuAreaProfile").css("display", "none");
            $("#narrowMenuProfile").css("border-bottom", "1px solid #ffffff");
            // Admin
            $("#narrowMenuAreaAdmin").css("display", "none");
            $("#narrowMenuAdmin").css("border-bottom", "1px solid #ffffff");
            // Monitoring
            $("#narrowMenuAreaMonitoring").css("display", "none");
            $("#narrowMenuMonitoring").css("border-bottom", "1px solid #ffffff");
            // Report
            $("#narrowMenuReport").css("border-bottom", "0");
            $("#narrowMenuAreaReport").css("display", "block");
            // About
            $("#narrowMenuAreaAbout").css("display", "none");
            $("#narrowMenuAbout").css("border-bottom", "1px solid #ffffff");
        } else if ($("#narrowMenuAreaReport").css("display") == "block") {
            $("#narrowMenuReport").css("border-bottom", "1px solid #ffffff");
            $("#narrowMenuAreaReport").css("display", "none");
        }
    });
    $("#narrowMenuAbout").click(function() {
        if($("#narrowMenuAreaAbout").css("display") == "none") {
            // Profile
            $("#narrowMenuAreaProfile").css("display", "none");
            $("#narrowMenuProfile").css("border-bottom", "1px solid #ffffff");
            // Admin
            $("#narrowMenuAreaAdmin").css("display", "none");
            $("#narrowMenuAdmin").css("border-bottom", "1px solid #ffffff");
            // Monitoring
            $("#narrowMenuAreaMonitoring").css("display", "none");
            $("#narrowMenuMonitoring").css("border-bottom", "1px solid #ffffff");
            // Report
            $("#narrowMenuAreaReport").css("display", "none");
            $("#narrowMenuReport").css("border-bottom", "1px solid #ffffff");
            // About
            $("#narrowMenuAbout").css("border-bottom", "0");
            $("#narrowMenuAreaAbout").css("display", "block");
        } else if ($("#narrowMenuAreaAbout").css("display") == "block") {
            $("#narrowMenuAbout").css("border-bottom", "1px solid #ffffff");
            $("#narrowMenuAreaAbout").css("display", "none");
        }
    });
    $("#wideMenuConfig").mouseover(function() {
        show_config_menu();
    });
    $("#wideMenuConfig").click(function() {
        show_config_menu();
    });
    $("#wideMenuMonitoring").mouseover(function() {
        show_monitoring_menu();
    });
    $("#wideMenuReport").mouseover(function() {
        show_report_menu();
    });
    $("#wideMenuHelp").mouseover(function() {
        show_help_menu();
    });
    $("#divOuterLayer1").click(function() {
        hide_menu();
    });
    $("#buttonMenuBack").click(function() {
        if($("#divMenuNarrow2").css("margin-top") == "-500px") {
            $("#buttonMenuBack").animate(
                {deg: 90},
                {duration: 150,
                 step: function(now) {
                     $(this).css({ transform: 'rotateY(' + now + 'deg)' });
                 }
            });
            setTimeout(function() {
              $("#imgMenuBack").attr("src", "/static/img/icon/icons8-back-48.png");
              $("#buttonMenuBack").animate(
                  {deg: 0},
                  {duration: 150,
                   step: function(now) {
                       $(this).css({ transform: 'rotateY(' + now + 'deg)' });
                   }
                  }
              );
            }, 150)
            $("#divMenuNarrow2").animate({marginTop: "0px"});
        }
        else if ($("#divMenuNarrow2").css("margin-top") == "0px") {
          $("#buttonMenuBack").animate(
              {deg: 90},
              {duration: 150,
               step: function(now) {
                   $(this).css({ transform: 'rotateY(' + now + 'deg)' });
               }
          });
          setTimeout(function() {
            $("#imgMenuBack").attr("src", "/static/img/icon/icons8-menu-48.png");
            $("#buttonMenuBack").animate(
                {deg: 0},
                {duration: 150,
                 step: function(now) {
                     $(this).css({ transform: 'rotateY(' + now + 'deg)' });
                 }
                }
            );
          }, 150)
            $("#divMenuNarrow2").animate({marginTop: "-500px"});
        }
    });
    $("#btnAskCancel").click(function() {
        CloseAskModal(false);
    });
    setInterval(myTimer, 1000); //milliseconds
});
