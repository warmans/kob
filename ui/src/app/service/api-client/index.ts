import { Inject, Injectable, Optional } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams, HttpResponse } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';

import {
  rpcAuthURL,
  rpcAuthor,
  rpcComment,
  rpcCommentList,
  rpcEntry,
  rpcEntryList,
  rpcJWT,
  rpcTag
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
  * Method CreateJWT
  * @return Full HTTP response as Observable
  */
  public CreateJWT(): Observable<HttpResponse<rpcJWT>> {
    let uri = `/auth/token`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcJWT>('get', uri, headers, params, null);
  }

  /**
  * Method CreateAuthURL
  * @return Full HTTP response as Observable
  */
  public CreateAuthURL(): Observable<HttpResponse<rpcAuthURL>> {
    let uri = `/auth/url`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcAuthURL>('get', uri, headers, params, null);
  }

  /**
  * Method ListEntries
  * @param page 
  * @param num_per_page 
  * @param keyword 
  * @return Full HTTP response as Observable
  */
  public ListEntries(page: string, num_per_page: string, keyword: string): Observable<HttpResponse<rpcEntryList>> {
    let uri = `/entry`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    if (page !== undefined && page !== null) {
      params = params.set('page', page + '');
    }
    if (num_per_page !== undefined && num_per_page !== null) {
      params = params.set('num_per_page', num_per_page + '');
    }
    if (keyword !== undefined && keyword !== null) {
      params = params.set('keyword', keyword + '');
    }
    return this.sendRequest<rpcEntryList>('get', uri, headers, params, null);
  }

  /**
  * Method CreateEntry
  * @return Full HTTP response as Observable
  */
  public CreateEntry(): Observable<HttpResponse<rpcEntry>> {
    let uri = `/entry`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcEntry>('post', uri, headers, params, null);
  }

  /**
  * Method ListEntryComments
  * @param entry_id 
  * @return Full HTTP response as Observable
  */
  public ListEntryComments(entry_id: string): Observable<HttpResponse<rpcCommentList>> {
    let uri = `/entry/${entry_id}/comment`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcCommentList>('get', uri, headers, params, null);
  }

  /**
  * Method CreateEntryComment
  * @param entry_id 
  * @return Full HTTP response as Observable
  */
  public CreateEntryComment(entry_id: string): Observable<HttpResponse<rpcComment>> {
    let uri = `/entry/${entry_id}/comment`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcComment>('post', uri, headers, params, null);
  }

  /**
  * Method UpdateEntryComment
  * @param entry_id 
  * @param id 
  * @return Full HTTP response as Observable
  */
  public UpdateEntryComment(entry_id: string, id: string): Observable<HttpResponse<rpcComment>> {
    let uri = `/entry/${entry_id}/comment/${id}`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcComment>('put', uri, headers, params, null);
  }

  /**
  * Method GetEntry
  * @param id 
  * @return Full HTTP response as Observable
  */
  public GetEntry(id: string): Observable<HttpResponse<rpcEntry>> {
    let uri = `/entry/${id}`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcEntry>('get', uri, headers, params, null);
  }

  /**
  * Method UpdateEntry
  * @param id 
  * @return Full HTTP response as Observable
  */
  public UpdateEntry(id: string): Observable<HttpResponse<rpcEntry>> {
    let uri = `/entry/${id}`;
    let headers = new HttpHeaders();
    let params = new HttpParams();
    return this.sendRequest<rpcEntry>('put', uri, headers, params, null);
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
