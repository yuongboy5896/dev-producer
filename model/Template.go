package model

// 模版管理
type TemplateInfo struct {
	Id             int64  `xorm:"pk autoincr" json:"Id"`
	TemplateName   string `xorm:"varchar(50)" json:"TemplateName"` //模版名称
	TemplateCode   string `xorm:"varchar(50)" json:"TemplateCode"` //模版编码
	TemplateType   string `xorm:"Text" json:"TemplateType"`        //模版类型
	TemplateText   string `xorm:"Text" json:"TemplateText"`        //模版yaml
	TemplateJekins string `xorm:"Text" json:"TemplateJekins"`      //模版jenkins模版
}
