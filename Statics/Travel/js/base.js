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

    queryAfter('', '', '', '');

    $('#leftCurrency').click(function () {
        selection(function (value) {
            let leftArr = value.split(" ");

            if (leftArr.length === 1) {
                leftArr[1] = leftArr[0]
            }

            $('#FromP').val(leftArr[0]);
            $('#FromC').val(leftArr[1]);
            $('#leftCurrency').html(leftArr[1]);

            queryAfter(leftArr[0], leftArr[1], $('#ToP').val(), $('#ToC').val());
        })
    });

    $('#rightCurrency').click(function() {
        selection(function (value) {
            let rightArr = value.split(" ");

            if (rightArr.length === 1) {
                rightArr[1] = rightArr[0]
            }

            $('#ToP').val(rightArr[0]);
            $('#ToC').val(rightArr[1]);
            $('#rightCurrency').html(rightArr[1]);

            queryAfter($('#FromP').val(), $('#FromC').val(), rightArr[0], rightArr[1]);
        });
    });

    function selection(callback = function() {}) {
        $('#currencyModel').modal('show');
        $('.index_selection_page_container').empty();

        new indexSelection(null, callback);
    }

    function queryAfter(fP, fC, tP, tC) {
        var map = {
            form_p: fP,
            form_c: fC,
            to_p: tP,
            to_c: tC
        };

        $.get("/travel/index", map, function (response) {
            response = response.data;

            $('#currencyClose').click();

            $('#from_title').html(response.from.city);

            var fromBody = "";

            if (response.from.rule.out_policy) {
                fromBody = response.from.rule.out_policy;
            } else {
                fromBody = "暂无数据";
            }

            $('#from_data').html(fromBody);

            if (response.from.tag) {
                if (response.from.tag === '低风险') {
                    $('#from_tag')
                        .removeClass("border-danger text-danger")
                        .addClass("border-success text-success")
                        .html(response.from.tag);
                } else {
                    $('#from_tag')
                        .removeClass("border-success text-success")
                        .addClass("border-danger text-danger")
                        .html(response.from.tag);
                }
            }

            var fromFooter = "";

            if (response.from.rule.out_policy_resource) {
                fromFooter = "数据来源：" + response.from.rule.out_policy_resource + " | 更新时间：" +
                response.from.rule.out_policy_date
            }

            $('#from_footer').html(fromFooter);

            if (response.from.has_city_tel === true) {
                $('#fromTel').removeClass('d-none').addClass('d-flex');
            } else {
                $('#fromTel').removeClass('d-flex').addClass('d-none');
            }

            $('#to_title').html(response.to.city);

            var toBody = "";

            if (response.to.rule.into_policy) {
                toBody = response.to.rule.into_policy;
            } else {
                toBody = "暂无数据"
            }

            $('#to_data').html(toBody);

            if (response.to.tag) {
                if (response.to.tag === '低风险') {
                    $('#to_tag')
                        .removeClass("border-danger text-danger")
                        .addClass("border-success text-success")
                        .html(response.to.tag);
                } else {
                    $('#to_tag')
                        .removeClass("border-success text-success")
                        .addClass("border-danger text-danger")
                        .html(response.to.tag);
                }
            }

            var toFooter = "";

            if (response.to.rule.out_policy_resource) {
                toFooter = "数据来源：" + response.to.rule.out_policy_resource + " | 更新时间：" +
                    response.to.rule.out_policy_date;
            }

            $('#to_footer').html(toFooter);
        }, 'json');
    }

    $('#currencyClose').click(function () {
        $('#currencyModel').modal('hide');
    });

    $('#fromTel').click(function () {
        $('#telModel').modal('show');

        var map = {
            province: $('#FromP').val(),
            city: $('#FromC').val()
        };

        $('#telContent').html('');

        var listHtml = '';

        $.get("/travel/tel", map, function (response) {
            listHtml += '<div class="list-group">';

            response.data.tels.forEach(function (value) {
                listHtml += '<a href="tel:' + value.phone + '" class="list-group-item list-group-item-action flex-column align-items-start">' +
                    '<h6 class="mb-1 font-weight-bold">' + value.name + '</h6>' +
                    '<div class="text-black-50" style="font-size: 13px;">地址: ' + value.address + '</div>' +
                    '<small class="text-black-50">电话: ' + value.phone + '</small>' +
                    '</a>';
            });

            listHtml += '</div>';

            $('#telContent').html(listHtml);
        }, 'json');
    });

    $('#telClose').click(function () {
        $('#telModel').modal('hide');
    });
});

/*//初始化echarts实例
var myChart = echarts.init(document.getElementById('myEcharts'));
// 指定图表的配置项和数据

//获取数据
//https://view.inews.qq.com/g2/getOnsInfo?name=disease_h5&callback=jQuery34102581268431257997_1582545445186&_=1582545445187
function getData() {
    $.ajax({
        url: "https://view.inews.qq.com/g2/getOnsInfo?name=disease_h5",
        dataType: "jsonp",
        success: function (data) {
            var res = data.data || "";
            res = JSON.parse(res);
            var newArr = [];
            if (res) {
                var province = res.areaTree[0].children;
                for (var i = 0; i < province.length; i++) {
                    var json = {
                        name: province[i].name,
                        value: province[i].total.confirm
                    };
                    newArr.push(json)
                }

                myChart.setOption({
                    title: {
                        text: '中国疫情图',
                        left: 'center'
                    },
                    tooltip: {
                        trigger: 'item'
                    },
                    legend: {
                        orient: 'vertical',
                        left: 'left',
                        data: ['中国疫情图']
                    },
                    visualMap: {
                        type: 'piecewise',
                        pieces: [
                            { min: 1000, max: 1000000, label: '大于等于1000人', color: '#372a28' },
                            { min: 500, max: 999, label: '确诊500-999人', color: '#4e160f' },
                            { min: 100, max: 499, label: '确诊100-499人', color: '#974236' },
                            { min: 10, max: 99, label: '确诊10-99人', color: '#ee7263' },
                            { min: 1, max: 9, label: '确诊1-9人', color: '#f5bba7' },
                        ],
                        color: ['#E0022B', '#E09107', '#A3E00B']
                    },
                    roamController: {
                        show: true,
                        left: 'right',
                        mapTypeControl: {
                            'china': true
                        }
                    },
                    series: [
                        {
                            name: '确诊数',
                            type: 'map',
                            mapType: 'china',
                            roam: false,
                            label: {
                                show: true,
                                color: 'rgb(249, 249, 249)'
                            },
                            data: newArr
                        }
                    ]
                });
            }
        }

    })
}
getData();*/