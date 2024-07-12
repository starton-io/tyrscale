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


/**
 * 
 * @export
 */
export const HealthcheckHealthCheckType = {
    EthBlockNumberType: 'eth_block_number'
} as const;
export type HealthcheckHealthCheckType = typeof HealthcheckHealthCheckType[keyof typeof HealthcheckHealthCheckType];


export function instanceOfHealthcheckHealthCheckType(value: any): boolean {
    for (const key in HealthcheckHealthCheckType) {
        if (Object.prototype.hasOwnProperty.call(HealthcheckHealthCheckType, key)) {
            if (HealthcheckHealthCheckType[key] === value) {
                return true;
            }
        }
    }
    return false;
}

export function HealthcheckHealthCheckTypeFromJSON(json: any): HealthcheckHealthCheckType {
    return HealthcheckHealthCheckTypeFromJSONTyped(json, false);
}

export function HealthcheckHealthCheckTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): HealthcheckHealthCheckType {
    return json as HealthcheckHealthCheckType;
}

export function HealthcheckHealthCheckTypeToJSON(value?: HealthcheckHealthCheckType | null): any {
    return value as any;
}
