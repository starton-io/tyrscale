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
    middleware?: Array<Plugin>;
    /**
     * 
     * @type {Array<Plugin>}
     * @memberof Plugins
     */
    requestInterceptor?: Array<Plugin>;
    /**
     * 
     * @type {Array<Plugin>}
     * @memberof Plugins
     */
    responseInterceptor?: Array<Plugin>;
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
        
        'middleware': json['Middleware'] == null ? undefined : ((json['Middleware'] as Array<any>).map(PluginFromJSON)),
        'requestInterceptor': json['RequestInterceptor'] == null ? undefined : ((json['RequestInterceptor'] as Array<any>).map(PluginFromJSON)),
        'responseInterceptor': json['ResponseInterceptor'] == null ? undefined : ((json['ResponseInterceptor'] as Array<any>).map(PluginFromJSON)),
    };
}

export function PluginsToJSON(value?: Plugins | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'Middleware': value['middleware'] == null ? undefined : ((value['middleware'] as Array<any>).map(PluginToJSON)),
        'RequestInterceptor': value['requestInterceptor'] == null ? undefined : ((value['requestInterceptor'] as Array<any>).map(PluginToJSON)),
        'ResponseInterceptor': value['responseInterceptor'] == null ? undefined : ((value['responseInterceptor'] as Array<any>).map(PluginToJSON)),
    };
}
