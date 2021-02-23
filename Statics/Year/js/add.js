alert(plus.device.uuid);
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

    //点击按钮触发此方法，用于图片上传
    $('#backgroundImg, #blessingImg').change(function () {
        _this = this;

        if(typeof FileReader == 'undefined') {
            alert("当前浏览器不支持FileReader接口");
            return;
        }
        var file = this.files[0];
        let reader = new FileReader() ;//新建一个FileReader对象
        reader.readAsDataURL(file) ;//将读取的文件转换成base64格式
        reader.onload = function(e) {
            var imageType = file.type.replace('image/', '.');
            var bytes = window.atob(e.target.result.split(',')[1]); //去掉url的头，并转换为byte
            //处理异常,将ascii码小于0的转换为大于0
            var ab = new ArrayBuffer(bytes.length);
            var ia = new Uint8Array(ab);
            for (var i = 0; i < bytes.length; i++) {
                ia[i] = bytes.charCodeAt(i);
            }
            //获取图片的blob对象，因为上传至七牛云需要blob对象
            var blobData = new Blob([ab], {
                type: 'image/png'
            });

            var imageID = "";

            if (_this.id == 'backgroundImg') {
                imageID = 'backImg';
            } else {
                imageID = 'blessImg';
            }

            // 图片上传至七牛云
            if (window.sessionStorage.getItem("qnt")) {
                var qnToken = window.sessionStorage.getItem("qnt");

                qiniuUpload(
                    blobData,
                    new Date().getTime() + Math.random(1000, 9999) + imageType,
                    $("#" + imageID),
                    qnToken
                );
            } else {
                $.ajax({
                    url: "/year/token",
                    type: 'post',
                    dataType: 'json',
                    success: function (res) {
                        window.sessionStorage.setItem("qnt", res.token);

                        qiniuUpload(
                            blobData,
                            new Date().getTime() + Math.random(1000, 9999) + imageType,
                            $("#" + imageID),
                            res.token
                        );
                    }
                });
            }
        }
    });

    /**七牛云上传图片
     * cardImageUrl-----需要上传图片的bolb地址
     * fileName-----上传图片的原文件名
     * nodeObj------显示图片的盒子
     * token------七牛token
     * */
    function qiniuUpload(cardImageUrl, fileName, nodeObj, token) {
        //获取七牛云token，这个需要请求后台接口获得
        const options = {
            quality: 0.92,
            noCompressIfLarger: true
        };

        qiniu.compressImage(cardImageUrl, options).then(data => {
            //拿到token之后，请求七牛云，将图片上传至七牛云
            var observable = qiniu.upload(
                data.dist, //上传图片的blob对象
                fileName, //图片名
                token, //token
                {
                    fname: fileName,
                    params: {}, //用来放置自定义变量
                    mimeType: null
                }, {
                    useCdnDomain: true,
                    region: null
                }
            );

            observable.subscribe({
                complete(res) {
                    //图片上传成功之后，获取图片地址，添加到img标签的src内，连同其他标签元素一起追加到图片显示区域
                    nodeObj.attr('src', "http://qiniu.18x6.com/" + res.key);
                }
            });
        });
    }

    // 提交表单
    $('#submit').click(function (e) {
        if (e.preventDefault) {
            e.preventDefault();
        } else {
            window.event.returnValue = false;
        }

        if ($('#alias').val() == "") {
            showAlert("", "请填写别名", "warning");
            return;
        }

        if ($('#title').val() == "") {
            showAlert("", "请填写标题", "warning");
            return;
        }

        if ($('#content').val() == "") {
            showAlert("", "请填写祝福语", "warning");
            return;
        }

        if ($('#backImgBg').attr('bgsrc') == "") {
            showAlert("", "请选择背景图", "warning");
            return;
        }

        if ($('#fuImg').attr('src') == "" && $('#blessImg').attr('src') == "") {
            showAlert("", "请选择福字或上传祝福图", "warning");
            return;
        }

        _this = $(this);

        _this.prop("disabled", true);
        _this.html(
            `<span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span> Loading...`
        );

        var url = window.location.href;
        var index = url.lastIndexOf("\/");
        str = url.substring(index + 1,url.length);

        $.ajax({
            url: "/year/add",
            type: "post",
            data: {
                open_id: str,
                alias: $('#alias').val(),
                title: $('#title').val(),
                content: $('#content').val(),
                background_image: $('#backImgBg').attr('bgsrc'),
                blessing_image: $('#blessImg').attr('src'),
                fu_image: $('#fuImg').attr('src')
            },
            dataType: "json",
            success: function (resp) {
                if (resp.code == 200) {
                    showAlert("", resp.msg, "success", function () {
                        location.href = "/year/list/" + str;
                    });
                } else {
                    showAlert("", resp.msg, "error");
                }

                _this.prop("disabled", false);
                _this.html('提交我的拜年贴');
            }
        })
    });

    function showAlert(title, text, type, callback) {
        swal({
            title: title,
            text: text,
            type: type,
            showCancelButton: false,  // 是否显示取消按钮
            confirmButtonColor: "#007BFF", // 确认按钮颜色
            confirmButtonText: "确定",
            closeOnConfirm: true, // 是否点击确认后直接关闭
            closeOnCancel: false // 是否点击取消后直接关闭
        }, callback);
    }

    $('#eyeAll').click(function () {
        var url = window.location.href;
        var index = url.lastIndexOf("\/");
        str = url.substring(index + 1,url.length);

        window.location.href = "/year/list/" + str;
    });
});

// 内容库选择
function blessingText(value) {
    $('#content').val("");
    $('#content').val(value);
    $('#blessingWordModel').modal('hide');
}

// 背景图选择
function backgroundImage(value) {
    $('#backImgBg').css({
        "background": "url('" + value + "')"
    });
    $('#backImgBg').attr('bgsrc', value);
    $('#backgroundModel').modal('hide');
}

// 福字选择
function fuImage(value) {
    $('#fuImg').attr('src', value);
    $('#fuModel').modal('hide');
}
