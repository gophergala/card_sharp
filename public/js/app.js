var Current = {};
$(document).ready(function() {
  if ($('meta[name="role"]').attr('value') == "presenter" ){
    PresenterLoop($('meta[name="lobby"]').attr('value'));
  }
  if ($('meta[name="role"]').attr('value') == "player" ){
    PlayerLoop($('meta[name="lobby"]').attr('value'),$('meta[name="player"]').attr('value'))
    if (PlayerStartup[$('meta[name="initial"]').attr('value')]) {
      PlayerStartup[$('meta[name="initial"]').attr('value')]()
    }
  }
});

var ws;

function PresenterLoop(lobbyid) {
  ws = new WebSocket("ws://localhost:8101/presenter/"+lobbyid)
  ws.onopen = function(event) {
    ws.send("hello");
  }
  ws.onmessage = function(e) {
    ev = JSON.parse(e.data)

    console.log(ev)
    if (PresenterEvents[ev["Type"]]) {
      PresenterEvents[ev.Type](ev)
    } else {
      console.log("UNKOWN EVENT")
    }
  }
  var jid = $('meta[name="judge"]').attr("value");
  console.log("#king_"+jid)
  $(document.getElementById("king_"+jid)).removeClass("hide")
}

var PresenterEvents = {};

PresenterEvents["new_player"] = function(ev) {
  console.log(ev)
  $("#player_listing").append(ev.HTML)
}
PresenterEvents["player_commit"] = function(ev) {
  $(document.getElementById("item_"+ev.Data)).addClass("active");
}

PresenterEvents["game_start"] = function(ev) {
  // lazy, but should work
  location.reload()
}
PresenterEvents["sync"] = function(ev) {
  // just because
  location.reload()
}
PresenterEvents["player_ready"] = function(ev) {
  
}

PresenterEvents["start_judge"] = function(ev) {
  location.reload()
}

PresenterEvents["set_points"] = function(ev) {

}


function PlayerLoop(lobbyid, playerid) {
  ws = new WebSocket("ws://localhost:8101/player/" + lobbyid + "/"+playerid)
  ws.onmessage = function(evt) {
    ev = JSON.parse(evt.data)
    if (PlayerEvents[ev["Type"]]) {
      PlayerEvents[ev.Type](ev)
    } else {
      console.log("UNKOWN EVENT")
    }
  }
}

var PlayerEvents = {};

PlayerEvents["game_ready"] = function(ev) {
  $("body").html(ev.HTML);
  $("#dothing").removeClass("disabled")
  $("#dothing").click(function() {
    $.post("/players/start_game")
  })
}

PlayerEvents["game_start"] = function(ev) {
  $("body").html(ev.HTML);
  PlayerStartup["play"]();
}
PlayerEvents["sync"] = function(ev) {
  location.reload()
}
PlayerEvents["new_hand"] = function(ev) {
  $("body").html(ev.HTML);
}
var PlayerStartup = {};

PlayerStartup["play"] = function() {
  $(".play_card").click(function(ev) {
    console.log(ev)
    $(".play_card").removeClass("btn-info")
    $(ev.currentTarget).addClass("btn-info")
    Current["card"] = ev.currentTarget.innerHTML;
    $("#dothing").removeClass("disabled")
    $("#dothing").html("Play Card")
  })
  $("#dothing").click(function() {
    $("#dothing").addClass("disabled");
    $("#dothing").html("Playing Card");
    $.post("/players/make_move", {card: Current["card"]})
  })
}

PlayerStartup["play_wait"] = function() {

}
PlayerStartup["judge-wait"] = function() {
  $(".card-list").hide()
}

