package main

import (
	"context"
	"errors"
	"et_test/mocks"
	"et_test/proto/et_test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer_GetThumbnail(t *testing.T) {
	t.Parallel()

	type args struct {
		serverInterface *mocks.ServerInterface
	}

	type testCase struct {
		name    string
		req     *et_test.GetThumbnailRequest
		exp     *et_test.GetThumbnailResponse
		prepare func(args *args)
		err     error
	}

	testCases := []testCase{
		{
			name: "success",
			req: &et_test.GetThumbnailRequest{
				Id: "11",
			},
			prepare: func(args *args) {
				args.serverInterface.On("GetThumbnail", context.Background(), &et_test.GetThumbnailRequest{
					Id: "11",
				}).Return(&et_test.GetThumbnailResponse{
					Image: []byte{1, 2, 3, 4},
				}, nil)
			},
			exp: &et_test.GetThumbnailResponse{
				Image: []byte{1, 2, 3, 4},
			},
			err: nil,
		},
		{
			name: "fail",
			req: &et_test.GetThumbnailRequest{
				Id: "12",
			},
			prepare: func(args *args) {
				args.serverInterface.On("GetThumbnail", context.Background(), &et_test.GetThumbnailRequest{
					Id: "12",
				}).Return(nil, errors.New("GetThumbnailRequest failed"))
			},
			exp: nil,
			err: errors.New("GetThumbnailRequest failed"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			a := assert.New(t)
			client := new(mocks.ServerInterface)
			tc.prepare(&args{serverInterface: client})

			resp, err := client.GetThumbnail(context.Background(), tc.req)
			if tc.err != nil {
				a.EqualError(err, tc.err.Error())
			} else {
				a.Equal(tc.exp, resp)
			}
		})
	}
}
