// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CountAccountByUserID(ctx context.Context, userID int64) (int32, error)
	CreateAccount(ctx context.Context, arg *CreateAccountParams) error
	CreateApplication(ctx context.Context, arg *CreateApplicationParams) error
	CreateFriendRelation(ctx context.Context, arg *CreateFriendRelationParams) (int64, error)
	CreateFile(ctx context.Context, arg *CreateFileParams) (*File, error)
	CreateGroupNotify(ctx context.Context, arg *CreateGroupNotifyParams) (*GroupNotify, error)
	CreateGroupRelation(ctx context.Context, arg *CreateGroupRelationParams) error
	CreateMsg(ctx context.Context, arg *CreateMsgParams) (*Message, error)
	CreateSetting(ctx context.Context, arg *CreateSettingParams) error
	CreateRelationSetting(ctx context.Context, arg *CreateRelationSettingParams) (*RelationSetting, error)
	CreateUser(ctx context.Context, arg *CreateUserParams) (*User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteAccountsByUserID(ctx context.Context, userID int64) ([]int64, error)
	DeleteApplication(ctx context.Context, arg *DeleteApplicationParams) error
	DeleteFriendRelationsByAccountID(ctx context.Context, accountID int64) error
	DeleteFriendRelationsByAccountIDs(ctx context.Context, accountIds []int64) error
	DeleteFileByID(ctx context.Context, arg *DeleteFileByIDParams) error
	DeleteGroupNotify(ctx context.Context, id int64) error
	DeleteRelation(ctx context.Context, id int64) error
	DeleteSetting(ctx context.Context, arg *DeleteSettingParams) error
	DeleteRelationSetting(ctx context.Context, arg *DeleteRelationSettingParams) error
	DeleteUser(ctx context.Context, id int64) error
	ExistEmail(ctx context.Context, email string) (bool, error)
	ExistsAccountByID(ctx context.Context, id int64) (bool, error)
	ExistsAccountByNameAndUserID(ctx context.Context, arg *ExistsAccountByNameAndUserIDParams) (bool, error)
	ExistsApplicationByID(ctx context.Context, arg *ExistsApplicationByIDParams) (bool, error)
	ExistsApplicationByIDWithLock(ctx context.Context, arg *ExistsApplicationByIDWithLockParams) (bool, error)
	ExistsFriendRelation(ctx context.Context, arg *ExistsFriendRelationParams) (bool, error)
	ExistsFriendSetting(ctx context.Context, arg *ExistsFriendSettingParams) (bool, error)
	ExistsSetting(ctx context.Context, arg *ExistsSettingParams) (bool, error)
	ExistsUserByID(ctx context.Context, id int64) (bool, error)
	GetAccountByID(ctx context.Context, arg *GetAccountByIDParams) (*GetAccountByIDRow, error)
	GetAccountsByName(ctx context.Context, arg *GetAccountsByNameParams) ([]*GetAccountsByNameRow, error)
	GetAccountsByUserID(ctx context.Context, userID int64) ([]*GetAccountsByUserIDRow, error)
	GetAllEmails(ctx context.Context) ([]string, error)
	GetFileKeyByID(ctx context.Context, arg *GetFileKeyByIDParams) (interface{}, error)
	GetApplicationByID(ctx context.Context, arg *GetApplicationByIDParams) (*Application, error)
	GetApplications(ctx context.Context, arg *GetApplicationsParams) ([]*GetApplicationsRow, error)
	GetFriendPinSettingsOrderByPinTime(ctx context.Context, accountID int64) ([]*GetFriendPinSettingsOrderByPinTimeRow, error)
	GetFriendRelationByID(ctx context.Context, id int64) (*GetFriendRelationByIDRow, error)
	GetFriendSettingsByName(ctx context.Context, arg *GetFriendSettingsByNameParams) ([]*GetFriendSettingsByNameRow, error)
	GetFriendSettingsOrderByName(ctx context.Context, accountID int64) ([]*GetFriendSettingsOrderByNameRow, error)
	GetFriendShowSettingsOrderByShowTime(ctx context.Context, accountID int64) ([]*GetFriendShowSettingsOrderByShowTimeRow, error)
	GetFileByRelationID(ctx context.Context, relationID sql.NullInt64) ([]*File, error)
	GetGroupNotifyByID(ctx context.Context, relationID sql.NullInt64) (*GroupNotify, error)
	GetGroupRelationByID(ctx context.Context, id int64) (*GetGroupRelationByIDRow, error)
	GetMsgByID(ctx context.Context, id int64) (*Message, error)
	GetMsgsByRelationIDAndTime(ctx context.Context, arg *GetMsgsByRelationIDAndTimeParams) ([]*GetMsgsByRelationIDAndTimeRow, error)
	GetPinMsgsByRelationID(ctx context.Context, arg *GetPinMsgsByRelationIDParams) ([]*GetPinMsgsByRelationIDRow, error)
	GetRlyMsgsInfoByMsgID(ctx context.Context, arg *GetRlyMsgsInfoByMsgIDParams) ([]*GetRlyMsgsInfoByMsgIDRow, error)
	GetSettingByID(ctx context.Context, arg *GetSettingByIDParams) (*Setting, error)
	GetRelationSetting(ctx context.Context, arg *GetRelationSettingParams) (*RelationSetting, error)
	GetTopMsgByRelationID(ctx context.Context, relationID int64) (*GetTopMsgByRelationIDRow, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	HasReadMsg(ctx context.Context, arg *HasReadMsgParams) (bool, error)
	UpdateAccount(ctx context.Context, arg *UpdateAccountParams) error
	UpdateApplication(ctx context.Context, arg *UpdateApplicationParams) error
	UpdateGroupNotify(ctx context.Context, arg *UpdateGroupNotifyParams) (*GroupNotify, error)
	UpdateGroupRelation(ctx context.Context, arg *UpdateGroupRelationParams) error
	UpdateMsgPin(ctx context.Context, arg *UpdateMsgPinParams) error
	UpdateMsgReads(ctx context.Context, arg *UpdateMsgReadsParams) error
	UpdateMsgRevoke(ctx context.Context, arg *UpdateMsgRevokeParams) error
	UpdateMsgTopFalseByMsgID(ctx context.Context, id int64) error
	UpdateMsgTopFalseByRelationID(ctx context.Context, relationID int64) error
	UpdateMsgTopTrueByMsgID(ctx context.Context, id int64) error
	UpdateRelationSetting(ctx context.Context, arg *UpdateRelationSettingParams) (*RelationSetting, error)
	UpdateSettingDisturb(ctx context.Context, arg *UpdateSettingDisturbParams) error
	UpdateSettingLeader(ctx context.Context, arg *UpdateSettingLeaderParams) error
	UpdateSettingNickName(ctx context.Context, arg *UpdateSettingNickNameParams) error
	UpdateSettingPin(ctx context.Context, arg *UpdateSettingPinParams) error
	UpdateUser(ctx context.Context, arg *UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
