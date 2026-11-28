package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// 1：定义配置
type Config struct {
	APIKey  string
	BaseURL string
	Model   string
}

type Message struct{
	Role string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct{
	Model string `json:"model"`
	Messages []Message `json:"messages"`
}

func reqiureEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New("missing required env:" + key)
	}
	val = strings.TrimSpace(val)
	if val == "" {
		return "", errors.New("empty required env:" + key)
	}
	return val, nil
}

// 2:读取配置的函数
func loadConfig() (Config, error) {
	// 读取APIKEY
	apiKey, err := reqiureEnv("DEEPSEEK_API_KEY")
	if err != nil {
		return Config{}, err
	}
	// 读URL
	BaseURL, err := reqiureEnv("DEEPSEEK_BASE_URL")
	if err != nil {
		return Config{}, err
	}
	// 读取模型
	model, err := reqiureEnv("MODEL")
	if err != nil {
		return Config{}, err
	}
	// 把所有信息拼起来并返回config，如果有一项没有填入就直接报错
	return Config{
		APIKey:  apiKey,
		BaseURL: BaseURL,
		Model:   model,
	}, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf(
		"BaseURL: %s\n"+
			"model: %s\n"+
			"APIKey length: %d\n",
		cfg.BaseURL,
		cfg.Model,
		len(cfg.APIKey),
	)
}

// 开始写模型接口
func sendChat(cfg Config, prompt string)(string,error){
reqbody := []ChatRequest{
		Model: cfg.Model,
		Messages: (
			Role:"user",
			Content: prompt
		)
	}
reqBytes,err := json.Marshal(reqBody)
return "", err
if err != nil {
	

// 写chat/completions URL
endpoint := strings.TrimRight(cfg.BaseURL,"/") + "/chat/completions"
}


