package entity

type OperationType uint8

const (
	OperationTypeAdd OperationType = iota
	OperationTypeUpdate
	OperationTypeDelete

	// TODO OperationTypeClear
)
