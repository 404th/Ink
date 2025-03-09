package postService

import (
	"context"

	"github.com/404th/Ink/config"
	"github.com/404th/Ink/internal/storage"
	"github.com/404th/Ink/model"
	"go.uber.org/zap"
)

type postService struct {
	cfg   *config.Config
	sugar *zap.SugaredLogger
	strg  storage.PostPgI
}

func NewPostService(cfg *config.Config, sugar *zap.SugaredLogger, strg storage.PostPgI) *postService {
	return &postService{
		cfg:   cfg,
		sugar: sugar,
		strg:  strg,
	}
}

func (p *postService) CreatePost(ctx context.Context, req *model.CreatePostRequest) (resp *model.Post, err error) {
	resp = &model.Post{}

	p.sugar.Infoln("CreatePost ", config.InfoSplitter, "req", config.InfoSplitter, req)
	resp, err = p.strg.CreatePost(ctx, req)
	if err != nil {
		p.sugar.Infoln("GetUser ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return
	}

	return
}

func (p *postService) GetAllPosts(ctx context.Context, req *model.GetAllPostsRequest) (resp *model.GetAllPostsResponse, err error) {
	resp = &model.GetAllPostsResponse{}

	p.sugar.Infoln("GetAllPosts ", config.InfoSplitter, "req", config.InfoSplitter, req)
	resp, err = p.strg.GetAllPosts(ctx, req)
	if err != nil {
		p.sugar.Infoln("GetUser ", config.ErrorSplitter, "err", config.ErrorSplitter, err.Error())
		return
	}

	return
}
