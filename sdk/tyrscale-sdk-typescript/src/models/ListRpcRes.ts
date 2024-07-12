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
import type { Rpc } from './Rpc';
import {
    RpcFromJSON,
    RpcFromJSONTyped,
    RpcToJSON,
} from './Rpc';

/**
 * 
 * @export
 * @interface ListRpcRes
 */
export interface ListRpcRes {
    /**
     * 
     * @type {Array<Rpc>}
     * @memberof ListRpcRes
     */
    items?: Array<Rpc>;
}

/**
 * Check if a given object implements the ListRpcRes interface.
 */
export function instanceOfListRpcRes(value: object): value is ListRpcRes {
    return true;
}

export function ListRpcResFromJSON(json: any): ListRpcRes {
    return ListRpcResFromJSONTyped(json, false);
}

export function ListRpcResFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListRpcRes {
    if (json == null) {
        return json;
    }
    return {
        
        'items': json['items'] == null ? undefined : ((json['items'] as Array<any>).map(RpcFromJSON)),
    };
}

export function ListRpcResToJSON(value?: ListRpcRes | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'items': value['items'] == null ? undefined : ((value['items'] as Array<any>).map(RpcToJSON)),
    };
}

