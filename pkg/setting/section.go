package setting

import "time"

type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettings struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettings struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// ReadSection 读取配置段信息
func (s *Setting) ReadSection(k string, v interface{}) error {

	/*
		You also have the option of Unmarshalling all or a specific value to a struct, map, etc.
		There are two methods to do this:

		Unmarshal(rawVal interface{}) : error
		UnmarshalKey(key string, rawVal interface{}) : error
	*/

	//解构配置文件信息组
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
