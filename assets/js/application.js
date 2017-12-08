require('expose-loader?$!expose-loader?jQuery!jquery');
require("bootstrap-sass/assets/javascripts/bootstrap.js");
require("./jquery.simple.websocket.min.js");
require("./hammer.min.js");

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
});


window.monitorRPM = function() {
  var webSocket = $.simpleWebSocket({ url: 'ws://127.0.0.1:3000/rpm_monitor' });

  // reconnected listening
  webSocket.listen(function(message) {
    $("#rpm").html(message.rpm.toFixed(2));
  });
}
