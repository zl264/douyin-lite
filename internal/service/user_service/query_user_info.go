package user_service

import (
	"errors"
)

type UserInfo struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowingCount  int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

func QueryUserInfo(userId int64) (*UserInfo, error) {
	return NewQueryUserInfoFlow(userId).Do()
}

type QueryUserInfoFlow struct {
	userId   int64
	userInfo *UserInfo
}

func NewQueryUserInfoFlow(userId int64) *QueryUserInfoFlow {
	return &QueryUserInfoFlow{
		userId: userId,
	}
}

func (f *QueryUserInfoFlow) Do() (*UserInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.prepareUserInfo()
	if err != nil {
		return nil, err
	}
	return f.userInfo, nil
}

func (f *QueryUserInfoFlow) checkParam() error {
	if f.userId < 0 {
		return errors.New("user id should be larger than 0")
	}
	return nil
}

func (f *QueryUserInfoFlow) prepareUserInfo() error {
	userInfo, err := QueryAUserInfo2(f.userId)
	if err != nil {
		return err
	}
	f.userInfo = userInfo
	return nil
}