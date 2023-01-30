// Code generated by hertz generator.

package api

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/nanakura/go-ramda"
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
		c.JSON(consts.StatusBadRequest, api.Response{StatusCode: 1, StatusMsg: err.Error()})
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
		c.JSON(consts.StatusOK, api.Response{StatusCode: 1, StatusMsg: err.Error()})
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
	resp := &api.UserRegisterResp{}
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	registerResponse, err := rpc.UserRpcClient.Register(context.Background(), &userservice.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	resp = &api.UserRegisterResp{
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
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
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
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	if loginResponse == nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
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
	resp := &api.UserResp{}
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}

	userId, _ := strconv.Atoi(req.UserID)
	info, err := rpc.UserRpcClient.Info(context.Background(), &userservice.DouyinUserRequest{UserId: int64(userId), Token: req.Token})
	if err != nil {
		hlog.Infof("获取用户信息时error occur: %v", err)
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	resp = &api.UserResp{
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
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	resp := new(api.PublishActionResp)
	actionResponse, err := rpc.VideoRpcClient.PublishAction(context.Background(), &videoservice.DouyinPublishActionRequest{
		Token: req.Token,
		Data: ramda.Map(func(in int8) byte {
			return byte(in)
		})(req.Data),
		Title: req.Title,
	})
	if err != nil || actionResponse == nil {
		resp.StatusCode = 1
		resp.StatusMessage = fmt.Sprintf("上传失败: %v", err)
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.StatusCode = int64(actionResponse.StatusCode)
	resp.StatusMessage = actionResponse.StatusMsg
	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}

	resp := new(api.PublishListResp)
	userId, err := strconv.Atoi(req.UserID)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	publishListResponse, err := rpc.VideoRpcClient.PublishList(context.Background(), &videoservice.DouyinPublishListRequest{
		Token:  req.Token,
		UserId: int64(userId),
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	if resp.VideoList == nil {
		resp.VideoList = []*api.Video{}
	}
	resp.VideoList = utils2.CastUserserviceVideoToApiVideo(publishListResponse.VideoList)
	c.JSON(consts.StatusOK, resp)
}

// FavoriteAction .点赞接口
// @router /douyin/favorite/action [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
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
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
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
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	userId, _ := strconv.Atoi(req.UserID)
	r, err := rpc.VideoRpcClient.FavoriteList(context.Background(), &videoservice.DouyinFavoriteListRequest{
		UserId: int64(userId),
		Token:  req.Token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
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
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	videoId, _ := strconv.Atoi(req.VideoID)
	act, _ := strconv.Atoi(req.ActionType)
	// 删除操作
	if act == 2 {
		CommentId, _ := strconv.Atoi(*req.CommentID)
		info, err := rpc.VideoRpcClient.CommentAction(context.Background(), &videoservice.DouyinCommentActionRequest{
			Token:      req.Token,
			VideoId:    int64(videoId),
			ActionType: int32(act),
			CommentId:  int64(CommentId),
		})
		if err != nil {
			c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
			return
		}
		resp := &api.CommentActionResp{
			StatusCode:    int64(info.StatusCode),
			StatusMessage: info.StatusMsg,
		}
		c.JSON(consts.StatusOK, resp)
	} else {
		// 评论操作
		info, err := rpc.VideoRpcClient.CommentAction(context.Background(), &videoservice.DouyinCommentActionRequest{
			Token:       req.Token,
			VideoId:     int64(videoId),
			CommentText: *req.CommentText,
			ActionType:  int32(act),
		})
		if err != nil {
			c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
			return
		}

		user := &api.User{
			ID:            info.Comment.User.Id,
			Name:          info.Comment.User.Name,
			FollowCount:   info.Comment.User.FollowCount,
			FollowerCount: info.Comment.User.FollowerCount,
		}
		resp := &api.CommentActionResp{
			StatusCode:    int64(info.StatusCode),
			StatusMessage: info.StatusMsg,
			Comment: &api.Comment{
				ID:         info.Comment.Id,
				User:       user,
				Content:    info.Comment.Content,
				CreateDate: info.Comment.CreateDate,
			},
		}

		c.JSON(consts.StatusOK, resp)
	}
}

// CommentList .
// @router /douyin/comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	videoId, _ := strconv.Atoi(req.VideoID)
	resp, err := rpc.VideoRpcClient.CommentList(context.Background(), &videoservice.DouyinCommentListRequest{
		Token:   req.Token,
		VideoId: int64(videoId),
	})

	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
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
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	// 将 字符串 转为 int64
	userId, _ := strconv.ParseInt(req.UserID, 10, 64)
	resp, err := rpc.UserRpcClient.FollowList(ctx, &userservice.DouyinRelationFollowListRequest{
		UserId: userId,
		Token:  req.Token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
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
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	userId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
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
	var resp api.RelationFriendListResp
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	userId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	result, err := rpc.UserRpcClient.FriendList(ctx, &userservice.DouyinRelationFriendListRequest{
		UserId: userId,
		Token:  req.Token,
	})
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
		resp.UserList = nil
		c.JSON(consts.StatusOK, resp)
	}
	c.JSON(consts.StatusOK, result)
}
