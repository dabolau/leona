package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 获取环境变量的本地数据库连接字符串
// mongodb://username:password@localhost:27017
func GetEnvLocalhostMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return fmt.Sprintf("mongodb://%s:%s@%s", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_ADDRESS"))
}

// 获取环境变量的容器数据库连接字符串
// mongodb://username:password@mongo:27017
func GetEnvDockerMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return fmt.Sprintf("mongodb://%s:%s@%s", os.Getenv("DOCKER_MONGO_USERNAME"), os.Getenv("DOCKER_MONGO_PASSWORD"), os.Getenv("DOCKER_MONGO_ADDRESS"))
}

// 获取环境变量的数据库连接字符串
func GetEnvMongoURI() string {
	// 容器中存在的文件
	const DOCKERENV_FILE string = "/.dockerenv"
	// 判断文件是否存在，存在表示在容器中，不存在表示在本地
	if IsExists(DOCKERENV_FILE) {
		return GetEnvDockerMongoURI()
	}
	return GetEnvLocalhostMongoURI()
}
