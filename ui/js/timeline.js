$(function() {
    var template = Handlebars.compile($("#shout-template").html());
    
    var refreshTimeline = function() {
        var list = $("#shout-list").html("");
        $.ajax({
            type: "GET",
            url: "/api/statuses",
            data: "count=10",
            success: function(msg) {
                list.append(template({"items": msg.content}));
            }
        });
    };

    refreshTimeline();

    $("#shout-btn").click(function() {
        $.ajax({
            type: "POST",
            url: "/api/statuses",
            data: "shout="+ $("#shout-str").val(),
            success: function(msg) {
                refreshTimeline();
                $("#shout-str").val("")
            },
            error: function(XMLHttpRequest, textStatus, errorThrown) {
                var res = XMLHttpRequest.responseJSON;
                alert(res.meta.code + " / " + res.meta.message);
            }
        });
    });
});
