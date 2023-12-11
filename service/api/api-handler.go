package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	rt.router.POST("/session", rt.postSession)
	rt.router.PUT("/user/:idUser", rt.putUsername)
	rt.router.GET("/user/:idUser", rt.getUser)
	rt.router.GET("/user/:idUser/following/", rt.getStream)
	rt.router.PUT("/user/:idUser/following/:followedUserId", rt.putFollow)
	rt.router.DELETE("/user/:idUser/following/:followedUserId", rt.deleteFollow)
	rt.router.GET("/user/:idUser/banned/", rt.getBan)
	rt.router.PUT("/user/:idUser/banned/:bannedUserId", rt.putBan)
	rt.router.DELETE("/user/:idUser/banned/:bannedUserId", rt.deleteBan)
	rt.router.POST("/user/:idUser/images", rt.postImage)
	rt.router.DELETE("/user/:idUser/images/:imageId", rt.deleteImage)
	rt.router.GET("/user/:idUser/images/:imageId", rt.getImage)
	rt.router.PUT("/user/:idUser/images/:imageId/like/:likeUserId", rt.putLike)
	rt.router.DELETE("/user/:idUser/images/:imageId/like/:likeUserId", rt.deleteLike)
	rt.router.POST("/user/:idUser/images/:imageId/comments", rt.postComments)
	rt.router.DELETE("/user/:idUser/images/:imageId/comments/:commentId", rt.deleteComments)

	rt.router.GET("/liveness", rt.liveness)
	return rt.router
}
