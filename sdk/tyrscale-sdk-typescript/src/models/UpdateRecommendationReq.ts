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
import type { StrategyName } from './StrategyName';
import {
    StrategyNameFromJSON,
    StrategyNameFromJSONTyped,
    StrategyNameToJSON,
} from './StrategyName';

/**
 * 
 * @export
 * @interface UpdateRecommendationReq
 */
export interface UpdateRecommendationReq {
    /**
     * 
     * @type {string}
     * @memberof UpdateRecommendationReq
     */
    networkName: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateRecommendationReq
     */
    routeUuid: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateRecommendationReq
     */
    schedule: string;
    /**
     * 
     * @type {StrategyName}
     * @memberof UpdateRecommendationReq
     */
    strategy: StrategyName;
}

/**
 * Check if a given object implements the UpdateRecommendationReq interface.
 */
export function instanceOfUpdateRecommendationReq(value: object): value is UpdateRecommendationReq {
    if (!('networkName' in value) || value['networkName'] === undefined) return false;
    if (!('routeUuid' in value) || value['routeUuid'] === undefined) return false;
    if (!('schedule' in value) || value['schedule'] === undefined) return false;
    if (!('strategy' in value) || value['strategy'] === undefined) return false;
    return true;
}

export function UpdateRecommendationReqFromJSON(json: any): UpdateRecommendationReq {
    return UpdateRecommendationReqFromJSONTyped(json, false);
}

export function UpdateRecommendationReqFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateRecommendationReq {
    if (json == null) {
        return json;
    }
    return {
        
        'networkName': json['network_name'],
        'routeUuid': json['route_uuid'],
        'schedule': json['schedule'],
        'strategy': StrategyNameFromJSON(json['strategy']),
    };
}

export function UpdateRecommendationReqToJSON(value?: UpdateRecommendationReq | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'network_name': value['networkName'],
        'route_uuid': value['routeUuid'],
        'schedule': value['schedule'],
        'strategy': StrategyNameToJSON(value['strategy']),
    };
}

