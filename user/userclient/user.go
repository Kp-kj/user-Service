// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"user/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddAdminRequest                       = user.AddAdminRequest
	AddAdminResponse                      = user.AddAdminResponse
	AddBlackListRequest                   = user.AddBlackListRequest
	AddBlackListResponse                  = user.AddBlackListResponse
	AddUserInfoRequest                    = user.AddUserInfoRequest
	AddUserInfoResponse                   = user.AddUserInfoResponse
	AdminLoginRequest                     = user.AdminLoginRequest
	AdminLoginResponse                    = user.AdminLoginResponse
	CheckTodayInviteRequest               = user.CheckTodayInviteRequest
	CheckTodayInviteResponse              = user.CheckTodayInviteResponse
	CheckTwitterIdRequest                 = user.CheckTwitterIdRequest
	CheckTwitterIdResponse                = user.CheckTwitterIdResponse
	CreateHelpCategoryRequest             = user.CreateHelpCategoryRequest
	CreateHelpCategoryResponse            = user.CreateHelpCategoryResponse
	CreateHelpCategoryTranslationRequest  = user.CreateHelpCategoryTranslationRequest
	CreateHelpCategoryTranslationResponse = user.CreateHelpCategoryTranslationResponse
	CreateHelpDocumentRequest             = user.CreateHelpDocumentRequest
	CreateHelpDocumentResponse            = user.CreateHelpDocumentResponse
	CreateHelpDocumentTranslationRequest  = user.CreateHelpDocumentTranslationRequest
	CreateHelpDocumentTranslationResponse = user.CreateHelpDocumentTranslationResponse
	CreateInviteRequest                   = user.CreateInviteRequest
	CreateInviteResponse                  = user.CreateInviteResponse
	CreateNoticeRequest                   = user.CreateNoticeRequest
	CreateNoticeResponse                  = user.CreateNoticeResponse
	CreateNotificationRequest             = user.CreateNotificationRequest
	CreateNotificationResponse            = user.CreateNotificationResponse
	CreateSystemNotificationRequest       = user.CreateSystemNotificationRequest
	CreateSystemNotificationResponse      = user.CreateSystemNotificationResponse
	CreateUserRequest                     = user.CreateUserRequest
	CreateUserResponse                    = user.CreateUserResponse
	DeleteHelpCategoryRequest             = user.DeleteHelpCategoryRequest
	DeleteHelpCategoryResponse            = user.DeleteHelpCategoryResponse
	DeleteHelpCategoryTranslationRequest  = user.DeleteHelpCategoryTranslationRequest
	DeleteHelpCategoryTranslationResponse = user.DeleteHelpCategoryTranslationResponse
	DeleteHelpDocumentRequest             = user.DeleteHelpDocumentRequest
	DeleteHelpDocumentResponse            = user.DeleteHelpDocumentResponse
	DeleteHelpDocumentTranslationRequest  = user.DeleteHelpDocumentTranslationRequest
	DeleteHelpDocumentTranslationResponse = user.DeleteHelpDocumentTranslationResponse
	EditHelpCategoryRequest               = user.EditHelpCategoryRequest
	EditHelpCategoryResponse              = user.EditHelpCategoryResponse
	EditHelpCategoryTranslationRequest    = user.EditHelpCategoryTranslationRequest
	EditHelpCategoryTranslationResponse   = user.EditHelpCategoryTranslationResponse
	EditHelpDocumentRequest               = user.EditHelpDocumentRequest
	EditHelpDocumentResponse              = user.EditHelpDocumentResponse
	EditHelpDocumentTranslationRequest    = user.EditHelpDocumentTranslationRequest
	EditHelpDocumentTranslationResponse   = user.EditHelpDocumentTranslationResponse
	EditSystemNotificationRequest         = user.EditSystemNotificationRequest
	EditSystemNotificationResponse        = user.EditSystemNotificationResponse
	GetHelpCategoriesRequest              = user.GetHelpCategoriesRequest
	GetHelpCategoriesResponse             = user.GetHelpCategoriesResponse
	GetHelpCategoryTranslationsRequest    = user.GetHelpCategoryTranslationsRequest
	GetHelpCategoryTranslationsResponse   = user.GetHelpCategoryTranslationsResponse
	GetHelpDocumentTranslationsRequest    = user.GetHelpDocumentTranslationsRequest
	GetHelpDocumentTranslationsResponse   = user.GetHelpDocumentTranslationsResponse
	GetHelpDocumentsRequest               = user.GetHelpDocumentsRequest
	GetHelpDocumentsResponse              = user.GetHelpDocumentsResponse
	GetNotificationsRequest               = user.GetNotificationsRequest
	GetNotificationsResponse              = user.GetNotificationsResponse
	GetSystemNotificationsRequest         = user.GetSystemNotificationsRequest
	GetSystemNotificationsResponse        = user.GetSystemNotificationsResponse
	GetUserNotificationsRequest           = user.GetUserNotificationsRequest
	GetUserNotificationsResponse          = user.GetUserNotificationsResponse
	HelpCategory                          = user.HelpCategory
	HelpDocument                          = user.HelpDocument
	Notice                                = user.Notice
	Notification                          = user.Notification
	OnlineCountResponse                   = user.OnlineCountResponse
	QueryBlackListRequest                 = user.QueryBlackListRequest
	QueryBlackListResponse                = user.QueryBlackListResponse
	QueryHelpCategoryRequest              = user.QueryHelpCategoryRequest
	QueryHelpCategoryResponse             = user.QueryHelpCategoryResponse
	QueryHelpDocumentRequest              = user.QueryHelpDocumentRequest
	QueryHelpDocumentResponse             = user.QueryHelpDocumentResponse
	QueryRecordNoticeRequest              = user.QueryRecordNoticeRequest
	QueryRecordNoticeResponse             = user.QueryRecordNoticeResponse
	QuerySystemNotificationRequest        = user.QuerySystemNotificationRequest
	QuerySystemNotificationResponse       = user.QuerySystemNotificationResponse
	QueryUserRequest                      = user.QueryUserRequest
	QueryUserResponse                     = user.QueryUserResponse
	RecordNoticeRequest                   = user.RecordNoticeRequest
	RecordNoticeResponse                  = user.RecordNoticeResponse
	RegisterCountResponse                 = user.RegisterCountResponse
	RemoveAdminRequest                    = user.RemoveAdminRequest
	RemoveAdminResponse                   = user.RemoveAdminResponse
	RemoveBlackListRequest                = user.RemoveBlackListRequest
	RemoveBlackListResponse               = user.RemoveBlackListResponse
	Request                               = user.Request
	Response                              = user.Response
	SystemNotification                    = user.SystemNotification
	TotalCategory                         = user.TotalCategory
	TotalDocument                         = user.TotalDocument
	UserListResponse                      = user.UserListResponse
	UserResponseList                      = user.UserResponseList

	User interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		RegisterCount(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RegisterCountResponse, error)
		OnlineCount(ctx context.Context, in *Request, opts ...grpc.CallOption) (*OnlineCountResponse, error)
		CheckTwitterId(ctx context.Context, in *CheckTwitterIdRequest, opts ...grpc.CallOption) (*CheckTwitterIdResponse, error)
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
		CreateInvite(ctx context.Context, in *CreateInviteRequest, opts ...grpc.CallOption) (*CreateInviteResponse, error)
		CheckTodayInvite(ctx context.Context, in *CheckTodayInviteRequest, opts ...grpc.CallOption) (*CheckTodayInviteResponse, error)
		AddUserInfo(ctx context.Context, in *AddUserInfoRequest, opts ...grpc.CallOption) (*AddUserInfoResponse, error)
		QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error)
		GetUserList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*UserListResponse, error)
		AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error)
		AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error)
		RemoveAdmin(ctx context.Context, in *RemoveAdminRequest, opts ...grpc.CallOption) (*RemoveAdminResponse, error)
		AddBlackList(ctx context.Context, in *AddBlackListRequest, opts ...grpc.CallOption) (*AddBlackListResponse, error)
		QueryBlackList(ctx context.Context, in *QueryBlackListRequest, opts ...grpc.CallOption) (*QueryBlackListResponse, error)
		RemoveBlackList(ctx context.Context, in *RemoveBlackListRequest, opts ...grpc.CallOption) (*RemoveBlackListResponse, error)
		GetHelpCategories(ctx context.Context, in *GetHelpCategoriesRequest, opts ...grpc.CallOption) (*GetHelpCategoriesResponse, error)
		CreateHelpCategory(ctx context.Context, in *CreateHelpCategoryRequest, opts ...grpc.CallOption) (*CreateHelpCategoryResponse, error)
		DeleteHelpCategory(ctx context.Context, in *DeleteHelpCategoryRequest, opts ...grpc.CallOption) (*DeleteHelpCategoryResponse, error)
		EditHelpCategory(ctx context.Context, in *EditHelpCategoryRequest, opts ...grpc.CallOption) (*EditHelpCategoryResponse, error)
		CreateHelpCategoryTranslation(ctx context.Context, in *CreateHelpCategoryTranslationRequest, opts ...grpc.CallOption) (*CreateHelpCategoryTranslationResponse, error)
		DeleteHelpCategoryTranslation(ctx context.Context, in *DeleteHelpCategoryTranslationRequest, opts ...grpc.CallOption) (*DeleteHelpCategoryTranslationResponse, error)
		GetHelpCategoryTranslations(ctx context.Context, in *GetHelpCategoryTranslationsRequest, opts ...grpc.CallOption) (*GetHelpCategoryTranslationsResponse, error)
		EditHelpCategoryTranslation(ctx context.Context, in *EditHelpCategoryTranslationRequest, opts ...grpc.CallOption) (*EditHelpCategoryTranslationResponse, error)
		QueryHelpCategory(ctx context.Context, in *QueryHelpCategoryRequest, opts ...grpc.CallOption) (*QueryHelpCategoryResponse, error)
		GetHelpDocuments(ctx context.Context, in *GetHelpDocumentsRequest, opts ...grpc.CallOption) (*GetHelpDocumentsResponse, error)
		CreateHelpDocument(ctx context.Context, in *CreateHelpDocumentRequest, opts ...grpc.CallOption) (*CreateHelpDocumentResponse, error)
		DeleteHelpDocument(ctx context.Context, in *DeleteHelpDocumentRequest, opts ...grpc.CallOption) (*DeleteHelpDocumentResponse, error)
		EditHelpDocument(ctx context.Context, in *EditHelpDocumentRequest, opts ...grpc.CallOption) (*EditHelpDocumentResponse, error)
		CreateHelpDocumentTranslation(ctx context.Context, in *CreateHelpDocumentTranslationRequest, opts ...grpc.CallOption) (*CreateHelpDocumentTranslationResponse, error)
		DeleteHelpDocumentTranslation(ctx context.Context, in *DeleteHelpDocumentTranslationRequest, opts ...grpc.CallOption) (*DeleteHelpDocumentTranslationResponse, error)
		GetHelpDocumentTranslations(ctx context.Context, in *GetHelpDocumentTranslationsRequest, opts ...grpc.CallOption) (*GetHelpDocumentTranslationsResponse, error)
		EditHelpDocumentTranslation(ctx context.Context, in *EditHelpDocumentTranslationRequest, opts ...grpc.CallOption) (*EditHelpDocumentTranslationResponse, error)
		QueryHelpDocument(ctx context.Context, in *QueryHelpDocumentRequest, opts ...grpc.CallOption) (*QueryHelpDocumentResponse, error)
		CreateSystemNotification(ctx context.Context, in *CreateSystemNotificationRequest, opts ...grpc.CallOption) (*CreateSystemNotificationResponse, error)
		EditSystemNotification(ctx context.Context, in *EditSystemNotificationRequest, opts ...grpc.CallOption) (*EditSystemNotificationResponse, error)
		GetSystemNotifications(ctx context.Context, in *GetSystemNotificationsRequest, opts ...grpc.CallOption) (*GetSystemNotificationsResponse, error)
		QuerySystemNotification(ctx context.Context, in *QuerySystemNotificationRequest, opts ...grpc.CallOption) (*QuerySystemNotificationResponse, error)
		// recordNotice 通知记录
		CreateNotice(ctx context.Context, in *CreateNoticeRequest, opts ...grpc.CallOption) (*CreateNoticeResponse, error)
		RecordNotice(ctx context.Context, in *RecordNoticeRequest, opts ...grpc.CallOption) (*RecordNoticeResponse, error)
		QueryRecordNotice(ctx context.Context, in *QueryRecordNoticeRequest, opts ...grpc.CallOption) (*QueryRecordNoticeResponse, error)
		// 新增用户消息通知
		CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error)
		// 获取用户消息通知
		GetUserNotifications(ctx context.Context, in *GetUserNotificationsRequest, opts ...grpc.CallOption) (*GetUserNotificationsResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUser) RegisterCount(ctx context.Context, in *Request, opts ...grpc.CallOption) (*RegisterCountResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RegisterCount(ctx, in, opts...)
}

