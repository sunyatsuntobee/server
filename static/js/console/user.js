$(function() {
  $("a.submit").click(onSubmit);
  $("input#inputAvatar").change(onAvatarChange);
});

function onSubmit() {
  $.ajax({
    url: "/api/users/" + $("input#inputID").val(),
    method: "PUT",
    data: {
      id: $("input#inputID").val(),
      username: $("#inputUsername").val(),
      phone: $("#inputPhone").val(),
      password: null,
      location: $("#inputLocation").val(),
      create_time: null,
      vip: null,
      avatar_url: $("img.icon").attr("src"),
      camera: $("#inputCamera").val(),
      description: $("#inputDescription").val(),
      occupation: $("#inputOccupation").val(),
      collage: $("#inputCollage").val()
    }
  }).done(function() {
    alert("修改成功");
    window.location = "/users?id=" + $("#inputID").val();
  });
}

function onAvatarChange() {
  var file = $("#inputAvatar")[0].files[0];
  $("label[for='inputAvatar']").html(file.name);
  var reader = new FileReader();
  reader.onload = function(e) {
    $(".croppie-container").children().remove();
    var avatarCroppie = $("div.avatar-croppie").croppie({
      viewport: {
        width: 184,
        height: 184
      }
    });
    avatarCroppie.croppie("bind", {
      url: e.target.result
    });
    $("div.avatar-croppie").show();
    $("button.upload").off("click").click(function() {
      avatarCroppie.croppie("result", "base64").then(function(base64) {
        $("img.icon").attr("src", base64);
        $("div.croppie-container").hide();
      });
    });
  }
  reader.readAsDataURL(file);
  this.value = null;
}
