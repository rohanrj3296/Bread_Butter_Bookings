package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/config"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig
	m :=routes(&app)
	switch v := m.(type){
	case *chi.Mux:
		//do nothing; test passed
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux which should be returned by routes function, type is %T",v))

	}
}
