package config
type Config struct {
	HttpServer HttpServer `yaml:HttpServer`
	Database Database `yaml:"Database"`
}
type HttpServer struct {
	Port int `yaml:"Port"`
}
type Database struct {
	Main DbNode `yaml:"Main"`
}
type DbNode struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port'`
	User string `yaml:"User"`
	Password string `yaml:"Password"`
	Name string `yaml:"Name"`
}
