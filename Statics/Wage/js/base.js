var customConfig = false;

$("#customCheck").click(function () {
    if ($(this).prop("checked")) {
        customConfig = true;
        $('#yanglao_base').val($('#moneyAll').val())
        $('#yiliao_base').val($('#moneyAll').val())
        $('#shiye_base').val($('#moneyAll').val())
        $('#gongshang_base').val($('#moneyAll').val())
        $('#shengyu_base').val($('#moneyAll').val())
        $('#gongjijin_base').val($('#moneyAll').val())
        $("#customForm").show()
    } else {
        customConfig = false;
        $("#customForm").hide()
    }
});

var bingValue = $('#bingTu').val();
var zheValue = $('#zheXian').val();

if (bingValue) {
    let bingTu = JSON.parse(bingValue);

    var geRen = echarts.init(document.getElementById('ge-ren'));
    var qiYe = echarts.init(document.getElementById('qi-ye'));

    let geRenData = [
        {value: bingTu[0][0], name: '养老保险金'},
        {value: bingTu[0][1], name: '医疗保险金'},
        {value: bingTu[0][2], name: '生育保险金'},
        {value: bingTu[0][3], name: '失业保险金'},
        {value: bingTu[0][4], name: '工伤保险金'},
        {value: bingTu[0][5], name: '基本住房公积金'},
        {value: bingTu[0][6], name: '补充住房公积金'}
    ];

    let qiYeData = [
        {value: bingTu[1][0], name: '养老保险金'},
        {value: bingTu[1][1], name: '医疗保险金'},
        {value: bingTu[1][2], name: '生育保险金'},
        {value: bingTu[1][3], name: '失业保险金'},
        {value: bingTu[1][4], name: '工伤保险金'},
        {value: bingTu[1][5], name: '基本住房公积金'},
        {value: bingTu[1][6], name: '补充住房公积金'}
    ];

    geRen.setOption({
        title: {text: '个人税前工资去向', left: 'center'},
        tooltip: {trigger: 'item', formatter: '{b} : {c} ({d}%)'},
        series: [{type: 'pie', radius: '55%', center: ['50%', '60%'], data: geRenData}]
    });

    qiYe.setOption({
        title: {text: '公司税前工资去向', left: 'center'},
        tooltip: {trigger: 'item', formatter: '{b} : {c} ({d}%)'},
        series: [{type: 'pie', radius: '55%', center: ['50%', '60%'], data: qiYeData}]
    });

    geRen.resize();
    qiYe.resize();
}

if (zheValue) {
    let zheXian = JSON.parse(zheValue);

    var quXian = echarts.init(document.getElementById('qu-xian'));

    quXian.setOption({
        tooltip: {trigger: 'axis'},
        legend: {data: ['税收', '累计税收', '税后工资']},
        grid: {left: '1%', right: '2%', bottom: '1%', containLabel: true},
        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月', '十月', '十一月', '十二月']
        },
        yAxis: {type: 'value'},
        series: [
            {name: '税收', type: 'line', data: zheXian[0]},
            {name: '累计税收', type: 'line', data: zheXian[1]},
            {name: '税后工资', type: 'line', data: zheXian[2]}
        ]
    });

    quXian.resize();
}

$('#mathButton').click(function () {
    money = $("#moneyAll").val();
    if(!money){
        alert("正确填写个人税前工资")
    }else{
        city = "beijing"
        $("#mathform").attr('action','/city/'+city+'/'+money)
        $("#mathform").submit()
    }
});
