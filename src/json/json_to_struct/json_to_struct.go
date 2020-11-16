package main

import (
	"encoding/json"
	"fmt"
)

type DebugInfo struct {
	Aqi string
	date string
	fl string
	fx string
	high string
	low string
	notice string
	sunrise string
	sunset string
	typ string
}

func (dbgInfo DebugInfo) String() string {
	return fmt.Sprintf("{Aqi: %s, date: %s, fl: %s, fx: %s, high: %s, low: %s, notice: %s, sunrise: %s, sunset: %s, typ: %s,}", dbgInfo.Aqi , dbgInfo.date, dbgInfo.fl, dbgInfo.fx, dbgInfo.high, dbgInfo.low, dbgInfo.notice, dbgInfo.sunrise, dbgInfo.sunset, dbgInfo.typ)
}

func main()  {
	data := `[{"date":"22日星期四","sunrise":"06:17","high":"高温 17.0℃","low":"低温 1.0℃","sunset":"18:27","aqi":"98","fx":"西南风","fl":"<3级","typ":"晴","notice":"愿你拥有比阳光明媚的心情"},{"date":"23日星期五","sunrise":"06:16","high":"高温 18.0℃","low":"低温 5.0℃","sunset":"18:28","aqi":"118","fx":"无持续风向","fl":"<3级","typ":"多云","notice":"阴晴之间，谨防紫外线侵扰"},{"date":"24日星期六","sunrise":"06:14","high":"高温 21.0℃","low":"低温 7.0℃","sunset":"18:29","aqi":"52","fx":"西南风","fl":"<3级","typ":"晴","notice":"愿你拥有比阳光明媚的心情"},{"date":"25日星期日","sunrise":"06:13","high":"高温 22.0℃","low":"低温 7.0℃","sunset":"18:30","aqi":"71","fx":"西南风","fl":"<3级","typ":"晴","notice":"愿你拥有比阳光明媚的心情"},{"date":"26日星期一","sunrise":"06:11","high":"高温 21.0℃","low":"低温 8.0℃","sunset":"18:31","aqi":"97","fx":"西南风","fl":"<3级","typ":"多云","notice":"阴晴之间，谨防紫外线侵扰"}]`
	var dbgInfos []DebugInfo
	json.Unmarshal([]byte(data), &dbgInfos)
	fmt.Println(dbgInfos)
}