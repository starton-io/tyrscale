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
import type { Recommendation } from './Recommendation';
import {
    RecommendationFromJSON,
    RecommendationFromJSONTyped,
    RecommendationToJSON,
} from './Recommendation';

/**
 * 
 * @export
 * @interface ListRecommendationRes
 */
export interface ListRecommendationRes {
    /**
     * 
     * @type {Array<Recommendation>}
     * @memberof ListRecommendationRes
     */
    items?: Array<Recommendation>;
}

/**
 * Check if a given object implements the ListRecommendationRes interface.
 */
export function instanceOfListRecommendationRes(value: object): value is ListRecommendationRes {
    return true;
}

export function ListRecommendationResFromJSON(json: any): ListRecommendationRes {
    return ListRecommendationResFromJSONTyped(json, false);
}

export function ListRecommendationResFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListRecommendationRes {
    if (json == null) {
        return json;
    }
    return {
        
        'items': json['items'] == null ? undefined : ((json['items'] as Array<any>).map(RecommendationFromJSON)),
    };
}

export function ListRecommendationResToJSON(value?: ListRecommendationRes | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'items': value['items'] == null ? undefined : ((value['items'] as Array<any>).map(RecommendationToJSON)),
    };
}

