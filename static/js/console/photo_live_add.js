const API_HOST = "http://localhost:9090/api";

$(function() {
  $.get(API_HOST + "/users", function(data) {
    for (var i = 0; i < data.length; i++) {
      var html =
      "<option value=" + data[i].id + ">" +
        data[i].username +
      "</option>";
      $("#inputManager").append(html);
      $("#inputPhotographerManager").append(html);
      $(".inputSupervisor").append(html);
    }
    var userSelectHTML = $("#supervisor-group").html();
    $("button").click(function() {
      $("#supervisor-group").append(userSelectHTML);
    });
  });

  $("#inputOrganization").on("change", onInputOrganizationChange);
  $("#inputActivity").on("change", onInputActivityChange);
  $("a.submit").click(onSubmit);
});

function clearSelect(select) {
  select.empty();
  select.append("<option val=-1>请选择</option>");
}

function onInputOrganizationChange() {
  clearSelect($("#inputActivity"));
  var oid = $(this).val();
  if (oid == -1) return;
  $.get(API_HOST + "/activities?oid=" + oid, function(data) {
    for (var i = 0; i < data.length; i++) {
      $("#inputActivity").append(
        "<option value="+ data[i].id + ">" + data[i].name + "</option>"
      );
    }
  });
}

function onInputActivityChange() {
  clearSelect($("#inputActivityStage"));
  var aid = $(this).val();
  if (aid == -1) return;
  $.get(API_HOST + "/activitystages?aid=" + aid, function(data) {
    for (var i = 0; i < data.length; i++) {
      $("#inputActivityStage").append(
        "<option value="+ data[i].id + ">" +
          "#" + data[i].stage_num + " " + data[i].content +
        "</option>"
      );
    }
  });
}

function onSubmit() {
  //TODO submit
}
