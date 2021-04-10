package web

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/fs"
	"net/http"
)

//getFileSystem returns either the compiled in fs or the debug fs
func getFileSystem() http.FileSystem {
	webDir := viper.GetString("web.directory")
	logrus.Printf("Directory: %s", webDir)

	if "" != webDir {
		logrus.Printf("serving static files from %s", webDir)
		logrus.Print("using live mode")
		return http.Dir(webDir)
	}

	logrus.Print("serving static files from binary")
	fsys, err := fs.Sub(embededFiles, "dist")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
