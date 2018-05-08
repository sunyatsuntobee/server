$(function() {
  initActivityStages();
});

function initActivityStages() {
  for (var i = 0; i < $("td.td-collapse").length; i++) {
    var that = $($("td.td-collapse")[i]);
    $.ajax({
      url: "/api/activitystages?aid=" + that.children("input").val(),
      method: "GET",
    }).done(function(data) {
      for (var i = 0; i < data.length; i++) {
        var startTime = moment(data[i].start_time);
        var endTime = moment(data[i].end_time);
        var html =
          "<div>" +
            "<h5>阶段 #"+data[i].stage_num+"</h5>" +
            "<h6 class='text-muted'>时间 "+
              startTime.format("YYYY-MM-DD HH:MM:SS")+"~"+
              endTime.format("YYYY-MM-DD HH:MM:SS")+
            "</h6>" +
            "<h6 class='text-muted'>地点 "+data[i].location+"</h6>" +
            "<p>"+data[i].content+"</p>" +
          "</div>";
        that.children("div.collapse").append(html);
      }
    });
  }
}
