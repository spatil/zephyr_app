require('expose-loader?$!expose-loader?jQuery!jquery');
require("bootstrap-sass/assets/javascripts/bootstrap.js");
require("./jquery.simple.websocket.min.js");
require("./hammer.min.js");
require("./jquery.numpad.js");

$(() => {
});

$(document).ready(function() {
  /*
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
  */
  var bd = document.getElementsByTagName("body")[0];
  var mc = new Hammer.Manager(bd);

  mc.add( new Hammer.Tap({event: 'tripletap', taps: 3 }));
  mc.add( new Hammer.Tap({event: 'doubletap', taps: 2 }));
  mc.add( new Hammer.Tap({event: 'singletap' }));

  mc.get('doubletap').recognizeWith(['doubletap', 'singletap']);
  mc.get('doubletap').recognizeWith('singletap');
  mc.get('singletap').requireFailure('doubletap');

  mc.on("singletap doubletap tripletap", function(ev) {
    if(ev.type == "tripletap" && window.location.pathname == "/dashboard") {
      window.location.href = "/settings";
    }
  });

  $(".tracks").on("click", "button", function() {
    $(".tracks button").removeClass("active");
    $(this).addClass("active");
  })

  $(document).on("click", ".move_arm", function(e) {
    if($(".tracks").find(".active").length == 0) {
      alert("Please select track");
    } else {
      var track = $(".tracks").find(".active").data("no");
      $.get("/settings/play", {action: 1, track: track});
    }
  });

  $(document).on("click", ".play_track", function(e) {
    if($(".tracks").find(".active").length == 0) {
      alert("Please select track");
    } else {
      var track = $(".tracks").find(".active").data("no");
      $.get("/settings/play", {action: 0, track: track});
    }
  });

  $(document).on("click", "#setPw", function() {
    var pw = $("#pw").val();
    if(pw == ""){
      alert("Please enter pulse width");
    } else {  
      $.post("/settings/change_pw", {pw: pw})
    }
  });

  $(".rpms").on("click", "button", function() {
    $(".rpms button").removeClass("active");
    $(this).addClass("active");
  })

  $(document).on("click", "#start_platter", function() {
    if($(".rpms").find(".active").length == 0) {
      alert("Please select RPM");
    } else {  
      var rpm = $(".rpms").find(".active").data("rpm");
      window.location.href = "/platter/"+ rpm;
    }
  });


	$(document).on("mousedown", ".move-left, .move-right", function() {
    console.log("down");
		var dir = $(this).data("dir");
		$.get("/move_arm/"+ dir);	
	});

	$(document).on("mouseup", ".move-left, .move-right", function() {
    console.log("up");
    $.get("/arm_motor/stop");	
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
  var webSocket = $.simpleWebSocket({ url: 'ws://127.0.0.1:3000/settings/rpm_monitor' });

  // reconnected listening
  webSocket.listen(function(message) {
    $("#rpm").html(message.rpm.toFixed(2));
  });
}

