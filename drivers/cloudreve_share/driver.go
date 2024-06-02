package cloudreve_share

import (
	"context"
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/errs"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/pkg/utils"
	"net/http"
	"net/url"
	"strings"
)

type CloudreveShare struct {
	model.Storage
	Addition
}

func (d *CloudreveShare) Config() driver.Config {
	return config
}

func (d *CloudreveShare) GetAddition() driver.Additional {
	return &d.Addition
}

func (d *CloudreveShare) Init(ctx context.Context) error {
	if d.Cookie != "" {
		return nil
	}
	// removing trailing slash
	d.Address = strings.TrimSuffix(d.Address, "/")
	return d.request(http.MethodGet, "/share/info/"+d.Key, nil, nil)
	//if err != nil {
	//	return err
	//}
	//err = d.request(http.MethodGet, "/share/info/"+d.Key, nil, nil)
	//if err != nil {
	//	return err
	//}
	//return d.request(http.MethodGet, "/site/config", nil, nil)
}

func (d *CloudreveShare) Drop(ctx context.Context) error {
	return nil
}

func (d *CloudreveShare) List(ctx context.Context, dir model.Obj, args model.ListArgs) ([]model.Obj, error) {
	var r DirectoryResp
	err := d.request(http.MethodGet, "/share/list/"+d.Key+url.QueryEscape(dir.GetPath()), nil, &r)
	if err != nil {
		return nil, err
	}

	return utils.SliceConvert(r.Objects, func(src Object) (model.Obj, error) {
		return objectToObj(src, model.Thumbnail{}), nil
	})
}

func (d *CloudreveShare) Link(ctx context.Context, file model.Obj, args model.LinkArgs) (*model.Link, error) {
	var dUrl string
	err := d.request(http.MethodPut, "/share/download/"+d.Key+"?path="+url.QueryEscape(file.GetPath()), nil, &dUrl)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(dUrl, "/api") {
		dUrl = d.Address + dUrl
	}
	return &model.Link{
		URL: dUrl,
	}, nil
}

func (d *CloudreveShare) MakeDir(ctx context.Context, parentDir model.Obj, dirName string) error {
	return errs.NotSupport
}

func (d *CloudreveShare) Move(ctx context.Context, srcObj, dstDir model.Obj) error {
	return errs.NotSupport
}

func (d *CloudreveShare) Rename(ctx context.Context, srcObj model.Obj, newName string) error {
	return errs.NotSupport
}

func (d *CloudreveShare) Copy(ctx context.Context, srcObj, dstDir model.Obj) error {
	return errs.NotSupport
}

func (d *CloudreveShare) Remove(ctx context.Context, obj model.Obj) error {
	return errs.NotSupport
}

func (d *CloudreveShare) Put(ctx context.Context, dstDir model.Obj, stream model.FileStreamer, up driver.UpdateProgress) error {
	return errs.NotSupport
}

var _ driver.Driver = (*CloudreveShare)(nil)
