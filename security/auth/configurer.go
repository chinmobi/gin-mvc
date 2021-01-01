// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/security"
)

// AuthenticationProcessorBase
type AuthenticationProcessorBase struct {
	manager      AuthManager
	authHandler  security.AuthHandler
}

type AuthProcessorBase = AuthenticationProcessorBase

func (ap *AuthProcessorBase) AuthManager() AuthManager {
	return ap.manager
}

func (ap *AuthProcessorBase) AuthHandler() security.AuthHandler {
	return ap.authHandler
}

// ProcessorConfigurer
type ProcessorConfigurer struct {
	processorBase  *AuthProcessorBase
	manager        ProviderManager
	handlerSet     security.AuthHandlerSet
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

	processor := &AuthProcessorBase{
		manager: &pc.manager,
		authHandler: &pc.handlerSet,
	}

	pc.processorBase = processor

	return processor
}

func (pc *ProcessorConfigurer) AddProvider(provider ...AuthProvider) {
	pc.manager.AddProvider(provider...)
}

func (pc *ProcessorConfigurer) AddSuccessFunc(onSuccess ...security.OnAuthSuccessFunc) {
	pc.handlerSet.AddSuccessFunc(onSuccess...)
}

func (pc *ProcessorConfigurer) AddFailureFunc(onFailure ...security.OnAuthFailureFunc) {
	pc.handlerSet.AddFailureFunc(onFailure...)
}

func (pc *ProcessorConfigurer) ProviderManager() *ProviderManager {
	return &pc.manager
}

func (pc *ProcessorConfigurer) AuthHandlerSet() *security.AuthHandlerSet {
	return &pc.handlerSet
}

func (pc *ProcessorConfigurer) Reset() {
	pc.manager.Clear()
	pc.handlerSet.Clear()
}
