package user

import (
	"go-projects-server/internal/app/user/dao"
	"go-projects-server/internal/app/user/service"

	"github.com/google/wire"
)

var Set = wire.NewSet(dao.UserDaoSet, service.UserServiceSet)
