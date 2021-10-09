package infrastructure

import (
	"sns-sample/src/interface/controllers/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	c := NewConfig()
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}
	r.cors()
	r.setRouting()

	return r
}

// cors 対応
func (r *Routing) cors() {
	c := NewConfig()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = c.CORS.AllowOrigins
	r.Gin.Use(cors.New(corsConfig))
}

func (r *Routing) setRouting() {

	MeController := product.NewMeController(r.DB)
	TokensController := product.NewTokensController(r.DB)
	TweetsController := product.NewTweetsController(r.DB)
	UsersController := product.NewUsersController(r.DB)
	UserSearchesController := product.NewUserSearchesController(r.DB)

	v1 := r.Gin.Group("v1/product")
	{

		/*
		 * Me
		 */
		v1.GET("/me", func(c *gin.Context) { MeController.Get(c) })

		/*
		 * Tokens
		 */
		v1.POST("/tokens", func(c *gin.Context) { TokensController.Post(c) })
		v1.POST("/tokens/refresh", func(c *gin.Context) { TokensController.PostRefresh(c) })

		/*
		 * Tweets
		 */
		v1.GET("/tweets", func(c *gin.Context) { TweetsController.GetList(c) })
		v1.POST("/tweets", func(c *gin.Context) { TweetsController.Post(c) })

		v1.GET("/tweets/:id", func(c *gin.Context) { TweetsController.Get(c) })
		v1.PATCH("/tweets/:id", func(c *gin.Context) { TweetsController.Patch(c) })
		v1.DELETE("/tweets/:id", func(c *gin.Context) { TweetsController.Delete(c) })

		/*
		 * Users
		 */
		v1.GET("/users", func(c *gin.Context) { UsersController.Get(c) })
		v1.POST("/users", func(c *gin.Context) { UsersController.Post(c) })
		//v1.PATCH("/users", func(c *gin.Context) { UsersController.Patch(c) })

		/*
		 * Users Search
		 */
		v1.GET("/search", func(c *gin.Context) { UserSearchesController.GetList(c) })

		v1.GET("/search/:id", func(c *gin.Context) { UserSearchesController.Get(c) })
	}
}

func (r *Routing) Run(port string) {
	r.Gin.Run(port)
}
