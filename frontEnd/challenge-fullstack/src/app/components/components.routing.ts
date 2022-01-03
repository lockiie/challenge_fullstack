import { Routes } from '@angular/router';
import { AddDevelopersComponent } from './developers/add-developers/add-developers.component';
import { DevelopersComponent } from './developers/developers.component';
import { EditDevelopersComponent } from './developers/edit-developers/edit-developers.component';
import { AddLevelsComponent } from './levels/add-levels/add-levels.component';
import { EditLevelsComponent } from './levels/edit-levels/edit-levels.component';
import { LevelsComponent } from './levels/levels.component';


export const ComponentsRoutes: Routes = [
  {
    path: '',
    children: [
      {
        path: 'levels/list',
        component: LevelsComponent,
        data: {
          title: 'List of levels',
          urls: [
            { title: 'Levels', url: '/levels' },
            { title: 'List' }
          ]
        }
      },
      {
        path: 'levels/add',
        component: AddLevelsComponent,
        data: {
          title: 'Add levels',
          urls: [
            { title: 'Levels', url: '/levels' },
            { title: 'Add' }
          ]
        }
      },
      {
        path: 'levels/edit/:id',
        component: EditLevelsComponent,
        data: {
          title: 'Edit levels',
          urls: [
            { title: 'Levels', url: '/levels' },
            { title: 'Edit' }
          ]
        }
      },
      {
        path: 'developers/list',
        component: DevelopersComponent,
        data: {
          title: 'List of developers',
          urls: [
            { title: 'Developers', url: '/developers' },
            { title: 'List' }
          ]
        }
      },
      {
        path: 'developers/add',
        component: AddDevelopersComponent,
        data: {
          title: 'Add developers',
          urls: [
            { title: 'Developers', url: '/developers' },
            { title: 'Add' }
          ]
        }
      },
      {
        path: 'developers/edit/:id',
        component: EditDevelopersComponent,
        data: {
          title: 'Edit developers',
          urls: [
            { title: 'Developers', url: '/developers' },
            { title: 'Edit' }
          ]
        }
      },
    ]
  }
];