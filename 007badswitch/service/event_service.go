package service

import (
	"github.com/kaysun/code_taste/007badswitch/entity"
)

// EventService 事件服务，实现Processor 处理器接口
type EventService struct {
	entity.OperationType
}

// Process 处理
func (l EventService) Process() {
	switch l.OperationType {
	case entity.OperationTypeAdd:
	case entity.OperationTypeUpdate:
	case entity.OperationTypeDelete:
	// TODO case entity.OperationTypeClear:
	default:
	}
}
