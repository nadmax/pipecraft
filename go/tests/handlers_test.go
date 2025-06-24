package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/nadmax/pipecraft/src/api"
	"github.com/nadmax/pipecraft/src/routes"
)

func setupTestServer() *httptest.Server {
	a := api.NewAPI()
	return httptest.NewServer(routes.SetupRoutes(a))
}

func TestCreateUser(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	body := `{"name":"Alice","email":"alice@example.com"}`
	resp, err := http.Post(server.URL+"/users", "application/json", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}

	var user map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatal(err)
	}

	if user["name"] != "Alice" || user["email"] != "alice@example.com" {
		t.Fatal("unexpected user data")
	}
}

func TestGetAllUsers(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	http.Post(server.URL+"/users", "application/json", strings.NewReader(`{"name":"Bob","email":"bob@example.com"}`))

	resp, err := http.Get(server.URL + "/users")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var users []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		t.Fatal(err)
	}

	if len(users) == 0 {
		t.Fatal("expected at least one user")
	}
}

func TestGetUserByID(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	res, _ := http.Post(server.URL+"/users", "application/json", strings.NewReader(`{"name":"Carl","email":"carl@example.com"}`))
	defer res.Body.Close()

	var user map[string]interface{}
	json.NewDecoder(res.Body).Decode(&user)
	id := int(user["id"].(float64))

	resp, err := http.Get(server.URL + "/users/" + strconv.Itoa(id))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var fetched map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&fetched)

	if fetched["id"] != float64(id) {
		t.Fatalf("expected id %d, got %v", id, fetched["id"])
	}
}

func TestUpdateUser(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	res, _ := http.Post(server.URL+"/users", "application/json", strings.NewReader(`{"name":"Dave","email":"dave@example.com"}`))
	defer res.Body.Close()

	var user map[string]interface{}
	json.NewDecoder(res.Body).Decode(&user)
	id := int(user["id"].(float64))

	update := `{"name":"David","email":"david@new.com"}`
	req, _ := http.NewRequest(http.MethodPut, server.URL+"/users/"+strconv.Itoa(id), strings.NewReader(update))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var updated map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&updated)

	if updated["name"] != "David" {
		t.Fatalf("expected name David, got %v", updated["name"])
	}
}

func TestDeleteUser(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	res, _ := http.Post(server.URL+"/users", "application/json", strings.NewReader(`{"name":"Eve","email":"eve@example.com"}`))
	defer res.Body.Close()

	var user map[string]interface{}
	json.NewDecoder(res.Body).Decode(&user)
	id := int(user["id"].(float64))

	req, _ := http.NewRequest(http.MethodDelete, server.URL+"/users/"+strconv.Itoa(id), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", resp.StatusCode)
	}

	getResp, _ := http.Get(server.URL + "/users/" + strconv.Itoa(id))
	if getResp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", getResp.StatusCode)
	}
}

func TestInvalidJSON(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, _ := http.Post(server.URL+"/users", "application/json", strings.NewReader(`{"name":}`))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	req, _ := http.NewRequest(http.MethodPatch, server.URL+"/users", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", resp.StatusCode)
	}
}

func TestInvalidUserID(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, _ := http.Get(server.URL + "/users/abc")
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestUpdateNonExistentUser(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	update := `{"name":"Ghost","email":"ghost@example.com"}`
	req, _ := http.NewRequest(http.MethodPut, server.URL+"/users/999", strings.NewReader(update))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}

func TestDeleteNonExistentUser(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	req, _ := http.NewRequest(http.MethodDelete, server.URL+"/users/999", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}
