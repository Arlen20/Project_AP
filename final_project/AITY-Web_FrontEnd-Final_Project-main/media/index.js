$(".images img").on("mouseover", function () {
    $(this).addClass("targetImage");
    $(this).animate({
        width: '520px',
        height: '500px',
        bottom: '75px',
        right: '75px'
    }, 300);
    $(".images img").not(this).animate({
        width: '295px',
        height: '275px',
        bottom: '-37.5px',
        right: '-37.5px'
    }, 300);
});

$(".images img").on("mouseout", function () {
    $(this).animate({
        width: '370px',
        height: '350px',
        bottom: '0',
        right: '0'
    }, 300);
    $(".images img").not(this).animate({
        width: '370px',
        height: '350px',
        bottom: '0',
        right: '0'
    }, 300);
    $(this).removeClass("targetImage");
});


let audio = new Audio('media/mouseClick.mp3');
$("#lgbtn").on("click", function () {
    audio.play();
});
$("#rgbtn").on("click", function () {
    audio.play();
});

