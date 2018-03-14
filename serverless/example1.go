package main
import "github.com/aws/aws-lambda-go/lambda"

type MessageEvent struct {Msg string}

func MessageHandler(event MessageEvent) (string, error) {
   return "Your message is  " + event.Msg  , nil
}

func main() {
  lambda.Start(MessageHandler)
}

