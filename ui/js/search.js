$(function() {
    var template = Handlebars.compile($("#shout-template").html());
    var refreshTimeline = function() {
        $("#search-word").html($("#keyword").val());

        var list = $("#shout-list").html("");
        $.ajax({
            type: "GET",
            url: "/api/search",
            data: "q=" + $("#keyword").val(),
            success: function(msg) {
                list.append(template({"items": msg.content}));
            }
        });
    };

    var getUserName = function() {
        $.ajax({
            type: "GET",
            url: "/api/user",
            data: "",
            success: function(msg) {
                $("#username").text(msg.content)
            },
        });
    };

    $("#search-btn").click(function() {
        refreshTimeline()
    });

    $("body").ready(function(){
        if (window.location.hash != "") {
            $("#keyword").val(window.location.hash.slice(1))
            refreshTimeline();
        }
        getUserName();
    })
});
