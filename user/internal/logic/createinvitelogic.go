package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"user/internal/model"
	"user/internal/svc"
	"user/user"
)

type CreateInviteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInviteLogic {
	return &CreateInviteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证邀请关系是否存在
func (l *CreateInviteLogic) checkIfInviteExists(userId, inviteId int64) error {
	isExists, err := l.svcCtx.Invitation.CheckIfInviteExists(l.ctx, inviteId, userId)
	if err != nil {
		return err
	} else if isExists {
		return errors.New("invite already exists")
	}
	return nil
}

// 创建邀请关系
func (l *CreateInviteLogic) createInvitation(userId, inviteId int64) (*model.Invitation, error) {
	// 生成邀请关系ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	invitation := &model.Invitation{
		InvitationId: node.Generate().Int64(), // 邀请关系ID
		InviterId:    inviteId,                // 邀请人ID
		InviteeId:    userId,                  // 被邀请人ID
	}

	_, err = l.svcCtx.Invitation.Insert(l.ctx, invitation)
	if err != nil {
		return nil, err
	}

	return invitation, nil
}

// 维护邀请关系树
func (l *CreateInviteLogic) maintainInvitationTree(parentId int64) error {
	parentTree, err := l.svcCtx.InvitationTree.FindParentID(l.ctx, parentId)
	if err != nil {
		fmt.Println(123)
		fmt.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}

	invitationTree := &model.InvitationTree{
		InvitationTreeId: parentTree.InvitationTreeId, // 邀请关系树ID
		UserId:           parentTree.UserId,           // 用户ID
		ParentId:         parentTree.ParentId,         // 父级ID
		Depth:            parentTree.Depth + 1,        // 更新深度
	}

	_, err = l.svcCtx.InvitationTree.Insert(l.ctx, invitationTree)
	if err != nil {
		return err
	}

	l.svcCtx.InvitationTree.Update(l.ctx, invitationTree)

	return nil
}

// 创建邀请关系树
func (l *CreateInviteLogic) createInvitationTree(userId, inviteId int64) error {

	// 维护之前的邀请关系树
	if inviteId != 0 {
		err := l.maintainInvitationTree(inviteId)
		if err != nil {
			return err
		}
	}
	// 生成邀请关系树ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}

	invitationTree := &model.InvitationTree{
		InvitationTreeId: node.Generate().Int64(), // 邀请关系树ID
		ParentId:         userId,                  // 父级ID
		UserId:           inviteId,                // 用户ID
		Depth:            0,                       // 深度
	}

	_, err = l.svcCtx.InvitationTree.Insert(l.ctx, invitationTree)
	if err != nil {
		return err
	}

	return nil
}

func (l *CreateInviteLogic) CreateInvite(in *user.CreateInviteRequest) (*user.CreateInviteResponse, error) {
	// todo: 不是很安全，需要改进 需要加事务   5.27前
	// todo: 加入根据推特接口来查找用户ID
	// todo: 获取用户信息来

	// 将 in.UserId 转为 int64 类型
	userId, err := strconv.ParseInt(in.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	// 将 in.UserId 转为 int64 类型
	inviteId, err := strconv.ParseInt(in.InviteId, 10, 64)
	if err != nil {
		return nil, err
	}

	// 验证邀请人是否存在
	if err := l.checkIfInviteExists(userId, inviteId); err != nil {
		return nil, err
	}

	// 创建邀请关系
	invitation, err := l.createInvitation(userId, inviteId)
	if err != nil {
		return nil, err
	}

	// 创建邀请关系树
	err = l.createInvitationTree(userId, inviteId)
	if err != nil {
		return nil, err
	}

	return &user.CreateInviteResponse{
		InvitationId: strconv.FormatInt(invitation.InvitationId, 10),
	}, nil
}
