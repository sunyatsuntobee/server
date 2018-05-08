$(function() {
  $("tbody tr").click(function() {
    window.location = "/users?id=" + $(this).children("input").val();
  });
});
