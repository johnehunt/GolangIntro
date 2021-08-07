package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestGETPeople(t *testing.T) {
    t.Run("Get a individual person", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/api/people/1", nil)
        response := httptest.NewRecorder()

        serve(response, request)

        respBodyString := response.Body.String()
        requiredResult := "{\"id\":0,\"name\":\"John\",\"age\":57}"

        if respBodyString != requiredResult {
            t.Errorf("Error Got %q, expected %q", respBodyString, requiredResult)
        }
    })
}
