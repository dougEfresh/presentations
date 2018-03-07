func main() {
  lambda.Start(HelloWorld)
}

func Start(handler interface{}) {
	port := os.Getenv("_LAMBDA_SERVER_PORT")
	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	wrappedHandler := newHandler(handler)
	function := new(Function)
	function.handler = wrappedHandler
	err = rpc.Register(function)
	if err != nil {
		log.Fatal("failed to register handler function")
	}
	rpc.Accept(lis)
	log.Fatal("accept should not have returned")
}

