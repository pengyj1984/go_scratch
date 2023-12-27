/*
 */
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type DailyData struct {
	FxDate         string `json:"fxDate"`         // 预报日期
	Sunrise        string `json:"sunrise"`        // 日出时间，在高纬度地区可能为空
	Sunset         string `json:"sunset"`         // 日落时间，在高纬度地区可能为空
	Moonrise       string `json:"moonrise"`       // 当天月升时间，可能为空
	Moonset        string `json:"moonset"`        // 当天月落时间，可能为空
	MoonPhase      string `json:"moonPhase"`      // 月相名称
	MoonPhaseIcon  string `json:"moonPhaseIcon"`  // 月相图标代码，另请参考天气图标项目
	TempMax        string `json:"tempMax"`        // 预报当天最高温度
	TempMin        string `json:"tempMin"`        // 预报当天最低温度
	IconDay        string `json:"iconDay"`        // 预报白天天气状况的图标代码，另请参考天气图标项目
	TextDay        string `json:"textDay"`        // 预报白天天气状况文字描述，包括阴晴雨雪等天气状态的描述
	IconNight      string `json:"iconNight"`      // 预报夜间天气状况的图标代码，另请参考天气图标项目
	TextNight      string `json:"textNight"`      // 预报晚间天气状况文字描述，包括阴晴雨雪等天气状态的描述
	Wind360Day     string `json:"wind360day"`     // 预报白天风向360角度
	WindDirDay     string `json:"windDirDay"`     // 预报白天风向
	WindScaleDay   string `json:"windScaleDay"`   // 预报白天风力等级
	WindSpeedDay   string `json:"windSpeedDay"`   // 预报白天风速，公里/小时
	Wind360Night   string `json:"wind360Night"`   // 预报夜间风向360角度
	WindDirNight   string `json:"windDirNight"`   // 预报夜间当天风向
	WindScaleNight string `json:"windScaleNight"` // 预报夜间风力等级
	WindSpeedNight string `json:"windSpeedNight"` // 预报夜间风速，公里/小时
	Precip         string `json:"precip"`         // 预报当天总降水量，默认单位：毫米
	UvIndex        string `json:"uvIndex"`        // 紫外线强度指数
	Humidity       string `json:"humidity"`       // 相对湿度，百分比数值
	Pressure       string `json:"pressure"`       // 大气压强，默认单位：百帕
	Vis            string `json:"vis"`            // 能见度，默认单位：公里
	Cloud          string `json:"cloud"`          // 云量，百分比数值。可能为空
}

type ReferData struct {
	Sources []string `json:"sources"` // 原始数据来源，或数据源说明，可能为空
	License []string `json:"license"` // 数据许可或版权声明，可能为空
}

type Data struct {
	Code       string      `json:"code"`       // 请参考状态码
	UpdateTime string      `json:"updateTime"` // 当前API的最近更新时间
	fxLink     string      `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Daily      []DailyData `json:"daily"`
	Refer      ReferData   `json:"refer"`
}

func main() {
	url := fmt.Sprintf("https://devapi.qweather.com/v7/weather/7d?key=%s&location=%s", "3fec77b74e8d4d3a9b739bd7b88611ad", "101040100")
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("get error:", err.Error())
		return
	}

	httpRsp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Println("do error:", err.Error())
		return
	}
	defer httpRsp.Body.Close()

	rspBody, err := io.ReadAll(httpRsp.Body)
	if err != nil {
		log.Println("read error:", err.Error())
		return
	}

	log.Println("rspBody =", string(rspBody))
}
