package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

//ID datastore object
type ID struct {
	Counter int64
}

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}
func handle(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("id")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if len(name) > 0 {

		ctx := appengine.NewContext(r)

		var count int64
		err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
			var err1 error
			count, err1 = increment(ctx, name)
			return err1
		}, nil)

		if err != nil {
			fmt.Fprintf(w, "error: %q", err.Error())
		} else {
			fmt.Fprintf(w, "%d", count)
		}
	} else {
		fmt.Fprintf(w, "please add id parameter, like ?id=test")
	}
}

func increment(ctx context.Context, name string) (int64, error) {

	digit := 10000
	rand.Seed(time.Now().UTC().UnixNano())
	seed := rand.Intn(digit)

	var id ID
	key := datastore.NewKey(ctx, name, strconv.Itoa(seed), 0, nil)
	if err := datastore.Get(ctx, key, &id); err != nil && err != datastore.ErrNoSuchEntity {
		return 0, err
	}

	id.Counter++
	if id.Counter <= 0 {
		id.Counter = 1
	}

	if _, err := datastore.Put(ctx, key, &id); err != nil {
		return 0, err
	}
	return id.Counter*int64(digit) + int64(seed), nil
}
