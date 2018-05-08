$(function() {
  $("a.edit").click();

  sortStageNumbers();
  $("a.submit").click(onSubmit);
  $("button.del").click(onStageDelete);
  $("button.add").click(onStageAdd);
  $("input#logo").change(onLogoChange);
  $("input#poster").change(onPosterChange);
});

function onPosterChange() {
  var file = $("input#poster")[0].files[0];
  $("label[for='poster']").html(file.name);
  var reader = new FileReader();
  reader.onload = function(e) {
    $("img.poster").attr("src", e.target.result);
  }
  reader.readAsDataURL(file);
  this.value = null;
}

function onLogoChange() {
  var file = $("input#logo")[0].files[0];
  $("label[for='logo']").html(file.name);
  var reader = new FileReader();
  reader.onload = function(e) {
    $(".logo-croppie").children().remove();
    var logoCroppie = $(".logo-croppie").croppie({
      viewport: {
        width: 184,
        height: 184
      }
    });
    logoCroppie.croppie("bind", {
      url: e.target.result
    });
    $(".croppie-container").show();
    $("button.logo-upload").off("click").click(function() {
      logoCroppie.croppie("result", "base64").then(function(base64) {
        $("img.logo").attr("src", base64);
        $(".croppie-container").hide();
      });
    });
  }
  reader.readAsDataURL(file);
  this.value = null;
}

function onSubmit() {
  var id = $("input#activity-id").val();
  $.ajax({
    url: "/api/activities/" + id,
    method: "PUT",
    data: {
      id: $("input#activity-id").val(),
      name: $("input#activity-name").val(),
      description: $("textarea#activity-description").val(),
      category: $("input#activity-category").val(),
      poster_url: $("img.poster").attr("src"),
      logo_url: $("img.logo").attr("src"),
      organization_id: $("input#organization-id").val()
    }
  });
  $.ajax({
    url: "/api/activities/" + id + "/stages",
    method: "DELETE"
  }).done(function() {
    for (var i = 0; i < $("div.activity-stage").length; i++) (function(i) {
      var that = $($("div.activity-stage")[i]);
      $.ajax({
        url: "/api/activities/" + id + "/stages",
        method: "POST",
        data: {
          id: null,
          stage_num: that.children("h5.stage-num").val(),
          start_time: that.children("input.start-time").val(),
          end_time: that.children("input.end-time").val(),
          location: that.children("input.location").val(),
          content: that.children("textarea.content").val(),
          activity_id: id
        }
      }).done(function() {
        if (i == $("div.activity-stage").length - 1) {
          alert("修改成功");
          window.location = "/activities?oid=" + $("input#organization-id").val();
        }
      })
    })(i);
  });
}

function sortStageNumbers() {
  for (var i = 0; i < $("h5.stage-num").length; i++) {
    $($("h5.stage-num")[i]).html("阶段 #" + (i + 1));
  }
}

function onStageDelete() {
  $(this).parent().remove();
  sortStageNumbers();
}

function onStageAdd() {
  var stagesTotal = $("div.activity-stage").length + 1;
  var html =
    "<div class='activity-stage border p-3'>" +
      "<h5 class='stage-num' value="+stagesTotal+">阶段 #"+stagesTotal+"</h5>" +
      "<label for=''>阶段活动地点</label>" +
      "<input type='text' class='form-control'>" +
      "<label for=''>开始时间</label>" +
      "<input type='datetime-local' class='form-control'>" +
      "<label for=''>结束时间</label>" +
      "<input type='datetime-local' class='form-control'>" +
      "<label for=''>活动内容</label>" +
        "<textarea class='form-control' rows='4'></textarea>" +
      "<button type='button'" +
        "class='btn btn-outline-secondary text-danger mt-3 col del'>" +
        "删除" +
      "</button>" +
    "</div>";
  $(this).before(html);
  $("button.del").click(onStageDelete);
}
