package appwrite

import (
	"strings"
)

// Projects service
type Projects struct {
	client Client
}

func NewProjects(clt Client) *Projects {
	return &Projects{
		client: clt,
	}
}

// List
func (srv *Projects) List(Search string, Limit int, Offset int, Cursor string, CursorDirection string, OrderType string) (*ClientResponse, error) {
	path := "/projects"

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

// Create
func (srv *Projects) Create(ProjectId string, Name string, TeamId string, Description string, Logo string, Url string, LegalName string, LegalCountry string, LegalState string, LegalCity string, LegalAddress string, LegalTaxId string) (*ClientResponse, error) {
	path := "/projects"

	params := map[string]interface{}{
		"projectId": ProjectId,
		"name": Name,
		"teamId": TeamId,
		"description": Description,
		"logo": Logo,
		"url": Url,
		"legalName": LegalName,
		"legalCountry": LegalCountry,
		"legalState": LegalState,
		"legalCity": LegalCity,
		"legalAddress": LegalAddress,
		"legalTaxId": LegalTaxId,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// Get
func (srv *Projects) Get(ProjectId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// Update
func (srv *Projects) Update(ProjectId string, Name string, Description string, Logo string, Url string, LegalName string, LegalCountry string, LegalState string, LegalCity string, LegalAddress string, LegalTaxId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}")

	params := map[string]interface{}{
		"name": Name,
		"description": Description,
		"logo": Logo,
		"url": Url,
		"legalName": LegalName,
		"legalCountry": LegalCountry,
		"legalState": LegalState,
		"legalCity": LegalCity,
		"legalAddress": LegalAddress,
		"legalTaxId": LegalTaxId,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// Delete
func (srv *Projects) Delete(ProjectId string, Password string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}")

	params := map[string]interface{}{
		"password": Password,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateAuthLimit
func (srv *Projects) UpdateAuthLimit(ProjectId string, Limit int) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/auth/limit")

	params := map[string]interface{}{
		"limit": Limit,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// UpdateAuthStatus
func (srv *Projects) UpdateAuthStatus(ProjectId string, Method string, Status bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{method}", Method)
	path := r.Replace("/projects/{projectId}/auth/{method}")

	params := map[string]interface{}{
		"status": Status,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// ListDomains
func (srv *Projects) ListDomains(ProjectId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/domains")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateDomain
func (srv *Projects) CreateDomain(ProjectId string, Domain string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/domains")

	params := map[string]interface{}{
		"domain": Domain,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetDomain
func (srv *Projects) GetDomain(ProjectId string, DomainId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{domainId}", DomainId)
	path := r.Replace("/projects/{projectId}/domains/{domainId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// DeleteDomain
func (srv *Projects) DeleteDomain(ProjectId string, DomainId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{domainId}", DomainId)
	path := r.Replace("/projects/{projectId}/domains/{domainId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateDomainVerification
func (srv *Projects) UpdateDomainVerification(ProjectId string, DomainId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{domainId}", DomainId)
	path := r.Replace("/projects/{projectId}/domains/{domainId}/verification")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// ListKeys
func (srv *Projects) ListKeys(ProjectId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/keys")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateKey
func (srv *Projects) CreateKey(ProjectId string, Name string, Scopes []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/keys")

	params := map[string]interface{}{
		"name": Name,
		"scopes": Scopes,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetKey
func (srv *Projects) GetKey(ProjectId string, KeyId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{keyId}", KeyId)
	path := r.Replace("/projects/{projectId}/keys/{keyId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateKey
func (srv *Projects) UpdateKey(ProjectId string, KeyId string, Name string, Scopes []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{keyId}", KeyId)
	path := r.Replace("/projects/{projectId}/keys/{keyId}")

	params := map[string]interface{}{
		"name": Name,
		"scopes": Scopes,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// DeleteKey
func (srv *Projects) DeleteKey(ProjectId string, KeyId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{keyId}", KeyId)
	path := r.Replace("/projects/{projectId}/keys/{keyId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateOAuth2
func (srv *Projects) UpdateOAuth2(ProjectId string, Provider string, AppId string, Secret string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/oauth2")

	params := map[string]interface{}{
		"provider": Provider,
		"appId": AppId,
		"secret": Secret,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// ListPlatforms
func (srv *Projects) ListPlatforms(ProjectId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/platforms")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreatePlatform
func (srv *Projects) CreatePlatform(ProjectId string, Type string, Name string, Key string, Store string, Hostname string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/platforms")

	params := map[string]interface{}{
		"type": Type,
		"name": Name,
		"key": Key,
		"store": Store,
		"hostname": Hostname,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetPlatform
func (srv *Projects) GetPlatform(ProjectId string, PlatformId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{platformId}", PlatformId)
	path := r.Replace("/projects/{projectId}/platforms/{platformId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdatePlatform
func (srv *Projects) UpdatePlatform(ProjectId string, PlatformId string, Name string, Key string, Store string, Hostname string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{platformId}", PlatformId)
	path := r.Replace("/projects/{projectId}/platforms/{platformId}")

	params := map[string]interface{}{
		"name": Name,
		"key": Key,
		"store": Store,
		"hostname": Hostname,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// DeletePlatform
func (srv *Projects) DeletePlatform(ProjectId string, PlatformId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{platformId}", PlatformId)
	path := r.Replace("/projects/{projectId}/platforms/{platformId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}

// UpdateServiceStatus
func (srv *Projects) UpdateServiceStatus(ProjectId string, Service string, Status bool) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/service")

	params := map[string]interface{}{
		"service": Service,
		"status": Status,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", path, headers, params)
}

// GetUsage
func (srv *Projects) GetUsage(ProjectId string, Range string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/usage")

	params := map[string]interface{}{
		"range": Range,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// ListWebhooks
func (srv *Projects) ListWebhooks(ProjectId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/webhooks")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// CreateWebhook
func (srv *Projects) CreateWebhook(ProjectId string, Name string, Events []interface{}, Url string, Security bool, HttpUser string, HttpPass string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId)
	path := r.Replace("/projects/{projectId}/webhooks")

	params := map[string]interface{}{
		"name": Name,
		"events": Events,
		"url": Url,
		"security": Security,
		"httpUser": HttpUser,
		"httpPass": HttpPass,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", path, headers, params)
}

// GetWebhook
func (srv *Projects) GetWebhook(ProjectId string, WebhookId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{webhookId}", WebhookId)
	path := r.Replace("/projects/{projectId}/webhooks/{webhookId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// UpdateWebhook
func (srv *Projects) UpdateWebhook(ProjectId string, WebhookId string, Name string, Events []interface{}, Url string, Security bool, HttpUser string, HttpPass string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{webhookId}", WebhookId)
	path := r.Replace("/projects/{projectId}/webhooks/{webhookId}")

	params := map[string]interface{}{
		"name": Name,
		"events": Events,
		"url": Url,
		"security": Security,
		"httpUser": HttpUser,
		"httpPass": HttpPass,
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", path, headers, params)
}

// DeleteWebhook
func (srv *Projects) DeleteWebhook(ProjectId string, WebhookId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{projectId}", ProjectId, "{webhookId}", WebhookId)
	path := r.Replace("/projects/{projectId}/webhooks/{webhookId}")

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", path, headers, params)
}
