import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

import { SimplemdeModule, SIMPLEMDE_CONFIG } from 'ng2-simplemde'

import { Routing } from './app.routing';
import { AppComponent } from './component/app/app.component';
import { IndexComponent } from './component/page/index/index.component';
import { EntryComponent } from './component/page/entry/entry.component';
import { AuthComponent } from './component/page/auth/auth.component';
import { ApiClientService } from './service/api-client/index';
import { EditEntryComponent } from './component/page/edit-entry/edit-entry.component';

declare var window: any;

@NgModule({
  declarations: [
    AppComponent,
    IndexComponent,
    EntryComponent,
    AuthComponent,
    EditEntryComponent,
  ],
  imports: [
    Routing,
    BrowserModule,
    HttpClientModule,
    SimplemdeModule.forRoot({provide: SIMPLEMDE_CONFIG, useValue: {}})
  ],
  providers: [
    ApiClientService,
    {provide: 'domain', useValue: "//"+window.location.hostname+(window.location.port ? ":"+window.location.port : "")}
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
