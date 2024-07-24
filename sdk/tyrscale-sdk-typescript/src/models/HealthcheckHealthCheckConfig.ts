/* tslint:disable */
/* eslint-disable */
/**
 * Tyrscale Manager API
 * This is the manager service for Tyrscale
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@starton.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { HealthcheckHealthCheckType } from './HealthcheckHealthCheckType';
import {
    HealthcheckHealthCheckTypeFromJSON,
    HealthcheckHealthCheckTypeFromJSONTyped,
    HealthcheckHealthCheckTypeToJSON,
} from './HealthcheckHealthCheckType';
import type { HealthcheckRequest } from './HealthcheckRequest';
import {
    HealthcheckRequestFromJSON,
    HealthcheckRequestFromJSONTyped,
    HealthcheckRequestToJSON,
} from './HealthcheckRequest';

/**
 * 
 * @export
 * @interface HealthcheckHealthCheckConfig
 */
export interface HealthcheckHealthCheckConfig {
    /**
     * 
     * @type {boolean}
     * @memberof HealthcheckHealthCheckConfig
     */
    combinedWithCircuitBreaker?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof HealthcheckHealthCheckConfig
     */
    enabled?: boolean;
    /**
     * 
     * @type {number}
     * @memberof HealthcheckHealthCheckConfig
     */
    interval?: number;
    /**
     * 
     * @type {HealthcheckRequest}
     * @memberof HealthcheckHealthCheckConfig
     */
    request?: HealthcheckRequest;
    /**
     * 
     * @type {number}
     * @memberof HealthcheckHealthCheckConfig
     */
    timeout?: number;
    /**
     * 
     * @type {HealthcheckHealthCheckType}
     * @memberof HealthcheckHealthCheckConfig
     */
    type?: HealthcheckHealthCheckType;
}

/**
 * Check if a given object implements the HealthcheckHealthCheckConfig interface.
 */
export function instanceOfHealthcheckHealthCheckConfig(value: object): value is HealthcheckHealthCheckConfig {
    return true;
}

export function HealthcheckHealthCheckConfigFromJSON(json: any): HealthcheckHealthCheckConfig {
    return HealthcheckHealthCheckConfigFromJSONTyped(json, false);
}

export function HealthcheckHealthCheckConfigFromJSONTyped(json: any, ignoreDiscriminator: boolean): HealthcheckHealthCheckConfig {
    if (json == null) {
        return json;
    }
    return {
        
        'combinedWithCircuitBreaker': json['combined_with_circuit_breaker'] == null ? undefined : json['combined_with_circuit_breaker'],
        'enabled': json['enabled'] == null ? undefined : json['enabled'],
        'interval': json['interval'] == null ? undefined : json['interval'],
        'request': json['request'] == null ? undefined : HealthcheckRequestFromJSON(json['request']),
        'timeout': json['timeout'] == null ? undefined : json['timeout'],
        'type': json['type'] == null ? undefined : HealthcheckHealthCheckTypeFromJSON(json['type']),
    };
}

export function HealthcheckHealthCheckConfigToJSON(value?: HealthcheckHealthCheckConfig | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'combined_with_circuit_breaker': value['combinedWithCircuitBreaker'],
        'enabled': value['enabled'],
        'interval': value['interval'],
        'request': HealthcheckRequestToJSON(value['request']),
        'timeout': value['timeout'],
        'type': HealthcheckHealthCheckTypeToJSON(value['type']),
    };
}

