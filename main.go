package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// 1：定义配置
type Config struct {
APIKey string
BaseURL string
Model string	
}


func reqiureEnv(Key string) (string,error){
val.ok :=os.LookupEnv(key)
if !ok {return Config(errors.New(missing required env.))}
val :=strings.TrimSpace(val)
if val == "" {return Config(errors.New(empty required env.))}
return val,nil
}

// 2:读取配置的函数
func loadConfig()(Config,error){
// 读取APIKEY，如果OK就放行，不OK就throw err
rawKey,ok :=os.LookupEnv("DEEPSEEK_API_KEY")
apikey := strings.TrimSpace(rawKey)
// 读URL
baseURL := strings.TrimSpace(os.Getenv("DEEPSEEK_BASE_URL"))
// 读取模型
model := strings.TrimSpace(os.Getenv("MODEL")) 
// 把所有信息拼起来并返回config，如果有一项没有填入就直接报错
 return {Config{"DEEPSEEK_API_KEY","DEEPSEEK_BASE_URL"}}
 error nil
}

func (main)  {
cfg,err := loadConfig()
if errr
}