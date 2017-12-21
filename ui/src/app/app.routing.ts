import {ModuleWithProviders} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {IndexComponent} from './component/page/index/index.component';

export const appRoutes: Routes = [
  {
    path: '',
    redirectTo: '/index',
    pathMatch: 'full'
  },
  {
    path: 'index',
    component: IndexComponent,

  }, {
    path: '**',
    redirectTo: '/index',//todo: 404 page?
  }
];

export const Routing: ModuleWithProviders = RouterModule.forRoot(appRoutes);