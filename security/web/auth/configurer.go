// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/security/auth"
)

// AuthenticationProcessorBase
type AuthenticationProcessorBase struct {
	manager      auth.AuthManager
	authHandler  AuthHandler
}

type AuthProcessorBase = AuthenticationProcessorBase

func (ap *AuthProcessorBase) AuthManager() auth.AuthManager {
	return ap.manager
}

func (ap *AuthProcessorBase) AuthHandler() AuthHandler {
	return ap.authHandler
}

// ProcessorConfigurer
type ProcessorConfigurer struct {
	manager        auth.ProviderManager
	handlerSet     AuthHandlerSet
	processorBase  *AuthProcessorBase
}

func NewProcessorConfigurer() *ProcessorConfigurer {
	configurer := &ProcessorConfigurer{
	}
	return configurer
}

func (pc *ProcessorConfigurer) ProcessorBase() *AuthProcessorBase {
	if pc.processorBase != nil {
		return pc.processorBase
	}

	base := &AuthProcessorBase{
		manager: &pc.manager,
		authHandler: &pc.handlerSet,
	}

	pc.processorBase = base

	return base
}

func (pc *ProcessorConfigurer) AddProvider(provider ...auth.AuthProvider) {
	pc.manager.AddProvider(provider...)
}

func (pc *ProcessorConfigurer) AddSuccessFunc(onSuccess ...OnAuthSuccessFunc) {
	pc.handlerSet.AddSuccessFunc(onSuccess...)
}

func (pc *ProcessorConfigurer) AddFailureFunc(onFailure ...OnAuthFailureFunc) {
	pc.handlerSet.AddFailureFunc(onFailure...)
}

func (pc *ProcessorConfigurer) ProviderManager() *auth.ProviderManager {
	return &pc.manager
}

func (pc *ProcessorConfigurer) AuthHandlerSet() *AuthHandlerSet {
	return &pc.handlerSet
}

func (pc *ProcessorConfigurer) Reset() {
	pc.manager.Clear()
	pc.handlerSet.Clear()
}
