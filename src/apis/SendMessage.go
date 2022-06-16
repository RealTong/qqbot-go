package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notify-bot/src/dto"
	"notify-bot/src/service"
)

func DroneWebHook(w http.ResponseWriter, r *http.Request) {
	// 释放资源
	defer r.Body.Close()

	length := r.ContentLength

	body := make([]byte, length)

	r.Body.Read(body)

	requestBody := dto.DroneWebHookEntity{}
	json.Unmarshal(body, &requestBody)

	fmt.Println(requestBody)
	fmt.Println(requestBody.Event)

	service.SuccessNotion(&requestBody)

	fmt.Fprintf(w, "%#v\n", requestBody) // 格式化输出结果
	//转换为json格式
	jsonStr, _ := json.Marshal(requestBody)
	//打印转换为json的结果
	fmt.Println(string(jsonStr))

	w.Write([]byte("Hi Drone the message from Golang build by shutong !"))
}
