// tslint:disable
/**
 * Connthass API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    Entry,
    EntryFromJSON,
    EntryToJSON,
} from './';

/**
 * 
 * @export
 * @interface InlineObject1
 */
export interface InlineObject1 {
    /**
     * 
     * @type {Array<Entry>}
     * @memberof InlineObject1
     */
    entries?: Array<Entry>;
}

export function InlineObject1FromJSON(json: any): InlineObject1 {
    return {
        'entries': !exists(json, 'entries') ? undefined : (json['entries'] as Array<any>).map(EntryFromJSON),
    };
}

export function InlineObject1ToJSON(value?: InlineObject1): any {
    if (value === undefined) {
        return undefined;
    }
    return {
        'entries': value.entries === undefined ? undefined : (value.entries as Array<any>).map(EntryToJSON),
    };
}


