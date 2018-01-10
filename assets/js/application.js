require('expose-loader?$!expose-loader?jQuery!jquery');
require("bootstrap-sass/assets/javascripts/bootstrap.js");
require("./jquery.simple.websocket.min.js");
require("./hammer.min.js");
require("./jquery.numpad.js");

$(() => {
});

$(document).ready(function() {
  var tags = document.getElementsByTagName("a");

  $.each(tags, function(r) {
    Hammer(this).on("tap", function(e) {
      //singletap stuff
      window.location.href = $(this).attr("href");
    });
    Hammer(this).on("doubletap", function(e) {
       e.preventDefault();
    });
  });


  $(document).on("click", "#setPw", function() {
    var pw = $("#pw").val();
    if(pw == ""){
      alert("Please enter pulse width");
    } else {  
      $.post("/change_pw", {pw: pw})
    }
  });
});

window.loadKeypad = function() {
  $.fn.numpad.defaults.gridTpl = '<table class="table modal-content"></table>';
  $.fn.numpad.defaults.backgroundTpl = '<div class="modal-backdrop in"></div>';
  $.fn.numpad.defaults.displayTpl = '<input type="text" class="form-control  input-lg" />';
  $.fn.numpad.defaults.buttonNumberTpl =  '<button type="button" class="btn btn-default btn-lg"></button>';
  $.fn.numpad.defaults.buttonFunctionTpl = '<button type="button" class="btn btn-lg" style="width: 100%;"></button>';
  $.fn.numpad.defaults.onKeypadCreate = function(){$(this).find('.done').addClass('btn-primary');};
  $("#pw").numpad();
}


window.monitorRPM = function() {
  var webSocket = $.simpleWebSocket({ url: 'ws://127.0.0.1:3000/rpm_monitor' });

  // reconnected listening
  webSocket.listen(function(message) {
    $("#rpm").html(message.rpm.toFixed(2));
  });
}

window.monitorSteps = function() {
  var webSocket = $.simpleWebSocket({ url: 'ws://192.168.1.14:3000/step_monitor' });

  // reconnected listening
  webSocket.listen(function(message) {
    $("#steps").html(message.steps.toFixed(0));
  });
}
