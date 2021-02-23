package Model

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type MyJsonName struct {
	Code int64 `json:"code"`
	Data struct {
		From struct {
			City       string `json:"city"`
			FromCity   string `json:"fromCity"`
			HasCityTel bool   `json:"has_city_tel"`
			Icon       string `json:"icon"`
			Notice     struct {
				Bobao struct {
					Desc  string `json:"desc"`
					Icon  string `json:"icon"`
					Title string `json:"title"`
					URL   string `json:"url"`
				} `json:"bobao"`
				Jingqu struct {
					Desc  string `json:"desc"`
					Icon  string `json:"icon"`
					Title string `json:"title"`
					URL   string `json:"url"`
				} `json:"jingqu"`
			} `json:"notice"`
			Province     string `json:"province"`
			ProvinceData struct {
				Danger struct {
					Zero int64 `json:"0"`
					One  int64 `json:"1"`
					Two  int64 `json:"2"`
				} `json:"danger"`
				Deadline string `json:"deadline"`
				List     []struct {
					Name       string `json:"name"`
					Present    string `json:"present"`
					SureNewCnt string `json:"sure_new_cnt"`
					SureNewHid string `json:"sure_new_hid"`
					SureNewLoc string `json:"sure_new_loc"`
				} `json:"list"`
				URL string `json:"url"`
			} `json:"provinceData"`
			Rule struct {
				OutPolicy         string `json:"out_policy"`
				OutPolicyDate     string `json:"out_policy_date"`
				OutPolicyResource string `json:"out_policy_resource"`
			} `json:"rule"`
			Tag string `json:"tag"`
		} `json:"from"`
		To struct {
			City        string `json:"city"`
			Icon        string `json:"icon"`
			HasCityTel  bool   `json:"has_city_tel"`
			IsDanger    bool   `json:"isDanger"`
			IsMidDanger bool   `json:"isMidDanger"`
			LocData     struct {
				Danger struct {
					One int64 `json:"1"`
					Two int64 `json:"2"`
				} `json:"danger"`
				Deadline string `json:"deadline"`
				List     []struct {
					Name       string `json:"name"`
					Present    string `json:"present"`
					SureNewCnt string `json:"sure_new_cnt"`
					SureNewHid string `json:"sure_new_hid"`
					SureNewLoc string `json:"sure_new_loc"`
				} `json:"list"`
				URL string `json:"url"`
			} `json:"locData"`
			Province string `json:"province"`
			Rule     struct {
				City                    string `json:"city"`
				ID                      string `json:"id"`
				IntoPolicy              string `json:"into_policy"`
				IntoPolicyDate          string `json:"into_policy_date"`
				IntoPolicyResource      string `json:"into_policy_resource"`
				OutPolicy               string `json:"out_policy"`
				OutPolicyDate           string `json:"out_policy_date"`
				OutPolicyResource       string `json:"out_policy_resource"`
				OutsideBorder           string `json:"outside_border"`
				OutsideBorderDate       string `json:"outside_border_date"`
				OutsideBorderResource   string `json:"outside_border_resource"`
				Province                string `json:"province"`
				TrafficAirplane         string `json:"traffic_airplane"`
				TrafficAirplaneDate     string `json:"traffic_airplane_date"`
				TrafficAirplaneResource string `json:"traffic_airplane_resource"`
				TrafficBus              string `json:"traffic_bus"`
				TrafficBusDate          string `json:"traffic_bus_date"`
				TrafficBusResource      string `json:"traffic_bus_resource"`
				TrafficCar              string `json:"traffic_car"`
				TrafficCarDate          string `json:"traffic_car_date"`
				TrafficCarResource      string `json:"traffic_car_resource"`
				TrafficShip             string `json:"traffic_ship"`
				TrafficShipDate         string `json:"traffic_ship_date"`
				TrafficShipResource     string `json:"traffic_ship_resource"`
				TrafficTrain            string `json:"traffic_train"`
				TrafficTrainDate        string `json:"traffic_train_date"`
				TrafficTrainResource    string `json:"traffic_train_resource"`
				WithinBorder            string `json:"within_border"`
				WithinBorderDate        string `json:"within_border_date"`
				WithinBorderResource    string `json:"within_border_resource"`
			} `json:"rule"`
			Tag    string `json:"tag"`
			ToCity string `json:"toCity"`
		} `json:"to"`
		Zfb []struct {
			IconPic   string `json:"icon_pic"`
			ID        string `json:"id"`
			Link      string `json:"link"`
			RankScore string `json:"rank_score"`
			Title     string `json:"title"`
		} `json:"zfb"`
	} `json:"data"`
	Msg    string `json:"msg"`
	Status int64  `json:"status"`
}

type MyTel struct {
	Code int64 `json:"code"`
	Data struct {
		Tels []struct {
			Address  string `json:"address"`
			City     string `json:"city"`
			District string `json:"district"`
			Id       string `json:"id"`
			Name     string `json:"name"`
			Phone    string `json:"phone"`
			Province string `json:"province"`
		} `json:"tels"`
	} `json:"data"`
	Msg    string `json:"msg"`
	Status int64  `json:"status"`
}

func DoMchrApi(url string) ([]byte, error) {
	client := http.Client{Timeout: time.Second * 60}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", "https://vt.sm.cn", url), nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		} else {
			return body, nil
		}
	}
	return nil, nil
}
