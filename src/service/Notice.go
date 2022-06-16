package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"notify-bot/src/dto"
)

func errorNotice() string {
	return "error"
}
func warnNotion() string {
	return "warn"
}
func infoNotion() string {
	return "info"
}
func SuccessNotion(entity *dto.DroneWebHookEntity) {
	jsonStr, _ := json.Marshal(dto.GroupMsg{
		GroupId: 1011111,
		//Msg:     "[CQ:face,id=124]构建成功 地址:" + entity.Link,
		Msg: "[CQ:share,url=" + entity.Link + ",title=Drone]构建成功",
	})
	post, _ := http.Post("http://localhost:7778/send_group_msg", "application/json", bytes.NewBuffer(jsonStr))
	fmt.Printf("响应体长度是: %s", post.ContentLength)
}
func failNotion() string {
	return "fail"
}
