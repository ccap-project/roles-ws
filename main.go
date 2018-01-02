/*
 *
 * Copyright (c) 2016, 2017, 2018 Alexandre Biancalana <ale@biancalanas.net>.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in the
 *       documentation and/or other materials provided with the distribution.
 *     * Neither the name of the <organization> nor the
 *       names of its contributors may be used to endorse or promote products
 *       derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

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

