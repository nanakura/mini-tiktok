// Code generated by hertz generator.

package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "mini_tiktok/cmd/api/biz/model/api"
	"mini_tiktok/cmd/api/biz/rpc"
	userservice "mini_tiktok/kitex_gen/userservice"
	"mini_tiktok/kitex_gen/videoservice"
	utils2 "mini_tiktok/pkg/utils"
	"strconv"
	"time"
)

// Feed .
// @router /douyin/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FeedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var latestDate int64
	if req.LatestTime == nil {
		latestDate = time.Now().Unix() * 1000
	} else {
		latestDate, _ = strconv.ParseInt(*req.LatestTime, 10, 64)
	}
	feedResponse, err := rpc.VideoRpcClient.Feed(context.Background(),
		&videoservice.DouyinFeedRequest{
			LatestTime: latestDate,
		},
	)
	resp := new(api.FeedResp)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = fmt.Sprintf("获取视频失败：%v", err)
		c.JSON(consts.StatusOK, resp)
		return
	}

	resp.StatusCode = 0
	resp.NextTime = time.Now().Unix()
	if feedResponse == nil || feedResponse.VideoList == nil {
		resp.StatusMessage = fmt.Sprintf("获取视频失败, error：%v", err)
		resp.VideoList = []*api.Video{}
	} else {
		resp.VideoList = utils2.CastUserserviceVideoToApiVideo(feedResponse.VideoList)
	}
	hlog.Infof("feed: %+v", resp)
	c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router /douyin/user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	username := c.Query("username")
	password := c.Query("password")
	hlog.Info("start call login rpc api")
	hlog.Infof("name: %v, pass: %v", username, password)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	registerResponse, err := rpc.UserRpcClient.Register(context.Background(), &userservice.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"code": 0, "message": err.Error()})
		return
	}
	resp := &api.UserRegisterResp{
		StatusCode:    int64(registerResponse.StatusCode),
		StatusMessage: registerResponse.StatusMsg,
		UserID:        registerResponse.UserId,
		Token:         registerResponse.Token,
	}

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	username := c.Query("username")
	password := c.Query("password")
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	hlog.Info("start call login rpc api")
	hlog.Infof("name: %v, pass: %v", username, password)
	loginResponse, err := rpc.UserRpcClient.Login(context.Background(), &userservice.DouyinUserLoginRequest{
		Username: username,
		Password: password,
	})
	hlog.Info("call login rpc api end")
	if err != nil {
		hlog.Error("error occur", err)
		c.JSON(consts.StatusOK, utils.H{"code": 0, "message": err.Error()})
		return
	}
	if loginResponse == nil {
		c.JSON(consts.StatusOK, utils.H{
			"status": "nil",
		})
		return
	}
	resp := &api.UserLoginResp{
		StatusCode:    int64(loginResponse.StatusCode),
		StatusMessage: loginResponse.StatusMsg,
		UserID:        loginResponse.UserId,
		Token:         loginResponse.Token,
	}
	hlog.Infof("get resp: %+v", loginResponse)

	c.JSON(consts.StatusOK, resp)
}

// User .
// @router /douyin/user [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userId, _ := strconv.Atoi(req.UserID)
	info, err := rpc.UserRpcClient.Info(context.Background(), &userservice.DouyinUserRequest{UserId: int64(userId), Token: req.Token})
	if err != nil {
		hlog.Infof("获取用户信息时error occur: %v", err)
		c.JSON(consts.StatusOK, utils.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	resp := &api.UserResp{
		StatusCode:    int64(info.StatusCode),
		StatusMessage: info.StatusMsg,
		User: &api.User{
			ID:            info.User.Id,
			Name:          info.User.Name,
			FollowCount:   info.User.FollowCount,
			FollowerCount: info.User.FollowerCount,
			IsFollow:      info.User.IsFollow,
		},
	}

	c.JSON(consts.StatusOK, resp)
}

// PublishAction .
// @router /douyin/publish/action [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.PublishActionResp)

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.PublishListResp)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteAction .点赞接口
// @router /douyin/favorite/action [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	videoId, _ := strconv.Atoi(req.VideoID)
	actionType, _ := strconv.Atoi(req.ActionType)
	r, err := rpc.VideoRpcClient.FavoriteAction(context.Background(), &videoservice.DouyinFavoriteActionRequest{
		Token:      req.Token,
		VideoId:    int64(videoId),
		ActionType: int32(actionType),
	})

	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	resp := &api.FavoriteActionResp{
		StatusCode:    0,
		StatusMessage: r.StatusMsg,
	}

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList
// @router /douyin/favorite/list [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userId, _ := strconv.Atoi(req.UserID)
	r, err := rpc.VideoRpcClient.FavoriteList(context.Background(), &videoservice.DouyinFavoriteListRequest{
		UserId: int64(userId),
		Token:  req.Token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}

	resp := &videoservice.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		VideoList:  r.VideoList,
	}
	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CommentActionResp)

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.CommentListResp)

	c.JSON(consts.StatusOK, resp)
}

// RelationAction .
// @router /douyin/relation/action [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	// 关注操作
	var err error
	var req api.RelationActionReq
	var resp api.RelationActionResp
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusMessage = err.Error()
		resp.StatusCode = 1
		return
	}
	toUserId, toUserIdErr := strconv.ParseInt(req.ToUserID, 10, 64)
	actionType, actionTypeErr := strconv.ParseInt(req.ActionType, 10, 32)
	if toUserIdErr != nil || actionTypeErr != nil {
		resp.StatusCode = 0
		resp.StatusMessage = "传入的参数错误"
		c.JSON(consts.StatusOK, resp)
		return
	}
	// userService 完成 关注操作
	result, err := rpc.UserRpcClient.Action(ctx, &userservice.DouyinRelationActionRequest{
		Token:      req.Token,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
	})

	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.StatusCode = 0
	resp.StatusMessage = result.StatusMsg
	c.JSON(consts.StatusOK, resp)
}

// RelationFollowList .
// @router /douyin/relation/follow/list [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	// 关注列表
	var err error
	var req api.RelationFollowListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	hlog.Info("start call login rpc api")
	// 将 字符串 转为 int64
	userId, _ := strconv.ParseInt(req.UserID, 10, 64)
	resp, err := rpc.UserRpcClient.FollowList(ctx, &userservice.DouyinRelationFollowListRequest{
		UserId: userId,
		Token:  req.Token,
	})
	hlog.Info("call login rpc api end")
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status": "nil"})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowerList .
// @router /douyin/relation/follower/list [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFollowerListReq
	var resp api.RelationFollowerListResp
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	userId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		c.String(consts.StatusOK, "请求参数错误")
	}
	result, err := rpc.UserRpcClient.FollowerList(ctx, &userservice.DouyinRelationFollowerListRequest{
		UserId: userId,
		Token:  req.Token,
	})
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = "请求失败"
		c.JSON(consts.StatusOK, resp)
	}

	c.JSON(consts.StatusOK, result)
}

// RelationFriendList .
// @router /douyin/relation/friend/list [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFriendListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RelationFriendListResp)

	c.JSON(consts.StatusOK, resp)
}
