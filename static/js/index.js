$(function() {
  var curPage = window.location.pathname;
  $("a.nav-link[href='"+curPage+"']").addClass("active");
});