func (m *defaultUser) OnlineCount(ctx context.Context, in *Request, opts ...grpc.CallOption) (*OnlineCountResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.OnlineCount(ctx, in, opts...)
}

func (m *defaultUser) CheckTwitterId(ctx context.Context, in *CheckTwitterIdRequest, opts ...grpc.CallOption) (*CheckTwitterIdResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CheckTwitterId(ctx, in, opts...)
}

func (m *defaultUser) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUser) CreateInvite(ctx context.Context, in *CreateInviteRequest, opts ...grpc.CallOption) (*CreateInviteResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateInvite(ctx, in, opts...)
}

func (m *defaultUser) CheckTodayInvite(ctx context.Context, in *CheckTodayInviteRequest, opts ...grpc.CallOption) (*CheckTodayInviteResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CheckTodayInvite(ctx, in, opts...)
}

func (m *defaultUser) AddUserInfo(ctx context.Context, in *AddUserInfoRequest, opts ...grpc.CallOption) (*AddUserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddUserInfo(ctx, in, opts...)
}

func (m *defaultUser) QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.QueryUser(ctx, in, opts...)
}

func (m *defaultUser) GetUserList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*UserListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserList(ctx, in, opts...)
}

func (m *defaultUser) AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*AddAdminResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddAdmin(ctx, in, opts...)
}

