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

    $("#shout-btn").click(function() {
        $.ajax({
            type: "GET",
            url: "/api/token",
            success: function(msg1) {
                $.ajax({
                    type: "POST",
                    url: "/api/statuses",
                    data: "shout="+ $("#shout-str").val()+"&token="+msg1.content,
                    success: function(msg) {
                        refreshTimeline();
                        $("#shout-str").val("")
                    },
                    error: function(XMLHttpRequest, textStatus, errorThrown) {
                        var res = XMLHttpRequest.responseJSON;
                        alert(res.meta.code + " / " + res.meta.message);
                    }
                });
            }
        });
    });

    $("body").ready(function(){
        refreshTimeline();
        getUserName();
    })
});
