const CURRENT_URL_OBJECT = new URL(window.location.href);
const MAIN_URL = `${CURRENT_URL_OBJECT.origin}${CURRENT_URL_OBJECT.pathname}`;
const QUERY_STRING_PARAMS = new URLSearchParams(window.location.search);
var RelocationUrl = null;


function getCookie(cname) {
  let name = cname + "=";
  let decodedCookie = decodeURIComponent(document.cookie);
  let ca = decodedCookie.split(';');
  for(let i = 0; i <ca.length; i++) {
    let c = ca[i];
    while (c.charAt(0) == ' ') {
      c = c.substring(1);
    }
    if (c.indexOf(name) == 0) {
      return c.substring(name.length, c.length);
    }
  }
  return "";
}


function OpenAboutModal() {
    if($("#statModal").val() === "") {
        $("#divModalAlert").css("display","none");
        $("#divModalAsk").css("display","none");
        $("#divModalFront").css("display","block");
        $("#divModal").css("display","block");
        $("#divModalAbout").css("display","block");
        $("#statModal").val("about");
        $("#divModalAbout").animate(
          {deg: 90},
          {
            duration: 500,
            start:function() {
            },
            step: function(now) {
                now2 = now-90;
                $(this).css({ transform: 'rotateY(' + now2 + 'deg)' });
            }
          });
    }
}


function CloseAboutModal() {
    if($("#statModal").val() === "alert") {
        $("#divModalFront").css("display","none");
        $("#divModalAlert").animate(
            {deg: 0},
            {
                duration: 500,
                start:function() {
                },
                step: function(now) {
                    now2 = now-90;
                    console.log(now2);
                    $(this).css({ transform: 'rotateY(' + now2 + 'deg)' });
                },
                done:function() {
                    $("#divModalFront").css("display","none");
                    $("#divModal").css("display","none");
                    if(RelocationUrl) {
                        window.location.href = RelocationUrl;
                    }
                }
            }
        );
    } else if($("#statModal").val() === "ask") {
        $("#divModalAlert").css({transform: "rotateY(-90deg)"});
        $("#divModalAsk").css("display","block");
        CloseAskModal(true);
    }
    $("#statModal").val("");
}



function OpenAlertModal(title, message) {
    $("#divModalAbout").css("display","none");
    $("#divModalAsk").css("display","none");
    $("#divModalFront").css("display","block");
    $("#divModalAlertTitle").html(title);
    $("#divModalAlertMessage").html(message);
    if($("#statModal").val() === "") {
        $("#statModal").val("alert");
        $("#divModal").css("display","block");
        $("#divModalFront").focus();
        $("#divModalAlert").css("display","block");
        $("#divModalAlert").animate(
          {deg: 90},
          {
            duration: 500,
            start:function() {
            },
            step: function(now) {
                now2 = now-90;
                $(this).css({ transform: 'rotateY(' + now2 + 'deg)' });
            }
          });
    } else if($("#statModal").val() === "ask") {
        $("#divModalAlert").css({transform: "rotateY(0deg)"});
        $("#divModalAlert").css("display","block");
    }
}


function CloseAlertModal() {
    if($("#statModal").val() === "alert") {
        $("#divModalFront").css("display","none");
        $("#divModalAlert").animate(
            {deg: 0},
            {
                duration: 500,
                start:function() {
                },
                step: function(now) {
                    now2 = now-90;
                    console.log(now2);
                    $(this).css({ transform: 'rotateY(' + now2 + 'deg)' });
                },
                done:function() {
                    $("#divModalFront").css("display","none");
                    $("#divModal").css("display","none");
                    if(RelocationUrl) {
                        window.location.href = RelocationUrl;
                    }
                }
            }
        );
    } else if($("#statModal").val() === "ask") {
        $("#divModalAlert").css({transform: "rotateY(-90deg)"});
        $("#divModalAsk").css("display","block");
        CloseAskModal(true);
    } else if($("#statModal").val() === "about") {
        $("#divModalFront").css("display","none");
        $("#divModalAbout").animate(
            {deg: 0},
            {
                duration: 500,
                start:function() {
                },
                step: function(now) {
                    now2 = now-90;
                    console.log(now2);
                    $(this).css({ transform: 'rotateY(' + now2 + 'deg)' });
                },
                done:function() {
                    $("#divModalFront").css("display","none");
                    $("#divModal").css("display","none");
                }
            }
        );
    }
    $("#statModal").val("");
}


function OpenAskModal(title, message) {
    $("#statModal").val("ask");
    $("#divModalAbout").css("display","none");
    $("#divModalAlert").css("display","none");
    $("#divModalAskTitle").html(title);
    $("#divModalAskMessage").html(message);
    $("#divModal").css("display","block");
    $("#divModalAsk").css("display","block");
    $("#divModalAsk").animate(
      {deg: 90},
      {
        duration: 500,
        start:function() {
        },
        step: function(now) {
            now2 = now-90;
            $(this).css({ transform: 'rotateY(' + now2 + 'deg)' });
        }
      });
}


function CloseAskModal(relocate) {
    $("#divModalAsk").animate(
      {deg: 0},
      {
        duration: 500,
        start:function() {
      },
      step: function(now) {
          now2 = now-90;
          $(this).css({ transform: 'rotateY(' + now2 + 'deg)' });
      },
      done:function() {
         $("#divModalAsk").css("display","none");
         $("#divModal").css("display","none");
         if(RelocationUrl && relocate) {
            window.location.href = RelocationUrl;
         }
      }
    });
    $("#statModal").val("");
}


$(document).ready(function(){
    if (!navigator.cookieEnabled) {
        window.location.href = "/cookies-warning";
    }
    let documenth = $(document).height();
    let documentw = $(document).width();
    if ( documenth > 250 ) {
        var top_alert_modal = 0;
        var tol_about_modal = 0;
        if (documentw < 728) {
            top_alert_modal = (documenth - 500) / 2;
            top_about_modal = (documenth - 650) / 2;
        } else {
            top_alert_modal = (documenth - 250) / 2;
            top_about_modal = (documenth - 238) / 2;
        }
        console.log("documenth = "+documenth);
        console.log("top_alert_modal = "+top_alert_modal);
        $("#divModalAlert").css("margin-top", top_alert_modal+"px")
        $("#divModalAsk").css("margin-top", top_alert_modal+"px")
        $("#divModalAbout").css("margin-top", top_about_modal+"px")
    }
    $("#divModalFront").click(function() {
        CloseAlertModal();
    });
    $("#divModalFront").keypress(function() {
        CloseAlertModal();
    });
});
