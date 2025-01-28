package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Fenroe/shortform/internal/database"
)

func TestHandlerCreateURL(t *testing.T) {
	queries, cleanup := dbTestSetup()
	defer cleanup()

	// Add repeat to db for a test later on
	_, err := queries.CreateURL(
		context.Background(),
		database.CreateURLParams{
			ID:        "repeat",
			ExpiredAt: time.Now().Add(time.Hour * 24),
			Dest:      "https://www.google.com/",
		},
	)
	if err != nil {
		t.Errorf("error setting up test database, %v\n", err.Error())
	}
	// Mock config
	cfg := &apiConfig{
		DB: queries,
	}

	type testInput struct {
		ID        string
		ExpiredAt int64
		Dest      string
	}

	type testOutput struct {
		Code    int
		Message string
	}

	cases := []struct {
		Name   string
		Input  testInput
		Output testOutput
	}{
		{
			Name: "Happy path",
			Input: testInput{
				ID:        "happy-path",
				ExpiredAt: time.Now().Add(time.Hour).Unix(),
				Dest:      "https://www.example.com",
			},
			Output: testOutput{
				Code:    http.StatusCreated,
				Message: "URL created successfully",
			},
		},
	}

	for _, c := range cases {
		var body createURLParams
		body.Dest = &c.Input.Dest
		body.ID = &c.Input.ID
		body.ExpiredAt = &c.Input.ExpiredAt
		data, err := json.Marshal(body)
		if err != nil {
			t.Errorf("error in testing environment, %v\n", err.Error())
		}
		w := httptest.NewRecorder()
		r, err := http.NewRequest("POST", "/url", bytes.NewBuffer(data))
		if err != nil {
			t.Errorf("error in testing environment, %v\n", err.Error())
		}
		cfg.handlerCreateURL(w, r)
		var res createURLResponse
		if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
			t.Errorf("error in testing environment, %v\n", err.Error())
		}
		if w.Code != c.Output.Code {
			t.Errorf("Invalid response code")
		}
		if res.Message != c.Output.Message {
			t.Errorf("oh no!")
		}
		t.Logf("Test case %s: Success! \n", c.Name)
	}
}
