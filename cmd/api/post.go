package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/stefanaki/shelter/internal/db"
	"github.com/stefanaki/shelter/internal/dto"
	"github.com/stefanaki/shelter/internal/store"
)

const PostContextKey = "post"

func (a *api) createPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload dto.CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		a.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		a.badRequest(w, r, err)
		return
	}

	post, err := a.store.Posts.Create(ctx, payload)
	if err != nil {
		a.internalServerError(w, r, err)
		return
	}

	writeJSON(w, http.StatusCreated, dto.Post{
		ID:        hex.EncodeToString(post.ID.Bytes[:]),
		Title:     post.Title,
		Content:   post.Content,
		Tags:      post.Tags,
		UserID:    hex.EncodeToString(post.UserID.Bytes[:]),
		CreatedAt: post.CreatedAt.Time,
		UpdatedAt: post.UpdatedAt.Time,
	})
}

func (a *api) retrievePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r)

	writeJSON(w, http.StatusOK, dto.Post{
		ID:        hex.EncodeToString(post.ID.Bytes[:]),
		Title:     post.Title,
		Content:   post.Content,
		Tags:      post.Tags,
		UserID:    hex.EncodeToString(post.UserID.Bytes[:]),
		CreatedAt: post.CreatedAt.Time,
		UpdatedAt: post.UpdatedAt.Time,
	})
}

func (a *api) listPostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	ctx := r.Context()

	comments, err := a.store.Comments.ListByPostID(ctx, postID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrorNotFound):
			a.notFound(w, r, err)
		default:
			a.internalServerError(w, r, err)
		}
		return
	}

	var res []dto.Comment
	for _, c := range *comments {
		res = append(res, dto.Comment{
			ID:        store.PostgresUUIDColumnToString(c.ID),
			PostID:    store.PostgresUUIDColumnToString(c.PostID),
			UserID:    store.PostgresUUIDColumnToString(c.UserID),
			Content:   c.Content,
			CreatedAt: c.CreatedAt.Time,
		})
	}

	writeJSON(w, 200, res)
}

func (a *api) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r)

	err := a.store.Posts.Delete(r.Context(), store.PostgresUUIDColumnToString(post.ID))
	if err != nil {
		switch {
		case errors.Is(err, store.ErrorInvalidInput):
			a.badRequest(w, r, err)
		default:
			a.internalServerError(w, r, err)
		}
		return
	}

	writeJSON(w, http.StatusOK, "OK")
}

func (a *api) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r)

	var payload dto.UpdatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		a.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		a.badRequest(w, r, err)
		return
	}

	updatedPost, err := a.store.Posts.Update(r.Context(), store.PostgresUUIDColumnToString(post.ID), payload)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			a.notFound(w, r, err)
		default:
			a.internalServerError(w, r, err)
		}
		return
	}

	writeJSON(w, http.StatusOK, updatedPost)
}

func (a *api) withPostContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx := r.Context()
		post, err := a.store.Posts.Retrieve(ctx, id)

		if err != nil {
			switch {
			case errors.Is(err, store.ErrorNotFound):
				a.notFound(w, r, err)
			case errors.Is(err, store.ErrorInvalidInput):
				a.badRequest(w, r, err)
			default:
				a.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, PostContextKey, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPostFromContext(r *http.Request) *db.Post {
	post, _ := r.Context().Value(PostContextKey).(*db.Post)

	return post
}
