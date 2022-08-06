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
	CreateFile(ctx context.Context, arg *CreateFileParams) (*File, error)
	CreateFriendRelation(ctx context.Context, arg *CreateFriendRelationParams) (int64, error)
	CreateGroupNotify(ctx context.Context, arg *CreateGroupNotifyParams) (*CreateGroupNotifyRow, error)
	CreateGroupRelation(ctx context.Context, arg *CreateGroupRelationParams) (int64, error)
	CreateMsg(ctx context.Context, arg *CreateMsgParams) (*CreateMsgRow, error)
	CreateSetting(ctx context.Context, arg *CreateSettingParams) error
	CreateUser(ctx context.Context, arg *CreateUserParams) (*User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteAccountsByUserID(ctx context.Context, userID int64) ([]int64, error)
	DeleteApplication(ctx context.Context, arg *DeleteApplicationParams) error
	DeleteFileByID(ctx context.Context, id int64) error
	DeleteFriendRelationsByAccountID(ctx context.Context, accountID int64) error
	DeleteFriendRelationsByAccountIDs(ctx context.Context, accountIds []int64) error
	DeleteGroup(ctx context.Context, relationID int64) error
	DeleteGroupNotify(ctx context.Context, id int64) error
	DeleteRelation(ctx context.Context, id int64) error
	DeleteSetting(ctx context.Context, arg *DeleteSettingParams) error
	DeleteUser(ctx context.Context, id int64) error
	ExistEmail(ctx context.Context, email string) (bool, error)
	ExistsAccountByID(ctx context.Context, id int64) (bool, error)
	ExistsAccountByNameAndUserID(ctx context.Context, arg *ExistsAccountByNameAndUserIDParams) (bool, error)
	ExistsApplicationByID(ctx context.Context, arg *ExistsApplicationByIDParams) (bool, error)
	ExistsApplicationByIDWithLock(ctx context.Context, arg *ExistsApplicationByIDWithLockParams) (bool, error)
	ExistsFriendRelation(ctx context.Context, arg *ExistsFriendRelationParams) (bool, error)
	ExistsFriendSetting(ctx context.Context, arg *ExistsFriendSettingParams) (bool, error)
	ExistsIsLeader(ctx context.Context, arg *ExistsIsLeaderParams) (bool, error)
	ExistsSetting(ctx context.Context, arg *ExistsSettingParams) (bool, error)
	ExistsUserByID(ctx context.Context, id int64) (bool, error)
	GetAccountByID(ctx context.Context, arg *GetAccountByIDParams) (*GetAccountByIDRow, error)
	GetAccountsByName(ctx context.Context, arg *GetAccountsByNameParams) ([]*GetAccountsByNameRow, error)
	GetAccountsByUserID(ctx context.Context, userID int64) ([]*GetAccountsByUserIDRow, error)
	GetAllEmails(ctx context.Context) ([]string, error)
	GetAllGroupRelation(ctx context.Context) ([]int64, error)
	GetAllRelationOnRelation(ctx context.Context) ([]int64, error)
	GetAllRelationsOnFile(ctx context.Context) ([]sql.NullInt64, error)
	GetApplicationByID(ctx context.Context, arg *GetApplicationByIDParams) (*Application, error)
	GetApplications(ctx context.Context, arg *GetApplicationsParams) ([]*GetApplicationsRow, error)
	GetAvatar(ctx context.Context, accountID sql.NullInt64) (*GetAvatarRow, error)
	GetFileByRelationID(ctx context.Context, relationID sql.NullInt64) ([]*File, error)
	GetFileByRelationIDIsNUll(ctx context.Context) ([]*GetFileByRelationIDIsNUllRow, error)
	GetFileKeyByID(ctx context.Context, id int64) (string, error)
	GetFriendPinSettingsOrderByPinTime(ctx context.Context, accountID int64) ([]*GetFriendPinSettingsOrderByPinTimeRow, error)
	GetFriendRelationByID(ctx context.Context, id int64) (*GetFriendRelationByIDRow, error)
	GetFriendSettingsByName(ctx context.Context, arg *GetFriendSettingsByNameParams) ([]*GetFriendSettingsByNameRow, error)
	GetFriendSettingsOrderByName(ctx context.Context, accountID int64) ([]*GetFriendSettingsOrderByNameRow, error)
	GetFriendShowSettingsOrderByShowTime(ctx context.Context, accountID int64) ([]*GetFriendShowSettingsOrderByShowTimeRow, error)
	GetGroupAvatar(ctx context.Context, relationID sql.NullInt64) (*File, error)
	GetGroupMembers(ctx context.Context, relationID int64) ([]int64, error)
	GetGroupNotifyByID(ctx context.Context, relationID sql.NullInt64) ([]*GetGroupNotifyByIDRow, error)
	GetGroupPinSettingsOrderByPinTime(ctx context.Context, accountID int64) ([]*GetGroupPinSettingsOrderByPinTimeRow, error)
	GetGroupRelationByID(ctx context.Context, id int64) (*GetGroupRelationByIDRow, error)
	GetGroupShowSettingsOrderByShowTime(ctx context.Context, accountID int64) ([]*GetGroupShowSettingsOrderByShowTimeRow, error)
	GetMsgByID(ctx context.Context, id int64) (*GetMsgByIDRow, error)
	GetMsgsByRelationIDAndTime(ctx context.Context, arg *GetMsgsByRelationIDAndTimeParams) ([]*GetMsgsByRelationIDAndTimeRow, error)
	GetPinMsgsByRelationID(ctx context.Context, arg *GetPinMsgsByRelationIDParams) ([]*GetPinMsgsByRelationIDRow, error)
	GetRlyMsgsInfoByMsgID(ctx context.Context, arg *GetRlyMsgsInfoByMsgIDParams) ([]*GetRlyMsgsInfoByMsgIDRow, error)
	GetSettingByID(ctx context.Context, arg *GetSettingByIDParams) (*Setting, error)
	GetTopMsgByRelationID(ctx context.Context, relationID int64) (*GetTopMsgByRelationIDRow, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	HasReadMsg(ctx context.Context, arg *HasReadMsgParams) (bool, error)
	TransferIsSelfFalse(ctx context.Context, arg *TransferIsSelfFalseParams) error
	TransferIsSelfTrue(ctx context.Context, arg *TransferIsSelfTrueParams) error
	UpdateAccount(ctx context.Context, arg *UpdateAccountParams) error
	UpdateApplication(ctx context.Context, arg *UpdateApplicationParams) error
	UpdateGroupAvatar(ctx context.Context, arg *UpdateGroupAvatarParams) error
	UpdateGroupNotify(ctx context.Context, arg *UpdateGroupNotifyParams) (*UpdateGroupNotifyRow, error)
	UpdateGroupRelation(ctx context.Context, arg *UpdateGroupRelationParams) error
	UpdateMsgPin(ctx context.Context, arg *UpdateMsgPinParams) error
	UpdateMsgReads(ctx context.Context, arg *UpdateMsgReadsParams) error
	UpdateMsgRevoke(ctx context.Context, arg *UpdateMsgRevokeParams) error
	UpdateMsgTopFalseByMsgID(ctx context.Context, id int64) error
	UpdateMsgTopFalseByRelationID(ctx context.Context, relationID int64) error
	UpdateMsgTopTrueByMsgID(ctx context.Context, id int64) error
	UpdateSettingDisturb(ctx context.Context, arg *UpdateSettingDisturbParams) error
	UpdateSettingLeader(ctx context.Context, arg *UpdateSettingLeaderParams) error
	UpdateSettingNickName(ctx context.Context, arg *UpdateSettingNickNameParams) error
	UpdateSettingPin(ctx context.Context, arg *UpdateSettingPinParams) error
	UpdateUser(ctx context.Context, arg *UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
