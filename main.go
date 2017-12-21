package main

import (

  "encoding/json" 
  "net/http"
  "time"
  "github.com/gorilla/mux"

  "roles-ws/role"
  "roles-ws/config"
)

var ( SystemConfig *config.Configuration )

func main() {

  SystemConfig = config.ReadFile("etc/system.conf")

  main_router := mux.NewRouter()
  api_router := main_router.PathPrefix("/api/v1").Subrouter()

  // Role endpoints
  //api_router.Methods("POST").Path("/roles").HandlerFunc(ListRoles)

  // Role endpoints
  api_router.Methods("GET").Path("/roles").HandlerFunc(ListRoles)
  api_router.Methods("GET").Path("/roles/{role_id}/params/{version}").HandlerFunc(GetRoleParams)
  api_router.Methods("GET").Path("/roles/{role_id}/meta/{version}").HandlerFunc(GetRoleMeta)

  server := &http.Server{
          Handler:      main_router,
          Addr:         "0.0.0.0:8000",
          // Good practice: enforce timeouts for servers you create!
          WriteTimeout: 15 * time.Second,
          ReadTimeout:  15 * time.Second,
      }

  role.Init(SystemConfig)

  SystemConfig.Log.Fatal(server.ListenAndServe())
}


func ListRoles(w http.ResponseWriter, r *http.Request) {
  projects, err := role.List(SystemConfig, SystemConfig.GitlabState.GroupID)

  if err != nil {
    SystemConfig.Log.Fatal(err)
  } else {
    //fmt.Fprintln(w, "showing url", projects)
    json.NewEncoder(w).Encode(projects)
  }
}

func GetRoleMeta(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)

  params, _ := role.GetMeta(SystemConfig, vars["role_id"], vars["version"])
  w.Write(params)
}

func GetRoleParams(w http.ResponseWriter, r *http.Request) {

  vars := mux.Vars(r)

  params, _ := role.GetParams(SystemConfig, vars["role_id"], vars["version"])
  w.Write(params)
  //params, err := role.GetParams(SystemConfig, vars["role_id"], vars["version"])
}

