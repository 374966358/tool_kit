package Cortroller

import (
	"encoding/json"
	"fmt"
	"reflect"
	"tools_more/Wage/Conf"

	"tools_more/Wage/Model"
)

type UserInfo struct {
	Wage      UserWage
	List      map[int]Model.Shui
	C_pay_all float64
	P_pay_all float64
	ZheXian   string
	BingTu    string
	Formdata  Conf.Formdata
}
type UserWage struct {
	Yanglao     Model.Wxyj
	Yiliao      Model.Wxyj
	Shiye       Model.Wxyj
	Gongshang   Model.Wxyj
	Shengyu     Model.Wxyj
	Gongjijin   Model.Wxyj
	BuGongjijin Model.Wxyj
}

func Math(formdata Conf.Formdata) UserInfo {
	UserInfo := UserInfo{}
	UserWage := UserWage{}
	if formdata.Money <= 0 {
		return UserInfo
	}
	wage := formdata.Money               //工资
	Fujiakoukuan := formdata.Param.Extra //附加扣款

	//医疗和生育上下限计算
	Yiliao_base := Model.MathJine(formdata.Param.Is_config, formdata.Param.Yiliao_base, wage, formdata.Param.Yiliao_height, formdata.Param.Yiliao_low)
	UserWage.Yiliao = Model.WxyjAction(Yiliao_base, formdata.Param.P_yiliao_rate, formdata.Param.C_yiliao_rate)
	UserWage.Shengyu = Model.WxyjAction(Yiliao_base, formdata.Param.P_shengyu_rate, formdata.Param.C_shengyu_rate) //生育医疗缴费上下限
	UserWage.Shengyu.P_pay += 3                                                                                    //生育个人缴费3元
	//工伤上下限计算
	Gongshang_base := Model.MathJine(formdata.Param.Is_config, formdata.Param.Gongshang_base, wage, formdata.Param.Gongshang_height, formdata.Param.Gongshang_low)
	UserWage.Gongshang = Model.WxyjAction(Gongshang_base, formdata.Param.P_gongshang_rate, formdata.Param.C_gongshang_rate)
	//失业 上下限计算
	Shiye_base := Model.MathJine(formdata.Param.Is_config, formdata.Param.Shiye_base, wage, formdata.Param.Other_height, formdata.Param.Other_low)
	UserWage.Shiye = Model.WxyjAction(Shiye_base, formdata.Param.P_shiye_rate, formdata.Param.C_shiye_rate)
	//养老 上下限计算
	Yanglao_base := Model.MathJine(formdata.Param.Is_config, formdata.Param.Yanglao_base, wage, formdata.Param.Other_height, formdata.Param.Other_low)
	UserWage.Yanglao = Model.WxyjAction(Yanglao_base, formdata.Param.P_yanglao_rate, formdata.Param.C_yanglao_rate)
	//公积金
	Gongjijin_base := Model.MathJine(formdata.Param.Is_config, formdata.Param.Gongjijin_base, wage, formdata.Param.Other_height, formdata.Param.Other_low)
	UserWage.Gongjijin = Model.WxyjAction(Gongjijin_base, formdata.Param.P_gongjijin_rate, formdata.Param.C_gongjijin_rate)
	if formdata.Param.Is_jin_extra == 1 {
		UserWage.BuGongjijin = Model.WxyjAction(wage, formdata.Param.Jin_extra, formdata.Param.Jin_extra)
	}

	UserInfo.C_pay_all =
		Model.DecimalRoundFix2(
			UserWage.Yanglao.C_pay +
				UserWage.Yiliao.C_pay +
				UserWage.Shiye.C_pay +
				UserWage.Gongshang.C_pay +
				UserWage.Shengyu.C_pay +
				UserWage.Gongjijin.C_pay +
				UserWage.BuGongjijin.C_pay)
	UserInfo.P_pay_all =
		Model.DecimalRoundFix2(
			UserWage.Yanglao.P_pay +
				UserWage.Yiliao.P_pay +
				UserWage.Shiye.P_pay +
				UserWage.Gongshang.P_pay +
				UserWage.Shengyu.P_pay +
				UserWage.Gongjijin.P_pay +
				UserWage.BuGongjijin.P_pay)
	UserInfo.Wage = UserWage
	UserInfo.Formdata = formdata
	UserInfo.List = Model.GeshuiAction(wage, UserInfo.P_pay_all, Fujiakoukuan)

	// 用于显示折线图数据
	var ShuiShowTu = []float64{}
	var LeiJiShuiShowTu = []float64{}
	var ShuiHouGongZiTu = []float64{}

	for _, v := range UserInfo.List {
		ShuiShowTu = append(ShuiShowTu, v.Shui)
		LeiJiShuiShowTu = append(LeiJiShuiShowTu, v.LeijiShui)
		ShuiHouGongZiTu = append(ShuiHouGongZiTu, v.Shuihou)
	}

	ZheXian := [3][]float64{ShuiShowTu, LeiJiShuiShowTu, ShuiHouGongZiTu}

	zheXianData, _ := json.Marshal(ZheXian)

	UserInfo.ZheXian = string(zheXianData)

	// 用于显示圆饼图数据
	BingTu := [2][7]float64{{
		UserWage.Yanglao.C_pay,
		UserWage.Yiliao.C_pay,
		UserWage.Shengyu.C_pay,
		UserWage.Shiye.C_pay,
		UserWage.Gongshang.C_pay,
		UserWage.Gongjijin.C_pay,
		UserWage.BuGongjijin.C_pay,
	}, {
		UserWage.Yanglao.P_pay,
		UserWage.Yiliao.P_pay,
		UserWage.Shengyu.P_pay,
		UserWage.Shiye.P_pay,
		UserWage.Gongshang.P_pay,
		UserWage.Gongjijin.P_pay,
		UserWage.BuGongjijin.P_pay,
	}}

	BingTuData, _ := json.Marshal(BingTu)

	UserInfo.BingTu = string(BingTuData)

	return UserInfo

	// 循环计算指定字段
	/*a := traverse("C_pay", UserWage)
	fmt.Println(a)*/

	// 之前的代码
	/*value := reflect.ValueOf(UserWage)
	for num := 0; num < value.NumField(); num++ {
		ccc := value.Field(num).Interface()
		fmt.Printf("%T", ccc)
	}*/
}

/**
 * name 计算字段
 * target 结构体
 */
func traverse(name string, target interface{}) float64 {
	var value float64 = 0
	// 反射获取结构体value
	sVal := reflect.ValueOf(target)
	// 反射获取结构体key
	sType := reflect.TypeOf(target)

	// 如果是指针类型获取对应的值信息
	if sType.Kind() == reflect.Ptr {
		sVal = sVal.Elem()
		sType = sType.Elem()
	}
	num := sVal.NumField()
	for i := 0; i < num; i++ {
		//判断字段是否为结构体类型，或者是否为指向结构体的指针类型
		if sVal.Field(i).Kind() == reflect.Struct || (sVal.Field(i).Kind() == reflect.Ptr && sVal.Field(i).Elem().Kind() == reflect.Struct) {
			value = traverse(name, sVal.Field(i).Interface()) + value
		} else {
			f := sType.Field(i)
			val := sVal.Field(i).Float()
			if f.Name == name {
				value = value + float64(val)
			}
			fmt.Printf("%5s %v = %v\n", f.Name, f.Type, val)
		}
	}

	return value
}
