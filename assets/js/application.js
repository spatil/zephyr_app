require('expose-loader?$!expose-loader?jQuery!jquery');
require("bootstrap-sass/assets/javascripts/bootstrap.js");
require("./jquery.simple.websocket.min.js")

$(() => {
  
});


window.monitorRPM = function() {
  var webSocket = $.simpleWebSocket({ url: 'ws://127.0.0.1:3000/rpm_monitor' });
  
  // reconnected listening
  webSocket.listen(function(message) {
    $("#rpm").html(message.rpm.toFixed(2));
  });
}
