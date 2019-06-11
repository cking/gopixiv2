package v2

import (
	"context"
	"fmt"
	"github.com/kanosaki/gopixiv2"
)

type API struct {
	Illust *Illust
	IllustBookmark *IllustBookmark
}

func NewAPI(client pixiv.Session) *API {
	return &API{
		Illust: NewIllust(client),
		IllustBookmark´: NewIllustBookmark(client),
	}
}

func (a *API) DoGet(ctx context.Context, path string, params map[string]string) (interface{}, error) {
	dispatchers := map[string]func() (interface{}, error){
	}
	if d, ok := dispatchers[path]; ok {
		return d()
	} else {
		return nil, fmt.Errorf("not found (at /v2): %s", path)
	}
}
