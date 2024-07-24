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
import type { Plugin } from './Plugin';
import {
    PluginFromJSON,
    PluginFromJSONTyped,
    PluginToJSON,
} from './Plugin';

/**
 * 
 * @export
 * @interface Plugins
 */
export interface Plugins {
    /**
     * 
     * @type {Array<Plugin>}
     * @memberof Plugins
     */
    interceptorRequest?: Array<Plugin>;
    /**
     * 
     * @type {Array<Plugin>}
     * @memberof Plugins
     */
    interceptorResponse?: Array<Plugin>;
    /**
     * 
     * @type {Array<Plugin>}
     * @memberof Plugins
     */
    middleware?: Array<Plugin>;
}

/**
 * Check if a given object implements the Plugins interface.
 */
export function instanceOfPlugins(value: object): value is Plugins {
    return true;
}

export function PluginsFromJSON(json: any): Plugins {
    return PluginsFromJSONTyped(json, false);
}

export function PluginsFromJSONTyped(json: any, ignoreDiscriminator: boolean): Plugins {
    if (json == null) {
        return json;
    }
    return {
        
        'interceptorRequest': json['interceptor_request'] == null ? undefined : ((json['interceptor_request'] as Array<any>).map(PluginFromJSON)),
        'interceptorResponse': json['interceptor_response'] == null ? undefined : ((json['interceptor_response'] as Array<any>).map(PluginFromJSON)),
        'middleware': json['middleware'] == null ? undefined : ((json['middleware'] as Array<any>).map(PluginFromJSON)),
    };
}

export function PluginsToJSON(value?: Plugins | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'interceptor_request': value['interceptorRequest'] == null ? undefined : ((value['interceptorRequest'] as Array<any>).map(PluginToJSON)),
        'interceptor_response': value['interceptorResponse'] == null ? undefined : ((value['interceptorResponse'] as Array<any>).map(PluginToJSON)),
        'middleware': value['middleware'] == null ? undefined : ((value['middleware'] as Array<any>).map(PluginToJSON)),
    };
}

