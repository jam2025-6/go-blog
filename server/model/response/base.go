package response

type Captcha struct {
	CaptchaID string `json:"captcha_id"`
	PicPath   string `json:"pic_path"`
}

type SendEmailVerificationCode struct {
	Email     string `json:"email" binding:"required,email"`
	Captcha   string `json:"captcha" binding:"required,len=6"`
	CaptchaID string `json:"captcha_id" binding:"required"`
}
