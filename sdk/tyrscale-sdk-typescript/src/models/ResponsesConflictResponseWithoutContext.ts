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
 * @interface ResponsesConflictResponseWithoutContext
 */
export interface ResponsesConflictResponseWithoutContext {
    /**
     * 
     * @type {number}
     * @memberof ResponsesConflictResponseWithoutContext
     */
    code?: number;
    /**
     * 
     * @type {string}
     * @memberof ResponsesConflictResponseWithoutContext
     */
    message?: string;
    /**
     * 
     * @type {number}
     * @memberof ResponsesConflictResponseWithoutContext
     */
    status?: number;
}

/**
 * Check if a given object implements the ResponsesConflictResponseWithoutContext interface.
 */
export function instanceOfResponsesConflictResponseWithoutContext(value: object): value is ResponsesConflictResponseWithoutContext {
    return true;
}

export function ResponsesConflictResponseWithoutContextFromJSON(json: any): ResponsesConflictResponseWithoutContext {
    return ResponsesConflictResponseWithoutContextFromJSONTyped(json, false);
}

export function ResponsesConflictResponseWithoutContextFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResponsesConflictResponseWithoutContext {
    if (json == null) {
        return json;
    }
    return {
        
        'code': json['code'] == null ? undefined : json['code'],
        'message': json['message'] == null ? undefined : json['message'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}

export function ResponsesConflictResponseWithoutContextToJSON(value?: ResponsesConflictResponseWithoutContext | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'code': value['code'],
        'message': value['message'],
        'status': value['status'],
    };
}

