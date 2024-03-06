// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

//go:generate mockgen -package client -destination=mocks.go github.com/gardener/gardener-extension-provider-gcp/pkg/gcp/client Factory,DNSClient,ComputeClient

package client
