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

// JSONRpcRequest is the object of JSON RPC Request Object for Zabbix.
type JSONRpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    string      `json:"auth,omitempty"`
	ID      int32       `json:"id"`
}

// JSONRpcResponse is the object of JSON RPC Response Object for Zabbix.
type JSONRpcResponse struct {
	Jsonrpc string             `json:"jsonrpc"`
	Error   *ZabtonZabbixError `json:"error"`
	Result  interface{}        `json:"result"`
	ID      int32              `json:"id"`
}

// ZabtonZabbixError is the object of Zabbix general error response.
type ZabtonZabbixError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (e *ZabtonZabbixError) Error() string {
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

// Version return zabbix server api version.
func (api *API) Version() (version string, err error) {
	res, err := api.request("apiinfo.version", map[string]string{})
	if err != nil {
		return "", err
	}

	if res.Error != nil && res.Error.Code != 0 {
		return "", res.Error
	}

	var ok bool
	if version, ok = res.Result.(string); ok {
		return
	}

	return "", &ZabtonZabbixError{-1, "", "api response type error(string)"}
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

	var ok bool
	if auth, ok = res.Result.(string); ok {
		api.Auth = auth
		return
	}

	return "", &ZabtonZabbixError{-1, "", "api response type error(string)"}
}

// GetHost allows to retrieve hosts according to the given parameters.
func (api *API) GetHost(params interface{}) (hosts []interface{}, err error) {
	res, err := api.request("host.get", params)

	if err != nil {
		return nil, err
	}

	if res.Error != nil && res.Error.Code != 0 {
		return nil, res.Error
	}

	var ok bool
	if hosts, ok = res.Result.([]interface{}); ok {
		return
	}

	return nil, &ZabtonZabbixError{-1, "", "api response type error([]interface{})"}
}

// request requests api to zabbix server.
func (api *API) request(method string, params interface{}) (response JSONRpcResponse, err error) {
	var id int32
	id = 1
	req := JSONRpcRequest{"2.0", method, params, api.Auth, id}
	jsondata, err := json.Marshal(req)
	if err != nil {
		return JSONRpcResponse{}, err
	}

	request, err := http.NewRequest("POST", api.URL, bytes.NewReader(jsondata))
	if err != nil {
		return JSONRpcResponse{}, err
	}

	request.ContentLength = int64(len(jsondata))
	request.Header.Add("Content-Type", "application/json-rpc")
	request.Header.Add("User-Agent", "github.com/hiracy/zabton")

	res, err := api.client.Do(request)

	if err != nil {
		return JSONRpcResponse{}, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return JSONRpcResponse{}, err
	}

	if err := json.Unmarshal(b, &response); err != nil {
		return JSONRpcResponse{}, err
	}

	return
}
