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
  CreateRouteReq,
  ResponsesBadRequestResponse,
  ResponsesCreatedSuccessResponseCreateRouteRes,
  ResponsesDefaultSuccessResponseListRouteRes,
  ResponsesDefaultSuccessResponseWithoutData,
  ResponsesInternalServerErrorResponse,
  UpdateRouteReq,
} from '../models/index';
import {
    CreateRouteReqFromJSON,
    CreateRouteReqToJSON,
    ResponsesBadRequestResponseFromJSON,
    ResponsesBadRequestResponseToJSON,
    ResponsesCreatedSuccessResponseCreateRouteResFromJSON,
    ResponsesCreatedSuccessResponseCreateRouteResToJSON,
    ResponsesDefaultSuccessResponseListRouteResFromJSON,
    ResponsesDefaultSuccessResponseListRouteResToJSON,
    ResponsesDefaultSuccessResponseWithoutDataFromJSON,
    ResponsesDefaultSuccessResponseWithoutDataToJSON,
    ResponsesInternalServerErrorResponseFromJSON,
    ResponsesInternalServerErrorResponseToJSON,
    UpdateRouteReqFromJSON,
    UpdateRouteReqToJSON,
} from '../models/index';

export interface CreateRouteRequest {
    route: CreateRouteReq;
}

export interface DeleteRouteRequest {
    uuid: string;
}

export interface ListRoutesRequest {
    host?: string;
    loadBalancerStrategy?: ListRoutesLoadBalancerStrategyEnum;
    path?: string;
    uuid?: string;
}

export interface UpdateRouteRequest {
    uuid: string;
    route: UpdateRouteReq;
}

/**
 * 
 */
export class RoutesApi extends runtime.BaseAPI {

    /**
     * Create a route
     * Create a route
     */
    async createRouteRaw(requestParameters: CreateRouteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesCreatedSuccessResponseCreateRouteRes>> {
        if (requestParameters['route'] == null) {
            throw new runtime.RequiredError(
                'route',
                'Required parameter "route" was null or undefined when calling createRoute().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/routes`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateRouteReqToJSON(requestParameters['route']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesCreatedSuccessResponseCreateRouteResFromJSON(jsonValue));
    }

    /**
     * Create a route
     * Create a route
     */
    async createRoute(requestParameters: CreateRouteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesCreatedSuccessResponseCreateRouteRes> {
        const response = await this.createRouteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete a route
     * Delete a route
     */
    async deleteRouteRaw(requestParameters: DeleteRouteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesDefaultSuccessResponseWithoutData>> {
        if (requestParameters['uuid'] == null) {
            throw new runtime.RequiredError(
                'uuid',
                'Required parameter "uuid" was null or undefined when calling deleteRoute().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/routes/{uuid}`.replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters['uuid']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesDefaultSuccessResponseWithoutDataFromJSON(jsonValue));
    }

    /**
     * Delete a route
     * Delete a route
     */
    async deleteRoute(requestParameters: DeleteRouteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesDefaultSuccessResponseWithoutData> {
        const response = await this.deleteRouteRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get list routes
     * Get list routes
     */
    async listRoutesRaw(requestParameters: ListRoutesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesDefaultSuccessResponseListRouteRes>> {
        const queryParameters: any = {};

        if (requestParameters['host'] != null) {
            queryParameters['host'] = requestParameters['host'];
        }

        if (requestParameters['loadBalancerStrategy'] != null) {
            queryParameters['loadBalancerStrategy'] = requestParameters['loadBalancerStrategy'];
        }

        if (requestParameters['path'] != null) {
            queryParameters['path'] = requestParameters['path'];
        }

        if (requestParameters['uuid'] != null) {
            queryParameters['uuid'] = requestParameters['uuid'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/routes`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesDefaultSuccessResponseListRouteResFromJSON(jsonValue));
    }

    /**
     * Get list routes
     * Get list routes
     */
    async listRoutes(requestParameters: ListRoutesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesDefaultSuccessResponseListRouteRes> {
        const response = await this.listRoutesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Update a route
     * Update a route
     */
    async updateRouteRaw(requestParameters: UpdateRouteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesDefaultSuccessResponseWithoutData>> {
        if (requestParameters['uuid'] == null) {
            throw new runtime.RequiredError(
                'uuid',
                'Required parameter "uuid" was null or undefined when calling updateRoute().'
            );
        }

        if (requestParameters['route'] == null) {
            throw new runtime.RequiredError(
                'route',
                'Required parameter "route" was null or undefined when calling updateRoute().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/routes/{uuid}`.replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters['uuid']))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: UpdateRouteReqToJSON(requestParameters['route']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesDefaultSuccessResponseWithoutDataFromJSON(jsonValue));
    }

    /**
     * Update a route
     * Update a route
     */
    async updateRoute(requestParameters: UpdateRouteRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesDefaultSuccessResponseWithoutData> {
        const response = await this.updateRouteRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const ListRoutesLoadBalancerStrategyEnum = {
    BalancerWeightRoundRobin: 'weight-round-robin',
    BalancerLeastLoad: 'least-load',
    BalancerPriority: 'failover-priority'
} as const;
export type ListRoutesLoadBalancerStrategyEnum = typeof ListRoutesLoadBalancerStrategyEnum[keyof typeof ListRoutesLoadBalancerStrategyEnum];
