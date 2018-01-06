import { Inject, Injectable, Optional } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';

import {
  ActivityAction,
  rpcActivity,
  rpcActivityList,
  rpcAuthURL,
  rpcAuthor,
  rpcComment,
  rpcCommentList,
  rpcCreateEntryCommentRequest,
  rpcCreateEntryRequest,
  rpcEntry,
  rpcEntryList,
  rpcJWT,
  rpcSearchResult,
  rpcSearchResultList,
  rpcTag,
  rpcUpdateEntryCommentRequest,
  rpcUpdateEntryRequest
} from './models';

/**
* Created with angular4-swagger-client-generator.
*/
@Injectable()
export class ApiClientService {

  private domain = 'http://localhostundefined';

  constructor(private http: HttpClient, @Optional() @Inject('domain') domain: string ) {
    if (domain) {
      this.domain = domain;
    }
  }

  /**
  * Method ListActivity
  * @return Full HTTP response as Observable
  */
  public ListActivity(): Observable<HttpResponse<rpcActivityList>> {
    let uri = `/api/v1/activity`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcActivityList>('get', uri, headers, params, null);
  }

  /**
  * Method CreateJWT
  * @param id 
  * @param sub 
  * @param name 
  * @param given_name 
  * @param family_name 
  * @param profile 
  * @param picture 
  * @param email 
  * @param email_verified 
  * @param gender 
  * @return Full HTTP response as Observable
  */
  public CreateJWT(id: string, sub: string, name: string, given_name: string, family_name: string, profile: string, picture: string, email: string, email_verified: boolean, gender: string): Observable<HttpResponse<rpcJWT>> {
    let uri = `/api/v1/auth/token`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    if (id !== undefined && id !== null) {
      params = params.set('id', id + '');
    }
    if (sub !== undefined && sub !== null) {
      params = params.set('sub', sub + '');
    }
    if (name !== undefined && name !== null) {
      params = params.set('name', name + '');
    }
    if (given_name !== undefined && given_name !== null) {
      params = params.set('given_name', given_name + '');
    }
    if (family_name !== undefined && family_name !== null) {
      params = params.set('family_name', family_name + '');
    }
    if (profile !== undefined && profile !== null) {
      params = params.set('profile', profile + '');
    }
    if (picture !== undefined && picture !== null) {
      params = params.set('picture', picture + '');
    }
    if (email !== undefined && email !== null) {
      params = params.set('email', email + '');
    }
    if (email_verified !== undefined && email_verified !== null) {
      params = params.set('email_verified', email_verified + '');
    }
    if (gender !== undefined && gender !== null) {
      params = params.set('gender', gender + '');
    }
    return this.sendRequest<rpcJWT>('get', uri, headers, params, null);
  }

  /**
  * Method CreateAuthURL
  * @return Full HTTP response as Observable
  */
  public CreateAuthURL(): Observable<HttpResponse<rpcAuthURL>> {
    let uri = `/api/v1/auth/url`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcAuthURL>('get', uri, headers, params, null);
  }

  /**
  * Method ListEntries
  * @param page 
  * @param num_per_page 
  * @return Full HTTP response as Observable
  */
  public ListEntries(page: string, num_per_page: string): Observable<HttpResponse<rpcEntryList>> {
    let uri = `/api/v1/entry`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    if (page !== undefined && page !== null) {
      params = params.set('page', page + '');
    }
    if (num_per_page !== undefined && num_per_page !== null) {
      params = params.set('num_per_page', num_per_page + '');
    }
    return this.sendRequest<rpcEntryList>('get', uri, headers, params, null);
  }

  /**
  * Method CreateEntry
  * @param body 
  * @return Full HTTP response as Observable
  */
  public CreateEntry(body: rpcCreateEntryRequest): Observable<HttpResponse<rpcEntry>> {
    let uri = `/api/v1/entry`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcEntry>('post', uri, headers, params, JSON.stringify(body));
  }

  /**
  * Method ListEntryComments
  * @param entry_id 
  * @return Full HTTP response as Observable
  */
  public ListEntryComments(entry_id: string): Observable<HttpResponse<rpcCommentList>> {
    let uri = `/api/v1/entry/${entry_id}/comment`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcCommentList>('get', uri, headers, params, null);
  }

  /**
  * Method CreateEntryComment
  * @param entry_id 
  * @param body 
  * @return Full HTTP response as Observable
  */
  public CreateEntryComment(entry_id: string, body: rpcCreateEntryCommentRequest): Observable<HttpResponse<rpcComment>> {
    let uri = `/api/v1/entry/${entry_id}/comment`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcComment>('post', uri, headers, params, JSON.stringify(body));
  }

  /**
  * Method UpdateEntryComment
  * @param entry_id 
  * @param id 
  * @param body 
  * @return Full HTTP response as Observable
  */
  public UpdateEntryComment(entry_id: string, id: string, body: rpcUpdateEntryCommentRequest): Observable<HttpResponse<rpcComment>> {
    let uri = `/api/v1/entry/${entry_id}/comment/${id}`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcComment>('put', uri, headers, params, JSON.stringify(body));
  }

  /**
  * Method GetEntry
  * @param id 
  * @return Full HTTP response as Observable
  */
  public GetEntry(id: string): Observable<HttpResponse<rpcEntry>> {
    let uri = `/api/v1/entry/${id}`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcEntry>('get', uri, headers, params, null);
  }

  /**
  * Method UpdateEntry
  * @param id 
  * @param body 
  * @return Full HTTP response as Observable
  */
  public UpdateEntry(id: string, body: rpcUpdateEntryRequest): Observable<HttpResponse<rpcEntry>> {
    let uri = `/api/v1/entry/${id}`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcEntry>('put', uri, headers, params, JSON.stringify(body));
  }

  /**
  * Method GetMe
  * @return Full HTTP response as Observable
  */
  public GetMe(): Observable<HttpResponse<rpcAuthor>> {
    let uri = `/api/v1/me`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcAuthor>('get', uri, headers, params, null);
  }

  /**
  * Method Search
  * @param query 
  * @return Full HTTP response as Observable
  */
  public Search(query: string): Observable<HttpResponse<rpcSearchResultList>> {
    let uri = `/api/v1/search`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    if (query !== undefined && query !== null) {
      params = params.set('query', query + '');
    }
    return this.sendRequest<rpcSearchResultList>('get', uri, headers, params, null);
  }

  private sendRequest<T>(method: string, uri: string, headers: HttpHeaders, params: HttpParams, body: any): Observable<HttpResponse<T>> {
    if (method === 'get') {
      return this.http.get<T>(this.domain + uri, { headers: headers.set('Accept', 'application/json'), params: params, observe: 'response' });
    } else if (method === 'put') {
      return this.http.put<T>(this.domain + uri, body, { headers: headers.set('Content-Type', 'application/json'), params: params, observe: 'response' });
    } else if (method === 'post') {
      return this.http.post<T>(this.domain + uri, body, { headers: headers.set('Content-Type', 'application/json'), params: params, observe: 'response' });
    } else if (method === 'delete') {
      return this.http.delete<T>(this.domain + uri, { headers: headers, params: params, observe: 'response' });
    } else {
      console.error('Unsupported request: ' + method);
      return Observable.throw('Unsupported request: ' + method);
    }
  }
}
