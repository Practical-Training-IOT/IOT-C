package main

import (
	user "IOT-C/kitex_gen/iot/user"
	"context"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// Register implements the UserImpl interface.
func (s *UserImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterRes, err error) {
	// TODO: Your code here...
	return
}
