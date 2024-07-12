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
import type { ListRpcRes } from './ListRpcRes';
import {
    ListRpcResFromJSON,
    ListRpcResFromJSONTyped,
    ListRpcResToJSON,
} from './ListRpcRes';

/**
 * 
 * @export
 * @interface ResponsesDefaultSuccessResponseListRpcRes
 */
export interface ResponsesDefaultSuccessResponseListRpcRes {
    /**
     * 
     * @type {number}
     * @memberof ResponsesDefaultSuccessResponseListRpcRes
     */
    code?: number;
    /**
     * 
     * @type {ListRpcRes}
     * @memberof ResponsesDefaultSuccessResponseListRpcRes
     */
    data?: ListRpcRes;
    /**
     * 
     * @type {string}
     * @memberof ResponsesDefaultSuccessResponseListRpcRes
     */
    message?: string;
    /**
     * 
     * @type {number}
     * @memberof ResponsesDefaultSuccessResponseListRpcRes
     */
    status?: number;
}

/**
 * Check if a given object implements the ResponsesDefaultSuccessResponseListRpcRes interface.
 */
export function instanceOfResponsesDefaultSuccessResponseListRpcRes(value: object): value is ResponsesDefaultSuccessResponseListRpcRes {
    return true;
}

export function ResponsesDefaultSuccessResponseListRpcResFromJSON(json: any): ResponsesDefaultSuccessResponseListRpcRes {
    return ResponsesDefaultSuccessResponseListRpcResFromJSONTyped(json, false);
}

export function ResponsesDefaultSuccessResponseListRpcResFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResponsesDefaultSuccessResponseListRpcRes {
    if (json == null) {
        return json;
    }
    return {
        
        'code': json['code'] == null ? undefined : json['code'],
        'data': json['data'] == null ? undefined : ListRpcResFromJSON(json['data']),
        'message': json['message'] == null ? undefined : json['message'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}

export function ResponsesDefaultSuccessResponseListRpcResToJSON(value?: ResponsesDefaultSuccessResponseListRpcRes | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'code': value['code'],
        'data': ListRpcResToJSON(value['data']),
        'message': value['message'],
        'status': value['status'],
    };
}

