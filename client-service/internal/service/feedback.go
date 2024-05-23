package service

import (
	"client-service/internal/entities"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"

	"client-service/internal/repository"
)

type FeedbackService struct {
	upgrader     websocket.Upgrader
	conns        map[*websocket.Conn]string
	mutex        sync.Mutex
	feedbackRepo *repository.FeedbackRepository
}

func NewFeedbackService(feedbackRepo *repository.FeedbackRepository) *FeedbackService {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conns := make(map[*websocket.Conn]string, 8)

	return &FeedbackService{upgrader: upgrader, conns: conns, mutex: sync.Mutex{}, feedbackRepo: feedbackRepo}
}

func (s *FeedbackService) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	s.mutex.Lock()
	s.conns[conn] = "connected"
	s.mutex.Unlock()

	feedbackItems, err := s.feedbackRepo.GetAllFeedbackItems(r.Context())
	if err != nil {
		conn.Close()

		s.mutex.Lock()
		delete(s.conns, conn)
		s.mutex.Unlock()

		return
	}

	for _, v := range feedbackItems {
		if err := conn.WriteJSON(v); err != nil {
			conn.Close()

			s.mutex.Lock()
			delete(s.conns, conn)
			s.mutex.Unlock()

			return
		}
	}

	for {
		var a entities.FeedbackItem
		err := conn.ReadJSON(&a)
		if err != nil {
			conn.Close()

			s.mutex.Lock()
			delete(s.conns, conn)
			s.mutex.Unlock()

			return
		}
		fmt.Println("RECEINVED")

		f, err := s.feedbackRepo.CreateFeedbackItem(r.Context(), a.Email, a.Text)
		if err != nil {
			conn.Close()

			s.mutex.Lock()
			delete(s.conns, conn)
			s.mutex.Unlock()

			return
		}

		for k, _ := range s.conns {
			if err := k.WriteJSON(f); err != nil {
				conn.Close()

				s.mutex.Lock()
				delete(s.conns, conn)
				s.mutex.Unlock()

				return
			}
		}
	}
}
