package dao

import (
	"dev-producer/model"
	"dev-producer/tool"
	"fmt"
)

type JenkinsDao struct {
	*tool.Orm
}

//实例化Dao对象
func NewJenkinsDao() *JenkinsDao {
	return &JenkinsDao{tool.DbEngine}
}

//jenkins的数据库插入操作
func (jd *JenkinsDao) InsertJenkinsJobs(jobs []model.JenkinsJob) int64 {
	result, err := jd.Insert(&jobs)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

//从数据库中查询所有jenkins job 列表
func (jd *JenkinsDao) QueryJenkinsJobs(daoPage *model.DaoPage) ([]model.JenkinsJob, error) {
	var jenkinsJobs []model.JenkinsJob
	if nil == daoPage {
		if err := jd.Engine.Desc("buildnum").Find(&jenkinsJobs); err != nil {
			return nil, err
		}
	} else {
		if err := jd.Engine.Desc("buildnum").Where("").Limit(daoPage.Pagenum, daoPage.Pagesize).Find(&jenkinsJobs); err != nil {
			return nil, err
		}
	}
	return jenkinsJobs, nil
}
