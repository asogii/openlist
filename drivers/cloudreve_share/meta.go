package cloudreve_share

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	// Usually one of two
	driver.RootPath
	// define other
	Address  string `json:"address" required:"true"`
	Key      string `json:"key" required:"true"`
	CustomUA string `json:"custom_ua"`
	Cookie   string `json:"cookie"`
}

var config = driver.Config{
	Name:              "Cloudreve Share",
	DefaultRoot:       "/",
	CheckStatus:       false,
	Alert:             "",
	NoOverwriteUpload: true,
	NoUpload:          true,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &CloudreveShare{}
	})
}
