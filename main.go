package main

import (
	"crypto/tls"
	"fmt"
	_ "github.com/olongfen/gorm-gin-admin/docs"
	_ "github.com/olongfen/gorm-gin-admin/src/models"
	_ "github.com/olongfen/gorm-gin-admin/src/pkg/gredis"
	"github.com/olongfen/gorm-gin-admin/src/router"
	_ "github.com/olongfen/gorm-gin-admin/src/router"
	"github.com/olongfen/gorm-gin-admin/src/setting"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {

	go func() {
		// 开启服务
		s := &http.Server{
			Addr:           setting.Settings.Serve.ServerAddr + ":" + setting.Settings.Serve.ServerPort,
			Handler:        router.Engine,
			ReadTimeout:    60 * time.Second,
			WriteTimeout:   60 * time.Second,
			MaxHeaderBytes: 1 << 20, // 10M
		}
		logrus.Println("server listen on: ", s.Addr)
		if setting.Settings.Serve.IsTLS { // 开启tls
			TLSConfig := &tls.Config{
				MinVersion:               tls.VersionTLS11,
				CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				PreferServerCipherSuites: true,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				},
			}

			TLSProto := make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)

			s.TLSConfig = TLSConfig
			s.TLSNextProto = TLSProto

			if err := s.ListenAndServeTLS(setting.Settings.Serve.TLS.Cert, setting.Settings.Serve.TLS.Key); err != nil {
				logrus.Fatal(err)
			}
		} else {
			if err := s.ListenAndServe(); err != nil {
				logrus.Fatal(err)
			}
		}
	}()

	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
EXIT:
	for {
		sig := <-sc
		fmt.Printf("signal[%s]\n", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	fmt.Println("exit")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}
