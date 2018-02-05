const API_HOST = "/api"

$(function() {
  $("a.submit").click(onSubmit);
  $("a.edit").click(function() {
    $("input.inputDepartment").removeClass("col-sm-12").addClass("col-sm-11");
    $("select.inputSupervisor").removeClass("col-sm-12").addClass("col-sm-11");
  });
  $("input#inputLogo").change(function() {
    var file = $("#inputLogo")[0].files[0];
    $("label[for='inputLogo']").html(file.name);
    var reader = new FileReader();
    reader.onload = function(e) {
      $("img.icon").attr("src", e.target.result);
      var avatarCroppie = $("img.icon").croppie({
        viewport: {
          width: 184,
          height: 184
        }
      });
      $("button.upload").off("click").click(function() {
        avatarCroppie.croppie("result", "base64").then(function(base64) {
          $("div.form-avatar").prepend(
            "<img class='form-control icon mb-4 mx-auto'>"
          );
          $("img.icon").attr("src", base64);
          $(".croppie-container").remove();
        });
      });
    }
    reader.readAsDataURL(file);
    this.value = null;
  });
  $.get(API_HOST + "/users", function(data){
    for (var i = 0; i < data.length; i++) {
      for (var j = 0; j < $("select.inputSupervisor").length; j++) {
        if ($("select.inputSupervisor").val() == data[i].id) {
          continue;
        }
        $($("select.inputSupervisor")[j]).append(
          "<option value=" + data[i].id + ">" +
            data[i].username + " " + data[i].phone +
          "</option>"
        );
      }
    }
    $("button.addContactor").click(function() {
      $("div.contactor-group").append(
        "<div class='row mx-auto w-100'>" +
        $("div.contactor-group").children().html() +
        "</div>"
      );
      $("button.del").click(function() {
        $(this).parent().remove();
      });
    });
  });
  $("button.del").click(function() {
    $(this).parent().remove();
  });
  $("button.addDepartment").click(function() {
    $("div.department-group").append(
      "<div class='row mx-auto w-100'>" +
      "<input class='form-control mb-2 col-sm-11 " +
        "inputDepartment' type='text' placeholder='部门名称'>" +
      "<button class='btn btn-danger col h-100 ml-2 del'>" +
        "删除" +
      "</button>" +
      "</div>"
    );
    $("button.del").click(function() {
      $(this).parent().remove();
    });
  });
});

function onSubmit() {
  var finishes = 0;
  $.ajax({
    url: API_HOST + "/organizations",
    method: "PUT",
    data: {
      id: $("input[name='id']").val(),
      name: $("#inputOrganizationName").val(),
      phone: $("input[name='phone']").val(),
      password: $("input[name='password']").val(),
      collage: $("#inputOrganizationCollage").val(),
      logo_url: $("img.icon").attr("src"),
      description: $("#inputOrganizationDescription").val()
    }
  });
  $.ajax({
    url: API_HOST + "/organizations/" + $("input[name='id']").val() +
      "/departments",
    method: "DELETE"
  }).done(function() {
    for (var i = 0; i < $("input.inputDepartment").length; i++) {
      $.ajax({
        url: API_HOST + "/organizations/" + $("input[name='id']").val() +
          "/departments",
        method: "POST",
        data: {
          id: null,
          name: $("input.inputDepartment")[i].value,
          organization_id: $("input[name='id']").val()
        }
      });
    }
  });
  $.ajax({
    url: API_HOST + "/organizations/" + $("input[name='id']").val() +
      "/contacts",
    method: "DELETE"
  }).done(function() {
    for (var i = 0; i < $("select.inputSupervisor").length; i++) (function(i) {
      $.ajax({
        url: API_HOST + "/organizations/" + $("input[name='id']").val() +
          "/contacts",
        method: "POST",
        data: {
          id: null,
          organization_id: $("input[name='id']").val(),
          contact_id: $("select.inputSupervisor")[i].value
        }
      }).done(function() {
        if (i == $("select.inputSupervisor").length - 1) {
          alert("修改成功");
          window.location = "/organizations?id=" + $("input[name='id']").val();
        }
      });
    })(i);
  });
}