func (m *defaultUser) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminLoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AdminLogin(ctx, in, opts...)
}

func (m *defaultUser) RemoveAdmin(ctx context.Context, in *RemoveAdminRequest, opts ...grpc.CallOption) (*RemoveAdminResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RemoveAdmin(ctx, in, opts...)
}

func (m *defaultUser) AddBlackList(ctx context.Context, in *AddBlackListRequest, opts ...grpc.CallOption) (*AddBlackListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddBlackList(ctx, in, opts...)
}

func (m *defaultUser) QueryBlackList(ctx context.Context, in *QueryBlackListRequest, opts ...grpc.CallOption) (*QueryBlackListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.QueryBlackList(ctx, in, opts...)
}

func (m *defaultUser) RemoveBlackList(ctx context.Context, in *RemoveBlackListRequest, opts ...grpc.CallOption) (*RemoveBlackListResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RemoveBlackList(ctx, in, opts...)
}

func (m *defaultUser) GetHelpCategories(ctx context.Context, in *GetHelpCategoriesRequest, opts ...grpc.CallOption) (*GetHelpCategoriesResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetHelpCategories(ctx, in, opts...)
}

func (m *defaultUser) CreateHelpCategory(ctx context.Context, in *CreateHelpCategoryRequest, opts ...grpc.CallOption) (*CreateHelpCategoryResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateHelpCategory(ctx, in, opts...)
}

