package main

import (
	"fmt"
	"os"

	"github.com/DevOps-Ben11/minitwit/backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {

	fmt.Println("Connecting to sqlite...")
	sql, err := gorm.Open(sqlite.Open(os.Args[1]), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("Error connecting to sqlite")
		panic(err)
	}

	err = sql.AutoMigrate(
		&model.User{},
		&model.Follower{},
		&model.Message{},
		&model.KeyVal{},
	)

	if err != nil {
		fmt.Println("Error migrating sqlite")
		panic(err)
	}

	fmt.Println("Successfully connected to sqlite")

	fmt.Println("Connecting to Postgres...")
	conString := os.Getenv("PSQL_CON_STR")
	psql, err := gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to postgres")
		panic(err)
	}

	err = psql.AutoMigrate(
		&model.User{},
		&model.Follower{},
		&model.Message{},
		&model.KeyVal{},
	)

	if err != nil {
		fmt.Println("Error migrating psql")
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")

	fmt.Println("Fetching users...")

	users := []model.User{}

	err = sql.Select("username", "email", "pw_hash").Find(&users).Error

	if err != nil {
		fmt.Println("Error fetching users")
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Fetched %v users", len(users)))

	fmt.Println("Inserting into postgres...")

	//for _, u := range users {
	err = psql.Save(&users).Error

	if err != nil {
		fmt.Println("Error saving users")
		panic(err)
	}
	//}

	fmt.Println("Successfully inserted users")

	fmt.Println("Fetching messages...")

	messages := []model.Message{}

	err = sql.Select("author_id", "text", "pub_date", "flagged").Find(&messages).Error

	if err != nil {
		fmt.Println("Error fetching messages")
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Fetched %v messages", len(messages)))

	fmt.Println("Inserting into postgres...")

	//for _, m := range messages {
	err = psql.Save(&messages).Error

	if err != nil {
		fmt.Println("Error saving messages")
		panic(err)
	}
	// }

	fmt.Println("Successfully inserted messages")

	fmt.Println("Fetching follows...")

	follows := []model.Follower{}

	err = sql.Find(&follows).Error

	if err != nil {
		fmt.Println("Error fetching follows")
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Fetched %v follows", len(follows)))

	fmt.Println("Inserting into postgres...")

	// for _, f := range follows {

	err = psql.Save(&follows).Error

	if err != nil {
		fmt.Println("Error saving follows")
		panic(err)
	}
	// }

	fmt.Println("Successfully inserted follows")

	fmt.Println("Fetching keyvals...")

	keyvals := []model.KeyVal{}

	err = sql.Find(&keyvals).Error

	if err != nil {
		fmt.Println("Error fetching keyvals")
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Fetched %v keyvals", len(keyvals)))

	fmt.Println("Inserting into postgres...")

	for _, k := range keyvals {
		err = psql.Save(&k).Error

		if err != nil {
			fmt.Println("Error saving keyvals")
			panic(err)
		}

	}

	fmt.Println("Successfully inserted keyvals")

	fmt.Println("All done! Ready to Go :))")
}
