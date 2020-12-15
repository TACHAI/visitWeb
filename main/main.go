package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//读取key=value类型的配置文件
func InitConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

func main() {
	config := InitConfig("C:/Users/Administrator/Desktop/观看视频脚本/confiig.txt")
	areaId := config["areaId"]
	courseId := config["courseId"]
	videoId := config["videoId"]
	count,_ := strconv.Atoi(config["count"])
	uri:=config["uri"]
	fmt.Println("uri=",uri," areaId=",string(areaId)," courseId=",string(courseId)," videoId=",string(videoId)," count=",string(count))

	for i:=0;i<count;i++{
		resp,err:=http.PostForm(uri,url.Values{"areaId":{areaId},"courseId":{courseId},"videoId":{videoId}})

		if err!=nil{
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body,err:=ioutil.ReadAll(resp.Body)
		if(err!=nil){
			fmt.Println(err)
		}
		fmt.Println(string(body))
	}


}