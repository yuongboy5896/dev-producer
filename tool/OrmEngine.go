package tool

import (
	"dev-producer/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine(database.Driver, conn)

	if err != nil {
		return nil, err
	}
	engine.SetMapper(core.SameMapper{})
	engine.ShowSQL(database.ShowSql)
	err = engine.Sync2(new(model.SmsCode),
		new(model.ClusterInfo), new(model.Member),
		new(model.ProjectInfo), new(model.VirtualMachine),
		new(model.ModuleInfo), new(model.PipeLine),
		new(model.PipeLineHistory), new(model.PipeLineSimple),
		new(model.DeployEnv), new(model.DeployEnv),
		new(model.VcenterVm), new(model.ModuleForImageUrl),
		new(model.IpAlive))

	if err != nil {
		return nil, err
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
