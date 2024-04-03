export * from './networksApi';
import { NetworksApi } from './networksApi';
export * from './recommendationApi';
import { RecommendationApi } from './recommendationApi';
export * from './routesApi';
import { RoutesApi } from './routesApi';
export * from './rpcsApi';
import { RpcsApi } from './rpcsApi';
export * from './upstreamsApi';
import { UpstreamsApi } from './upstreamsApi';
import * as http from 'http';

export class HttpError extends Error {
    constructor (public response: http.IncomingMessage, public body: any, public statusCode?: number) {
        super('HTTP request failed');
        this.name = 'HttpError';
    }
}

export { RequestFile } from '../model/models';

export const APIS = [NetworksApi, RecommendationApi, RoutesApi, RpcsApi, UpstreamsApi];
