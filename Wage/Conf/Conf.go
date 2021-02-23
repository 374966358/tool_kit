package Conf

type Formdata struct {
	List  string
	City  string
	Money float64
	Param FormParam
}
type FormParam struct {
	Gongshang_base   float64 `form:"gongshang_base" binding:"numeric,max=100,min=100"`
	Shengyu_base     float64 `form:"shengyu_base" binding:"numeric,max=100,min=100"`
	Yanglao_base     float64 `form:"yanglao_base" binding:"numeric,max=100,min=100"`
	Shiye_base       float64 `form:"shiye_base" binding:"numeric,max=100,min=100"`
	Yiliao_base      float64 `form:"yiliao_base" binding:"numeric,max=100,min=100"`
	Gongjijin_base   float64 `form:"gongjijin_base" binding:"numeric,max=100,min=100"`
	P_gongshang_rate float64 `form:"p_gongshang_rate" binding:"numeric,max=100,min=100"`
	C_gongshang_rate float64 `form:"c_gongshang_rate" binding:"numeric,max=100,min=100"`
	C_gongjijin_rate float64 `form:"c_gongjijin_rate" binding:"numeric,max=100,min=100"`
	P_gongjijin_rate float64 `form:"p_gongjijin_rate" binding:"numeric,max=100,min=100"`
	C_yanglao_rate   float64 `form:"c_yanglao_rate" binding:"numeric,max=100,min=100"`
	P_yanglao_rate   float64 `form:"p_yanglao_rate" binding:"numeric,max=100,min=100"`
	P_yiliao_rate    float64 `form:"p_yiliao_rate" binding:"numeric,max=100,min=100"`
	C_yiliao_rate    float64 `form:"c_yiliao_rate" binding:"numeric,max=100,min=100"`
	P_shiye_rate     float64 `form:"p_shiye_rate" binding:"numeric,max=100,min=100"`
	C_shiye_rate     float64 `form:"c_shiye_rate" binding:"numeric,max=100,min=100"`
	C_shengyu_rate   float64 `form:"c_shengyu_rate" binding:"numeric,max=100,min=100"`
	P_shengyu_rate   float64 `form:"p_shengyu_rate" binding:"numeric,max=100,min=100"`
	Is_jin_extra     float64 `form:"is_jin_extra" binding:"numeric,max=100,min=100"`
	Jin_extra        float64 `form:"jin_extra" binding:"numeric,max=100,min=100"`
	Is_config        float64 `form:"is_config" binding:"numeric,max=100,min=100"`
	Extra            float64 `form:"extra" binding:"numeric,max=100,min=100"`
	Yiliao_height    float64
	Yiliao_low       float64
	Gongshang_height float64
	Gongshang_low    float64
	Other_height     float64
	Other_low        float64
}

func Beijing() FormParam {
	return FormParam{
		Gongshang_base:   4713.00,
		Shengyu_base:     5360.00,
		Yanglao_base:     3613.00,
		Shiye_base:       3613.00,
		Yiliao_base:      5360.00,
		Gongjijin_base:   1540,
		C_gongshang_rate: 0.5,
		P_gongshang_rate: 0,
		C_gongjijin_rate: 12,
		P_gongjijin_rate: 12,
		C_yanglao_rate:   16,
		P_yanglao_rate:   8,
		P_yiliao_rate:    2,
		C_yiliao_rate:    9.8,
		P_shiye_rate:     0.2,
		C_shiye_rate:     0.8,
		P_shengyu_rate:   0,
		C_shengyu_rate:   1,
		Is_jin_extra:     0,
		Jin_extra:        8,
		Is_config:        0,
		Extra:            0,
		Yiliao_height:    29732.00,
		Yiliao_low:       5360.00,
		Gongshang_height: 26541.00,
		Gongshang_low:    4713.00,
		Other_height:     26541.00,
		Other_low:        3613.00,
	}
}
func Other() FormParam {
	return FormParam{
		Gongshang_base:   4713.00,
		Shengyu_base:     5360.00,
		Yanglao_base:     3613.00,
		Shiye_base:       3613.00,
		Yiliao_base:      5360.00,
		Gongjijin_base:   1540,
		C_gongshang_rate: 0.5,
		P_gongshang_rate: 0,
		C_gongjijin_rate: 12,
		P_gongjijin_rate: 12,
		C_yanglao_rate:   16,
		P_yanglao_rate:   8,
		P_yiliao_rate:    2,
		C_yiliao_rate:    9.8,
		P_shiye_rate:     0.2,
		C_shiye_rate:     0.8,
		P_shengyu_rate:   0,
		C_shengyu_rate:   1,
		Is_jin_extra:     0,
		Jin_extra:        8,
		Is_config:        0,
		Extra:            0,
	}
}
