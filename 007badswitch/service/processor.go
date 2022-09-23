package service

import "github.com/kaysun/code_taste/007badswitch/entity"

// Processor 处理器接口
type Processor interface {
	// Process 处理
	Process(operationType entity.OperationType)
}
