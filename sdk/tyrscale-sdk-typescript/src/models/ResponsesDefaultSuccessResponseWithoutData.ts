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
 * @interface ResponsesDefaultSuccessResponseWithoutData
 */
export interface ResponsesDefaultSuccessResponseWithoutData {
    /**
     * 
     * @type {number}
     * @memberof ResponsesDefaultSuccessResponseWithoutData
     */
    code?: number;
    /**
     * 
     * @type {string}
     * @memberof ResponsesDefaultSuccessResponseWithoutData
     */
    message?: string;
    /**
     * 
     * @type {number}
     * @memberof ResponsesDefaultSuccessResponseWithoutData
     */
    status?: number;
}

/**
 * Check if a given object implements the ResponsesDefaultSuccessResponseWithoutData interface.
 */
export function instanceOfResponsesDefaultSuccessResponseWithoutData(value: object): value is ResponsesDefaultSuccessResponseWithoutData {
    return true;
}

export function ResponsesDefaultSuccessResponseWithoutDataFromJSON(json: any): ResponsesDefaultSuccessResponseWithoutData {
    return ResponsesDefaultSuccessResponseWithoutDataFromJSONTyped(json, false);
}

export function ResponsesDefaultSuccessResponseWithoutDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResponsesDefaultSuccessResponseWithoutData {
    if (json == null) {
        return json;
    }
    return {
        
        'code': json['code'] == null ? undefined : json['code'],
        'message': json['message'] == null ? undefined : json['message'],
        'status': json['status'] == null ? undefined : json['status'],
    };
}

export function ResponsesDefaultSuccessResponseWithoutDataToJSON(value?: ResponsesDefaultSuccessResponseWithoutData | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'code': value['code'],
        'message': value['message'],
        'status': value['status'],
    };
}

