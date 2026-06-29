// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package fastic

import (
	"github.com/amodemoli/fastic/core/tools"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// env file load structure =D
type Env struct {
	// on this section will saves .env file find status. =D
	LoadedFile bool

	// global and general var's
	DevelopemtMode bool
	Port           int // serv amd listen port

	// this section is for website cors
	AllowedDomains map[string]bool
	IsWildcard     bool
	AllowedMethods string
	AllowedHeaders string
	MaxAge         int

	// security header's section
	XFrameOptions           string
	ReferrerPolicy          string
	ContentSecurityPolicy   string
	StrictTransportSecurity string

	// server settions values
	MaxConnsPerIP      int
	MaxRequestBodySize int
	IdleTimeout        int // per seccond...
	ReadTimeout        int // per seccond...
	WriteTimeout       int // per seccond...
	Concurrency        int
	DisableKeepalive   bool
}

// this function maded for read and load env files and save env values to Env struct.
func LoadEnv() *Env {
	env := &Env{
		// set loaded file to true as default value
		LoadedFile: true,
	} // create new model of Env struct and save on env var

	// load .env file.
	if err := godotenv.Load(); err != nil {
		// if cannot load env file update env.LoadedFile value to false
		env.LoadedFile = false
	}

	// if "DEVELOPMENT_MODE" value is true return's true else if "DEVELOPMENT_MODE" is nil("") set default value to false
	env.DevelopemtMode = tools.ReadEnvValue("DEVELOPMENT_MODE", "false") == "true"

	// search for "PORT" value on env file. verify port number and chnage port as string to int,
	//  if port dont have default port patter, port set to 8000.
	port, err := strconv.Atoi(tools.ReadEnvValue("PORT", "8000"))
	if err != nil || port < 1 || port > 65535 { // pattern for change port to number and verify port number
		port = 8000 // set default port number
	}
	// save port number to struct =D
	env.Port = port

	// get allowed domains for cors orgin, this value is raw
	raw := tools.ReadEnvValue("ALLOWED_DOMAINS", "")

	// if raw domains list value is "*", user need to accept all domains for api
	if raw == "*" {
		env.IsWildcard = true
		env.AllowedDomains = nil
	} else if raw == "" && !env.DevelopemtMode { // if they cannot find raw value and development mode is off, they block all domains for security
		env.AllowedDomains = make(map[string]bool)
	} else { // if raw value has domain they get domain name or domains range and save it to struct
		env.AllowedDomains = make(map[string]bool) // create a map
		// split the domains and get domains
		for _, part := range strings.Split(raw, ",") {
			// remove suffix/prefix spaces from domains
			domain := strings.TrimSpace(part)
			if domain != "" { // if domain is nil dont save to map, else? save it on map
				env.AllowedDomains[domain] = true
			}
		}
	}

	// set other value's to struct, read values and set default values and save it to thyer struct section =D
	env.AllowedMethods = tools.ReadEnvValue("ALLOWED_METHODS", "GET, POST, PUT, DELETE, OPTIONS")
	env.AllowedHeaders = tools.ReadEnvValue("ALLOWED_HEADERS", "Content-Type, Authorization")
	// if cannot change string to int set's default value for MaxAge.
	if env.MaxAge, err = strconv.Atoi(tools.ReadEnvValue("MAX_AGE", "86400")); err != nil {
		env.MaxAge = 86400
	}
	env.XFrameOptions = tools.ReadEnvValue("X_FRAME_OPTIONS", "DENY")
	env.ReferrerPolicy = tools.ReadEnvValue("REFERRER_POLICY", "strict-origin-when-cross-origin")
	env.ContentSecurityPolicy = tools.ReadEnvValue("CONTENT_SECURITY_POLICY", "default-src 'self'")
	env.StrictTransportSecurity = tools.ReadEnvValue("STRICT_TRANSPORT_SECURITY", "max-age=63072000; includeSubDomains; preload")

	// write mac connetction per ip, change int to string
	env.MaxConnsPerIP, _ = strconv.Atoi(tools.ReadEnvValue("MAX_CONNS_PER_IP", "100"))
	if env.MaxConnsPerIP < 0 {
		env.MaxConnsPerIP = 0 // if number is - or 0 for unlimited
	}

	env.MaxRequestBodySize, _ = strconv.Atoi(tools.ReadEnvValue("MAX_REQUEST_BODY_SIZE", "4194304")) // 4 MG
	if env.MaxRequestBodySize < 0 {
		env.MaxRequestBodySize = 0 // if number is - or 0 means unlimited body size
	}

	env.IdleTimeout, _ = strconv.Atoi(tools.ReadEnvValue("IDLE_TIMEOUT", "60"))
	env.ReadTimeout, _ = strconv.Atoi(tools.ReadEnvValue("READ_TIMEOUT", "15"))
	env.WriteTimeout, _ = strconv.Atoi(tools.ReadEnvValue("WRITE_TIMEOUT", "15"))

	env.DisableKeepalive = tools.ReadEnvValue("DISABLE_KEEPALIVE", "false") == "true" // if DISABLE_KEEPALIVE is true returns true

	env.Concurrency, _ = strconv.Atoi(tools.ReadEnvValue("CONCURRENCY", "0"))
	if env.Concurrency < 0 {
		env.Concurrency = 0 // unlimited
	}

	return env
}

// this method maded for create ".env" file on your project,
// this method make's new ".env" file with default settings on caller function path. (call this function on main.go)
func (e *Env) Create() (ok bool) {

	// get the caller function path for create file on caller function path.
	_, filename, _, ok := runtime.Caller(1)
	if !ok { // if cannot get caller function path returns false
		return false
	}

	// use filepath.Dir for get full project path, for createing file
	projectRoot := filepath.Dir(filename)
	// next i used for filepath.Join for connect project root to .env file, this function maded for multy os (windos,linux,mac and more...)
	envPath := filepath.Join(projectRoot, ".env")

	// get file stats, dont need to close file only get error and drop path
	if _, err := os.Stat(envPath); err == nil {
		return true // if finded file, return true and back to caller function
	}

	// if cannot find file, start writing to file with 0644 perm,
	// this function need to byte for content (speedly)
	if err := os.WriteFile(envPath, []byte(EnvFileContent), 0644); err != nil {
		return false // if they have error (cannot write and create). return the application and return false value
	}

	return true // if created file return true =D
}
