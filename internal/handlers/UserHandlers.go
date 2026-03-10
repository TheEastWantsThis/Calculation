package handlers

import (
	"context"
	userservice "project/internal/userService"
	"project/internal/web/users"
	"time"
)

type UserHand struct {
	s userservice.UserService
}

func NewUserHandler(ser userservice.UserService) *UserHand {
	return &UserHand{s: ser}
}

func (h *UserHand) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	us := request.Body

	userToCreat := userservice.User{
		Email:    *us.Email,
		Password: *us.Password,
	}

	user, err := h.s.CreateUser(userToCreat)
	if err != nil {
		return nil, err
	}
	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}
	result := users.PostUsers201JSONResponse{
		Id:        &user.ID,
		Email:     &user.Email,
		Password:  &user.Password,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: deletedAt,
	}
	return result, err
}

func (h *UserHand) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.s.GetAllUsers()
	if err != nil {
		return nil, err
	}

	resp := users.GetUsers200JSONResponse{}

	var deletedAt *time.Time

	for _, user := range allUsers {

		if user.DeletedAt.Valid {
			deletedAt = &user.DeletedAt.Time
		}
		urs := users.Users{
			Id:        &user.ID,
			Email:     &user.Email,
			Password:  &user.Password,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: deletedAt,
		}
		resp = append(resp, urs)
	}
	return resp, nil

}

func (h *UserHand) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	oldser := request.Body

	update, err := h.s.UpdateUser(id, userservice.User{
		Email:    *oldser.Email,
		Password: *oldser.Password,
	})
	if err != nil {
		return nil, err
	}
	var deletedAt *time.Time
	if update.DeletedAt.Valid {
		deletedAt = &update.DeletedAt.Time
	}
	result := users.Users{
		Id:        &id,
		Email:     &update.Email,
		Password:  &update.Password,
		CreatedAt: &update.CreatedAt,
		UpdatedAt: update.UpdatedAt,
		DeletedAt: deletedAt,
	}

	return users.PatchUsersId200JSONResponse(result), nil

}

func (h *UserHand) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	if err := h.s.DeleteUser(id); err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}
