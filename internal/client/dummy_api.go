package client

import (
	"backend-task/bootstrap/config"
	"backend-task/internal/client/dto"
	"backend-task/pkg/http"
	"encoding/json"
	"fmt"
)

type DummyApi interface {
	GetUsers(page int) (*dto.DummyAPIUsersResponse, error)
	GetPosts(page int) (*dto.DummyAPIPostsResponse, error)
}

type DummyApiImpl struct {
	baseUrl string
	appID   string
}

// NewDummyApi Dummy Api Factory
func NewDummyApi(config *config.Config) DummyApi {
	return &DummyApiImpl{
		baseUrl: config.DummyApiUrl,
		appID:   config.DummyApiAppID,
	}
}

func (d *DummyApiImpl) GetUsers(page int) (*dto.DummyAPIUsersResponse, error) {
	resp, err := http.FetchExternalAPI(
		http.FetchExternalAPIParam{
			URL:     d.getUserEndpoint(),
			Headers: d.getDefaultHeaders(),
			Params: map[string]interface{}{
				"page":  page,
				"limit": 5,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	var user dto.DummyAPIUsersResponse
	err = json.Unmarshal(resp, &user)
	return &user, err
}

func (d *DummyApiImpl) GetPosts(page int) (*dto.DummyAPIPostsResponse, error) {
	resp, err := http.FetchExternalAPI(
		http.FetchExternalAPIParam{
			URL:     d.getPostEndpoint(),
			Headers: d.getDefaultHeaders(),
			Params: map[string]interface{}{
				"page": page,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	var posts dto.DummyAPIPostsResponse
	err = json.Unmarshal(resp, &posts)
	return &posts, err
}

func (d *DummyApiImpl) getDefaultHeaders() map[string]string {
	return map[string]string{
		"app-id": d.appID,
	}
}

func (d *DummyApiImpl) getUserEndpoint() string {
	return fmt.Sprintf("%s/user", d.baseUrl)
}

func (d *DummyApiImpl) getPostEndpoint() string {
	return fmt.Sprintf("%s/post", d.baseUrl)
}
