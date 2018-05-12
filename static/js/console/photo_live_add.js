const API_HOST = "/api";

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
    var userSelectHTML = $(".supervisor-group").html();
    $("button.add").click(function() {
      $(".supervisor-group").append(userSelectHTML);
      $("button.del").click(function() {
        $(this).parent().remove();
      });
    });
  });
  $("button.del").click(function() {
    $(this).parent().remove();
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
  clearSelect($("#inputActivityStage"));
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
  var supervisors = new Array();
  for (var i = 0; i < $(".supervisor-group div").length; i++) {
    supervisors[supervisors.length] =
      $($(".supervisor-group div")[i]).children("select").val();
  }
  $.post(
    API_HOST + "/photolives",
    {
      "expect_members": $("#inputExpectMembers").val(),
      "ad_progress": $("#inputAdProgress").val(),
      "activity_stage_id": $("#inputActivityStage").val(),
      "manager_id": $("#inputManager").val(),
      "photographer_manager_id": $("#inputPhotographerManager").val(),
      "supervisor_ids": supervisors.toString()
    },
    function(data) {
      alert("添加成功");
      window.location.reload();
    }
  );
}
