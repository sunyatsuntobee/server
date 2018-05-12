$(function() {
  $("button[type='submit']").click(function() {
    window.location.href = "/photos"
  });

  $("button[type='button']").click(function() {
    window.location.href = "/register"
  })
});
