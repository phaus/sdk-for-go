package appwrite

import (
	"strings"
)

// Account service
type Account struct {
	client Client
}

func NewAccount(clt Client) *Account {
	return &Account{
		client: clt,
	}
}

// Get get currently logged in user data as JSON object.
func (srv *Account) Get() (*ClientResponse, error) {
	path := "/account"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Create use this endpoint to allow a new user to register a new account in
// your project. After the user registration completes successfully, you can
// use the
// [/account/verfication](/docs/client/account#accountCreateVerification)
// route to start verifying the user email address. To allow the new user to
// login to their new account, you need to create a new [account
// session](/docs/client/account#accountCreateSession).
func (srv *Account) Create(UserId string, Email string, Password string, Name string) (*ClientResponse, error) {
	path := "/account"

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

// Delete delete a currently logged in user account. Behind the scene, the
// user record is not deleted but permanently blocked from any access. This is
// done to avoid deleted accounts being overtaken by new users with the same
// email address. Any user-related resources like documents or storage files
// should be deleted separately.
func (srv *Account) Delete() (*ClientResponse, error) {
	path := "/account"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateEmail update currently logged in user account email address. After
// changing user address, the user confirmation status will get reset. A new
// confirmation email is not sent automatically however you can use the send
// confirmation email endpoint again to send the confirmation email. For
// security measures, user password is required to complete this request.
// This endpoint can also be used to convert an anonymous account to a normal
// one, by passing an email address and a new password.
// 
func (srv *Account) UpdateEmail(Email string, Password string) (*ClientResponse, error) {
	path := "/account/email"

	params := map[string]interface{}{
		"email": Email,
		"password": Password,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// CreateJWT use this endpoint to create a JSON Web Token. You can use the
// resulting JWT to authenticate on behalf of the current user when working
// with the Appwrite server-side API and SDKs. The JWT secret is valid for 15
// minutes from its creation and will be invalid if the user will logout in
// that time frame.
func (srv *Account) CreateJWT() (*ClientResponse, error) {
	path := "/account/jwt"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetLogs get currently logged in user list of latest security activity logs.
// Each log returns user IP address, location and date and time of log.
func (srv *Account) GetLogs(Limit int, Offset int) (*ClientResponse, error) {
	path := "/account/logs"

	params := map[string]interface{}{
		"limit": Limit,
		"offset": Offset,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateName update currently logged in user account name.
func (srv *Account) UpdateName(Name string) (*ClientResponse, error) {
	path := "/account/name"

	params := map[string]interface{}{
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdatePassword update currently logged in user password. For validation,
// user is required to pass in the new password, and the old password. For
// users created with OAuth and Team Invites, oldPassword is optional.
func (srv *Account) UpdatePassword(Password string, OldPassword string) (*ClientResponse, error) {
	path := "/account/password"

	params := map[string]interface{}{
		"password": Password,
		"oldPassword": OldPassword,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// GetPrefs get currently logged in user preferences as a key-value object.
func (srv *Account) GetPrefs() (*ClientResponse, error) {
	path := "/account/prefs"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdatePrefs update currently logged in user account preferences. The object
// you pass is stored as is, and replaces any previous value. The maximum
// allowed prefs size is 64kB and throws error if exceeded.
func (srv *Account) UpdatePrefs(Prefs interface{}) (*ClientResponse, error) {
	path := "/account/prefs"

	params := map[string]interface{}{
		"prefs": Prefs,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// CreateRecovery sends the user an email with a temporary secret key for
// password reset. When the user clicks the confirmation link he is redirected
// back to your app password reset URL with the secret key and email address
// values attached to the URL query string. Use the query string params to
// submit a request to the [PUT
// /account/recovery](/docs/client/account#accountUpdateRecovery) endpoint to
// complete the process. The verification link sent to the user's email
// address is valid for 1 hour.
func (srv *Account) CreateRecovery(Email string, Url string) (*ClientResponse, error) {
	path := "/account/recovery"

	params := map[string]interface{}{
		"email": Email,
		"url": Url,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// UpdateRecovery use this endpoint to complete the user account password
// reset. Both the **userId** and **secret** arguments will be passed as query
// parameters to the redirect URL you have provided when sending your request
// to the [POST /account/recovery](/docs/client/account#accountCreateRecovery)
// endpoint.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
func (srv *Account) UpdateRecovery(UserId string, Secret string, Password string, PasswordAgain string) (*ClientResponse, error) {
	path := "/account/recovery"

	params := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
		"password": Password,
		"passwordAgain": PasswordAgain,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// GetSessions get currently logged in user list of active sessions across
// different devices.
func (srv *Account) GetSessions() (*ClientResponse, error) {
	path := "/account/sessions"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateSession allow the user to login into their account by providing a
// valid email and password combination. This route will create a new session
// for the user.
func (srv *Account) CreateSession(Email string, Password string) (*ClientResponse, error) {
	path := "/account/sessions"

	params := map[string]interface{}{
		"email": Email,
		"password": Password,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// DeleteSessions delete all sessions from the user account and remove any
// sessions cookies from the end client.
func (srv *Account) DeleteSessions() (*ClientResponse, error) {
	path := "/account/sessions"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// CreateAnonymousSession use this endpoint to allow a new user to register an
// anonymous account in your project. This route will also create a new
// session for the user. To allow the new user to convert an anonymous account
// to a normal account, you need to update its [email and
// password](/docs/client/account#accountUpdateEmail) or create an [OAuth2
// session](/docs/client/account#accountCreateOAuth2Session).
func (srv *Account) CreateAnonymousSession() (*ClientResponse, error) {
	path := "/account/sessions/anonymous"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// CreateMagicURLSession sends the user an email with a secret key for
// creating a session. When the user clicks the link in the email, the user is
// redirected back to the URL you provided with the secret key and userId
// values attached to the URL query string. Use the query string parameters to
// submit a request to the [PUT
// /account/sessions/magic-url](/docs/client/account#accountUpdateMagicURLSession)
// endpoint to complete the login process. The link sent to the user's email
// address is valid for 1 hour. If you are on a mobile device you can leave
// the URL parameter empty, so that the login completion will be handled by
// your Appwrite instance by default.
func (srv *Account) CreateMagicURLSession(UserId string, Email string, Url string) (*ClientResponse, error) {
	path := "/account/sessions/magic-url"

	params := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"url": Url,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// UpdateMagicURLSession use this endpoint to complete creating the session
// with the Magic URL. Both the **userId** and **secret** arguments will be
// passed as query parameters to the redirect URL you have provided when
// sending your request to the [POST
// /account/sessions/magic-url](/docs/client/account#accountCreateMagicURLSession)
// endpoint.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
func (srv *Account) UpdateMagicURLSession(UserId string, Secret string) (*ClientResponse, error) {
	path := "/account/sessions/magic-url"

	params := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// CreateOAuth2Session allow the user to login to their account using the
// OAuth2 provider of their choice. Each OAuth2 provider should be enabled
// from the Appwrite console first. Use the success and failure arguments to
// provide a redirect URL's back to your app when login is completed.
// 
// If there is already an active session, the new session will be attached to
// the logged-in account. If there are no active sessions, the server will
// attempt to look for a user with the same email address as the email
// received from the OAuth2 provider and attach the new session to the
// existing user. If no matching user is found - the server will create a new
// user..
// 
func (srv *Account) CreateOAuth2Session(Provider string, Success string, Failure string, Scopes []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{provider}", Provider)
	path := r.Replace("/account/sessions/oauth2/{provider}")

	params := map[string]interface{}{
		"success": Success,
		"failure": Failure,
		"scopes": Scopes,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetSession use this endpoint to get a logged in user's session using a
// Session ID. Inputting 'current' will return the current session being used.
func (srv *Account) GetSession(SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateSession
func (srv *Account) UpdateSession(SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// DeleteSession use this endpoint to log out the currently logged in user
// from all their account sessions across all of their different devices. When
// using the Session ID argument, only the unique session ID provided is
// deleted.
// 
func (srv *Account) DeleteSession(SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	path := r.Replace("/account/sessions/{sessionId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// CreateVerification use this endpoint to send a verification message to your
// user email address to confirm they are the valid owners of that address.
// Both the **userId** and **secret** arguments will be passed as query
// parameters to the URL you have provided to be attached to the verification
// email. The provided URL should redirect the user back to your app and allow
// you to complete the verification process by verifying both the **userId**
// and **secret** parameters. Learn more about how to [complete the
// verification process](/docs/client/account#accountUpdateVerification). The
// verification link sent to the user's email address is valid for 7 days.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md),
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
// 
func (srv *Account) CreateVerification(Url string) (*ClientResponse, error) {
	path := "/account/verification"

	params := map[string]interface{}{
		"url": Url,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// UpdateVerification use this endpoint to complete the user email
// verification process. Use both the **userId** and **secret** parameters
// that were attached to your app URL to verify the user email ownership. If
// confirmed this route will return a 200 status code.
func (srv *Account) UpdateVerification(UserId string, Secret string) (*ClientResponse, error) {
	path := "/account/verification"

	params := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}
