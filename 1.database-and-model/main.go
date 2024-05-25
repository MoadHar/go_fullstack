package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	chapters1 "myfitness.dev/app/gen"
)

func main() {
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		GetAsString("DB_USER", "postgres"),
		GetAsString("DB_PASSWORD", "postgres_pass"),
		GetAsString("DB_HOST", "localhost"),
		GetAsInt("DB_PORT", 5432),
		GetAsString("DB_NAME", "postgres"),
	)

	// Open database
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}

	// connectivity check
	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping:", err)
	}

	// create the store
	st := chapters1.New(db)
	ctx := context.Background()

	_, err = st.CreateUsers(ctx, chapters1.CreateUsersParams{
		UserName:     "testuser",
		Name:         "test",
		PassWordHash: "hash",
	})

	if err != nil {
		log.Fatalln("Error creating user:", err)
	}

	eid, err := st.CreateExercise(ctx, "Exercise1")
	if err != nil {
		log.Fatalln("Error creating exercise:", err)
	}

	set, err := st.CreateSet(ctx, chapters1.CreateSetParams{
		ExerciseID: eid,
		Weight:     100,
	})
	if err != nil {
		log.Fatalln("[- ERROR ] creating the set:", err)
	}

	set, err = st.UpdateSet(ctx, chapters1.UpdateSetParams{
		ExerciseID: eid,
		SetID:      set.SetID,
		Weight:     3000,
	})
	if err != nil {
		log.Fatalln("[- ERROR ] updating set:", err)
	}

	log.Println("Done!")

	u, err := st.ListUsers(ctx)
	for i, usr := range u {
		fmt.Println(i, " --> ", usr)
	}
}
