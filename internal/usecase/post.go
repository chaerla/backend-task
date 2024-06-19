package usecase

import (
	"backend-task/internal/client"
	"backend-task/pkg/logger"
	"fmt"
	"strings"
	"time"
)

type PostService interface {
	GetManyAndPrint(page int)
}

type PostServiceImpl struct {
	dummyApi client.DummyApi
}

func NewPostService(dummyApi client.DummyApi) PostService {
	return &PostServiceImpl{
		dummyApi: dummyApi,
	}
}

func (s *PostServiceImpl) GetManyAndPrint(page int) {
	logger.Log.Debug(fmt.Sprintf("Printing posts from page: %d", page))
	res, err := s.dummyApi.GetPosts(page)
	if err != nil {
		logger.Log.Error(err.Error())
		return
	}
	for _, post := range res.Data {
		logger.Log.Info(
			fmt.Sprintf(
				`
Posted by %s %s: 

%s

Likes: %d | Tags: %s
Date posted: %s
`,
				post.Owner.FirstName, post.Owner.LastName, post.Text,
				post.Likes, strings.Join(post.Tags, ", "),
				post.PublishDate.Format(time.RFC3339),
			),
		)
	}
	logger.Log.Debug(fmt.Sprintf("Finished printing posts from page: %d", page))
}
