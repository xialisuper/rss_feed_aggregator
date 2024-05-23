package main

import (
	"net/http"
	"strconv"

	"github.com/xialisuper/rss_feed_aggregator/internal/database"
)

func (cfg *apiConfig) handleGetPostsByUserID(w http.ResponseWriter, r *http.Request) {

	// get user info  from context
	user := r.Context().Value(userKey).(database.User)

	// get limit and offset from query params

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	// 将字符串转换成int
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		// 错误处理
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		// 错误处理
		offset = 0
	}

	// get post by user id
	posts, err := cfg.DB.GetPostsByUserID(r.Context(), database.GetPostsByUserIDParams{
		UserID: user.ID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// respond with posts
	respondWithJSON(w, http.StatusOK, posts)

}
