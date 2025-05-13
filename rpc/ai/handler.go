package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/config"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/database"
	ai "github.com/Practical-Training-IOT/IOT-C/kitex_gen/iot/ai"
	"io"
	"net/http"
	"time"
)

// AiImpl implements the last service interface defined in the IDL.
type AiImpl struct{}

var Model = "deepseek-r1"

type OllamaResponse struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	DoneReason         string `json:"done_reason,omitempty"`
	Context            []int  `json:"context,omitempty"`
	TotalDuration      int64  `json:"total_duration,omitempty"`
	LoadDuration       int64  `json:"load_duration,omitempty"`
	PromptEvalCount    int    `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration int    `json:"prompt_eval_duration,omitempty"`
	EvalCount          int    `json:"eval_count,omitempty"`
	EvalDuration       int64  `json:"eval_duration,omitempty"`
}

// OneHistory implements the AiImpl interface.
func (s *AiImpl) OneHistory(ctx context.Context, req *ai.OneHistoryRequest) (resp *ai.OneHistoryResponse, err error) {
	// TODO: Your code here...
	var aiScen database.AiScene
	err = config.DB.Where("id = ? and user_id=?", req.Id, req.UserId).Find(&aiScen).Error
	if err != nil {
		return
	}
	var a []database.Ai
	config.DB.Where("ai_scen_id = ?", req.Id).Find(&a)
	var sli []*ai.ChatOneResponse
	for _, v := range a {
		one := &ai.ChatOneResponse{
			Message: v.Res,
			Model:   v.Model,
			Req:     v.Req,
		}
		sli = append(sli, one)
	}
	one := &ai.OneHistoryResponse{
		Chat: sli,
		Id:   aiScen.ID,
	}
	return one, nil
}

// Chat implements the AiImpl interface.
func (s *AiImpl) Chat(ctx context.Context, req *ai.ChatRequest) (resp *ai.ChatResponse, err error) {
	// TODO: Your code here...
	fmt.Println("开始聊天了", Model)
	requestBody := map[string]string{
		"model":  Model,
		"prompt": req.Message,
	}

	jsonBody, _ := json.Marshal(requestBody)

	fmt.Println(string(jsonBody))

	// 发起请求
	post, err := http.Post("http://117.27.231.112:11434/api/generate", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer post.Body.Close()

	var fullResponse string
	decoder := json.NewDecoder(post.Body)

	for {
		var r OllamaResponse
		if err := decoder.Decode(&r); err == io.EOF {
			break
		} else if err != nil {

		}
		fullResponse += r.Response
		fmt.Println(r.Response)
		if r.Done {
			break
		}
	}

	db := config.DB.Begin()

	var aiScene database.AiScene

	var aiSceneId int64

	if req.Id == 0 {
		aiScene = database.AiScene{
			Title:     req.Message,
			CreatedAt: time.Now(),
			UserId:    int64(req.UserId),
		}
		err = db.Create(&aiScene).Error
		if err != nil {
			db.Rollback()
			return
		}
		aiSceneId = aiScene.ID
	} else {
		aiSceneId = int64(req.Id)
	}

	a := database.Ai{
		Model:    Model,
		Req:      req.Message,
		Res:      fullResponse,
		Title:    req.Message,
		AiScenId: aiSceneId,
		UserId:   int64(req.UserId),
	}

	err = db.Create(&a).Error

	if err != nil {
		db.Rollback()
		return
	}

	db.Commit()

	return &ai.ChatResponse{
		Message: fullResponse,
		Model:   Model,
		Id:      aiSceneId,
	}, nil
}

// Change implements the AiImpl interface.
func (s *AiImpl) Change(ctx context.Context, req *ai.ChangeRequest) (resp *ai.ChangeResponse, err error) {
	// TODO: Your code here...
	Model = req.Message
	fmt.Println(Model)
	return &ai.ChangeResponse{}, nil
}

// History implements the AiImpl interface.
func (s *AiImpl) History(ctx context.Context, req *ai.HistoryRequest) (resp *ai.HistoryResponse, err error) {
	// TODO: Your code here...
	var a []database.AiScene
	err = config.DB.Where("user_id=?", req.UserId).Find(&a).Error
	if err != nil {
		return
	}
	var aiSli []*ai.History
	for _, i := range a {
		one := &ai.History{
			Id:        int32(i.ID),
			Title:     i.Title,
			CreatedAt: i.CreatedAt.String(),
		}
		aiSli = append(aiSli, one)
	}
	// 返回响应给前端
	return &ai.HistoryResponse{List: aiSli}, nil
}
