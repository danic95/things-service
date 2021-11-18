package service

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/danic95/things-service/goutils/dbfactory"
	"github.com/danic95/things-service/goutils/settings"
	"github.com/danic95/things-service/internal/things"
	"github.com/danic95/things-service/internal/things/entity"
	"github.com/danic95/things-service/internal/things/repository/cache"
	"github.com/danic95/things-service/internal/things/repository/sqlite"
)

var service things.Service

func TestMain(m *testing.M) {

	ctx := context.Background()

	config := &settings.Settings{
		Cache: settings.Cache{
			Enabled: false,
		},
		DB: settings.Database{
			Engine: "sqlite",
			Name:   "service",
		},
	}

	db, err := dbfactory.CreateSqliteConnection(config)
	if err != nil {
		log.Panic(err)
	}

	cach, err := cache.New(ctx, config)
	if err != nil {
		log.Panic(err)
	}

	repo := sqlite.New(db, cach)
	err = repo.PopulateSchema(ctx)
	if err != nil {
		log.Panic(err)
	}

	service = New(repo)

	// run the test
	code := m.Run()

	path, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	os.Remove(path + "\\" + config.DB.Name + ".db")
	os.Exit(code)
}

func TestCreateThing(t *testing.T) {
	// table driven tests

	t.Parallel()

	tm, _ := time.Parse("1995-01-01", "1995-01-01")

	testCases := []struct {
		Name          string
		Data          *entity.Thing
		ExpectedError error
	}{
		{
			Name: "Should fail without name",
			Data: &entity.Thing{Name: "", Email: "", Phone: "", Birth_date: tm,
				Start_day: 5},
			ExpectedError: ErrInvalidThing,
		},
		{
			Name: "Should fail without start day",
			Data: &entity.Thing{Name: "test", Email: "", Phone: "", Birth_date: tm,
				Start_day: 0},
			ExpectedError: ErrInvalidThing,
		},
		{
			Name: "Should create the thing",
			Data: &entity.Thing{Name: "test", Email: "test@uvltd.com", Phone: "+50499887766", Birth_date: tm,
				Start_day: 5},
			ExpectedError: nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			err := service.CreateThing(ctx, tc.Data)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestGetThing(t *testing.T) {
	t.Parallel()
	// table tests
	testCases := []struct {
		Name          string
		Data          int
		ExpectedError error
	}{
		{
			Name:          "Should fail if no item found",
			Data:          100,
			ExpectedError: ErrThingNotFound,
		},
		{
			Name:          "Should get one item",
			Data:          1,
			ExpectedError: nil,
		},
	}
	// create a thing
	tm, _ := time.Parse("1995-01-01", "1995-01-01")
	ctx := context.Background()
	service.CreateThing(ctx, &entity.Thing{Name: "test", Email: "test@uvltd.com", Phone: "+50499887766", Birth_date: tm,
		Start_day: 5})

	// tests the cases
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			_, err := service.GetThing(ctx, uint(tc.Data))

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

func TestGetThings(t *testing.T) {
	t.Parallel()
	// table tests
	testCases := []struct {
		Name          string
		ExpectedError error
	}{
		{
			Name:          "Should return all items",
			ExpectedError: nil,
		},
	}
	// create a thing
	tm, _ := time.Parse("1995-01-01", "1995-01-01")
	ctx := context.Background()
	service.CreateThing(ctx, &entity.Thing{Name: "test", Email: "test@uvltd.com", Phone: "+50499887766", Birth_date: tm,
		Start_day: 5})

	// tests the cases
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			_, err := service.GetThings(ctx)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

// -> create
// -> get
