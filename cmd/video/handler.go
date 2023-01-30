package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/nanakura/go-ramda"
	"mini_tiktok/cmd/video/ftpUtil"
	videoservice "mini_tiktok/kitex_gen/videoservice"
	"mini_tiktok/pkg/cache"
	"mini_tiktok/pkg/configs/config"
	"mini_tiktok/pkg/consts"
	"mini_tiktok/pkg/dal/model"
	"mini_tiktok/pkg/dal/query"
	"mini_tiktok/pkg/utils"
	jwtutil "mini_tiktok/pkg/utils"
	"strconv"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videoservice.DouyinPublishActionRequest) (resp *videoservice.DouyinPublishActionResponse, err error) {
	data := bytes.NewBufferString(string(req.Data))
	uuidv4, _ := uuid.NewUUID()
	path := fmt.Sprintf("%s.mp4", uuidv4.String())
	tv := query.Q.TVideo
	cliams, _ := utils.CheckToken(req.Token)
	userId := cliams.UserId
	playUrl := fmt.Sprintf("%s/%s", config.GlobalConfigs.StaticConfig.Url, path)
	err = tv.WithContext(context.Background()).
		Create(&model.TVideo{
			AuthorID:      userId,
			PlayURL:       playUrl,
			CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         req.Title,
			CreateDate:    time.Now(),
		})
	if err != nil {
		klog.Error("Error uploading file:", err)
		err = fmt.Errorf("视频保存失败：%w", err)
		return
	}
	if err = ftpUtil.FtpClient.Stor(path, data); err != nil {
		klog.Error("Error uploading file:", err)
		err = fmt.Errorf("视频保存失败：%w", err)
		return
	}

	return
}

type queryVideoListRes struct {
	ID            int64 // 视频id
	AuthorID      int64 `sql:"author_id"`
	Name          string
	FollowCount   int64
	FollowerCount int64
	Password      string
	PlayURL       string    // 视频链接
	CoverURL      string    // 视频封面链接
	FavoriteCount int64     // 点赞数
	CommentCount  int64     // 评论数
	IsFavorite    bool      // 是否已点赞(0为未点赞, 1为已点赞)
	Title         string    // 视频标题
	CreateDate    time.Time // 视频上传时间
}

func CastQueryVideoListtoVideoServiceVideo(from []queryVideoListRes) []*videoservice.Video {
	return ramda.Map(func(model queryVideoListRes) *videoservice.Video {
		return &videoservice.Video{
			Id: model.ID,
			Author: &videoservice.User{
				Id:            model.AuthorID,
				Name:          model.Name,
				FollowCount:   model.FollowCount,
				FollowerCount: model.FollowerCount,
			},
			PlayUrl:       model.PlayURL,
			CoverUrl:      model.CoverURL,
			FavoriteCount: model.FavoriteCount,
			CommentCount:  model.CommentCount,
			IsFavorite:    model.IsFavorite,
			Title:         model.Title,
		}
	})(from)
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videoservice.DouyinFeedRequest) (resp *videoservice.DouyinFeedResponse, err error) {
	latestTime := req.LatestTime
	// 值为0（默认值）则说明不限制最新时间
	tv := query.Q.TVideo.As("v")
	tu := query.Q.TUser.As("u")
	var resList []queryVideoListRes
	if latestTime == 0 {
		err = tv.WithContext(context.Background()).
			Select(
				tv.ID,
				tv.AuthorID,
				tu.ALL,
				tv.PlayURL, tv.CoverURL, tv.FavoriteCount,
				tv.CommentCount, tv.IsFavorite, tv.Title, tv.CreateDate,
			).
			LeftJoin(tu, tu.ID.EqCol(tv.AuthorID)).
			Order(tv.CreateDate.Desc()).
			Limit(10).Scan(&resList)

		if err != nil {
			return
		}
	} else {
		t := time.Unix(latestTime/1000, 0)
		err = tv.WithContext(context.Background()).
			Select(
				tv.ID,
				tv.AuthorID,
				tu.Name,
				tu.Password,
				tu.FollowCount,
				tu.FollowerCount,
				tv.PlayURL, tv.CoverURL, tv.FavoriteCount,
				tv.CommentCount, tv.IsFavorite, tv.Title, tv.CreateDate,
			).
			LeftJoin(tu, tu.ID.EqCol(tv.AuthorID)).
			Where(tv.CreateDate.Lt(t)).
			Order(tv.CreateDate.Desc()).
			Limit(10).
			Scan(&resList)
		if err != nil {
			return
		}
	}
	if resList == nil {
		resList = []queryVideoListRes{}
	}
	resp = &videoservice.DouyinFeedResponse{
		StatusCode: 0,
		VideoList:  CastQueryVideoListtoVideoServiceVideo(resList),
	}
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *videoservice.DouyinPublishListRequest) (resp *videoservice.DouyinPublishListResponse, err error) {
	userId := req.UserId
	tv := query.Q.TVideo.As("tv")
	tu := query.Q.TUser
	tu2 := query.Q.TUser.As("tu2")
	var resList []queryVideoListRes
	qCtx := context.Background()
	err = tv.WithContext(qCtx).
		Select(
			tv.ID,
			tv.AuthorID,
			tu2.Name,
			tu2.Password,
			tu2.FollowCount,
			tu2.FollowerCount,
			tv.PlayURL, tv.CoverURL, tv.FavoriteCount,
			tv.CommentCount, tv.IsFavorite, tv.Title, tv.CreateDate,
		).
		LeftJoin(tu.WithContext(qCtx).Select(tu.ALL).Where(tu.ID.Eq(userId)).As("tu2"), tu2.ID.EqCol(tv.AuthorID)).
		Order(tv.CreateDate.Desc()).
		Limit(10).
		Scan(&resList)

	if resList == nil {
		resList = []queryVideoListRes{}
	}
	resp = &videoservice.DouyinPublishListResponse{
		StatusCode: 0,
		VideoList:  CastQueryVideoListtoVideoServiceVideo(resList),
	}
	return
}

