package env

import "os"

func DBUrl() string {
	return os.Getenv("DBUrl")
}

func Env() string {
	return os.Getenv("Env")
}

func AppName() string {
	return os.Getenv("AppName")
}

func CorsOrigin() string {
	return os.Getenv("CorsOrigin")
}

func AppToken() string {
	return os.Getenv("AppToken")
}

func InternalToolToken() string {
	return os.Getenv("InternalToolToken")
}

func AppPort() string {
	return os.Getenv("AppPort")
}
