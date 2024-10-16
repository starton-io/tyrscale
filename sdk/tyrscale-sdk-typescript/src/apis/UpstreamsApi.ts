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
  ResponsesCreatedSuccessResponseUpstreamUpsertRes,
  ResponsesDefaultSuccessResponseListUpstreamRes,
  ResponsesDefaultSuccessResponseUpstreamUpsertRes,
  ResponsesInternalServerErrorResponse,
  Upstream,
} from '../models/index';
import {
    ResponsesBadRequestResponseFromJSON,
    ResponsesBadRequestResponseToJSON,
    ResponsesCreatedSuccessResponseUpstreamUpsertResFromJSON,
    ResponsesCreatedSuccessResponseUpstreamUpsertResToJSON,
    ResponsesDefaultSuccessResponseListUpstreamResFromJSON,
    ResponsesDefaultSuccessResponseListUpstreamResToJSON,
    ResponsesDefaultSuccessResponseUpstreamUpsertResFromJSON,
    ResponsesDefaultSuccessResponseUpstreamUpsertResToJSON,
    ResponsesInternalServerErrorResponseFromJSON,
    ResponsesInternalServerErrorResponseToJSON,
    UpstreamFromJSON,
    UpstreamToJSON,
} from '../models/index';

export interface DeleteUpstreamRequest {
    routeUuid: string;
    uuid: string;
}

export interface ListUpstreamsRequest {
    routeUuid: string;
}

export interface UpsertUpstreamRequest {
    routeUuid: string;
    upstream: Upstream;
}

/**
 * 
 */
export class UpstreamsApi extends runtime.BaseAPI {

    /**
     * Delete a upstream
     * Delete a upstream
     */
    async deleteUpstreamRaw(requestParameters: DeleteUpstreamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesDefaultSuccessResponseUpstreamUpsertRes>> {
        if (requestParameters['routeUuid'] == null) {
            throw new runtime.RequiredError(
                'routeUuid',
                'Required parameter "routeUuid" was null or undefined when calling deleteUpstream().'
            );
        }

        if (requestParameters['uuid'] == null) {
            throw new runtime.RequiredError(
                'uuid',
                'Required parameter "uuid" was null or undefined when calling deleteUpstream().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/routes/{route_uuid}/upstreams/{uuid}`.replace(`{${"route_uuid"}}`, encodeURIComponent(String(requestParameters['routeUuid']))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters['uuid']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesDefaultSuccessResponseUpstreamUpsertResFromJSON(jsonValue));
    }

    /**
     * Delete a upstream
     * Delete a upstream
     */
    async deleteUpstream(requestParameters: DeleteUpstreamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesDefaultSuccessResponseUpstreamUpsertRes> {
        const response = await this.deleteUpstreamRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get list upstreams
     * Get list upstreams
     */
    async listUpstreamsRaw(requestParameters: ListUpstreamsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesDefaultSuccessResponseListUpstreamRes>> {
        if (requestParameters['routeUuid'] == null) {
            throw new runtime.RequiredError(
                'routeUuid',
                'Required parameter "routeUuid" was null or undefined when calling listUpstreams().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/routes/{route_uuid}/upstreams`.replace(`{${"route_uuid"}}`, encodeURIComponent(String(requestParameters['routeUuid']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesDefaultSuccessResponseListUpstreamResFromJSON(jsonValue));
    }

    /**
     * Get list upstreams
     * Get list upstreams
     */
    async listUpstreams(requestParameters: ListUpstreamsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesDefaultSuccessResponseListUpstreamRes> {
        const response = await this.listUpstreamsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Create or update a upstream
     * Create or update a upstream
     */
    async upsertUpstreamRaw(requestParameters: UpsertUpstreamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesCreatedSuccessResponseUpstreamUpsertRes>> {
        if (requestParameters['routeUuid'] == null) {
            throw new runtime.RequiredError(
                'routeUuid',
                'Required parameter "routeUuid" was null or undefined when calling upsertUpstream().'
            );
        }

        if (requestParameters['upstream'] == null) {
            throw new runtime.RequiredError(
                'upstream',
                'Required parameter "upstream" was null or undefined when calling upsertUpstream().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/routes/{route_uuid}/upstreams`.replace(`{${"route_uuid"}}`, encodeURIComponent(String(requestParameters['routeUuid']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: UpstreamToJSON(requestParameters['upstream']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesCreatedSuccessResponseUpstreamUpsertResFromJSON(jsonValue));
    }

    /**
     * Create or update a upstream
     * Create or update a upstream
     */
    async upsertUpstream(requestParameters: UpsertUpstreamRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesCreatedSuccessResponseUpstreamUpsertRes> {
        const response = await this.upsertUpstreamRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
