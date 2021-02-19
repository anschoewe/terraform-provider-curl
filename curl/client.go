package curl

import (
	"crypto/tls"
	azureAuth "github.com/Azure/go-autorest/autorest/azure/auth"
	"log"
	"net/http"
	"time"
)

type OAuthOpts struct {
	ClientId     string
	ClientSecret string
	Resource     string
	TenantId     string
}

type ApiClientOpts struct {
	insecure              bool
	//headers               map[string]string
	timeout               int
	oauthConfig          OAuthOpts
}

type MyHttpClient struct {
	httpClient           *http.Client
	//headers               map[string]string
	timeout               int
	oauthConfig          *azureAuth.ClientCredentialsConfig
}

func NewClient(opts ApiClientOpts) (*MyHttpClient, error) {
	tlsConfig := &tls.Config{
		/* Disable TLS verification if requested */
		InsecureSkipVerify: opts.insecure,
	}

	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
		Proxy:           http.ProxyFromEnvironment,
	}

	client := MyHttpClient{
		httpClient: &http.Client{
			Timeout:   time.Second * time.Duration(opts.timeout),
			Transport: tr,
		},
		//headers:               opts.headers,
	}

	if opts.oauthConfig.ClientId != "" && opts.oauthConfig.ClientSecret != "" && opts.oauthConfig.TenantId != "" && opts.oauthConfig.Resource != "" {
		log.Printf("[INFO] ******* Creating Oauth Config")
		config := azureAuth.NewClientCredentialsConfig(
			opts.oauthConfig.ClientId,
			opts.oauthConfig.ClientSecret,
			opts.oauthConfig.TenantId)
		config.Resource = opts.oauthConfig.Resource
		client.oauthConfig = &config
	}

	return &client, nil
}