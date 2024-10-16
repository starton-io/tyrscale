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
  CreateRecommendationReq,
  ResponsesBadRequestResponse,
  ResponsesConflictResponseWithoutContext,
  ResponsesCreatedSuccessResponseCreateRecommendationRes,
  ResponsesDefaultSuccessResponseListRecommendationRes,
  ResponsesDefaultSuccessResponseWithoutData,
  ResponsesInternalServerErrorResponse,
  ResponsesNotFoundResponse,
  UpdateRecommendationReq,
} from '../models/index';
import {
    CreateRecommendationReqFromJSON,
    CreateRecommendationReqToJSON,
    ResponsesBadRequestResponseFromJSON,
    ResponsesBadRequestResponseToJSON,
    ResponsesConflictResponseWithoutContextFromJSON,
    ResponsesConflictResponseWithoutContextToJSON,
    ResponsesCreatedSuccessResponseCreateRecommendationResFromJSON,
    ResponsesCreatedSuccessResponseCreateRecommendationResToJSON,
    ResponsesDefaultSuccessResponseListRecommendationResFromJSON,
    ResponsesDefaultSuccessResponseListRecommendationResToJSON,
    ResponsesDefaultSuccessResponseWithoutDataFromJSON,
    ResponsesDefaultSuccessResponseWithoutDataToJSON,
    ResponsesInternalServerErrorResponseFromJSON,
    ResponsesInternalServerErrorResponseToJSON,
    ResponsesNotFoundResponseFromJSON,
    ResponsesNotFoundResponseToJSON,
    UpdateRecommendationReqFromJSON,
    UpdateRecommendationReqToJSON,
} from '../models/index';

export interface CreateRecommendationRequest {
    recommendation: CreateRecommendationReq;
}

export interface DeleteRecommendationRequest {
    routeUuid: string;
}

export interface UpdateRecommendationRequest {
    recommendation: UpdateRecommendationReq;
}

/**
 * 
 */
export class RecommendationsApi extends runtime.BaseAPI {

    /**
     * Create a recommendation
     * Create a recommendation
     */
    async createRecommendationRaw(requestParameters: CreateRecommendationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesCreatedSuccessResponseCreateRecommendationRes>> {
        if (requestParameters['recommendation'] == null) {
            throw new runtime.RequiredError(
                'recommendation',
                'Required parameter "recommendation" was null or undefined when calling createRecommendation().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/recommendations`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateRecommendationReqToJSON(requestParameters['recommendation']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesCreatedSuccessResponseCreateRecommendationResFromJSON(jsonValue));
    }

    /**
     * Create a recommendation
     * Create a recommendation
     */
    async createRecommendation(requestParameters: CreateRecommendationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesCreatedSuccessResponseCreateRecommendationRes> {
        const response = await this.createRecommendationRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete a recommendation
     * Delete a recommendation
     */
    async deleteRecommendationRaw(requestParameters: DeleteRecommendationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesDefaultSuccessResponseWithoutData>> {
        if (requestParameters['routeUuid'] == null) {
            throw new runtime.RequiredError(
                'routeUuid',
                'Required parameter "routeUuid" was null or undefined when calling deleteRecommendation().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/recommendations/{route_uuid}`.replace(`{${"route_uuid"}}`, encodeURIComponent(String(requestParameters['routeUuid']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesDefaultSuccessResponseWithoutDataFromJSON(jsonValue));
    }

    /**
     * Delete a recommendation
     * Delete a recommendation
     */
    async deleteRecommendation(requestParameters: DeleteRecommendationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesDefaultSuccessResponseWithoutData> {
        const response = await this.deleteRecommendationRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * List recommendation
     * List recommendation
     */
    async listRecommendationsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesDefaultSuccessResponseListRecommendationRes>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/recommendations`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesDefaultSuccessResponseListRecommendationResFromJSON(jsonValue));
    }

    /**
     * List recommendation
     * List recommendation
     */
    async listRecommendations(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesDefaultSuccessResponseListRecommendationRes> {
        const response = await this.listRecommendationsRaw(initOverrides);
        return await response.value();
    }

    /**
     * Update a recommendation
     * Update a recommendation
     */
    async updateRecommendationRaw(requestParameters: UpdateRecommendationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ResponsesDefaultSuccessResponseWithoutData>> {
        if (requestParameters['recommendation'] == null) {
            throw new runtime.RequiredError(
                'recommendation',
                'Required parameter "recommendation" was null or undefined when calling updateRecommendation().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/recommendations`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: UpdateRecommendationReqToJSON(requestParameters['recommendation']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ResponsesDefaultSuccessResponseWithoutDataFromJSON(jsonValue));
    }

    /**
     * Update a recommendation
     * Update a recommendation
     */
    async updateRecommendation(requestParameters: UpdateRecommendationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ResponsesDefaultSuccessResponseWithoutData> {
        const response = await this.updateRecommendationRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
