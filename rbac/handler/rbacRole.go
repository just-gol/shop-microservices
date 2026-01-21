package handler

import (
	"context"
	"rbac/models"
	pb "rbac/proto/rbacRole"
	"rbac/utils"
)

type RbacRole struct{}

func NewRole() *RbacRole {
	return &RbacRole{}
}
func (r *RbacRole) RoleGet(ctx context.Context, req *pb.RoleGetRequest, resp *pb.RoleGetResponse) error {
	var roleList []models.Role
	models.DB.Where("id = ?", req.Id).Find(&roleList)
	var tempList []*pb.RoleModel
	for i := 0; i < len(roleList); i++ {
		tempList = append(tempList, &pb.RoleModel{
			Id:          int64(roleList[i].Id),
			Title:       roleList[i].Title,
			Description: roleList[i].Description,
			Status:      int64(roleList[i].Status),
			AddTime:     int64(roleList[i].AddTime),
		})
	}
	resp.RoleList = tempList
	return nil
}
func (r *RbacRole) RoleAdd(ctx context.Context, req *pb.RoleAddRequest, resp *pb.RoleAddResponse) error {
	role := models.Role{
		Title:       req.Title,
		Description: req.Description,
		Status:      1,
		AddTime:     int(utils.GetUnix()),
	}
	err := models.DB.Create(&role).Error
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
	} else {
		resp.Success = true
		resp.Message = "ok"
	}
	return nil
}
func (r *RbacRole) RoleEdit(ctx context.Context, req *pb.RoleEditRequest, resp *pb.RoleEditResponse) error {
	return nil
}
func (r *RbacRole) RoleDelete(ctx context.Context, req *pb.RoleDeleteRequest, resp *pb.RoleDeleteResponse) error {
	return nil
}
