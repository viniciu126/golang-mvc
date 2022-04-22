package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Routes{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		Auth:     false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.FindAllUsers,
		Auth:     true,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodGet,
		Function: controllers.FindOneUser,
		Auth:     true,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateOneUser,
		Auth:     true,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DestroyUser,
		Auth:     true,
	},
	{
		URI:      "/users/{id}/follow",
		Method:   http.MethodPost,
		Function: controllers.Follow,
		Auth:     true,
	},
	{
		URI:      "/users/{id}/unfollow",
		Method:   http.MethodPost,
		Function: controllers.Unfollow,
		Auth:     true,
	},
	{
		URI:      "/users/{id}/followers",
		Method:   http.MethodGet,
		Function: controllers.SearchFollowers,
		Auth:     true,
	},
	{
		URI:      "/users/{id}/following",
		Method:   http.MethodGet,
		Function: controllers.SearchFollowing,
		Auth:     true,
	},
	{
		URI:      "/users/{id}/update-password",
		Method:   http.MethodPost,
		Function: controllers.UpdatePassword,
		Auth:     true,
	},
}
