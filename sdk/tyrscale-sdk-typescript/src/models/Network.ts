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
/**
 * 
 * @export
 * @interface Network
 */
export interface Network {
    /**
     * 
     * @type {string}
     * @memberof Network
     */
    blockchain: string;
    /**
     * 
     * @type {number}
     * @memberof Network
     */
    chainId: number;
    /**
     * 
     * @type {string}
     * @memberof Network
     */
    name: string;
}

/**
 * Check if a given object implements the Network interface.
 */
export function instanceOfNetwork(value: object): value is Network {
    if (!('blockchain' in value) || value['blockchain'] === undefined) return false;
    if (!('chainId' in value) || value['chainId'] === undefined) return false;
    if (!('name' in value) || value['name'] === undefined) return false;
    return true;
}

export function NetworkFromJSON(json: any): Network {
    return NetworkFromJSONTyped(json, false);
}

export function NetworkFromJSONTyped(json: any, ignoreDiscriminator: boolean): Network {
    if (json == null) {
        return json;
    }
    return {
        
        'blockchain': json['blockchain'],
        'chainId': json['chain_id'],
        'name': json['name'],
    };
}

export function NetworkToJSON(value?: Network | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'blockchain': value['blockchain'],
        'chain_id': value['chainId'],
        'name': value['name'],
    };
}

