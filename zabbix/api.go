package zabbix

import (
	"fmt"
	"time"
)

// API is the object of basic login information for zabbbix server.
type API struct {
	URL      string
	User     string
	Password string
	Auth     string
}

type rpcRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    string      `json:"auth,omitempty"`
	Id      int32       `json:"id"`
}

type RpcResponse struct {
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
	return &API{URL: url, User: user, Password: password}
}

// Login  actually login to zabbix server.
func (api *API) Login() (auth string, err error) {
	params := make(map[string]string)
	params["user"] = api.User
	params["passwlrd"] = api.Password

	res, err := api.request("user.login", params)
	if err != nil {
		return "", err
	}

	if res.Error.Code != 0 {
		return "", nil
	}

	auth = res.Result.(string)
	api.Auth = auth

	return "", nil
}

// request requests api to zabbix server.
func (api *API) request(method string, params interface{}) (RpcResponse, error) {
	res := new(RpcResponse)
	return *res, nil
}

// transactionId generates zabbix session id.
func transactionId() int64 {
	seed := time.Now().UnixNano() / 1000
	return seed - ((seed / 1000000000) * 1000000000)
}
