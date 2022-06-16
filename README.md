# Note


> 将请求中的body转换为json格式
```go
length := r.ContentLength
body := make([]byte, length)
r.Body.Read(body)

requestBody := dto.DroneWebHookEntity{}
json.Unmarshal(body, &requestBody)

```


> 将请求中的body转换为struct格式
```go
jsonStr, _ := json.Marshal(requestBody)

fmt.Println(string(jsonStr))
```