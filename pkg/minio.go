package pkg

import (
	"context"
	"fmt"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/config"
	"github.com/Practical-Training-IOT/IOT-C/common/basic/database"
	"github.com/Practical-Training-IOT/IOT-C/gateway/biz/handler/response"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime"
	"path/filepath"
	"time"
)

// MinIO 配置
const (
	minioEndpoint  = "117.27.231.112:9000"
	minioAccessKey = "VCXX3gIRycKf2Zc47Ppz"
	minioSecretKey = "XaS0WTo7kIVHFZdquqExddG75kJi6s316Ctxb1Wm"
	bucketName     = "dongbo" // 存储头像的 bucket 名称
)

func UploadAvatarHandler(ctx context.Context, c *app.RequestContext) {
	var id int
	if userIDInterface, exists := c.Get("userID"); exists {
		id = userIDInterface.(int)
	} else {
		responses := response.Response{
			Code: 400,
			Msg:  "未找到用户信息，请重新登录",
			Data: nil,
		}
		c.JSON(400, responses)
		return
	}
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(consts.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": "上传失败：" + err.Error(),
		})
		return
	}
	client, err := minio.New(minioEndpoint, &minio.Options{
		Creds: credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
	})
	if err != nil {
		c.JSON(consts.StatusOK, map[string]interface{}{
			"code":    -1,
			"message": "上传失败：" + err.Error(),
		})
		return
	}
	open, err := file.Open()
	defer open.Close()

	defLoc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(defLoc).Truncate(time.Second)
	timeUrl := now.Format("20060102150405")

	ext := filepath.Ext(file.Filename)
	contentType := mime.TypeByExtension(ext)
	objectName := fmt.Sprintf("%s/%s", timeUrl, file.Filename)
	client.PutObject(ctx, bucketName, objectName, open, file.Size, minio.PutObjectOptions{ContentType: mime.TypeByExtension(contentType)})

	sprintf := fmt.Sprintf("http://%s/%s/%s", minioEndpoint, bucketName, objectName)

	config.DB.Model(&database.UserC{}).Where("id=?", id).Update("profile_picture_url", sprintf)

	c.JSON(consts.StatusOK, map[string]interface{}{
		"code": 200,
		"data": map[string]string{
			"url": sprintf,
		},
		"message": "上传成功",
	})

}
