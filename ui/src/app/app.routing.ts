import { ModuleWithProviders } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { IndexComponent } from './component/page/index/index.component';
import { EditEntryComponent } from './component/page/edit-entry/edit-entry.component';

export const appRoutes: Routes = [
  {
    path: '',
    redirectTo: '/index',
    pathMatch: 'full'
  }, {
    path: 'index',
    component: IndexComponent,

  }, {
    path: 'entry/new',
    component: EditEntryComponent,
  }, {
    path: '**',
    redirectTo: '/index',//todo: 404 page?
  }
];

export const Routing: ModuleWithProviders = RouterModule.forRoot(appRoutes);