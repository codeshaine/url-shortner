package test

import (
	"log"
	"os"
	"testing"

	"github.com/codeshaine/url-shortner/db"
	"github.com/codeshaine/url-shortner/internal/utils"
	"github.com/joho/godotenv"
)

func beforeAll(t *testing.T) {
	t.Helper()
	//loading .evn file
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	os.Setenv("DB_NAME", "test_url_shortner")
	db.DbConnect()
}

func afterAll(t *testing.T) {
	t.Helper()
	_, err := db.Db.Exec("DROP TABLE urls")
	if err != nil {
		t.Errorf("Erroc occured while dropping table")
	}
	t.Log("dropped the table ")

	err = db.Db.Close()
	if err != nil {
		t.Errorf("Error occured while closing the service: %v", err)
	}
	t.Log("closed the connection")
}

func TestInsertingQuery(t *testing.T) {
	beforeAll(t)
	t.Run("inserting one valid record", func(t *testing.T) {
		shortUrl := utils.GenerateUnqueUrl()
		longurl := "http://test.shaines"
		_, err := db.InsertUrl(longurl, shortUrl)
		if err != nil {
			t.Error("Not inserted")
		}
	})
	t.Run("inserting two same valid record", func(t *testing.T) {
		shortUrl := utils.GenerateUnqueUrl()
		longurl := "http://test.dup.shaine.com"
		_, err := db.InsertUrl(longurl, shortUrl)
		if err != nil {
			t.Error("Not inserted")
		}
		t.Log("Inserted successfully:")
		shortUrl = utils.GenerateUnqueUrl()
		_, err = db.InsertUrl(longurl, shortUrl)
		if err == nil {
			t.Error("unexpected: inserted ")
		}

	})
	t.Run("inserting invalid record : should not insert record", func(t *testing.T) {
		shortUrl := utils.GenerateUnqueUrl()
		longurl := "http://test.'shaine"
		_, err := db.InsertUrl(longurl, shortUrl)
		if err == nil {
			t.Error("row got Inserted")
		}

	})
	t.Run("inserting empty record : should not insert record", func(t *testing.T) {
		shortUrl := utils.GenerateUnqueUrl()
		longurl := " "
		_, err := db.InsertUrl(longurl, shortUrl)
		if err == nil {
			t.Error("row got Inserted")
		}

	})

	afterAll(t)
}

func TestGettingLongUrl(t *testing.T) {
	beforeAll(t)
	t.Run("test should pass", func(t *testing.T) {
		shortUrl := utils.GenerateUnqueUrl()
		longurl := "http://test.shaine"
		_, err := db.InsertUrl(longurl, shortUrl)
		if err != nil {
			t.Error("Not inserted")
		}

		redirectData, reErr := db.GetLongUrl(shortUrl)
		if reErr != nil {
			t.Error("Not inserted")
		}
		if redirectData.LongUrl != longurl {
			t.Errorf("not a valid url")
		}
	})

	t.Run("testing for empty input: should fail", func(t *testing.T) {
		shortUrl := ""
		_, err := db.GetLongUrl(shortUrl)
		if err == nil {
			t.Error(" inserted")
		}

	})

	afterAll(t)
}
