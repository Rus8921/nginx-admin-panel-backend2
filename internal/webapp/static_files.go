package webapp

import (
	"io/fs"
	"net/http"
	"path"

	"github.com/gin-contrib/static"
	logger "github.com/sirupsen/logrus"
	"gitlab.pg.innopolis.university/antiddos/nginx-admin-panel-backend.git/web"
)

type ServerFileSystemType struct {
	http.FileSystem
}

func (f ServerFileSystemType) Exists(prefix string, _path string) bool {
	file, err := f.Open(path.Join(prefix, _path))
	if file != nil {
		defer func(file http.File) {
			err = file.Close()
			if err != nil {
				logger.Error("file not found", err)
			}
		}(file)
	}
	return err == nil
}

func MustFs(dir string) (serverFileSystem static.ServeFileSystem) {

	sub, err := fs.Sub(web.WebFS, path.Join("build", dir))

	if err != nil {
		logger.Error(err)
		return
	}

	serverFileSystem = ServerFileSystemType{
		http.FS(sub),
	}

	return
}
