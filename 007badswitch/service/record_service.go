package service

import (
	"github.com/kaysun/code_taste/007badswitch/entity"
)

/*
RecordService 记录服务，实现Processor 处理器接口
实际开发中，针对枚举进行switch的代码很可能散在整个项目中的任意地方，当枚举类型变化（比如添加），
要修改非常多的地方，称之为霰弹式修改（Shotgun Surgery）
如果每遇到某种变化，你都必须在许多不同的类内做出许多小修改，你所面临的坏味道就是霰弹式修改。
如果需要修改的代码散布四处，你不但很难找到它们，也很容易错过某个重要的修改。

switch可以考虑使用多态解决，可以借鉴002strategy，
实现“一旦需要修改，我们希望能够跳到系统的某一点，只在该处做修改。”
*/
type RecordService struct {
	entity.OperationType
}

// Process 处理
func (l RecordService) Process() {
	switch l.OperationType {
	case entity.OperationTypeAdd:
	case entity.OperationTypeUpdate:
	case entity.OperationTypeDelete:
	// TODO case entity.OperationTypeClear:
	default:
	}
}
