package handler

import (
	"context"

	pb "captcha/proto"
	"github.com/mojocn/base64Captcha"
	log "go-micro.dev/v5/logger"
	"image/color"
)

type Captcha struct{}

// Return a new handler
func New() *Captcha {
	return &Captcha{}
}

var store = base64Captcha.DefaultMemStore

func (e *Captcha) MakeCaptcha(ctx context.Context, req *pb.MakeCaptchaReq, resp *pb.MakeCaptchaResp) error {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          int(req.Height),
		Width:           int(req.Width),
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          int(req.Length),
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3, G: 102, B: 214, A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driver = driverString.ConvertFonts()

	// 新变量使用 := ,赋值使用 =
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := c.Generate()
	if err == nil {
		log.Infof("makeCaptcha captchaId:%v,answer:%vgenerated captcha", id, answer)
	} else {
		log.Error("failed to generate captcha")
	}
	resp.Id = id
	resp.B64S = b64s
	resp.Answer = answer
	return nil

}

func (e *Captcha) VerifyCaptcha(ctx context.Context, req *pb.VerifyCaptchaReq, resp *pb.VerifyCaptchaResp) error {
	log.Infof("verifyCaptcha captchaId:%v,answer:%vgenerated captcha", req.Id, req.Answer)
	if store.Verify(req.Id, req.Answer, true) {
		resp.VerifyResult = true
	} else {
		resp.VerifyResult = false
	}
	return nil
}
