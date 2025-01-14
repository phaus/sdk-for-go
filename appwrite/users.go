package appwrite

import (
	"strings"
)

// Users service
type Users struct {
	client Client
}

func NewUsers(clt Client) *Users {
	return &Users{
		client: clt,
	}
}

// List get a list of all the project's users. You can use the query params to
// filter your results.
func (srv *Users) List(Search string, Limit int, Offset int, Cursor string, CursorDirection string, OrderType string) (*ClientResponse, error) {
	path := "/users"

	params := map[string]interface{}{
		"search": Search,
		"limit": Limit,
		"offset": Offset,
		"cursor": Cursor,
		"cursorDirection": CursorDirection,
		"orderType": OrderType,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Create create a new user.
func (srv *Users) Create(UserId string, Email string, Password string, Name string) (*ClientResponse, error) {
	path := "/users"

	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"password": Password,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetUsage
func (srv *Users) GetUsage(Range string, Provider string) (*ClientResponse, error) {
	path := "/users/usage"

	params := map[string]interface{}{
		"range": Range,
		"provider": Provider,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Get get a user by its unique ID.
func (srv *Users) Get(UserId string) (*ClientResponse, error) {
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
func (srv *Users) Delete(UserId string) (*ClientResponse, error) {
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
func (srv *Users) UpdateEmail(UserId string, Email string) (*ClientResponse, error) {
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

// GetLogs get the user activity logs list by its unique ID.
func (srv *Users) GetLogs(UserId string, Limit int, Offset int) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/logs")

	params := map[string]interface{}{
		"limit": Limit,
		"offset": Offset,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateName update the user name by its unique ID.
func (srv *Users) UpdateName(UserId string, Name string) (*ClientResponse, error) {
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
func (srv *Users) UpdatePassword(UserId string, Password string) (*ClientResponse, error) {
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
func (srv *Users) GetPrefs(UserId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{userId}", UserId)
	path := r.Replace("/users/{userId}/prefs")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdatePrefs update the user preferences by its unique ID. The object you
// pass is stored as is, and replaces any previous value. The maximum allowed
// prefs size is 64kB and throws error if exceeded.
func (srv *Users) UpdatePrefs(UserId string, Prefs interface{}) (*ClientResponse, error) {
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
func (srv *Users) GetSessions(UserId string) (*ClientResponse, error) {
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
func (srv *Users) DeleteSessions(UserId string) (*ClientResponse, error) {
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
func (srv *Users) DeleteSession(UserId string, SessionId string) (*ClientResponse, error) {
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
func (srv *Users) UpdateStatus(UserId string, Status bool) (*ClientResponse, error) {
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
func (srv *Users) UpdateVerification(UserId string, EmailVerification bool) (*ClientResponse, error) {
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
