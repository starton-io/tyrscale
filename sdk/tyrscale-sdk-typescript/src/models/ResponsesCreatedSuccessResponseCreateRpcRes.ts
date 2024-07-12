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
import type { CreateRpcRes } from './CreateRpcRes';
import {
    CreateRpcResFromJSON,
    CreateRpcResFromJSONTyped,
    CreateRpcResToJSON,
} from './CreateRpcRes';

/**
 * 
 * @export
 * @interface ResponsesCreatedSuccessResponseCreateRpcRes
 */
export interface ResponsesCreatedSuccessResponseCreateRpcRes {
    /**
     * 
     * @type {number}
     * @memberof ResponsesCreatedSuccessResponseCreateRpcRes
     */
    code?: number;
    /**
     * 
     * @type {CreateRpcRes}
     * @memberof ResponsesCreatedSuccessResponseCreateRpcRes
     */
    data?: CreateRpcRes;
    /**
     * 
     * @type {string}
     * @memberof ResponsesCreatedSuccessResponseCreateRpcRes
     */
    message?: string;
    /**
     * 
     * @type {number}
     * @memberof ResponsesCreatedSuccessResponseCreateRpcRes
     */
    status?: number;
}

/**
 * Check if a given object implements the ResponsesCreatedSuccessResponseCreateRpcRes interface.
 */
export function instanceOfResponsesCreatedSuccessResponseCreateRpcRes(value: object): value is ResponsesCreatedSuccessResponseCreateRpcRes {
    return true;
}

export function ResponsesCreatedSuccessResponseCreateRpcResFromJSON(json: any): ResponsesCreatedSuccessResponseCreateRpcRes {
    return ResponsesCreatedSuccessResponseCreateRpcResFromJSONTyped(json, false);
}

export function ResponsesCreatedSuccessResponseCreateRpcResFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResponsesCreatedSuccessResponseCreateRpcRes {
    if (json == null) {
        return json;
    }
    return {
        
        'code': json['code'] == null ? undefined : json['code'],
        'data': json['data'] == null ? undefined : CreateRpcResFromJSON(json['data']),
        'message': json['message'] == null ? undefined : json['message'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}

export function ResponsesCreatedSuccessResponseCreateRpcResToJSON(value?: ResponsesCreatedSuccessResponseCreateRpcRes | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'code': value['code'],
        'data': CreateRpcResToJSON(value['data']),
        'message': value['message'],
        'status': value['status'],
    };
}
