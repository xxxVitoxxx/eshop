package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xxxVitoxxx/eshop/pkg/config"
	"github.com/xxxVitoxxx/eshop/pkg/http/rest"
	"github.com/xxxVitoxxx/eshop/pkg/logger"
	"github.com/xxxVitoxxx/eshop/pkg/numberplate"
	"github.com/xxxVitoxxx/eshop/storage/conn"
	"github.com/xxxVitoxxx/eshop/storage/mysqldb"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(numberPlateCmd)
}

var numberPlateCmd = &cobra.Command{
	Use:   "number-plate",
	Short: "number-plate",
	Long:  "number-plate",
	Run: func(cmd *cobra.Command, args []string) {
		zap.ReplaceGlobals(logger.NewRolling("./log/number_plate.log"))
		zap.L().Info("number plate service start")
		defer func() {
			if err := recover(); err != nil {
				zap.S().Fatal(err)
			}
		}()

		if err := config.Load(); err != nil {
			panic(err)
		}

		r := gin.Default()
		r.Use(gzip.Gzip(gzip.BestSpeed))

		db := conn.ConnectToMySQL()
		npHandler := rest.NewNumberPlateHandler(
			numberplate.NewService(
				mysqldb.NewConditionRepo(db),
			),
		)
		npHandler.Route(r)

		srv := &http.Server{
			Addr:    fmt.Sprintf(":%s", viper.GetString("number_plate.port")),
			Handler: r,
		}
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				zap.S().Fatal("listen: ", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		zap.S().Info("shutdown number plate service")

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			zap.S().Fatal("number plate service shutdown: ", err)
		}

		zap.S().Info("number plate service exiting")
	},
}
