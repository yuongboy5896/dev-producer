package model

// 模版管理
type Template struct {
	Id             int64  `xorm:"pk autoincr" json:"id"`
	TemplateName   string `xorm:"varchar(50)" json:"TemplateName"`
	TemplateCode   string `xorm:"varchar(50)" json:"TemplateCode"`
	TemplateText   string `xorm:"Text" json:"TemplateText"`
	TemplateJekins string `xorm:"Text" json:"TemplateJekins"`
}
