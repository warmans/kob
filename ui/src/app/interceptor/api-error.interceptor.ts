import {Injectable} from '@angular/core';
import {HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest} from '@angular/common/http';

import {Observable} from 'rxjs/Observable';
import 'rxjs/add/operator/do';
import {NotificationService} from '../service/notification/notification.service';
import {Router} from "@angular/router";

@Injectable()
export class APIErrorInterceptor implements HttpInterceptor {

  constructor(private notify: NotificationService, private router: Router) {
  }

  intercept(req: HttpRequest<any>,
            next: HttpHandler): Observable<HttpEvent<any>> {

    return next.handle(req).do(
      (evt: any) => {},
      (err: any) => {
        if (err instanceof HttpErrorResponse) {
          console.log(err);
          let msg = err.message || 'Unknown error';
          this.notify.addNotification('alert-danger', msg);

          if (err.status === 401 || err.status == 403) {
            //if something auth-y happened go to login
            this.router.navigateByUrl("/login");
          }
          return Observable.throw(err);
        }
      }
    );
  }
}
