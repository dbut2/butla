package auth

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func authContext(ctx context.Context, audience string) (context.Context, error) {
	ts, err := idtoken.NewTokenSource(ctx, audience)
	if err != nil {
		return nil, err
	}
	token, err := ts.Token()
	if err != nil {
		return nil, err
	}
	return metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token.AccessToken), nil
}

func AuthContext(ctx context.Context, audience string) context.Context {
	ctx, err := authContext(ctx, audience)
	if err != nil {
		fmt.Println("Big oopsies: ", err.Error())
		panic(err.Error())
	}
	return ctx
}

func Interceptor(audience string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, err := authContext(ctx, audience)
		if err != nil {
			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
