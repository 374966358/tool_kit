$(function () {
    wx.config({
        debug: false, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
        appId: 'wx18d4d0547eee571e', // 必填，公众号的唯一标识
        timestamp: $('#timestamp').val(), // 必填，生成签名的时间戳
        nonceStr: $('#nonceStr').val(), // 必填，生成签名的随机串
        signature: $('#signature').val(),// 必填，签名
        jsApiList: [
            'updateAppMessageShareData',
            'updateTimelineShareData',
            'onMenuShareTimeline',
            'onMenuShareAppMessage',
            'onMenuShareQQ',
        ] // 必填，需要使用的JS接口列表
    });

    wx.ready(function () {      //需在用户可能点击分享按钮前就先调用
        wx.onMenuShareTimeline({
            title: '预约·慈铭奥亚【免下车核酸检测】', // 分享标题
            link: 'http://6xcloud.com/yue', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
            imgUrl: 'http://6xcloud.com/static/Wage/image/notice.jpg', // 分享图标
            success: function () {
                // 用户点击了分享后执行的回调函数
            }
        });

        wx.onMenuShareAppMessage({
            title: '预约·慈铭奥亚【免下车核酸检测】', // 分享标题
            desc: '小时候，乡愁是一张小小的邮票。现在，乡愁是一份核酸检测报告。慈铭奥亚祝您回乡一路顺风', // 分享描述
            link: 'http://6xcloud.com/yue', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
            imgUrl: 'http://6xcloud.com/static/Wage/image/notice.jpg', // 分享图标
            type: '', // 分享类型,music、video或link，不填默认为link
            dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
            success: function () {
                // 用户点击了分享后执行的回调函数
            }
        });

        wx.updateAppMessageShareData({
            title: '预约·慈铭奥亚【免下车核酸检测】', // 分享标题
            desc: '小时候，乡愁是一张小小的邮票。现在，乡愁是一份核酸检测报告。慈铭奥亚祝您回乡一路顺风', // 分享描述
            link: 'http://6xcloud.com/yue', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
            imgUrl: 'http://6xcloud.com/static/Wage/image/notice.jpg', // 分享图标
            success: function (response) {
                // 设置成功
            }
        });

        wx.updateTimelineShareData({
            title: '预约·慈铭奥亚【免下车核酸检测】', // 分享标题
            link: 'http://6xcloud.com/yue', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
            imgUrl: '', // 分享图标
            success: function () {
                // 设置成功
            }
        });
    });
});

$('#staticBackdrop').on("show.bs.modal", function() {
    $("input[type=reset]").trigger("click");
});

$('#submit').click(function () {
    if ($('#name').val() == "") {
        showAlert("", "请填写姓名", "warning");
        return;
    }

    if ($('#idNum').val() == "") {
        showAlert("", "请填写证件号码", "warning");
        return;
    }

    if ($('#mobile').val() == "") {
        showAlert("", "请填写手机号", "warning");
        return;
    }

    if ($('#code').val() == "") {
        showAlert("", "请填写验证码", "warning");
        return;
    }

    if (!$('#exampleCheck').is(':checked')) {
        showAlert("", "请先勾选最终解释权", "warning");
        return;
    }

    _this = this;

    _this.prop("disabled", true);
    _this.html(
        `<span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span> Loading...`
    );

    $.ajax({
        url: "/yue/add",
        type: "post",
        data: $('#form').serialize(),
        dataType: "json",
        success: function (resp) {
            if (resp.code == 200) {
                showAlert("", resp.msg, "success");
                $('#staticBackdrop').modal("hide");
            } else {
                showAlert("", resp.msg, "error");
            }

            _this.prop("disabled", false);
            _this.html('确认提交');
        }
    })
});

function IsPhoneAvailable (numStr) {
    var myreg = /^[1][3,4,5,7,8][0-9]{9}$/;
    return myreg.test(numStr);
}

window.callback = function(res){
    // res（用户主动关闭验证码）= {ret: 2, ticket: null}
    // res（验证成功） = {ret: 0, ticket: "String", randstr: "String"}
    if (res.ret == 0) {
        var mobile = $('#mobile').val();

        if (!mobile) {
            showAlert("", "请填写手机号码", "warning");
            return;
        }

        if (!IsPhoneAvailable(mobile)) {
            showAlert("", "请填写正确的手机号码", "warning");
            return;
        }

        setTime($("#TencentCaptcha"));

        $.ajax({
            url: "/yue/sms",
            type: "post",
            data: {
                mobile: mobile
            }
        })
    }
};

var countdown = 60;

function setTime(obj) {
    if (countdown == 0) {
        obj.prop('disabled', false);
        obj.text("点击获取验证码");
        countdown = 60;//60秒过后button上的文字初始化,计时器初始化;
        return;
    } else {
        obj.prop('disabled', true);
        obj.text("("+countdown+"s)后重新发送") ;
        countdown--;
    }
    setTimeout(function() { setTime(obj) },1000) //每1000毫秒执行一次
}

function showAlert(title, text, type) {
    swal({
        title: title,
        text: text,
        type: type,
        showCancelButton: false,  // 是否显示取消按钮
        confirmButtonColor: "#007BFF", // 确认按钮颜色
        confirmButtonText: "确定",
        closeOnConfirm: true, // 是否点击确认后直接关闭
        closeOnCancel: false // 是否点击取消后直接关闭
    });
}
