$(function() {
  var category = getParameterByName("category");
  if (category == "未审核") {
    $('a[href="/photos?category=未审核"]').addClass("active");
    $('a[href="#"]').removeClass("active")
  }
});

function getParameterByName(name, url) {
  if (!url) url = window.location.href;
  name = name.replace(/[\[\]]/g, "\\$&");
  var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
    results = regex.exec(url);
  if (!results) return null;
  if (!results[2]) return '';
  return decodeURIComponent(results[2].replace(/\+/g, " "));
}
