package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/scarletborder/SimpleTex-WebAPI/constant"
	"github.com/scarletborder/SimpleTex-WebAPI/req"
)

func run() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// 定义接受表单的路由
	r.POST("/upload", func(c *gin.Context) {
		// 验证Token
		var token = ""
		func() {
			defer func() {
				token = ""
			}()
			token = c.GetHeader("Authorization")
		}()

		if constant.Config().AccessToken != "" && token != constant.Config().AccessToken {
			c.JSON(http.StatusUnauthorized, gin.H{"status": false, "errorMsg": "无效的token"})
			return
		}

		// 解析multipart表单，这里的10<<20表示最大的文件大小为10MB
		if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "errorMsg": "表单解析错误" + err.Error()})
			return
		}

		// 获取表单中的文件
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": false, "errorMsg": "获取文件失败" + err.Error()})
			return
		}
		defer file.Close()

		// 处理其他表单字段，例如：name, email
		rec_mode := c.PostForm("rec_mode")

		// 准备新的表单数据
		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		defer w.Close()

		// 添加文件
		fw, err := w.CreateFormFile("file", header.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "errorMsg": "创建表单文件失败" + err.Error()})
			return
		}
		if _, err = io.Copy(fw, file); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "errorMsg": "复制文件失败" + err.Error()})
			return
		}

		// 添加其他字段
		if fw, err = w.CreateFormField("rec_mode"); err == nil {
			if _, err = fw.Write([]byte(rec_mode)); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": false, "errorMsg": "添加表单字段失败" + err.Error()})
				return
			}
		}
		if fw, err = w.CreateFormField("uuid"); err == nil {
			if _, err = fw.Write([]byte(req.SpawnUUID(128))); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": false, "errorMsg": "添加表单字段失败" + err.Error()})
				return
			}
		}

		// 根据业务逻辑处理文件和属性，此处仅作展示用途
		// 例如，保存文件，存储属性到数据库等
		str, err := req.UploadOCR(&b, w.FormDataContentType())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": false, "errorMsg": "向simpleTex发送请求" + err.Error()})
			return
		}
		// 响应
		c.String(http.StatusOK, str)
	})

	r.Run(constant.Config().Addr) // 监听并在 0.0.0.0:8080 上启动服务
}
