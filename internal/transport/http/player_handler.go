package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	playerpb "github.com/danilopucci/opentibiagateway/internal/protogen/v1"
	"github.com/gorilla/mux"
)

type PlayerHandler struct {
	playerClient playerpb.PlayerServiceClient
}

func NewPlayerHandler(client playerpb.PlayerServiceClient) *PlayerHandler {
	return &PlayerHandler{
		playerClient: client,
	}
}

func (h *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid player id", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	resp, err := h.playerClient.GetPlayer(ctx, &playerpb.GetPlayerRequest{
		Identifier: &playerpb.GetPlayerRequest_Id{Id: int32(id)},
	})
	if err != nil {
		http.Error(w, "error getting player: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Player)
}
