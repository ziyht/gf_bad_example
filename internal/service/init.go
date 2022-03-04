package service

import (
	"context"
	"test/internal/service/internal/dao"
)

func Init(){
	dao.Alert.Init(context.Background())
}