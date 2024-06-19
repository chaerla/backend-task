package usecase

import (
	"backend-task/internal/client"
	"backend-task/pkg/logger"
	"fmt"
)

type UserService interface {
	GetManyAndPrint(page int)
}

type UserServiceImpl struct {
	dummyApi client.DummyApi
}

func NewUserService(dummyApi client.DummyApi) UserService {
	return &UserServiceImpl{
		dummyApi: dummyApi,
	}
}

func (s *UserServiceImpl) GetManyAndPrint(page int) {
	logger.Log.Debug(fmt.Sprintf("Printing users from page: %d", page))
	res, err := s.dummyApi.GetUsers(page)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	for _, user := range res.Data {
		logger.Log.Info(
			fmt.Sprintf(
				"User with ID %s:  %s. %s %s", user.ID, user.Title, user.FirstName,
				user.FirstName,
			),
		)
	}
	logger.Log.Debug(fmt.Sprintf("Finished printing users from page: %d", page))
}
