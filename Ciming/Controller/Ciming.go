package Controller

import (
	"strconv"
	"time"

	"tools_more/Core"
)

type User struct {
	Realname    string `gorm:"" form:"realname" binding:"required"`
	Gender      string `gorm:""`
	Mobile      string `gorm:"" form:"mobile"  binding:"required,numeric"`
	Idcard      string `gorm:"" form:"idcard"  binding:"required"`
	DetectTime  string `gorm:"" form:"detect_time" binding:"required"`
	DetectState int    `gorm:""`
	Excel       int    `gorm:""`
}

func (User) TableName() string {
	return "user"
}

func Add(info User) (bool, string) {
	var Tmp []User
	var gender = []byte(info.Idcard)
	genderInt, _ := strconv.Atoi(string(gender[16]))
	var sex string
	if genderInt%2 == 0 {
		sex = "女"
	} else {
		sex = "男"
	}
	var detectTimeCount int64
	Core.GetDB("")
	Core.DB.Model(&User{}).Where("detect_time = ?", info.DetectTime).Count(&detectTimeCount)
	if detectTimeCount > 10 {
		return false, "该日预约数量已满"
	}
	sql := "select * from user where detect_state = 1 and detect_time = ? and mobile = ? and Idcard = ? and realname = ? and gender = ?"
	Core.DB.Raw(sql, info.DetectTime, info.Mobile, info.Idcard, info.Realname, sex).Scan(&Tmp)
	if len(Tmp) > 0 {
		return false, "已经申请,请等待短信通知"
	}
	time := time.Now().Format("2006年01月02日 15:04:05")
	sql = "INSERT INTO `user` ( `realname`, `gender`, `mobile`, `Idcard`, `add_time`, `detect_time`, `detect_state`, `excel`) VALUES (?, ?, ?, ?, ?,?, ?, ?);"
	Core.DB.Exec(sql, info.Realname, sex, info.Mobile, info.Idcard, time, info.DetectTime, "1", "0")
	//定义收件人
	mailTo := []string{
		"evelyn6161@163.com",
		"261738244@qq.com",
		"374966358@qq.com",
	}
	//邮件主题为"Hello"
	subject := "新的核酸预约，请同步信息"
	// 邮件正文
	body := "申请人：" + info.Realname +
		"\r\n<br>性别：" + sex +
		"\r\n<br>身份证号：" + info.Idcard +
		"\r\n<br>电话号码：" + info.Mobile +
		"\r\n<br>预约时间：" + info.DetectTime +
		"\r\n<br>申请时间：" + time
	Core.SendMail(mailTo, subject, body)
	return true, "预约申请已经提交，审核后短信通知预约结果，请及时查收"
}

func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006年01月02日"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}
