$(function() {
    $("#search-btn").click(function() {
        $("#search-word").html($("#keyword").val());
        //値をエスケープして出力するときは以下
        //$("#search-word").text($("#keyword").val());
    });
});