// FavoriteAction 2023-1-27 @Auth by 李卓轩 version 1.0
// 赞操作
// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *videoservice.DouyinFavoriteActionRequest) (resp *videoservice.DouyinFavoriteActionResponse, err error) {
	fmt.Println("in 点赞" + strconv.FormatInt(req.VideoId, 10))
	// 通过 token 解析出当前用户
	claims, flag := jwtutil.CheckToken(req.Token)
	// 说明 token 已经过期
	if !flag {
		return nil, errors.New("token is expired")
	}

	//判断当前用户是否点赞
	result, err := cache.RedisCache.RedisClient.SIsMember(context.Background(), consts.FavoriteActionPrefix+strconv.FormatInt(req.VideoId, 10), strconv.FormatInt(claims.UserId, 10)).Result()
	if err != nil {
		err = fmt.Errorf("redis访问失败")
		return
	}
	//已点过赞，取消点赞
	if result {
		// redis数据库中删除关联
		_, err1 := cache.RedisCache.RedisClient.SRem(context.Background(), consts.FavoriteActionPrefix+strconv.FormatInt(req.VideoId, 10), strconv.FormatInt(claims.UserId, 10)).Result()
		if err1 != nil {
			err1 = fmt.Errorf("redis 取消点赞失败")
			return
		}
		resp = &videoservice.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "已取消点赞",
		}
		return
	}

	// 在数据库中查询点赞信息
	q := query.Q
	favorite := q.TFavorite
	first, _ := q.WithContext(context.Background()).TFavorite.Where(favorite.UserID.Eq(claims.UserId), favorite.VideoID.Eq(req.VideoId)).First()

	// 查询为空
	if first == nil {
		// 将点赞存入redis
		cache.RedisCache.RedisClient.SAdd(context.Background(), consts.FavoriteActionPrefix+strconv.FormatInt(req.VideoId, 10), strconv.FormatInt(claims.UserId, 10), 0)
		resp = &videoservice.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "已成功点赞",
		}
		return
	}

	// 查询数据库，数据库为已点赞，取消点赞
	if first.Status {
		_, err1 := q.WithContext(context.Background()).TFavorite.Update(favorite.Status, false)
		if err1 != nil {
			err1 = fmt.Errorf("更新数据库失败")
		}
		resp = &videoservice.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "已取消点赞",
		}
		return
	}

	resp = &videoservice.DouyinFavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  "已成功点赞",
	}
	return
}

