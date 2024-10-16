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
 * @interface CreateRpcRes
 */
export interface CreateRpcRes {
    /**
     * 
     * @type {string}
     * @memberof CreateRpcRes
     */
    uuid?: string;
}

/**
 * Check if a given object implements the CreateRpcRes interface.
 */
export function instanceOfCreateRpcRes(value: object): value is CreateRpcRes {
    return true;
}

export function CreateRpcResFromJSON(json: any): CreateRpcRes {
    return CreateRpcResFromJSONTyped(json, false);
}

export function CreateRpcResFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateRpcRes {
    if (json == null) {
        return json;
    }
    return {
        
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
    };
}

export function CreateRpcResToJSON(value?: CreateRpcRes | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'uuid': value['uuid'],
    };
}

