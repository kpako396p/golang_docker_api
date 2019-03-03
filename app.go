package main

import (
	connections "Go_Docker/modules"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

type Container struct {
	ID    string
	Image string
}

type KillContainer struct {
	ID string
}

func getContainers(w http.ResponseWriter, r *http.Request) {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		dockerContainer := Container{container.ID, container.Image}
		js, err := json.Marshal(dockerContainer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func killContainer(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var kill KillContainer
	err = json.Unmarshal(body, &kill)
	if err != nil {
		panic(err)
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	cli.ContainerKill(context.Background(), kill.ID, "SIGKILL")
	if err != nil {
		panic(err)
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get", getContainers).Methods("GET")
	r.HandleFunc("/kill", killContainer).Methods("POST")
	http.Handle("/", r)
	log.Println("Started", connections.DetectMode())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
