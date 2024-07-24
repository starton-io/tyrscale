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


import * as runtime from '../runtime';
import type {
  ResponsesBadRequestResponse,
  ResponsesDefaultSuccessResponsePluginListPluginsResponse,
  ResponsesInternalServerErrorResponse,
} from '../models/index';
import {
    ResponsesBadRequestResponseFromJSON,
    ResponsesBadRequestResponseToJSON,
    ResponsesDefaultSuccessResponsePluginListPluginsResponseFromJSON,
    ResponsesDefaultSuccessResponsePluginListPluginsResponseToJSON,
    ResponsesInternalServerErrorResponseFromJSON,
    ResponsesInternalServerErrorResponseToJSON,
} from '../models/index';

/**
 * 
 */
export class PluginsApi extends runtime.BaseAPI {

    /**
     * Get list plugins
     * Get list plugins
     */
    async listPluginsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesDefaultSuccessResponsePluginListPluginsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/plugins`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesDefaultSuccessResponsePluginListPluginsResponseFromJSON(jsonValue));
    }

    /**
     * Get list plugins
     * Get list plugins
     */
    async listPlugins(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesDefaultSuccessResponsePluginListPluginsResponse> {
        const response = await this.listPluginsRaw(initOverrides);
        return await response.value();
    }

}
