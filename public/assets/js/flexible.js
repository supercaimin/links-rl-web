(function flexible(window, document) {
    var width = $(window).width();
    console.log(width);
    if(width < 720) width = 760;
    var aa = (50 / 1920) * width;
    console.log(aa);
    $('html').css('font-size', aa);
}(window, document));