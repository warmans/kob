import { ModuleWithProviders } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { EntriesComponent } from './component/page/entries/entries.component';
import { EntriesIndexComponent } from './component/page/entries-index/entries-index.component';
import { EntriesEditComponent } from './component/page/entries-edit/entries-edit.component';
import { EntriesAcivityComponent } from './component/page/entries-acivity/entries-acivity.component';
import { LoginComponent } from './component/page/login/login.component';
import { LoginCompleteComponent } from './component/page/login-complete/login-complete.component';
import { AuthGuard } from './guard/auth.guard';

export const appRoutes: Routes = [{
  path: '',
  redirectTo: '/entries',
  pathMatch: 'full'
}, {
  path: 'login',
  component: LoginComponent,
 }, {
  path: 'login/complete',
  component: LoginCompleteComponent,
 }, {
  path: 'entries',
  component: EntriesComponent,
  canActivate: [AuthGuard],
  children: [
    {path: '', redirectTo: 'activity', pathMatch: 'prefix'},
    {path: 'activity', component: EntriesAcivityComponent},
    {path: 'index', component: EntriesIndexComponent},
    {path: 'new', component: EntriesEditComponent}
  ],
}, {
  path: '**',
  redirectTo: '/index', // todo: 404 page?
}];

export const Routing: ModuleWithProviders = RouterModule.forRoot(appRoutes);
