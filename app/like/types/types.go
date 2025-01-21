package types

import "zhihu/app/like/pb/like"

type LikeAction struct {
	UserId     int64  `json:"user_id"`
	BizId      string `json:"biz_id"`
	ObjId      int64  `json:"obj_id"`
	ActionType like.LikeActionRequest_ActionType
}
