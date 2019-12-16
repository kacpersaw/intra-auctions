package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/kacpersaw/intra-auctions/ldap"
	"gitlab.com/kacpersaw/intra-auctions/model"
	"gitlab.com/kacpersaw/intra-auctions/router"
	"net/http"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server [port]",
	Short: "Start wsb_projekt server",
	Long:  "",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		model.DB = model.InitDB()
		ldap.LDAP = ldap.InitLDAP()

		logrus.Info("Starting API on port :" + args[0])
		r := router.NewRouter()
		logrus.Fatal(http.ListenAndServe(
			":"+args[0],
			router.CommonMiddleware(
				router.CorsMiddleware(r),
			),
		))
	},
}
