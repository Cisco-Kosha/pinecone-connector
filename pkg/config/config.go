package config

import (
	"flag"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	apiKey           string
	apiKeyHeaderName string
	serverUrl        string
	authType         string
	environment      string
}

func Get() *Config {
	conf := &Config{}
	// creating a new flagset everytime the get function is called allows for different flagsets to exist
	// rather than a conflict to be created when generating a new config object (such as for tests)
	flags := flag.NewFlagSet("Passthrough Flag Set", flag.PanicOnError)

	flags.StringVar(&conf.authType, "authType", os.Getenv("AUTH_TYPE"), "Auth Type")
	flags.StringVar(&conf.apiKeyHeaderName, "apiKeyHeaderName", os.Getenv("API_KEY_HEADER_NAME"), "API Key Header Name")
	flags.StringVar(&conf.apiKey, "apiKey", os.Getenv("API_KEY"), "API Key")
	flags.StringVar(&conf.serverUrl, "serverUrl", os.Getenv("SERVER_URL"), "Server Url")
	flags.StringVar(&conf.environment, "environment", os.Getenv("ENVIRONMENT"), "Pinecone controller environment")

	var arguments []string
	arguments = append(arguments, "os.Environ")
	flags.Parse(arguments)

	return conf
}

func (c *Config) GetApiKey() string {
	return c.apiKey
}

func (c *Config) GetEnv() string {
	return c.environment
}

func (c *Config) GetApiKeyHeaderName() string {
	return c.apiKeyHeaderName
}

// GetAuthType returns the auth type accepted by the server
// Possible values include: API_KEY, BASIC_AUTH, HMAC
func (c *Config) GetAuthType() string {
	// convert all characters to upper case
	authType := strings.ToUpper(c.authType)
	// replace space, hyphen with underscore
	authType = strings.ReplaceAll(authType, " ", "_")
	authType = strings.ReplaceAll(authType, "%20", "_")
	authType = strings.ReplaceAll(authType, "-", "_")
	return authType
}

func (c *Config) GetServerURL() string {
	c.serverUrl = strings.TrimSuffix(c.serverUrl, "/")
	u, _ := url.Parse(c.serverUrl)
	if u.Scheme == "" {
		return "https://" + c.serverUrl
	} else {
		return c.serverUrl
	}
}

func (c *Config) GetServerHost() string {
	c.serverUrl = strings.TrimSuffix(c.serverUrl, "/")
	u, _ := url.Parse(c.serverUrl)
	if u.Scheme == "" {
		return u.Host
	} else {
		return u.Host
	}
}

func (c *Config) GetServerPath() string {
	c.serverUrl = strings.TrimSuffix(c.serverUrl, "/")
	u, _ := url.Parse(c.serverUrl)
	if u.Scheme == "" {
		return u.Path
	} else {
		return u.Path
	}
}
