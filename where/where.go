package where

import (
	"github.com/metafates/mangal/constant"
	"github.com/metafates/mangal/filesystem"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const EnvConfigPath = "MANGAL_CONFIG_PATH"

// mkdir creates a directory and all parent directories if they don't exist
// will return the path of the directory
func mkdir(path string) string {
	lo.Must0(filesystem.Get().MkdirAll(path, os.ModePerm))
	return path
}

// Config path
func Config() string {
	var path string

	if customDir, present := os.LookupEnv(EnvConfigPath); present {
		path = customDir
	} else {
		path = filepath.Join(lo.Must(os.UserConfigDir()), constant.Mangal)
	}

	return mkdir(path)
}

// Sources path
func Sources() string {
	return mkdir(filepath.Join(Config(), "sources"))
}

// Logs path
func Logs() string {
	return mkdir(filepath.Join(Config(), "logs"))
}

// History path to the file
func History() string {
	cacheDir := filepath.Join(lo.Must(os.UserCacheDir()), constant.CachePrefix)
	lo.Must0(filesystem.Get().MkdirAll(filepath.Dir(cacheDir), os.ModePerm))

	path := filepath.Join(mkdir(cacheDir), "history.json")

	exists := lo.Must(filesystem.Get().Exists(path))
	if !exists {
		lo.Must0(filesystem.Get().WriteFile(path, []byte("{}"), os.ModePerm))
	}

	return path
}

// Downloads path
func Downloads() string {
	path, err := filepath.Abs(viper.GetString(constant.DownloaderPath))

	if err != nil {
		path, err = os.Getwd()
		if err != nil {
			path = "."
		}
	}

	return mkdir(path)
}
