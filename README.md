# zabton



## Description

CLI tool for managing Zabbix with text base configuration.

## Usage

Edit your zabbix environment and json file path like this sample.

```
vi /your/dir/.zaton
```

```
export ZABTON_ZABBIX_URL="http://your.zabbix.com/api_jsonrpc.php"
export ZABTON_ZABBIX_USER="you"
export ZABTON_ZABBIX_PASSWORD="your_zabbix_password"
export ZABTON_LOG_LEVEL="info"
export ZABTON_FILE_PATH="/your/zabbix/config/path.json"
```

Load this file.

```
source /your/dir/.zaton
```

Pull from Zabbix Server.

```
zabton pull
```

Edit zabbix configuration.

```
vi /your/zabbix/config/path.json
```

Push to Zabbix Server.

```
zabton push
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/hiracy/zabton
```

## Contribution

1. Fork ([https://github.com/hiracy/zabton/fork](https://github.com/hiracy/zabton/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create a new Pull Request

## Author

[hiracy](https://github.com/hiracy)
