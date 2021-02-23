package Model

import (
	"math"
	"strconv"
)

// 只要是对外展示的一定注意首字母是大写的
type Wxyj struct {
	Wage   float64
	P_rate float64
	P_pay  float64
	C_rate float64
	C_pay  float64
}

func WxyjAction(wage float64, p_rate float64, c_rate float64) Wxyj {
	return Wxyj{
		Wage:   wage,
		P_rate: p_rate,
		P_pay:  DecimalRoundFix2(wage / 100 * p_rate),
		C_rate: c_rate,
		C_pay:  DecimalRoundFix2(wage / 100 * c_rate),
	}
}

type Shui struct {
	Gongzi        float64 //工资
	LeijiGongzi   float64 //历时月累计工资
	Wxyj          float64 //个人五险一金总金额
	LeijiWxyj     float64 //历时月累计五险一金总金额
	Fujia         float64 //附加额
	LeijiFujia    float64 //历时月累计附加总金额
	Koushuie      float64 //扣税额度
	LeijiKoushuie float64 //历时月累计扣税总金额
	Shui          float64 //当月所交税额
	LeijiShui     float64 //历时月累计总税额
	Shuilv        float64 //当月所执行税率
	Sukou         float64 //速算扣除数
	Shuihou       float64 //月税后工资
	Month         int     //月税后工资
}

func GeshuiAction(wage float64, wxyj float64, fujia float64) map[int]Shui {
	List := make(map[int]Shui)
	for month := 1; month < 13; month++ {
		month_num := float64(month)
		Shui_month := Shui{}
		Shui_month.Gongzi = wage
		Shui_month.Wxyj = wxyj
		Shui_month.Fujia = fujia
		if month == 1 {
			Shui_month.LeijiGongzi = wage
			Shui_month.LeijiWxyj = wxyj
			Shui_month.LeijiFujia = fujia
			Shui_month.LeijiKoushuie = 0
			Shui_month.LeijiShui = 0
		} else {
			//累计工资等于，上个月累计加上本月
			Shui_month.LeijiGongzi = DecimalRoundFix2(List[month-1].LeijiGongzi + wage)
			//累计五险一金等于，上个月累计加上本月
			Shui_month.LeijiWxyj = DecimalRoundFix2(List[month-1].LeijiWxyj + wxyj)
			//累计附加等于，上个月累计加上本月
			Shui_month.LeijiFujia = DecimalRoundFix2(List[month-1].LeijiFujia + fujia)
			//上月累计交税额
			Shui_month.LeijiKoushuie = DecimalRoundFix2(List[month-1].LeijiKoushuie)
			//上月累计已交税
			Shui_month.LeijiShui = DecimalRoundFix2(List[month-1].LeijiShui)
		}

		Shui_month.Koushuie = (Shui_month.Gongzi * month_num) - (5000 * month_num) - (Shui_month.Wxyj * month_num) - (Shui_month.Fujia * month_num)

		Shui_month.LeijiKoushuie += Shui_month.Gongzi //追加本月扣税额度
		switch {
		case Shui_month.Koushuie <= 36000:
			Shui_month.Shuilv = 3
			Shui_month.Sukou = 0
		case Shui_month.Koushuie <= 144000:
			Shui_month.Shuilv = 10
			Shui_month.Sukou = 2520
		case Shui_month.Koushuie <= 300000:
			Shui_month.Shuilv = 20
			Shui_month.Sukou = 16920
		case Shui_month.Koushuie <= 420000:
			Shui_month.Shuilv = 25
			Shui_month.Sukou = 31920
		case Shui_month.Koushuie <= 660000:
			Shui_month.Shuilv = 30
			Shui_month.Sukou = 52920
		case Shui_month.Koushuie <= 960000:
			Shui_month.Shuilv = 35
			Shui_month.Sukou = 85920
		default:
			Shui_month.Shuilv = 45
			Shui_month.Sukou = 181920
		}
		Shui_month.Shuilv = Shui_month.Shuilv / 100
		//本月要交税 = 扣税额*税率-速算扣除数-累计交过的税
		Shui_month.Shui = DecimalRoundFix2(Shui_month.Koushuie*Shui_month.Shuilv - Shui_month.Sukou - Shui_month.LeijiShui)
		if Shui_month.Shui < 0 {
			Shui_month.Shui = 0
		}
		Shui_month.LeijiShui = DecimalRoundFix2(Shui_month.Shui + Shui_month.LeijiShui)              //追加本月累计交税
		Shui_month.Shuihou = DecimalRoundFix2(Shui_month.Gongzi - Shui_month.Wxyj - Shui_month.Shui) //追加本月扣税额度
		Shui_month.Month = month
		List[month] = Shui_month
		//fmt.Println(Shui_month)
	}
	return List
}
func DecimalRoundFix2(f float64) float64 {
	f1 := math.Trunc(f*1e2+0.5) * 1e-2
	f1Str := strconv.FormatFloat(f1, 'f', 2, 64)
	value, _ := strconv.ParseFloat(f1Str, 64)
	return value
}
func MathJine(Is_config float64, Base float64, Wage float64, Height float64, Low float64) float64 {
	if Is_config != 1 && Base <= 0 {
		Base = Wage
	}
	if Base > Height {
		Base = Height
	}
	if Base < Low {
		Base = Low
	}
	return Base
}
