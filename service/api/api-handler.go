package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	rt.router.POST("/session", rt.wrap(rt.postSession))
	rt.router.GET("/user", rt.wrap(rt.getSearchUser))
	rt.router.PUT("/user/:idUser", rt.wrap(rt.putUsername))
	rt.router.GET("/user/:idUser", rt.wrap(rt.getUser))
	rt.router.GET("/user/:idUser/following/", rt.wrap(rt.getStream))
	rt.router.PUT("/user/:idUser/following/:followedUserId", rt.wrap(rt.putFollow))
	rt.router.DELETE("/user/:idUser/following/:followedUserId", rt.wrap(rt.deleteFollow))
	rt.router.GET("/user/:idUser/banned/", rt.wrap(rt.getBan))
	rt.router.PUT("/user/:idUser/banned/:bannedUserId", rt.wrap(rt.putBan))
	rt.router.DELETE("/user/:idUser/banned/:bannedUserId", rt.wrap(rt.deleteBan))
	rt.router.POST("/user/:idUser/images", rt.wrap(rt.postImage))
	rt.router.GET("/user/:idUser/images", rt.wrap(rt.getAllImage))
	rt.router.DELETE("/user/:idUser/images/:imageId", rt.wrap(rt.deleteImage))
	rt.router.GET("/user/:idUser/images/:imageId", rt.wrap(rt.getImage))
	rt.router.PUT("/user/:idUser/images/:imageId/like/:likeUserId", rt.wrap(rt.putLike))
	rt.router.DELETE("/user/:idUser/images/:imageId/like/:likeUserId", rt.wrap(rt.deleteLike))
	rt.router.GET("/user/:idUser/images/:imageId/like/", rt.wrap(rt.getLike))
	rt.router.POST("/user/:idUser/images/:imageId/comments/", rt.wrap(rt.postComments))
	rt.router.DELETE("/user/:idUser/images/:imageId/comments/:commentId", rt.wrap(rt.deleteComments))
	rt.router.GET("/user/:idUser/images/:imageId/comments/", rt.wrap(rt.getComments))

	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
