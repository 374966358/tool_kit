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
            title: '疫情出行各地政策查询', // 分享标题
            link: 'http://6xcloud.com/travel', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
            imgUrl: 'http://6xcloud.com/static/Wage/image/notice.jpg', // 分享图标
            success: function () {
                // 用户点击了分享后执行的回调函数
            }
        });

        wx.onMenuShareAppMessage({
            title: '疫情出行各地政策查询', // 分享标题
            desc: '疫情期间出行，先查询，再行动，即刻查询目的地的政策', // 分享描述
            link: 'http://6xcloud.com/travel', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
            imgUrl: 'http://6xcloud.com/static/Wage/image/notice.jpg', // 分享图标
            type: '', // 分享类型,music、video或link，不填默认为link
            dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
            success: function () {
                // 用户点击了分享后执行的回调函数
            }
        });

        wx.updateAppMessageShareData({
            title: '疫情出行各地政策查询', // 分享标题
            desc: '疫情期间出行，先查询，再行动，即刻查询目的地的政策', // 分享描述
            link: 'http://6xcloud.com/travel', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
            imgUrl: 'http://6xcloud.com/static/Wage/image/notice.jpg', // 分享图标
            success: function (response) {
                // 设置成功
            }
        });

        wx.updateTimelineShareData({
            title: '疫情出行各地政策查询', // 分享标题
            link: 'http://6xcloud.com/travel', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
            imgUrl: '', // 分享图标
            success: function () {
                // 设置成功
            }
        });
    });

    $('#createYear').click(function () {
        var url = window.location.href;
        var index = url.lastIndexOf("\/");
        str = url.substring(index + 1,url.length);

        window.location.href = "/year/from/" + str
    });
});

function listClick(id, state) {
    if (state == 0) {
        swal({
            title: "",
            text: "拜年贴未审核通过",
            type: "warning",
            showCancelButton: false,  // 是否显示取消按钮
            confirmButtonColor: "#007BFF", // 确认按钮颜色
            confirmButtonText: "确定",
            closeOnConfirm: true, // 是否点击确认后直接关闭
            closeOnCancel: false // 是否点击取消后直接关闭
        });
        return;
    }

    // id：表id
    window.location.href = "/year/item/" + id
}
