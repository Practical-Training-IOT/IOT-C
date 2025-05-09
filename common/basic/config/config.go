package config

type Configs struct {
	Nacos
}

type Nacos struct {
	NameSpace string
	User      string
	Password  string
	Host      string
	Port      uint64
	DataId    string
	Group     string
}

type AlarmNacosStruct struct {
	PostgreSQL struct {
		User     string `json:"User"`
		Password string `json:"Password"`
		Host     string `json:"Host"`
		Port     int    `json:"Port"`
		Datebase string `json:"Datebase"`
	} `json:"PostgreSQL"`
}
