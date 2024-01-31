package rpc

type Service interface {
	Add(args *AddRequest, reply *AddResponse) error
}

type AddRequest struct {
	X int
	Y int
}

// 响应结构体
type AddResponse struct {
	result int
}

// 实现服务接口的具体服务
type AddServiceImpl struct{}

// 实现服务方法
func (s *AddServiceImpl) Add(args *AddRequest, reply *AddResponse) error {
	reply.result = args.X + args.Y
	return nil
}
