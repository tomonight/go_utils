package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//MySQLClientForSource mysql数据库
const MySQLClientForSource = "mysql"

//ErrMySQLClientNotInit mysql客户端未初始化
var ErrMySQLClientNotInit = errors.New("mysql client not init")

//ErrMySQLClientNotConnected  mysql客户端未连接
var ErrMySQLClientNotConnected = errors.New("mysql client not connected")

//ErrMySQLUnSupportDriver 不支持的客户端驱动
var ErrMySQLUnSupportDriver = errors.New("unsupport driver name")

//ErrMySQLHandlerNil mysql处理器为空
var ErrMySQLHandlerNil = errors.New("mysql handler is nil")

var userRe = regexp.MustCompile(`User@Host: ([^\[]+|\[[^[]+\]).*?@ (\S*) \[(.*)\]`)

//MySQLClient 定义mysql客户端
type MySQLClient struct {
	username  string
	password  string
	host      string
	port      uint
	connected bool
	init      bool
	lastErr   string
	dbHandler *sql.DB
	driver    string
	clientID  string
	timeout   int //10000ms
}

//NewMySQLClient 新建MySQL客户端
func NewMySQLClient(host, username, password string, port uint, timeout ...int) *MySQLClient {
	client := new(MySQLClient)
	client.username = username
	client.password = password
	client.host = host
	client.port = port
	client.connected = false
	client.init = true
	if len(timeout) > 0 {
		client.timeout = timeout[0] //ms 30s
	}
	client.driver = MySQLClientForSource
	return client
}

//DoExecuteSQL 执行SQL语句
func (client *MySQLClient) DoExecuteSQL(sqls []string) (err error) {
	if !client.connected || client.dbHandler == nil {
		return ErrMySQLClientNotConnected
	}
	for _, sql := range sqls {
		_, err = client.dbHandler.Exec(sql)
		if err != nil {
			return
		}
	}
	return
}

//Invalid  判断参数是否合法
func (client *MySQLClient) Invalid() bool {
	if client.username == "" || client.password == "" ||
		client.host == "" || client.port == 0 {
		return true
	}
	return false
}

//SetClientID 设置客户端id
func (client *MySQLClient) SetClientID(clientID string) {
	client.clientID = clientID
}

//SetTimeOut setting timeout
func (client *MySQLClient) SetTimeOut(timeout int) {
	client.timeout = timeout
}

//GetClientID  获取客户端id
func (client *MySQLClient) GetClientID() string {
	return client.clientID
}

func (client *MySQLClient) connectMySQL() error {
	link := fmt.Sprintf("%s:%s@tcp(%s:%d)/?timeout=%dms", client.username, client.password, client.host, client.port, client.timeout)
	//beego.Debug("link string ", link)
	db, err := sql.Open("mysql", link)
	if err != nil {
		return err
	}
	client.connected = true
	client.dbHandler = db
	return nil
}

func (client *MySQLClient) userDatabase(dbName string) error {
	if dbName == "" {
		dbName = "mysql"
	}
	_, err := client.dbHandler.Exec("USE " + dbName)
	return err
}

//Connect 连接mysql
func (client *MySQLClient) Connect() error {
	if client.init && !client.connected {
		return client.connectMySQL()
	}
	return ErrMySQLClientNotConnected
}

//GetSQLDB 获取mysql原始连接信息
func (client *MySQLClient) GetSQLDB() (*sql.DB, error) {
	if client.connected {
		if client.dbHandler != nil {
			return client.dbHandler, nil
		}
		return nil, ErrMySQLClientNotConnected

	}
	return nil, ErrMySQLClientNotInit
}

//Ping 判断连接是否可用
func (client *MySQLClient) Ping() error {
	if !client.connected {
		return ErrMySQLClientNotConnected
	}
	return client.dbHandler.Ping()
}

//Close 关闭连接
func (client *MySQLClient) Close() error {
	if client.connected && client.dbHandler != nil {
		return client.dbHandler.Close()
	}
	return nil
}

//GetMySQLInfoBySQLToString 获取数据库的信息ToString
func (client *MySQLClient) GetMySQLInfoBySQLToString(sql string) (info string, err error) {
	if !client.connected || client.dbHandler == nil {
		err = ErrMySQLClientNotConnected
		return
	}
	err = client.dbHandler.QueryRow(sql).Scan(&info)
	return
}

