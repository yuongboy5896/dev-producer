package service

import (
	"dev-producer/dao"
	"dev-producer/model"
	"dev-producer/tool"
)

type ClusterInfoService struct {
}

func (cs *ClusterInfoService) Import(clusterInfo *model.ClusterInfo) *model.ClusterInfo {

	cd := dao.ClusterDao{tool.DbEngine}

	cd.InsertClusterInfo(clusterInfo)

	return clusterInfo
}
