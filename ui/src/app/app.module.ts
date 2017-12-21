import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { Routing } from './app.routing';
import { AppComponent } from './component/app/app.component';
import { IndexComponent } from './component/page/index/index.component';
import { EntryComponent } from './component/page/entry/entry.component';
import { AuthComponent } from './component/page/auth/auth.component';

@NgModule({
  declarations: [
    AppComponent,
    IndexComponent,
    EntryComponent,
    AuthComponent
  ],
  imports: [
    Routing,
    BrowserModule
  ],
  providers: [
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
