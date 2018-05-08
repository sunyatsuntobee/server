const API_HOST = "/api"

$(function() {
  $.get(API_HOST + "/photos?category=未审核", function(data) {
    var index = 0;
    initPhoto(data[index]);
    $("button[name='passSubmit']").click(onPass(data, index));
    $("button[name='denySubmit']").click(onDeny(data, index));
    $("button[name='next']").click(nextPhoto(data, index));
  });
});

function onPass(data, index) {

  return function() {
    data[index].photo.category = $("select[name='category']").val();
    data[index].photo.release_time = moment().format();
    $.ajax({
      url: API_HOST + "/photos",
      method: "PUT",
      data: data[index].photo
    }).done(function() {
      nextPhoto(data, index)();
    });
  }

}

function onDeny(data, index) {

  return function() {
    data[index].photo.reject_reason = $("textarea[name='reason']").html();
    data[index].photo.category = "未通过";
    $.ajax({
      url: API_HOST + "/photos",
      method: "PUT",
      data: data[index].photo
    }).done(function() {
      nextPhoto(data, index)();
    });
  }

}

function nextPhoto(data, index) {

  return function() {
    index++;
    if (index == data.length) {
      alert("审核完成");
      window.location = "/photos";
    }
    initPhoto(data[index])
    $("button[name='passSubmit']").off("click");
    $("button[name='passSubmit']").click(onPass(data, index));
    $("button[name='denySubmit']").off("click");
    $("button[name='denySubmit']").click(onDeny(data, index));
  }

}

function initPhoto(data) {
  $("h2").html(data.photo.took_location);
  $("h3").html(data.user.username);
  var tookTime = moment(data.photo.took_time);
  $("small").html("拍摄于 " + tookTime.format("YYYY-MM-DD HH:MM:SS"));
  $("img").attr("src", data.photo.url);
}
