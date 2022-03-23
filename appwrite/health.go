package appwrite

import (
)

// Health service
type Health struct {
	client Client
}

func NewHealth(clt Client) *Health {
	return &Health{
		client: clt,
	}
}

// Get check the Appwrite HTTP server is up and responsive.
func (srv *Health) Get() (*ClientResponse, error) {
	path := "/health"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetAntivirus check the Appwrite Antivirus server is up and connection is
// successful.
func (srv *Health) GetAntivirus() (*ClientResponse, error) {
	path := "/health/anti-virus"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetCache check the Appwrite in-memory cache server is up and connection is
// successful.
func (srv *Health) GetCache() (*ClientResponse, error) {
	path := "/health/cache"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetDB check the Appwrite database server is up and connection is
// successful.
func (srv *Health) GetDB() (*ClientResponse, error) {
	path := "/health/db"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetQueueCertificates get the number of certificates that are waiting to be
// issued against [Letsencrypt](https://letsencrypt.org/) in the Appwrite
// internal queue server.
func (srv *Health) GetQueueCertificates() (*ClientResponse, error) {
	path := "/health/queue/certificates"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetQueueFunctions
func (srv *Health) GetQueueFunctions() (*ClientResponse, error) {
	path := "/health/queue/functions"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetQueueLogs get the number of logs that are waiting to be processed in the
// Appwrite internal queue server.
func (srv *Health) GetQueueLogs() (*ClientResponse, error) {
	path := "/health/queue/logs"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetQueueUsage get the number of usage stats that are waiting to be
// processed in the Appwrite internal queue server.
func (srv *Health) GetQueueUsage() (*ClientResponse, error) {
	path := "/health/queue/usage"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetQueueWebhooks get the number of webhooks that are waiting to be
// processed in the Appwrite internal queue server.
func (srv *Health) GetQueueWebhooks() (*ClientResponse, error) {
	path := "/health/queue/webhooks"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetStorageLocal check the Appwrite local storage device is up and
// connection is successful.
func (srv *Health) GetStorageLocal() (*ClientResponse, error) {
	path := "/health/storage/local"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}

// GetTime check the Appwrite server time is synced with Google remote NTP
// server. We use this technology to smoothly handle leap seconds with no
// disruptive events. The [Network Time
// Protocol](https://en.wikipedia.org/wiki/Network_Time_Protocol) (NTP) is
// used by hundreds of millions of computers and devices to synchronize their
// clocks over the Internet. If your computer sets its own clock, it likely
// uses NTP.
func (srv *Health) GetTime() (*ClientResponse, error) {
	path := "/health/time"

	params := map[string]interface{}{
	}

	headers := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", path, headers, params)
}
