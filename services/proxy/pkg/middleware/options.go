package middleware

import (
	"net/http"
	"time"

	"github.com/owncloud/ocis/v2/services/proxy/pkg/user/backend"

	settingssvc "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/settings/v0"
	storesvc "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/store/v0"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	"github.com/owncloud/ocis/v2/services/proxy/pkg/config"
)

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	// Logger to use for logging, must be set
	Logger log.Logger
	// TokenManagerConfig for communicating with the reva token manager
	TokenManagerConfig config.TokenManager
	// PolicySelectorConfig for using the policy selector
	PolicySelector config.PolicySelector
	// HTTPClient to use for communication with the oidcAuth provider
	HTTPClient *http.Client
	// UP
	UserProvider backend.UserBackend
	// SettingsRoleService for the roles API in settings
	SettingsRoleService settingssvc.RoleService
	// OIDCProviderFunc to lazily initialize an oidc provider, must be set for the oidc_auth middleware
	OIDCProviderFunc func() (OIDCProvider, error)
	// OIDCIss is the oidcAuth-issuer
	OIDCIss string
	// RevaGatewayClient to send requests to the reva gateway
	RevaGatewayClient gateway.GatewayAPIClient
	// Store for persisting data
	Store storesvc.StoreService
	// PreSignedURLConfig to configure the middleware
	PreSignedURLConfig config.PreSignedURL
	// UserOIDCClaim to read from the oidc claims
	UserOIDCClaim string
	// UserCS3Claim to use when looking up a user in the CS3 API
	UserCS3Claim string
	// AutoprovisionAccounts when an accountResolver does not exist.
	AutoprovisionAccounts bool
	// EnableBasicAuth to allow basic auth
	EnableBasicAuth bool
	// UserinfoCacheSize defines the max number of entries in the userinfo cache, intended for the oidc_auth middleware
	UserinfoCacheSize int
	// UserinfoCacheTTL sets the max cache duration for the userinfo cache, intended for the oidc_auth middleware
	UserinfoCacheTTL time.Duration
	// CredentialsByUserAgent sets the auth challenges on a per user-agent basis
	CredentialsByUserAgent map[string]string
	// AccessTokenVerifyMethod configures how access_tokens should be verified but the oidc_auth middleware.
	// Possible values currently: "jwt" and "none"
	AccessTokenVerifyMethod string
	// JWKS sets the options for fetching the JWKS from the IDP
	JWKS config.JWKS
	// RoleQuotas hold userid:quota mappings. These will be used when provisioning new users.
	// The users will get as much quota as is set for their role.
	RoleQuotas map[string]uint64
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Logger provides a function to set the logger option.
func Logger(l log.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

// TokenManagerConfig provides a function to set the token manger config option.
func TokenManagerConfig(cfg config.TokenManager) Option {
	return func(o *Options) {
		o.TokenManagerConfig = cfg
	}
}

// PolicySelectorConfig provides a function to set the policy selector config option.
func PolicySelectorConfig(cfg config.PolicySelector) Option {
	return func(o *Options) {
		o.PolicySelector = cfg
	}
}

// HTTPClient provides a function to set the http client config option.
func HTTPClient(c *http.Client) Option {
	return func(o *Options) {
		o.HTTPClient = c
	}
}

// SettingsRoleService provides a function to set the role service option.
func SettingsRoleService(rc settingssvc.RoleService) Option {
	return func(o *Options) {
		o.SettingsRoleService = rc
	}
}

// OIDCProviderFunc provides a function to set the the oidc provider function option.
func OIDCProviderFunc(f func() (OIDCProvider, error)) Option {
	return func(o *Options) {
		o.OIDCProviderFunc = f
	}
}

// OIDCIss sets the oidcAuth issuer url
func OIDCIss(iss string) Option {
	return func(o *Options) {
		o.OIDCIss = iss
	}
}

// CredentialsByUserAgent sets UserAgentChallenges.
func CredentialsByUserAgent(v map[string]string) Option {
	return func(o *Options) {
		o.CredentialsByUserAgent = v
	}
}

// RevaGatewayClient provides a function to set the the reva gateway service client option.
func RevaGatewayClient(gc gateway.GatewayAPIClient) Option {
	return func(o *Options) {
		o.RevaGatewayClient = gc
	}
}

// Store provides a function to set the store option.
func Store(sc storesvc.StoreService) Option {
	return func(o *Options) {
		o.Store = sc
	}
}

// PreSignedURLConfig provides a function to set the PreSignedURL config
func PreSignedURLConfig(cfg config.PreSignedURL) Option {
	return func(o *Options) {
		o.PreSignedURLConfig = cfg
	}
}

// UserOIDCClaim provides a function to set the UserClaim config
func UserOIDCClaim(val string) Option {
	return func(o *Options) {
		o.UserOIDCClaim = val
	}
}

// UserCS3Claim provides a function to set the UserClaimType config
func UserCS3Claim(val string) Option {
	return func(o *Options) {
		o.UserCS3Claim = val
	}
}

// AutoprovisionAccounts provides a function to set the AutoprovisionAccounts config
func AutoprovisionAccounts(val bool) Option {
	return func(o *Options) {
		o.AutoprovisionAccounts = val
	}
}

// EnableBasicAuth provides a function to set the EnableBasicAuth config
func EnableBasicAuth(enableBasicAuth bool) Option {
	return func(o *Options) {
		o.EnableBasicAuth = enableBasicAuth
	}
}

// TokenCacheSize provides a function to set the TokenCacheSize
func TokenCacheSize(size int) Option {
	return func(o *Options) {
		o.UserinfoCacheSize = size
	}
}

// TokenCacheTTL provides a function to set the TokenCacheTTL
func TokenCacheTTL(ttl time.Duration) Option {
	return func(o *Options) {
		o.UserinfoCacheTTL = ttl
	}
}

// UserProvider sets the accounts user provider
func UserProvider(up backend.UserBackend) Option {
	return func(o *Options) {
		o.UserProvider = up
	}
}

// AccessTokenVerifyMethod set the mechanism for access token verification
func AccessTokenVerifyMethod(method string) Option {
	return func(o *Options) {
		o.AccessTokenVerifyMethod = method
	}
}

// JWKSOptions sets the options for fetching the JWKS from the IDP
func JWKSOptions(jo config.JWKS) Option {
	return func(o *Options) {
		o.JWKS = jo
	}
}

// RoleQuotas sets the role quota mapping setting
func RoleQuotas(roleQuotas map[string]uint64) Option {
	return func(o *Options) {
		o.RoleQuotas = roleQuotas
	}
}