func (m *defaultUser) DeleteHelpCategory(ctx context.Context, in *DeleteHelpCategoryRequest, opts ...grpc.CallOption) (*DeleteHelpCategoryResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteHelpCategory(ctx, in, opts...)
}

func (m *defaultUser) EditHelpCategory(ctx context.Context, in *EditHelpCategoryRequest, opts ...grpc.CallOption) (*EditHelpCategoryResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EditHelpCategory(ctx, in, opts...)
}

func (m *defaultUser) CreateHelpCategoryTranslation(ctx context.Context, in *CreateHelpCategoryTranslationRequest, opts ...grpc.CallOption) (*CreateHelpCategoryTranslationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateHelpCategoryTranslation(ctx, in, opts...)
}

func (m *defaultUser) DeleteHelpCategoryTranslation(ctx context.Context, in *DeleteHelpCategoryTranslationRequest, opts ...grpc.CallOption) (*DeleteHelpCategoryTranslationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteHelpCategoryTranslation(ctx, in, opts...)
}

func (m *defaultUser) GetHelpCategoryTranslations(ctx context.Context, in *GetHelpCategoryTranslationsRequest, opts ...grpc.CallOption) (*GetHelpCategoryTranslationsResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetHelpCategoryTranslations(ctx, in, opts...)
}

func (m *defaultUser) EditHelpCategoryTranslation(ctx context.Context, in *EditHelpCategoryTranslationRequest, opts ...grpc.CallOption) (*EditHelpCategoryTranslationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EditHelpCategoryTranslation(ctx, in, opts...)
}

