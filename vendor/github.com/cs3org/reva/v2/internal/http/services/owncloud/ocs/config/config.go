// Copyright 2018-2021 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package config

import (
	"github.com/cs3org/reva/v2/internal/http/services/owncloud/ocs/data"
	"github.com/cs3org/reva/v2/pkg/sharedconf"
)

// Config holds the config options that need to be passed down to all ocs handlers
type Config struct {
	Prefix                                string                            `mapstructure:"prefix"`
	Config                                data.ConfigData                   `mapstructure:"config"`
	Capabilities                          data.CapabilitiesData             `mapstructure:"capabilities"`
	GatewaySvc                            string                            `mapstructure:"gatewaysvc"`
	StorageregistrySvc                    string                            `mapstructure:"storage_registry_svc"`
	DefaultUploadProtocol                 string                            `mapstructure:"default_upload_protocol"`
	UserAgentChunkingMap                  map[string]string                 `mapstructure:"user_agent_chunking_map"`
	SharePrefix                           string                            `mapstructure:"share_prefix"`
	HomeNamespace                         string                            `mapstructure:"home_namespace"`
	AdditionalInfoAttribute               string                            `mapstructure:"additional_info_attribute"`
	CacheWarmupDriver                     string                            `mapstructure:"cache_warmup_driver"`
	CacheWarmupDrivers                    map[string]map[string]interface{} `mapstructure:"cache_warmup_drivers"`
	ResourceInfoCacheStore                string                            `mapstructure:"resource_info_cache_store"`
	ResourceInfoCacheNodes                []string                          `mapstructure:"resource_info_cache_nodes"`
	ResourceInfoCacheDatabase             string                            `mapstructure:"resource_info_cache_database"`
	ResourceInfoCacheTable                string                            `mapstructure:"resource_info_cache_table"`
	ResourceInfoCacheTTL                  int                               `mapstructure:"resource_info_cache_ttl"`
	ResourceInfoCacheSize                 int                               `mapstructure:"resource_info_cache_size"`
	UserIdentifierCacheTTL                int                               `mapstructure:"user_identifier_cache_ttl"`
	MachineAuthAPIKey                     string                            `mapstructure:"machine_auth_apikey"`
	SkipUpdatingExistingSharesMountpoints bool                              `mapstructure:"skip_updating_existing_shares_mountpoint"`
	EnableDenials                         bool                              `mapstructure:"enable_denials"`
}

// Init sets sane defaults
func (c *Config) Init() {
	if c.Prefix == "" {
		c.Prefix = "ocs"
	}

	if c.SharePrefix == "" {
		c.SharePrefix = "/Shares"
	}

	if c.DefaultUploadProtocol == "" {
		c.DefaultUploadProtocol = "tus"
	}

	if c.HomeNamespace == "" {
		c.HomeNamespace = "/users/{{.Id.OpaqueId}}"
	}

	if c.AdditionalInfoAttribute == "" {
		c.AdditionalInfoAttribute = "{{.Mail}}"
	}

	if c.UserIdentifierCacheTTL == 0 {
		c.UserIdentifierCacheTTL = 60
	}

	c.GatewaySvc = sharedconf.GetGatewaySVC(c.GatewaySvc)
}