// FavoriteList 2023-1-27 @Auth by 李卓轩 version 1.0
// 喜欢列表
// FavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteList(ctx context.Context, req *videoservice.DouyinFavoriteListRequest) (resp *videoservice.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	// 通过 token 解析出当前用户
	claims, flag := jwtutil.CheckToken(req.Token)
	// 说明 token 已经过期
	if !flag {
		return nil, errors.New("token is expired")
	}

	q := query.Q
	favorite := q.TFavorite
	// 查询数据库得到喜欢列表
	data, err := q.WithContext(context.Background()).TFavorite.Where(favorite.UserID.Eq(claims.UserId)).Find()
	ids := make([]int64, 10)
	//得到喜欢视频的所有id
	for _, fav := range data {
		ids = append(ids, fav.VideoID)
	}

	//查询所有的喜欢视频信息
	video := q.TVideo
	find, err := q.WithContext(context.Background()).TVideo.Where(video.ID.In(ids...)).Find()
	if err != nil {
		err = fmt.Errorf("查询失败")
	}
	var videos []*videoservice.Video
	//通过用用户id查询用户
	Tuser := q.TUser
	for _, videosInfo := range find {
		var vid videoservice.Video
		var usr videoservice.User
		vid.FavoriteCount = videosInfo.FavoriteCount
		vid.Id = videosInfo.ID
		vid.CoverUrl = videosInfo.CoverURL
		vid.PlayUrl = videosInfo.PlayURL
		vid.IsFavorite = videosInfo.IsFavorite
		vid.Title = videosInfo.Title
		first, _ := q.WithContext(context.Background()).TUser.Where(Tuser.ID.Eq(videosInfo.AuthorID)).First()
		usr.Id = first.ID
		usr.Name = first.Name
		usr.FollowCount = first.FollowerCount
		usr.FollowerCount = first.FollowerCount
		vid.Author = &usr
		videos = append(videos, &vid)
	}

	resp = &videoservice.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "成功",
		VideoList:  videos,
	}
	return
}

// CommentAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentAction(ctx context.Context, req *videoservice.DouyinCommentActionRequest) (resp *videoservice.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	// 评论操作
	queryUser := query.Q.TUser
	queryVideo := query.Q.TVideo
	queryComment := query.Q.TComment
	timeLayoutStr := "2006-01-02 15:04:05"
	// 解析 token 拿取用户id
	claims, flag := jwtutil.CheckToken(req.Token)
	if !flag {
		return nil, errors.New("token is expired")
	}
	// 判断视频是否存在
	_, err = queryVideo.WithContext(ctx).Where(queryVideo.ID.Eq(req.VideoId)).First()
	if err != nil {
		return nil, errors.New("video does not exist")
	}

	// 发布评论
	if req.ActionType == 1 {
		comment := &model.TComment{
			UserID:     claims.UserId,
			Content:    req.CommentText,
			VideoID:    req.VideoId,
			CreateDate: time.Now(),
		}

		err := queryComment.WithContext(ctx).Create(comment)
		user, _ := queryUser.WithContext(ctx).Select(queryUser.ID, queryUser.Name).
			Where(queryUser.ID.Eq(claims.UserId)).First()
		if err != nil {
			return nil, errors.New("add failure")
		}
		resp = &videoservice.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "评论成功",
			Comment: &videoservice.Comment{
				Id: comment.ID,
				User: &videoservice.User{
					Id:            user.ID,
					Name:          user.Name,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
				},
				Content:    comment.Content,
				CreateDate: comment.CreateDate.Format(timeLayoutStr),
			},
		}
		// 删除评论
	} else if req.ActionType == 2 {
		// 用户是否有此条评论
		_, err := queryComment.WithContext(ctx).Where(queryComment.ID.Eq(req.CommentId)).
			Where(queryComment.UserID.Eq(claims.UserId)).Delete()
		if err != nil {
			return nil, errors.New("comment does not exist")
		}
		resp = &videoservice.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "删除成功",
		}
	} else {
		return nil, errors.New("operation error")
	}

	return
}

// CommentList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentList(ctx context.Context, req *videoservice.DouyinCommentListRequest) (resp *videoservice.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	// 获取评论
	queryUser := query.Q.TUser
	queryComment := query.Q.TComment
	// 数据库取出的 result
	type Result struct {
		Content    string
		CreateDate string
		ID         int64
		userID     int64
		Name       string
	}
	var result []Result
	// 解析 token
	_, flag := jwtutil.CheckToken(req.Token)
	// 登录后以查看评论
	if !flag {
		return nil, errors.New("log in to view the comments")
	}
	// 查询视频下的评论
	// 运用 left join 联合查询
	err = queryComment.WithContext(ctx).
		Select(queryComment.Content, queryComment.CreateDate, queryUser.ID, queryUser.Name).LeftJoin(&queryUser, queryUser.ID.EqCol(queryComment.UserID)).Where(queryComment.VideoID.Eq(req.VideoId)).Scan(&result)
	if err != nil {
		return
	}
	var comment videoservice.Comment
	var comments []*videoservice.Comment
	for _, com := range result {
		// 序列化
		user := videoservice.User{
			Id:   com.userID,
			Name: com.Name,
		}
		comment.User = &user
		comment.Content = com.Content
		comment.CreateDate = com.CreateDate
		// 这里要再创建一个干净的变量，要不然会只传最后一个
		var com videoservice.Comment
		com = comment
		comments = append(comments, &com)
	}

	resp = &videoservice.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "the request succeeded",
		CommentList: comments,
	}

	return
}
