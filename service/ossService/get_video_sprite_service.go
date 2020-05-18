package ossService

import (
	"fmt"
	"giligili/model"
	"giligili/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// UploadTokenService 获得上传oss token的服务
type GetVideoSpriteService struct {
}

//Create 创建雪碧图，返回生成连接
func (service *GetVideoSpriteService) CreateSpritePic(videoId string) serializer.Response {
	video, err := model.GetVideoById(videoId)
	if err != nil {
		return serializer.Response{
			Status: 40004,
			Msg:    "目标视频未找到",
			Error:  "video error",
		}
	}
	srcdir := "upload/videos"
	srcdirFilter := "upload/videos/"
	videoName := strings.TrimPrefix(video.URL, srcdirFilter)
	saveas := "sprite"
	w := "320"
	h := "180"
	offset := "1"
	interval := "1"
	sprite := "5*5"
	capture_num := "25"

	srcdirEncode := url.QueryEscape(srcdir)
	videoNameEncode := url.QueryEscape(videoName)

	urlstring := "https://1847252403518609.cn-beijing.fc.aliyuncs.com/2016-08-15/proxy/video-sprite/video-sprite-maker" +
		"/?srcdir=" + srcdirEncode + "&video=" + videoNameEncode + "&saveas=" + saveas + "&w=" + w + "&h=" + h + "&offset=" + offset +
		"&interval=" + interval + "&sprite=" + sprite + "&capture_num=" + capture_num

	//一个正确的URL栗子
	//urlstring := "https://1847252403518609.cn-beijing.fc.aliyuncs.com/2016-08-15/proxy/video-sprite/video-sprite-maker/?" +
	//	"srcdir=upload%2Fvideos&video=01%20Opening_1.mp4&saveas=sprite&w=320&h=180&offset=1&interval=1&sprite=5*4&capture_num=20"

	req, _ := http.NewRequest("GET", urlstring, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Sprite Pic:" + string(body))
	time.Sleep(time.Duration(2) * time.Second)

	//创建oss连接实例
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 查看资源
	signedGetURL, err := bucket.SignURL(string(body), oss.HTTPGet, 60)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: 0,
		Data:   signedGetURL,
		Msg:    "",
		Error:  "",
	}
}
