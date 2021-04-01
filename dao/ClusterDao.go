package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type ClusterDao struct {
	*tool.Orm
}

//新集群的数据库插入操作
func (cd *ClusterDao) InsertClusterInfo(clusterInfo *model.ClusterInfo) int64 {
	result, err := cd.InsertOne(clusterInfo)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
