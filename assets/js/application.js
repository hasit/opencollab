require('expose-loader?$!expose-loader?jQuery!jquery');
require("bootstrap/dist/js/bootstrap.js");

$(document).ready(function () {
  var url = window.location;
  $('ul.nav a[href="' + url + '"]').parent().addClass('active');
  $('ul.nav a').filter(function () {
    return this.href == url;
  }).parent().addClass('active');
});