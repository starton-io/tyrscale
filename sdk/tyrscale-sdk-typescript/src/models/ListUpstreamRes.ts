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
import type { Upstream } from './Upstream';
import {
    UpstreamFromJSON,
    UpstreamFromJSONTyped,
    UpstreamToJSON,
} from './Upstream';

/**
 * 
 * @export
 * @interface ListUpstreamRes
 */
export interface ListUpstreamRes {
    /**
     * 
     * @type {Array<Upstream>}
     * @memberof ListUpstreamRes
     */
    items?: Array<Upstream>;
}

/**
 * Check if a given object implements the ListUpstreamRes interface.
 */
export function instanceOfListUpstreamRes(value: object): value is ListUpstreamRes {
    return true;
}

export function ListUpstreamResFromJSON(json: any): ListUpstreamRes {
    return ListUpstreamResFromJSONTyped(json, false);
}

export function ListUpstreamResFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListUpstreamRes {
    if (json == null) {
        return json;
    }
    return {
        
        'items': json['items'] == null ? undefined : ((json['items'] as Array<any>).map(UpstreamFromJSON)),
    };
}

export function ListUpstreamResToJSON(value?: ListUpstreamRes | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'items': value['items'] == null ? undefined : ((value['items'] as Array<any>).map(UpstreamToJSON)),
    };
}

