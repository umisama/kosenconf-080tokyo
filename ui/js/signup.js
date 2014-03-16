$(function() {
    $("#signup-btn").click(function() {
        $.ajax({
            type: "POST",
            url: "/api/user.json",
            data: "id=" + $("#id").val() +
                    "&screen_name="+ $("#screen_name").val() +
                    "&password=" + $("#password").val(),
            success: function(msg) {
                if (msg.meta.code === 200) {
                    window.location.replace("./index.html");
                } else {
                    alert("ログインできませんでした");
                }
            },
            error: function(XMLHttpRequest, textStatus, errorThrown) {
                var res = XMLHttpRequest.responseJSON;
                alert(res.meta.code + " / " + res.meta.message);
            }
        });
    });
});
