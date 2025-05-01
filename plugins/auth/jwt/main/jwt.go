package main

import (
	"filepatrol/internal/plugins"
	"filepatrol/plugins/auth/jwt/plugin"
)

// nolint:unused
func NewAuth() plugins.Auth { return &plugin_jwt.JWTAuth{} }
