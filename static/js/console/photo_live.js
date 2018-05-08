const API_HOST = "/api";

$(function() {
  $("a.edit").click(function() {
    $("select.need-enable.inputSupervisor").removeClass("col-sm-12")
      .addClass("col-sm-11");
    $.get(API_HOST + "/users", function(data) {
      for (var i = 0; i < data.length; i++) {
        var html =
        "<option value=" + data[i].id + ">" +
          data[i].username + " " + data[i].phone
        "</option>";
        for (var j = 0; j < $(".inputSupervisor").length; j++) {
          if ($(".inputSupervisor")[j].value != data[i].id)
            $($(".inputSupervisor")[j]).append(html);
        }
        if ($("#inputManager").val() != data[i].id) {
          $("#inputManager").append(html).val()
        }
        if ($("#inputPhotographerManager").val() != data[i].id) {
          $("#inputPhotographerManager").append(html)
        }
      }
      var userSelectHTML = $(".supervisor-group div.need-appear").html();
      $("button.add").click(function() {
        $(".supervisor-group").append(
          "<div class='row w-100 mx-auto'>" +
            userSelectHTML +
          "</div>"
        );
        $("button.del").click(function() {
          $(this).parent().remove();
        });
      });
    });
  });
  $("a.submit").click(onSubmit);
  $("button.btn-primary").click(function() {
    $(".supervisor-group").append(
      "<input class='form-control mb-2' type='text'" +
      "  value='abcdef' disabled>"
    );
  });
  $("button.del").click(function() {
    $(this).parent().remove();
  });
});

function onSubmit() {
  var supervisors = new Array();
  for (var i = 0; i < $(".supervisor-group div").length; i++) {
    supervisors[supervisors.length] =
      $($(".supervisor-group div")[i]).children("select").val();
  }
  $.ajax({
    url: API_HOST + "/photolives",
    type: "PUT",
    data: {
      "id": $("#photoLiveID").val(),
      "expect_members": $("#inputExpectMembers").val(),
      "ad_progress": $("#inputAdProgress").val(),
      "activity_stage_id": $("#inputActivityStage").val(),
      "manager_id": $("#inputManager").val(),
      "photographer_manager_id": $("#inputPhotographerManager").val(),
      "supervisor_ids": supervisors.toString()
    }
  }).done(function(data) {
    alert("修改成功");
    window.location.reload();
  }).fail(function(msg) {
    alert("修改失败");
  });
}
