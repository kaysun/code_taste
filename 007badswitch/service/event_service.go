package service

import (
	"github.com/kaysun/code_taste/007badswitch/entity"
)

// EventService 事件服务，实现Processor 处理器接口
type EventService struct {
}

// Process 处理
func (EventService) Process(operationType entity.OperationType) {
	switch operationType {
	case entity.OperationTypeAdd:
	case entity.OperationTypeUpdate:
	case entity.OperationTypeDelete:
	// TODO case entity.OperationTypeClear:
	default:
	}
}
