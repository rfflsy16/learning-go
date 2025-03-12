package config

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    ServerPort string
}

func LoadConfig() *Config {
    return &Config{
        DBHost:     "localhost",
        DBPort:     "3306", // Port default MariaDB
        DBUser:     "root",
        DBPassword: "mariadb",
        DBName:     "learning-go-DB",
        ServerPort: "8080",
    }
}