//GetMySQLInfoBySQLToInt 获取数据库的信息ToString
func (client *MySQLClient) GetMySQLInfoBySQLToInt(sql string) (info int, err error) {
	if !client.connected || client.dbHandler == nil {
		err = ErrMySQLClientNotConnected
		return
	}
	err = client.dbHandler.QueryRow(sql).Scan(&info)
	return
}

//NewSQLClient NewSQLClient
func (client *MySQLClient) NewSQLClient(queryStr, MySQLClientForSource string) (stmt *sql.Stmt, err error) {
	if client.connected && client.dbHandler != nil {
		switch client.driver {
		case MySQLClientForSource:
			stmt, err = client.dbHandler.Prepare(queryStr)
			return
		}
	}
	err = ErrMySQLClientNotConnected
	return
}

//QueryGlobalVariableMap db query
//show global variable
//like  show slave status
func (client *MySQLClient) QueryGlobalVariableMap(sql string) (map[string]string, error) {
	rows, err := client.dbHandler.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	status, err := ScanMap(rows)

	if err != nil {
		return nil, err
	}

	nullStringMap := make(map[string]string)

	for k, v := range status {
		value := ""
		if v.Valid {
			value = v.String
		}
		nullStringMap[strings.TrimPrefix(k, "@@")] = value
	}
	return nullStringMap, nil
}

//QueryGlobalVariableList QueryGlobalVariableList
func (client *MySQLClient) QueryGlobalVariableList(sql string) ([]map[string]string, error) {
	rows, err := client.dbHandler.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	statusList, err := ScanList(rows)

	if err != nil {
		return nil, err
	}
	list := []map[string]string{}
	for _, status := range statusList {
		nullStringMap := make(map[string]string)
		for k, v := range status {
			value := ""
			if v.Valid {
				value = v.String
			}
			nullStringMap[strings.ToLower(strings.TrimPrefix(k, "@@"))] = value
		}

		list = append(list, nullStringMap)
	}
	return list, nil
}

//QueryGlobalVariableMap db query
//show global variable into map
//like  show slave status
func (client *MySQLClient) QueryGlobalVariable(sql string) (string, error) {
	rows, err := client.dbHandler.Query(sql)
	var info string
	if err != nil {
		return "", err
	}
	for rows.Next() {
		err = rows.Scan(&info)
		if err != nil {
			return "", err
		}
	}
	return info, nil
}

//ScanMap 获取sql 查询数据
func ScanMap(rows *sql.Rows) (map[string]sql.NullString, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		err = rows.Err()
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	values := make([]interface{}, len(columns))

	for index := range values {
		values[index] = new(sql.NullString)
	}

	err = rows.Scan(values...)

	if err != nil {
		return nil, err
	}

	result := make(map[string]sql.NullString)

	for index, columnName := range columns {
		result[columnName] = *values[index].(*sql.NullString)
	}

	return result, nil
}

//ScanList 获取sql 查询数据
func ScanList(rows *sql.Rows) ([]map[string]sql.NullString, error) {
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	list := []map[string]sql.NullString{}
	for rows.Next() {
		values := make([]interface{}, len(columns))

		for index := range values {
			values[index] = new(sql.NullString)
		}

		err = rows.Scan(values...)

		if err != nil {
			return nil, err
		}

		result := make(map[string]sql.NullString)

		for index, columnName := range columns {
			result[columnName] = *values[index].(*sql.NullString)
		}

		list = append(list, result)
	}

	return list, nil
}

//ShowGlobalInfo ShowGlobalInfo
func (client *MySQLClient) ShowGlobalInfo(sql string) (info []map[string]string, err error) {
	stmt, err := client.NewSQLClient(sql, MySQLClientForSource)
	if err != nil {
		return info, err
	}
	if stmt == nil {
		return info, errors.New("connect is valid")
	}
	defer stmt.Close()
	info = []map[string]string{}
	rows, err := stmt.Query()
	if err != nil {
		return info, err
	}
	for rows.Next() {
		data := make(map[string]string)
		var key, value string
		err = rows.Scan(&key, &value)
		if err != nil {
			return info, err
		}
		data[key] = value
		info = append(info, data)
	}
	return info, nil

}
