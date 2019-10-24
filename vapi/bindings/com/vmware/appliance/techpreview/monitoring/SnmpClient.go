/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Interface file for service: Snmp
 * Used by client-side stubs.
 */
package monitoring


// ``Snmp`` interface provides methods SNMP agent operations.
type SnmpClient interface {

    // Restore settings to factory defaults.
    // @throws Error Generic error
	Reset() error

    // Start a disabled SNMP agent.
    // @throws Error Generic error
	Enable() error

    // Generate localized keys for secure SNMPv3 communications.
    //
    // @param configParam SNMP hash configuration.
    // @return SNMP hash result
    // @throws Error Generic error
	Hash(configParam SnmpSNMPHashConfig) (SnmpSNMPHashResults, error)

    // Get SNMP limits information.
    // @return SNMP limits structure
    // @throws Error Generic error
	Limits() (SnmpSNMPLimits, error)

    // Return an SNMP agent configuration.
    // @return SNMP config structure
    // @throws Error Generic error
	Get() (SnmpSNMPConfigReadOnly, error)

    // Stop an enabled SNMP agent.
    // @throws Error Generic error
	Disable() error

    // Set SNMP configuration.
    //
    // @param configParam SNMP configuration.
    // @throws Error Generic error
	Set(configParam SnmpSNMPConfig) error

    // Send a warmStart notification to all configured traps and inform destinations (see RFC 3418).
    // @return SNMP test result
    // @throws Error Generic error
	Test() (SnmpSNMPTestResults, error)

    // Generate diagnostics report for snmp agent.
    // @return SNMP stats
    // @throws Error Generic error
	Stats() (SnmpSNMPStats, error)
}
