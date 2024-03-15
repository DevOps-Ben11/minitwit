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

	psql.Transaction(func(tx *gorm.DB) error {
		fmt.Println("Successfully connected to postgres")

		fmt.Println("Fetching users...")
		var userCount int
		err = sql.Model(&model.User{}).Select("count(*)").Find(&userCount).Error

		fmt.Println(fmt.Sprintf("Fetched %v users", userCount))

		fmt.Println("Inserting into postgres...")

		for i := 0; i < userCount; i += 15000 {
			users := []model.User{}
			err = sql.Select("username", "email", "pw_hash").Limit(15000).Offset(i).Find(&users).Error

			if err != nil {
				fmt.Println("Error fetching users")
				return err
			}

			err = tx.Save(&users).Error

			if err != nil {
				fmt.Println(err)
				fmt.Println("Error saving users")
				return err
			}
		}

		fmt.Println("Successfully inserted users")

		fmt.Println("Fetching messages...")

		var messageCount int
		err = sql.Model(&model.Message{}).Select("count(*)").Find(&messageCount).Error

		fmt.Println(fmt.Sprintf("Fetched %v messages", messageCount))

		fmt.Println("Inserting into postgres...")

		// run messages into db in clusters of insertions
		for i := 0; i < messageCount; i += 15000 {
			messages := []model.Message{}
			err = sql.Select("author_id", "text", "pub_date", "flagged").Limit(15000).Offset(i).Find(&messages).Error

			if err != nil {
				fmt.Println("Error fetching messages")
				return err
			}

			err = tx.Save(&messages).Error

			if err != nil {
				fmt.Println(err)
				fmt.Println("Error saving messages")
				return err
			}
		}

		fmt.Println("Successfully inserted messages")

		fmt.Println("Fetching follows...")

		var followerCount int
		err = sql.Model(&model.Follower{}).Select("count(*)").Find(&followerCount).Error

		fmt.Println(fmt.Sprintf("Fetched %v followers", followerCount))

		fmt.Println("Inserting into postgres...")

		for i := 0; i < followerCount; i += 25000 {
			followers := []model.Follower{}
			err = sql.Limit(25000).Offset(i).Find(&followers).Error

			if err != nil {
				fmt.Println("Error fetching followers")
				return err
			}

			err = tx.Save(&followers).Error

			if err != nil {
				fmt.Println(err)
				fmt.Println("Error saving followers")
				return err
			}
		}

		fmt.Println("Successfully inserted follows")

		fmt.Println("Fetching keyvals...")

		keyvals := []model.KeyVal{}

		err = sql.Find(&keyvals).Error

		if err != nil {
			fmt.Println("Error fetching keyvals")
			return err
		}

		fmt.Println(fmt.Sprintf("Fetched %v keyvals", len(keyvals)))

		fmt.Println("Inserting into postgres...")

		for _, k := range keyvals {
			err = tx.Save(&k).Error

			if err != nil {
				fmt.Println("Error saving keyvals")
				return err
			}

		}

		fmt.Println("Successfully inserted keyvals")

		fmt.Println("All done! Ready to Go :))")

		return nil
	})
}
