package model

type ClusterInfo struct {
	Clusterid      int64  `xorm:"pk autoincr" json:"Clusterid"`
	Clustername    string `xorm:"varchar(50)" json:"Clustername"`
	Clustervesrion string `xorm:"varchar(30)" json:"Clustervesrion"`
	Clustercode    string `xorm:"varchar(30)" json:"Clustercode"`
	Clusterdepson  string `xorm:"varchar(50)" json:"Clusterdepson"`
	Clusterconfig  string `xorm:"varchar(50)" json:"Clusterconfig"`
	Create_time    int64  `xorm:"bigint" json:"create_time"`
}