func (m *defaultUser) QueryHelpCategory(ctx context.Context, in *QueryHelpCategoryRequest, opts ...grpc.CallOption) (*QueryHelpCategoryResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.QueryHelpCategory(ctx, in, opts...)
}

func (m *defaultUser) GetHelpDocuments(ctx context.Context, in *GetHelpDocumentsRequest, opts ...grpc.CallOption) (*GetHelpDocumentsResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetHelpDocuments(ctx, in, opts...)
}

func (m *defaultUser) CreateHelpDocument(ctx context.Context, in *CreateHelpDocumentRequest, opts ...grpc.CallOption) (*CreateHelpDocumentResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateHelpDocument(ctx, in, opts...)
}

func (m *defaultUser) DeleteHelpDocument(ctx context.Context, in *DeleteHelpDocumentRequest, opts ...grpc.CallOption) (*DeleteHelpDocumentResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteHelpDocument(ctx, in, opts...)
}

func (m *defaultUser) EditHelpDocument(ctx context.Context, in *EditHelpDocumentRequest, opts ...grpc.CallOption) (*EditHelpDocumentResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EditHelpDocument(ctx, in, opts...)
}

func (m *defaultUser) CreateHelpDocumentTranslation(ctx context.Context, in *CreateHelpDocumentTranslationRequest, opts ...grpc.CallOption) (*CreateHelpDocumentTranslationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateHelpDocumentTranslation(ctx, in, opts...)
}

func (m *defaultUser) DeleteHelpDocumentTranslation(ctx context.Context, in *DeleteHelpDocumentTranslationRequest, opts ...grpc.CallOption) (*DeleteHelpDocumentTranslationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteHelpDocumentTranslation(ctx, in, opts...)
}

func (m *defaultUser) GetHelpDocumentTranslations(ctx context.Context, in *GetHelpDocumentTranslationsRequest, opts ...grpc.CallOption) (*GetHelpDocumentTranslationsResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetHelpDocumentTranslations(ctx, in, opts...)
}

func (m *defaultUser) EditHelpDocumentTranslation(ctx context.Context, in *EditHelpDocumentTranslationRequest, opts ...grpc.CallOption) (*EditHelpDocumentTranslationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EditHelpDocumentTranslation(ctx, in, opts...)
}

func (m *defaultUser) QueryHelpDocument(ctx context.Context, in *QueryHelpDocumentRequest, opts ...grpc.CallOption) (*QueryHelpDocumentResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.QueryHelpDocument(ctx, in, opts...)
}

func (m *defaultUser) CreateSystemNotification(ctx context.Context, in *CreateSystemNotificationRequest, opts ...grpc.CallOption) (*CreateSystemNotificationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateSystemNotification(ctx, in, opts...)
}

func (m *defaultUser) EditSystemNotification(ctx context.Context, in *EditSystemNotificationRequest, opts ...grpc.CallOption) (*EditSystemNotificationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EditSystemNotification(ctx, in, opts...)
}

func (m *defaultUser) GetSystemNotifications(ctx context.Context, in *GetSystemNotificationsRequest, opts ...grpc.CallOption) (*GetSystemNotificationsResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetSystemNotifications(ctx, in, opts...)
}

func (m *defaultUser) QuerySystemNotification(ctx context.Context, in *QuerySystemNotificationRequest, opts ...grpc.CallOption) (*QuerySystemNotificationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.QuerySystemNotification(ctx, in, opts...)
}

// recordNotice 通知记录
func (m *defaultUser) CreateNotice(ctx context.Context, in *CreateNoticeRequest, opts ...grpc.CallOption) (*CreateNoticeResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateNotice(ctx, in, opts...)
}

func (m *defaultUser) RecordNotice(ctx context.Context, in *RecordNoticeRequest, opts ...grpc.CallOption) (*RecordNoticeResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.RecordNotice(ctx, in, opts...)
}

func (m *defaultUser) QueryRecordNotice(ctx context.Context, in *QueryRecordNoticeRequest, opts ...grpc.CallOption) (*QueryRecordNoticeResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.QueryRecordNotice(ctx, in, opts...)
}

// 新增用户消息通知
func (m *defaultUser) CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.CreateNotification(ctx, in, opts...)
}

// 获取用户消息通知
func (m *defaultUser) GetUserNotifications(ctx context.Context, in *GetUserNotificationsRequest, opts ...grpc.CallOption) (*GetUserNotificationsResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserNotifications(ctx, in, opts...)
}
