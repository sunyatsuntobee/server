$(function() {
  $("button[type='button1']").click(function() {
    getEle('show').innerHTML=hhhhhhhh;
  });

  $("button[type='button2']").click(function() {
    getEle('show').innerHTML=hhhhhhhh;
  })
});

function toggle(id){
  var tb=document.getElementById(id);
  if(tb.style.display=='none') tb.style.display='block';
  else tb.style.display='none';
}
