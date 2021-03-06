// Package filesystem implements the Azure ARM Filesystem service API version 2016-11-01.
//
// Creates an Azure Data Lake Store filesystem client.
package filesystem

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultAdlsFileSystemDNSSuffix is the default value for adls file system dns suffix
	DefaultAdlsFileSystemDNSSuffix = "azuredatalakestore.net"
)

// BaseClient is the base client for Filesystem.
type BaseClient struct {
	autorest.Client
	AdlsFileSystemDNSSuffix string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithoutDefaults(DefaultAdlsFileSystemDNSSuffix)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(adlsFileSystemDNSSuffix string) BaseClient {
	return BaseClient{
		Client:                  autorest.NewClientWithUserAgent(UserAgent()),
		AdlsFileSystemDNSSuffix: adlsFileSystemDNSSuffix,
	}
}
