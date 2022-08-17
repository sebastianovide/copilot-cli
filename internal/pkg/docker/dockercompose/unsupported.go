// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package dockercompose

// NOTE: This file is currently unused, it will be used in the next PR. Don't mind the comments for now!

// IgnoredKeys stores a list of keys in the Compose YAML that couldn't be processed,
// but are likely to not be significant enough to cause the converted application to
// fail. It's expected that this will eventually be displayed to the user.
type IgnoredKeys []string

/*
// ignoredServiceKeys lists out the keys on Compose services that are ignored in conversion.
//
// note: build keys are handled separately in convertBuildConfig
var ignoredServiceKeys = map[string]struct{}{
	"blkio_config":        {},
	"cpu_count":           {},
	"cpu_percent":         {},
	"cpu_shares":          {},
	"cpu_period":          {},
	"cpu_quota":           {},
	"cpu_rt_runtime":      {},
	"cpu_rt_period":       {},
	"cpus":                {},
	"cpuset":              {},
	"cap_add":             {},
	"cap_drop":            {},
	"cgroup_parent":       {},
	"device_cgroup_rules": {},
	"logging":             {},
	"mac_address":         {},
	"mem_limit":           {},
	"mem_reservation":     {},
	"mem_swappiness":      {},
	"memswap_limit":       {},
	"oom_kill_disable":    {},
	"oom_score_adj":       {},
	"pid":                 {},
	"pids_limit":          {},
	"profiles":            {},
	"pull_policy":         {},
	"runtime":             {},
	"security_opt":        {},
	"shm_size":            {},
	"stdin_open":          {},
	"storage_opt":         {},
	"sysctls":             {},
	"tmpfs":               {},
	"user":                {},
	"userns_mode":         {},
	"hostname":            {},
	"depends_on":          {},
	"restart":             {},
}

// fatalServiceKeys lists out the service keys that are unsupported and whose absence will
// break applications.
//
// note: build keys are handled separately in convertBuildConfig
// TODO(rclinard-amzn): Handle unsupported network keys when network conversion is implemented
var fatalServiceKeys = map[string]string{
	"credential_spec":   "",
	"devices":           "",
	"domainname":        "",
	"group_add":         "",
	"init":              "",
	"ipc":               "",
	"isolation":         "",
	"privileged":        "unsupported in Fargate",
	"external_links":    "",
	"working_dir":       "unsupported in Copilot manifests",
	"configs":           "unsupported, use secrets for similar functionality",
	"dns":               "unsupported in Copilot manifests",
	"dns_opt":           "unsupported in Copilot manifests",
	"dns_search":        "unsupported in Copilot manifests",
	"stop_grace_period": "unsupported in Copilot manifests",
	"stop_signal":       "unsupported in Copilot manifests",
	"volumes_from":      "sharing volumes is not yet supported",
}
*/