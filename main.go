package main

import (
	"context"
	user_center "web-IDE-back-end/proto/user-center"
	"web-IDE-back-end/rpc"
)

func main() {
	ctx := context.Background()
	rpc.Register(ctx, &user_center.RegisterRequest{
		
	})
}
