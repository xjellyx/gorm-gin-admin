package menu

type NodeMenu interface {
	GetName() string
	GetID()   uint
	GetParentID() uint
	GetData() interface{}
	IsRoot() bool
}
