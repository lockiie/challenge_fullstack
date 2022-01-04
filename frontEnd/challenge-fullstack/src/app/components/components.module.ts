import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { CommonModule } from '@angular/common';
import { ComponentsRoutes } from './components.routing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { MatSidenavModule } from '@angular/material/sidenav';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatChipsModule } from '@angular/material/chips';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatRadioModule } from '@angular/material/radio';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatTableModule } from '@angular/material/table'  
import { MatButtonModule } from '@angular/material/button';
import { MatNativeDateModule } from '@angular/material/core';
import { MatDialogModule } from '@angular/material/dialog';

import { LevelsComponent } from './levels/levels.component';
import { AddLevelsComponent } from './levels/add-levels/add-levels.component';
import { EditLevelsComponent } from './levels/edit-levels/edit-levels.component';



import { AddDevelopersComponent } from './developers/add-developers/add-developers.component';
import { DevelopersComponent } from './developers/developers.component';
import { EditDevelopersComponent } from './developers/edit-developers/edit-developers.component';




@NgModule({
    imports: [
        CommonModule,
        RouterModule.forChild(ComponentsRoutes),
        ReactiveFormsModule,
        FormsModule,
        MatDividerModule,
        MatCardModule,
        MatFormFieldModule,
        MatInputModule,
        MatIconModule,
        MatSelectModule,
        MatChipsModule,
        MatDatepickerModule,
        MatRadioModule,
        MatPaginatorModule,
        MatTableModule,
        MatSidenavModule,
        MatIconModule,
        MatListModule,
        MatToolbarModule,
        MatButtonModule,
        MatNativeDateModule,
        MatDialogModule
    ],
    declarations: [
        LevelsComponent,
        AddLevelsComponent,
        EditLevelsComponent,
        DevelopersComponent,
        AddDevelopersComponent,
        EditDevelopersComponent
    ],
    providers: [
        MatNativeDateModule ,
        MatDatepickerModule
      ],
})
export class ComponentsModule { }