package handler

import (
	"context"
	"rbac/models"
	pb "rbac/proto/rbacLogin"

	"github.com/sirupsen/logrus"
)

type RbacLogin struct{}

func New() *RbacLogin {
	return &RbacLogin{}
}

func (e *RbacLogin) Login(ctx context.Context, req *pb.LoginReq, resp *pb.LoginResp) error {
	models.Logger.WithFields(logrus.Fields{
		"method": "Login",
		"params": req,
	}).Info("Login")
	var managerList []models.Manager
	err := models.DB.Where("username=? and password=?", req.Username, req.Password).Find(&managerList).Error
	if err != nil || len(managerList) <= 0 {
		resp.IsLogin = false
	} else {
		resp.IsLogin = true
	}

	var tempList []*pb.ManagerModel
	for i := 0; i < len(managerList); i++ {
		tempList = append(tempList, &pb.ManagerModel{
			Id:       int64(managerList[i].Id),
			Username: managerList[i].Username,
			Password: managerList[i].Password,
			Mobile:   managerList[i].Mobile,
			Email:    managerList[i].Email,
			Status:   int64(managerList[i].Status),
			RoleId:   int64(managerList[i].RoleId),
			AddTime:  int64(managerList[i].AddTime),
			IsSuper:  int64(managerList[i].IsSuper),
		})
	}
	if len(managerList) > 0 {
		resp.IsLogin = true
	} else {
		resp.IsLogin = false
	}
	resp.Userlist = tempList
	return nil
}
