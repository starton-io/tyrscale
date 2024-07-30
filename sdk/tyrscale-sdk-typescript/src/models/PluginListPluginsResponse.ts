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
import type { PluginPluginList } from './PluginPluginList';
import {
    PluginPluginListFromJSON,
    PluginPluginListFromJSONTyped,
    PluginPluginListToJSON,
} from './PluginPluginList';

/**
 * 
 * @export
 * @interface PluginListPluginsResponse
 */
export interface PluginListPluginsResponse {
    /**
     * Map of plugin lists, keyed by an integer ID
     * @type {{ [key: string]: PluginPluginList; }}
     * @memberof PluginListPluginsResponse
     */
    plugins?: { [key: string]: PluginPluginList; };
}

/**
 * Check if a given object implements the PluginListPluginsResponse interface.
 */
export function instanceOfPluginListPluginsResponse(value: object): value is PluginListPluginsResponse {
    return true;
}

export function PluginListPluginsResponseFromJSON(json: any): PluginListPluginsResponse {
    return PluginListPluginsResponseFromJSONTyped(json, false);
}

export function PluginListPluginsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PluginListPluginsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'plugins': json['plugins'] == null ? undefined : (mapValues(json['plugins'], PluginPluginListFromJSON)),
    };
}

export function PluginListPluginsResponseToJSON(value?: PluginListPluginsResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'plugins': value['plugins'] == null ? undefined : (mapValues(value['plugins'], PluginPluginListToJSON)),
    };
}
