// Code generated by goctl. DO NOT EDIT.
package types

type CommentDeleteRequest struct {
	CommentId int64 `json:"comment_id"`
}

type CommentDeleteResponse struct {
	Status string `json:"status"`
}

type CommentInfo struct {
	Id             int64         `json:"id"`
	ObjId          int64         `json:"objId"`
	UserId         int64         `json:"userId"`
	BeReplayUserId int64         `json:"beReplayUserId"`
	Nickname       string        `json:"nickname"`
	Avatar         string        `json:"avatar"`
	Content        string        `json:"content"`
	AddTime        int64         `json:"addTime"`
	LikeNums       int64         `json:"likeNums"`
	IsLike         uint8         `json:"isLike"`
	SuperNickname  string        `json:"superNickname"`
	SuperCommentId int64         `json:"superCommentId"`
	Children       []CommentInfo `json:"children"`
}

type CommentListRequest struct {
	VideoId int64 `json:"video_id"`
}

type CommentListResponse struct {
	CommentList []CommentInfo `json:"comment_list"`
}

type CommentPublishRequest struct {
	VideoId        int64  `json:"video_id"`
	Content        string `json:"content"`
	SuperCommentId int64  `json:"superCommentId"`
	BeReplayUserId int64  `json:"beReplayUserId"`
}

type CommentPublishResponse struct {
	CommentId int64 `json:"comment_id"`
}

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailLoginResponse struct {
	UserId      int64  `json:"user_id"`
	AccessToken string `json:"token"`
}

type EmailRegisterRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type EmailRegisterResponse struct {
	UserId      int64  `json:"user_id"`
	AccessToken string `json:"token"`
}

type LikeActionRequest struct {
	ActionType int32  `json:"action_type"`
	BizId      string `json:"biz_id"`
	ObjId      int64  `path:"obj_id"` // 对象ID
}

type LikeActionResponse struct {
	Status string `json:"status"`
}

type PublishHandlerRequest struct {
	VideoUrl    string `json:"video_url"`
	CoverUrl    string `json:"cover_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PublishHandlerResponse struct {
	VideoId int64 `json:"video_id"`
}

type UploadUrlResponse struct {
	VideoUrl string `json:"video_url"`
	CoverUrl string `json:"cover_url"`
}

type UserInfoResponse struct {
	UserId        int64  `json:"user_id"`
	UserName      string `json:"name"`
	Email         string `json:"email"`
	Avatar        string `json:"avatar"`
	Signature     string `json:"signature"`
	FollowerCount int64  `json:"follower_count"`
	FollowedCount int64  `json:"followed_count"`
}

type UserLikeListRequest struct {
	UserId int64 `json:"user_id"`
}

type UserLikeListResponse struct {
	VideoList []VideoInfo `json:"video_list"`
}

type UserVideoListRequest struct {
	UserId int64 `json:"user_id"`
}

type UserVideoListResponse struct {
	VideoList []VideoInfo `json:"video_list"`
}

type VerificationRequest struct {
	Email string `json:"email"`
}

type VerificationResponse struct {
}

type VideoInfo struct {
	VideoId      int64  `json:"video_id"`
	AuthorId     int64  `json:"author_id"`
	AuthorName   string `json:"author_name"`
	AuthorAvatar string `json:"author_avatar"`
	VideoUrl     string `json:"video_url"`
	Title        string `json:"title"`
	CoverUrl     string `json:"cover_url"`
	Description  string `json:"description"`
	CommentCount int64  `json:"comment_count"`
	LikeCount    int64  `json:"like_count"`
}

type VideoListRequest struct {
	Page     int64 `json:"page"`
	Cursor   int64 `json:"cursor"`
	PageSize int64 `json:"page_size"`
	FeedType int32 `json:"feed_type"`
}

type VideoListResponse struct {
	VideoList []VideoInfo `json:"video_list"`
}
