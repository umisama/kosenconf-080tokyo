$(function() {
    $("#login-btn").click(function() {
        $.ajax({
            type: "GET",
            url: "/api/session.json",
            data: "id=" + $("#userid").val() + "&password=" + $("#password").val(),
            success: function(msg) {
                if (msg.meta.code === 200) {
                    window.location.replace("./timeline.html");
                } else {
                    alert("ログインできませんでした");
                }
            },
            error: function(XMLHttpRequest, testStatus, errorThrown) {
               var res = XMLHttpRequest.responseJSON;
               alert(res.meta.code + "/" + res.meta.message);
            }
        });
    });

    $("#del-session-btn").click(function() {
        $.ajax({
            type: "DELETE",
            url: "/api/session.json",
            success: function(msg) {
                if (msg.meta.code === 200) {
                    alert("セッション削除しました");
                }
            }
        });
    });
});
