import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { SimplemdeModule, SIMPLEMDE_CONFIG } from 'ng2-simplemde'

import { Routing } from './app.routing';
import { AppComponent } from './component/app/app.component';
import { AuthComponent } from './component/page/auth/auth.component';
import { ApiClientService } from './service/api-client/index';
import { EntriesComponent } from './component/page/entries/entries.component';
import { EntriesIndexComponent } from './component/page/entries-index/entries-index.component';
import { EntriesEditComponent } from './component/page/entries-edit/entries-edit.component';
import { EntriesAcivityComponent } from './component/page/entries-acivity/entries-acivity.component';
import { LoginComponent } from './component/page/login/login.component';
import { LoginCompleteComponent } from './component/page/login-complete/login-complete.component';
import { StateService } from './service/state/state.service';
import { AuthGuard } from './guard/auth.guard';
import { TokenInterceptor } from './interceptor/token.interceptor';

declare var window: any;

@NgModule({
  declarations: [
    AppComponent,
    AuthComponent,
    EntriesComponent,
    EntriesIndexComponent,
    EntriesEditComponent,
    EntriesAcivityComponent,
    LoginComponent,
    LoginCompleteComponent,
  ],
  imports: [
    Routing,
    BrowserModule,
    HttpClientModule,
    SimplemdeModule.forRoot({provide: SIMPLEMDE_CONFIG, useValue: {spellChecker: false}})
  ],
  providers: [
    ApiClientService,
    {provide: 'domain', useValue: '//' + window.location.hostname + (window.location.port ? ':' + window.location.port : '')},
    {provide: HTTP_INTERCEPTORS, useClass: TokenInterceptor, multi: true},
    StateService,
    AuthGuard,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
