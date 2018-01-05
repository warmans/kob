import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { StateService } from '../service/state/state.service';

@Injectable()
export class TokenInterceptor implements HttpInterceptor {

    private token: string;

    constructor(public state: StateService) {
        state.onStateChange.subscribe((newState) => {
            this.token = newState['token'];
        });
    }

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    request = request.clone({
      setHeaders: {
        Authorization: `Bearer ${this.token}`
      }
    });
    return next.handle(request);
  }
}
