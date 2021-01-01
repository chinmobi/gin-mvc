// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/security"
)

// AuthenticationProvider
type AuthenticationProvider interface {
	Authenticate(authentication security.Authentication) (security.Authentication, error)
	Supports(authentication security.Authentication) bool
}

type AuthProvider = AuthenticationProvider

// ProviderManager
type ProviderManager struct {
	providers []AuthProvider
}

func (pm *ProviderManager) GetProviders() []AuthProvider {
	return pm.providers
}

func (pm *ProviderManager) AddProvider(provider ...AuthProvider) {
	pm.providers = append(pm.providers, provider...)
}

func (pm *ProviderManager) Clear() {
	if pm.providers != nil {
		pm.providers = pm.providers[0:0]
	}
}

func (pm *ProviderManager) Authenticate(authentication security.Authentication) (security.Authentication, error) {
	for i, cnt := 0, len(pm.providers); i < cnt; i++ {
		provider := pm.providers[i]

		if !provider.Supports(authentication) {
			continue
		}

		authed, err := provider.Authenticate(authentication)
		if err != nil {
			return authentication, err
		}

		if authed != nil {
			if authed.IsAuthenticated() {
				return authed, nil
			}

			authentication = authed
		}
	}

	return authentication, nil
}
