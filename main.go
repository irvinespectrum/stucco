// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	//	"google.golang.org/appengine"
	//	"google.golang.org/appengine/datastore"
	//	"google.golang.org/appengine/user"
)

type ID struct {
	Name    string
	Seed    int
	Counter uint64
}

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}
func handle(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	names := strings.Split(r.FormValue("name"), ",")
	for i := 0; i < len(names); i++ {
		name := names[i]
		seed := rand.Intn(100)
		result := record(ctx, name, seed)
		fmt.Fprintf(w, "Hello, %d", result)
	}

	/*
		name := "store"

		e1 := ID{
			Name:    name,
			Seed:    seed,
			Counter: 01,
		}
		fmt.Fprintf(w, "Hello, %d", e1.Seed)

		/*
			ctx: = appengine.NewContext(r)

			key, err: = datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "employee", nil),  & e1)
			if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}

			var e2 Employee
			if err = datastore.Get(ctx, key,  & e2); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
			}
			fmt.Fprintf(w, "Stored and retrieved the Employee named %q", e2.Name) */
	//	fmt.Fprintln(w, "Hello, world! from salesamount5.com")
}

func record(ctx context.Context, name string, seed int) uint64 {

	q := datastore.NewQuery("ID").
		Filter("Name=", name).
		Filter("Seed=", seed)

	var ids []ID
	if _, err := q.GetAll(ctx, &ids); err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var counter uint64 = 0
	if len(ids) > 0 {
		id := ids[0]
		counter = id.Counter + 1
	} else {

		counter = 1
	}
	return counter*100 + uint64(seed)
}
