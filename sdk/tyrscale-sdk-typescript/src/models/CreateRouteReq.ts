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
import type { HealthcheckHealthCheckConfig } from './HealthcheckHealthCheckConfig';
import {
    HealthcheckHealthCheckConfigFromJSON,
    HealthcheckHealthCheckConfigFromJSONTyped,
    HealthcheckHealthCheckConfigToJSON,
} from './HealthcheckHealthCheckConfig';
import type { BalancerLoadBalancerStrategy } from './BalancerLoadBalancerStrategy';
import {
    BalancerLoadBalancerStrategyFromJSON,
    BalancerLoadBalancerStrategyFromJSONTyped,
    BalancerLoadBalancerStrategyToJSON,
} from './BalancerLoadBalancerStrategy';
import type { CircuitbreakerSettings } from './CircuitbreakerSettings';
import {
    CircuitbreakerSettingsFromJSON,
    CircuitbreakerSettingsFromJSONTyped,
    CircuitbreakerSettingsToJSON,
} from './CircuitbreakerSettings';

/**
 * 
 * @export
 * @interface CreateRouteReq
 */
export interface CreateRouteReq {
    /**
     * 
     * @type {CircuitbreakerSettings}
     * @memberof CreateRouteReq
     */
    circuitBreaker?: CircuitbreakerSettings;
    /**
     * 
     * @type {HealthcheckHealthCheckConfig}
     * @memberof CreateRouteReq
     */
    healthCheck?: HealthcheckHealthCheckConfig;
    /**
     * 
     * @type {string}
     * @memberof CreateRouteReq
     */
    host: string;
    /**
     * 
     * @type {BalancerLoadBalancerStrategy}
     * @memberof CreateRouteReq
     */
    loadBalancerStrategy: BalancerLoadBalancerStrategy;
    /**
     * 
     * @type {string}
     * @memberof CreateRouteReq
     */
    path?: string;
    /**
     * 
     * @type {string}
     * @memberof CreateRouteReq
     */
    uuid?: string;
}

/**
 * Check if a given object implements the CreateRouteReq interface.
 */
export function instanceOfCreateRouteReq(value: object): value is CreateRouteReq {
    if (!('host' in value) || value['host'] === undefined) return false;
    if (!('loadBalancerStrategy' in value) || value['loadBalancerStrategy'] === undefined) return false;
    return true;
}

export function CreateRouteReqFromJSON(json: any): CreateRouteReq {
    return CreateRouteReqFromJSONTyped(json, false);
}

export function CreateRouteReqFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateRouteReq {
    if (json == null) {
        return json;
    }
    return {
        
        'circuitBreaker': json['circuit_breaker'] == null ? undefined : CircuitbreakerSettingsFromJSON(json['circuit_breaker']),
        'healthCheck': json['health_check'] == null ? undefined : HealthcheckHealthCheckConfigFromJSON(json['health_check']),
        'host': json['host'],
        'loadBalancerStrategy': BalancerLoadBalancerStrategyFromJSON(json['load_balancer_strategy']),
        'path': json['path'] == null ? undefined : json['path'],
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
    };
}

export function CreateRouteReqToJSON(value?: CreateRouteReq | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'circuit_breaker': CircuitbreakerSettingsToJSON(value['circuitBreaker']),
        'health_check': HealthcheckHealthCheckConfigToJSON(value['healthCheck']),
        'host': value['host'],
        'load_balancer_strategy': BalancerLoadBalancerStrategyToJSON(value['loadBalancerStrategy']),
        'path': value['path'],
        'uuid': value['uuid'],
    };
}
