package appwrite

import (
	"strings"
)

// Teams service
type Teams struct {
	client Client
}

func NewTeams(clt Client) *Teams {
	return &Teams{
		client: clt,
	}
}

// List get a list of all the teams in which the current user is a member. You
// can use the parameters to filter your results.
// 
// In admin mode, this endpoint returns a list of all the teams in the current
// project. [Learn more about different API modes](/docs/admin).
func (srv *Teams) List(Search string, Limit int, Offset int, Cursor string, CursorDirection string, OrderType string) (*ClientResponse, error) {
	path := "/teams"

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

// Create create a new team. The user who creates the team will automatically
// be assigned as the owner of the team. Only the users with the owner role
// can invite new members, add new owners and delete or update the team.
func (srv *Teams) Create(TeamId string, Name string, Roles []interface{}) (*ClientResponse, error) {
	path := "/teams"

	params := map[string]interface{}{
		"teamId": TeamId,
		"name": Name,
		"roles": Roles,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// Get get a team by its ID. All team members have read access for this
// resource.
func (srv *Teams) Get(TeamId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Update update a team using its ID. Only members with the owner role can
// update the team.
func (srv *Teams) Update(TeamId string, Name string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}")

	params := map[string]interface{}{
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// Delete delete a team using its ID. Only team members with the owner role
// can delete the team.
func (srv *Teams) Delete(TeamId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// GetMemberships use this endpoint to list a team's members using the team's
// ID. All team members have read access to this endpoint.
func (srv *Teams) GetMemberships(TeamId string, Search string, Limit int, Offset int, Cursor string, CursorDirection string, OrderType string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/memberships")

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

// CreateMembership invite a new member to join your team. If initiated from
// the client SDK, an email with a link to join the team will be sent to the
// member's email address and an account will be created for them should they
// not be signed up already. If initiated from server-side SDKs, the new
// member will automatically be added to the team.
// 
// Use the 'url' parameter to redirect the user from the invitation email back
// to your app. When the user is redirected, use the [Update Team Membership
// Status](/docs/client/teams#teamsUpdateMembershipStatus) endpoint to allow
// the user to accept the invitation to the team. 
// 
// Please note that to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// the only valid redirect URL's are the once from domains you have set when
// adding your platforms in the console interface.
func (srv *Teams) CreateMembership(TeamId string, Email string, Roles []interface{}, Url string, Name string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	path := r.Replace("/teams/{teamId}/memberships")

	params := map[string]interface{}{
		"email": Email,
		"roles": Roles,
		"url": Url,
		"name": Name,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetMembership get a team member by the membership unique id. All team
// members have read access for this resource.
func (srv *Teams) GetMembership(TeamId string, MembershipId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateMembershipRoles modify the roles of a team member. Only team members
// with the owner role have access to this endpoint. Learn more about [roles
// and permissions](/docs/permissions).
func (srv *Teams) UpdateMembershipRoles(TeamId string, MembershipId string, Roles []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}")

	params := map[string]interface{}{
		"roles": Roles,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// DeleteMembership this endpoint allows a user to leave a team or for a team
// owner to delete the membership of any other team member. You can also use
// this endpoint to delete a user membership even if it is not accepted.
func (srv *Teams) DeleteMembership(TeamId string, MembershipId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateMembershipStatus use this endpoint to allow a user to accept an
// invitation to join a team after being redirected back to your app from the
// invitation email received by the user.
// 
// If the request is successful, a session for the user is automatically
// created.
// 
func (srv *Teams) UpdateMembershipStatus(TeamId string, MembershipId string, UserId string, Secret string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	path := r.Replace("/teams/{teamId}/memberships/{membershipId}/status")

	params := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}
