package appwrite

import (
	"strings"
)

// Users service
type Users struct {
	client Client
}

func NewUsers(clt Client) Users {  
	service := Users{
		client: clt,
	}
	return service
}

// List get a list of all the project's users. You can use the query params to
// filter your results.
func (srv *Users) List(Search string, Limit int, Offset int, OrderType string) (map[string]interface{}, error) {
	path := "/users"

	params := map[string]interface{}{
		"search": Search,
		"limit": Limit,
		"offset": Offset,
		"orderType": OrderType,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Create create a new user.
func (srv *Users) Create(Email string, Password string, Name string) (map[string]interface{}, error) {
	path := "/users"

	params := map[string]interface{}{
		"email": Email,
		"password": Password,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// Get get a user by its unique ID.
func (srv *Users) Get(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Delete delete a user by its unique ID.
func (srv *Users) Delete(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateEmail update the user email by its unique ID.
func (srv *Users) UpdateEmail(UserId string, Email string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/email")

	params := map[string]interface{}{
		"email": Email,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// GetLogs get a user activity logs list by its unique ID.
func (srv *Users) GetLogs(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/logs")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateName update the user name by its unique ID.
func (srv *Users) UpdateName(UserId string, Name string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/name")

	params := map[string]interface{}{
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdatePassword update the user password by its unique ID.
func (srv *Users) UpdatePassword(UserId string, Password string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/password")

	params := map[string]interface{}{
		"password": Password,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// GetPrefs get the user preferences by its unique ID.
func (srv *Users) GetPrefs(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdatePrefs update the user preferences by its unique ID. You can pass only
// the specific settings you wish to update.
func (srv *Users) UpdatePrefs(UserId string, Prefs interface{}) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")

	params := map[string]interface{}{
		"prefs": Prefs,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// GetSessions get the user sessions list by its unique ID.
func (srv *Users) GetSessions(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// DeleteSessions delete all user's sessions by using the user's unique ID.
func (srv *Users) DeleteSessions(UserId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/sessions")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// DeleteSession delete a user sessions by its unique ID.
func (srv *Users) DeleteSession(UserId string, SessionId string) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId, "{sessionId}", SessionId)
	path := r.Replace("/users/{userId}/sessions/{sessionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateStatus update the user status by its unique ID.
func (srv *Users) UpdateStatus(UserId string, Status int) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/status")

	params := map[string]interface{}{
		"status": Status,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdateVerification update the user email verification status by its unique
// ID.
func (srv *Users) UpdateVerification(UserId string, EmailVerification bool) (map[string]interface{}, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/verification")

	params := map[string]interface{}{
		"emailVerification": EmailVerification,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}
