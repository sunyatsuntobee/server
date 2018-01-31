$(function() {
  $("a.edit").click(function() {
    $("div.supervisor-group input").removeClass("col-sm-12").
      addClass("col-sm-11");
  });
  $("a.submit").click(function() {
    //TODO submit changes
  });
  $("button.btn-primary").click(function() {
    $(".supervisor-group").append(
      "<input class='form-control mb-2' type='text'" +
      "  value='abcdef' disabled>"
    );
  });
});
