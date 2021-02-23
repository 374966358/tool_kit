package Ciming

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"

	"tools_more/Ciming/Controller"
	"tools_more/Core"

	"github.com/gin-gonic/gin"

	"tools_more/Common"
)

type MobileCode struct {
	Code int32            `josn:"code"`
	Msg  string           `json:"msg"`
	Data map[string]int64 `json:"data"`
}

type Identity struct {
	ChargeStatus int32             `json:"chargeStatus"`
	Message      string            `json:"message"`
	Data         map[string]string `json:"data"`
	Code         int64             `json:"code"`
}
type Conf struct {
	Wx   *Common.WeChat
	Conf []string
}

func IncAction(r *gin.Engine) {
	r.GET("/yue", func(context *gin.Context) {
		var assign = Conf{}
		var url string

		if find := strings.Contains(context.Request.Proto, "HTTP"); find {
			url = fmt.Sprintf("http://%s%s", context.Request.Host, context.Request.URL.Path)
		} else {
			url = fmt.Sprintf("https://%s%s", context.Request.Host, context.Request.URL.Path)
		}
		assign.Wx = Common.GetWeChat(url)
		now := time.Now()
		//start := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
		//end := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day()+10)
		start := fmt.Sprintf("%d-%02d-%02d", now.Year(), 2, 27)
		end := fmt.Sprintf("%d-%02d-%02d", now.Year(), 3, 10)
		assign.Conf = Controller.GetBetweenDates(start, end)
		context.HTML(http.StatusOK, "ciming_index.html", assign)
	})

	r.POST("/yue/add", func(context *gin.Context) {
		var user Controller.User
		err := context.ShouldBind(&user)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "请正确提交信息"})
		} else {
			code := context.DefaultPostForm("code", "")
			rCode, _ := Core.RedisDb.Get(user.Mobile).Result()

			if rCode == "" {
				context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "请先获取验证码"})
				return
			}

			if code != rCode {
				context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "验证码错误"})
				return
			}

			if len(user.Idcard) < 18 {
				context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "请填写正确的证件号码"})
				return
			}

			if user.DetectTime <= time.Now().Format("2006年01月02") {
				context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "请选择正确的预约时间"})
				return
			}

			// 验证证件与姓名是否匹配
			// postData := Post("https://api-wuser.biya.in/check", map[string]string{"name": user.Realname, "idNum": user.Idcard}, "application/json")
			//
			// var identity Identity
			//
			// json.Unmarshal(postData, &identity)
			//
			// if identity.Data["result"] != "01" {
			// 	context.JSON(http.StatusOK, gin.H{"code": 500, "msg": "姓名与证件信息不符"})
			// 	return
			// }

			state, msg := Controller.Add(user)

			if state == true {
				Core.RedisDb.Del(user.Mobile)

				context.JSON(http.StatusOK, gin.H{"code": 200, "msg": msg})
			} else {
				context.JSON(http.StatusOK, gin.H{"code": 500, "msg": msg})
			}
		}
	})

	r.POST("/yue/sms", func(context *gin.Context) {
		var mobile = context.DefaultPostForm("mobile", "")

		if mobile == "" {
			context.JSON(200, map[string]string{"code": "500", "msg": "请填写电话号码"})
			return
		}

		if (Core.RedisDb.TTL(mobile).Val() + 60) > 900 {
			context.JSON(200, map[string]string{"code": "500", "msg": "60秒发送一次,请勿频繁发送"})
			return
		}

		toSms := fmt.Sprintf("+86%s", mobile)

		var code = Common.Random(6, 0)

		credential := common.NewCredential(
			"AKIDDA0adEjf8KFvJQnp8BxsWu3wXsvhtYaK",
			"dz5uSkmbAVaaImGUuhDshG9lcZHGBVji",
		)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
		client, _ := sms.NewClient(credential, "", cpf)

		request := sms.NewSendSmsRequest()

		request.PhoneNumberSet = common.StringPtrs([]string{toSms})
		request.TemplateParamSet = common.StringPtrs([]string{string(code), "15"})
		request.TemplateID = common.StringPtr("865943")
		request.SmsSdkAppid = common.StringPtr("1400484499")
		request.Sign = common.StringPtr("猫儿网")

		response, err := client.SendSms(request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", response.ToJsonString())
		Core.RedisDb.Set(mobile, string(code), 900*time.Second)
	})
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) []byte {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)

	return result
}
