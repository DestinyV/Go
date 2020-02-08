package main

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"lottery/middleware"
	"lottery/models"
	"lottery/oa"
	"lottery/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessionMiddleware())
	router.Use(middleware.AuthHandler("/lottery/auth", "/lottery/login"))
	routers.InitRouters(router)

	server := http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 60,
		WriteTimeout:      time.Second * 60,
		IdleTimeout:       time.Second * 60,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("启动服务器失败！%v", err)
		}
	}()

	shutdownGracefully(&server)
}

func sessionMiddleware() gin.HandlerFunc {
	store := cookie.NewStore([]byte("WindOSX"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "vankeytech.com",
		MaxAge:   604800,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	return sessions.Sessions("X-Auth-Token", store)
}

func shutdownGracefully(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch
	log.Printf("Server shutting down by signal: %v", sig)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
	closeDB()
	log.Println("Server exit")
}

func closeDB() {
	var err error
	err = models.DB.Close()
	if err != nil {
		log.Fatalf("关闭数据库失败")
	}
	err = oa.DB.Close()
	if err != nil {
		log.Fatalf("关闭数据库失败")
	}
}
