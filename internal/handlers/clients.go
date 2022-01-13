package handlers

import (
	"net/http"

	"usulroster/internal/usulclient"

	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ClientCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r)
		desc := r.FormValue("description")
		secretKey := r.FormValue("secretKey")
		if secretKey == viper.GetString("clientsecretkey") {
			log.Infof("Attempting Client Create; description: %v", desc)
			t, err := usulclient.CreateNewClient(desc)
			if err != nil {
				// TODO: 401
				render.JSON(w, r, err)
				return
			}
			render.JSON(w, r, t)
		}
	}
}

func ClientList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r)
		desc := r.FormValue("description")
		secretKey := r.FormValue("secretKey")
		if secretKey == viper.GetString("clientsecretkey") {
			log.Infof("Attempting Client List; description: %v", desc)
			t, err := usulclient.ListClients(desc)
			if err != nil {
				// TODO: 401
				render.JSON(w, r, err)
				return
			}
			render.JSON(w, r, t)
		}
	}
}

func RemoveClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r)
		clientid := r.FormValue("clientid")
		secretKey := r.FormValue("secretKey")
		if secretKey == viper.GetString("clientsecretkey") {
			log.Infof("Attempting Remove Client ; client id: %v", clientid)
			t, err := usulclient.RemoveClient(clientid)
			if err != nil {
				// TODO: 401
				render.JSON(w, r, err)
				return
			}
			render.JSON(w, r, t)
		}
	}
}
