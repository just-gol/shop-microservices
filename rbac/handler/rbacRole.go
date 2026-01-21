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
	db := models.DB
	if req.Id > 0 {
		db = db.Where("id = ?", req.Id)
	}
	db.Find(&roleList)
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

// RoleEdit 修改角色
func (r *RbacRole) RoleEdit(ctx context.Context, req *pb.RoleEditRequest, res *pb.RoleEditResponse) error {
	role := models.Role{Id: int(req.Id)}
	models.DB.Find(&role)
	role.Title = req.Title
	role.Description = req.Description

	err := models.DB.Save(&role).Error
	if err != nil {
		res.Success = false
		res.Message = "修改数据失败"
	} else {
		res.Success = true
		res.Message = "修改数据成功"
	}

	return nil
}

// RoleDelete 删除角色
func (r *RbacRole) RoleDelete(ctx context.Context, req *pb.RoleDeleteRequest, res *pb.RoleDeleteResponse) error {
	role := models.Role{Id: int(req.Id)}
	err := models.DB.Delete(&role).Error
	if err != nil {
		res.Success = false
		res.Message = "删除数据失败"
	} else {
		res.Success = true
		res.Message = "删除数据成功"
	}
	return nil
}
