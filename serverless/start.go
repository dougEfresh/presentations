func Start(handler interface{}) {
	port := os.Getenv("_SERVER_PORT")
	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
    wrappedHandler := newHandler(handler)
    function := new(Function)
    function.handler = new(wrappedHandler)
    rpc.Register(function)
	rpc.Accept(lis)
}

