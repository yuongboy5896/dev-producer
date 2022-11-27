package model

type DaoPage struct {
	Query    string `json:"query"`
	Pagenum  int    `json:"pagenum"`
	Pagesize int    `json:"pagesized"`
}
