package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/srnewbie/sample-server/client"
	"github.com/srnewbie/sample-server/config"
)

type (
	Handler struct {
		c         *client.Client
		f         *config.Config
		vl        sync.Mutex
		Questions map[string]*Question
		Votes     map[string][]*Vote
	}
	HTTPResponse struct {
		Data interface{} `json:"data"`
		Err  string      `json:"err"`
	}
	Question struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
		Creator string `json:"creator"`
	}
	Vote struct {
		ID         string `json:"id"`
		QuestionID string `json:"question_id"`
		Voter      string `json:"voter"`
	}
)

func New(f *config.Config, c *client.Client) *Handler {
	return &Handler{
		c:         c,
		f:         f,
		Questions: make(map[string]*Question),
		Votes:     make(map[string][]*Vote),
	}
}

func respond(w http.ResponseWriter, status int, data interface{}, err string) {
	response := &HTTPResponse{
		Data: data,
		Err:  err,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(response)
}

func (h *Handler) Health() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		respond(w, 200, "ok", "")
	}
}

func (h *Handler) CreateQuestions() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respond(w, 500, "", "error")
			return
		}
		question := Question{}
		err = json.Unmarshal(q, &question)
		if err != nil {
			respond(w, 500, "", "error")
			return
		}
		question.ID = uuid.New().String()
		h.Questions[question.ID] = &question
		h.c.Logger.Info("questions:", question)
		respond(w, 200, "ok", "")
	}
}

func (h *Handler) ListQuestions() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		respond(w, 200, h.Questions, "")
	}
}

func (h *Handler) UpVote() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respond(w, 500, "", "error")
			return
		}
		vote := Vote{}
		err = json.Unmarshal(v, &vote)
		if err != nil {
			respond(w, 500, "", "error")
			return
		}
		vote.ID = uuid.New().String()
		q, ok := h.Questions[vote.QuestionID]
		if !ok {
			respond(w, 400, "", "question not found")
			return
		}
		h.vl.Lock()
		defer h.vl.Unlock()
		h.Votes[q.ID] = append(h.Votes[q.ID], &vote)
		respond(w, 200, "ok", "")
	}
}
