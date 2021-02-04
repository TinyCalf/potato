package rpc

// Server RPC Server
type Server struct {
	opts *Options
	dispatcher *dispatcher
	acceptor iacceptor
}

// NewServer 创建一个新的server
func NewServer(opts *Options) *Server {
	s := &Server{}
	s.opts = opts
	s.dispatcher = newDispatcher(opts.Services)
	s.acceptor = newTCPAcceptor()
	s.acceptor.onData(func(data []byte) ([]byte, error){
		msg, err := unpack(data)
		if err != nil {
			return nil, err
		}
		return s.dispatcher.route(msg)
	})
	return s
}

// Start 启动RPC Server
func (s *Server) Start() {
	s.acceptor.listen(s.opts.Port)
}