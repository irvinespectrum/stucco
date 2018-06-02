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

		rand.Seed(time.Now().UTC().UnixNano())
		seed := rand.Intn(100)
		ctx := appengine.NewContext(r)
		count, err := process(ctx, name, seed)

		if err != nil {
			fmt.Fprintf(w, "error: %q", err.Error())
		} else {
			fmt.Fprintf(w, "%d", count)
		}
	} else {
		fmt.Fprintf(w, "please add id parameter, like ?id=test")
	}
}

func process(ctx context.Context, name string, seed int) (int64, error) {

	var count int64
	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		var err1 error
		count, err1 = increment(ctx, name, seed)
		return err1
	}, nil)
	return count, err
}

func increment(ctx context.Context, name string, seed int) (int64, error) {

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
	return id.Counter*100 + int64(seed), nil
}

/*
	q := datastore.NewQuery("ID").
		Filter("Name=", name).
		Filter("Seed=", seed)

	var ids []ID
	keys, err2 := q.GetAll(ctx, &ids)
	if err2 != nil {
		return err2
	}

	var key *datastore.Key
	if len(ids) > 0 {
		key = keys[0]
	}*/
/*
	if key != nil {
		id.Counter++
		if id.Counter <= 0 {
			id.Counter = 1
		}
		if _, err := datastore.Put(ctx, key, &id); err != nil {
			return 0, err
		}
		counter = id.Counter
	} else {
		id := ID{
			Name:    name,
			Seed:    seed,
			Counter: 1,
		}
		key := datastore.NewIncompleteKey(ctx, "ID", nil)
		if _, err := datastore.Put(ctx, key, &id); err != nil {
			return 0, err
		}
		counter = 1
	}
*/
//return counter*100 + int64(seed), nil
