package api

import (
	"context"

	"github.com/digilolnet/client3xui"
)

type APIClient struct {
	Client *client3xui.Client
}

func NewAPIClient(baseURL, username, password string) *APIClient {
	c := client3xui.New(client3xui.Config{
		Url:      baseURL,
		Username: username,
		Password: password,
	})
	return &APIClient{Client: c}
}

// func (c *APIClient) SendRequest(endpoint string, method string, payload interface{}) ([]byte, error) {
// 	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
// 	data, err := json.Marshal(payload)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return body, nil
// }

// Получить статус сервера
func (c *APIClient) GetServerStatus() (*client3xui.ServerStatusResponse, error) {
	status, err := c.Client.ServerStatus(context.Background())
	if err != nil {
		return nil, err
	}
	return status, nil
}

// Перезапустить панель
func (c *APIClient) RestartPanel() (*client3xui.ApiResponse, error) {
	res, err := c.Client.RestartPanel(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Получить список пользователей
func (c *APIClient) GetOnlineClients() ([]string, error) {
	return c.Client.GetOnlineClients(context.Background())
}

func (c *APIClient) CreateUSer(user map[string]interface{}) (string, error) {
	// user := client3xui.User{
	// 	Username: name,
	// 	Traffic:  int64(limit),
	// }
	// return c.Client.CreateUser(user)
	return "", nil
}

// Удалить пользователя
func (c *APIClient) DeleteUser(username string) error {
	//return c.Client.DeleteUser(name)
	return nil
}

// Сгенерировать ссылку для подключения
func (c *APIClient) GenerateUserLink(username string) (string, error) {
	// link, err := c.Client.GenerateLink(name)
	// if err != nil {
	// 	return "", err
	// }
	// return link, nil
	return "", nil
}
