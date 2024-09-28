package main

import (
	"context"
	"et_test/proto/et_test"
	"et_test/store"
	"fmt"
	"io"
	"log"
	"net/http"
)

//go:generate mockery --name ServerInterface
type ServerInterface interface {
	GetThumbnail(ctx context.Context, req *et_test.GetThumbnailRequest) (*et_test.GetThumbnailResponse, error)
}

type Server struct {
	et_test.UnimplementedThumbnailServiceServer
	store *store.Store
}

func NewServer(store store.Store) *Server {
	return &Server{store: &store}
}

func (s *Server) GetThumbnail(ctx context.Context, req *et_test.GetThumbnailRequest) (*et_test.GetThumbnailResponse, error) {
	responseChan := make(chan *et_test.GetThumbnailResponse)
	errChan := make(chan error)

	go func() { //для асинхронности
		storeResp, storeErr := s.store.CheckThumbnailInStore(req.Id)
		if storeErr == nil {
			responseChan <- &et_test.GetThumbnailResponse{Image: storeResp}
		}

		resp, err := http.Get(fmt.Sprintf("https://img.youtube.com/vi/%s/maxresdefault.jpg", req.Id))
		if err != nil {
			errChan <- fmt.Errorf("failed to download thumbnail: %v", err)
		}
		defer resp.Body.Close()

		imageData, err := io.ReadAll(resp.Body)
		if err != nil {
			errChan <- fmt.Errorf("failed to read thumbnail data: %v", err)
		}
		storeErr = s.store.SaveThumbnail(req.Id, imageData)
		if storeErr != nil {
			log.Print("failed to save thumbnail to store " + storeErr.Error())
		}

		responseChan <- &et_test.GetThumbnailResponse{
			Image: imageData,
		}
	}()

	select {
	case response := <-responseChan:
		return response, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
