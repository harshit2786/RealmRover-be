package cmd

import (
	"log"
	"net/http"
	"realmrovers/config"
	"realmrovers/db"
	// "github.com/gorilla/handlers"
	"realmrovers/handler"
	router "realmrovers/route"
	"realmrovers/services"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "http",
	Short: "CLI for HTTP",
	Long: "CLI which starts the server on 8000 port",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		dbc := db.ConnectDb(cfg)
		userservice := &services.UserService{Db: dbc , Cfg: cfg}
		userhandler := &handler.UserHandler{Service: userservice}
		r := router.NewRouter(userhandler)
		log.Printf("Server is running on port %s", cfg.Port)
		log.Fatal(http.ListenAndServe(":"+cfg.Port, r ))
	},
}


func init() {
	rootCmd.AddCommand(serverCmd)
}