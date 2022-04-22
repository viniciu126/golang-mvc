package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Routes{
	{
		URI:      "/posts",
		Method:   http.MethodPost,
		Function: controllers.StorePost,
		Auth:     true,
	},
	{
		URI:      "/posts",
		Method:   http.MethodGet,
		Function: controllers.FindAllPosts,
		Auth:     true,
	},
	{
		URI:      "/posts/{id}",
		Method:   http.MethodGet,
		Function: controllers.FindOnePost,
		Auth:     true,
	},
	{
		URI:      "/posts/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateOnePost,
		Auth:     true,
	},
	{
		URI:      "/posts/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DestroyPost,
		Auth:     true,
	},
	{
		URI:      "/users/{id}/posts",
		Method:   http.MethodGet,
		Function: controllers.FindPostsByUser,
		Auth:     true,
	},
	{
		URI:      "/posts/{id}/like",
		Method:   http.MethodPost,
		Function: controllers.LikePost,
		Auth:     true,
	},
	{
		URI:      "/posts/{id}/unlike",
		Method:   http.MethodPost,
		Function: controllers.UnlikePost,
		Auth:     true,
	},
}
