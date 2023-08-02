document.addEventListener("DOMContentLoaded", function () {
  const slider = document.getElementById("slider");
  const carousel = new bootstrap.Carousel(slider, {
    interval: 5000,
  });
});

$(document).ready(function () {
  $('.animated').waypoint(function () {
    $(this.element).addClass('active');
  }, {
    offset: '80%' // Cambia este valor según tus necesidades
  });
});