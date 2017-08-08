package zabbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// API is the object of basic login information for zabbbix server.
type API struct {
	URL      string
	User     string
	Password string
	Auth     string
	client   http.Client
}

type JsonRpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    string      `json:"auth,omitempty"`
	Id      int32       `json:"id"`
}

type JsonRpcResponse struct {
	Jsonrpc string       `json:"jsonrpc"`
	Error   *ZabbixError `json:"error"`
	Result  interface{}  `json:"result"`
	Id      int32        `json:"id"`
}

type ZabbixError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (e *ZabbixError) Error() string {
	return fmt.Sprintf(`code=%d message="%s" data="%s"`, e.Code, e.Message, e.Data)
}

// NewAPI creates new zabbix api object.
func NewAPI(url, user, password string) *API {
	if !strings.HasPrefix(url, "http://") &&
		!strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	return &API{url, user, password, "", http.Client{}}
}

// Login actually login to zabbix server.
func (api *API) Login() (auth string, err error) {
	params := make(map[string]string)
	params["user"] = api.User
	params["password"] = api.Password

	res, err := api.request("user.login", params)
	if err != nil {
		return "", err
	}

	if res.Error != nil && res.Error.Code != 0 {
		return "", res.Error
	}

	auth = res.Result.(string)
	api.Auth = auth

	return
}

// request requests api to zabbix server.
func (api *API) request(method string, params interface{}) (response JsonRpcResponse, err error) {
	var id int32
	id = 1
	req := JsonRpcRequest{"2.0", method, params, api.Auth, id}
	jsondata, err := json.Marshal(req)
	if err != nil {
		return JsonRpcResponse{}, err
	}

	request, err := http.NewRequest("POST", api.URL, bytes.NewReader(jsondata))
	if err != nil {
		return JsonRpcResponse{}, err
	}

	request.ContentLength = int64(len(jsondata))
	request.Header.Add("Content-Type", "application/json-rpc")
	request.Header.Add("User-Agent", "github.com/hiracy/zabton")

	res, err := api.client.Do(request)

	if err != nil {
		return JsonRpcResponse{}, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return JsonRpcResponse{}, err
	}

	if err := json.Unmarshal(b, &response); err != nil {
		return JsonRpcResponse{}, err
	}

	return
}
