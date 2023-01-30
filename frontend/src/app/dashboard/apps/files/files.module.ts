import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { FilesComponent } from './files.component';
import { FilesRoutingModule } from './files-routing.module';


@NgModule({
    declarations: [
        FilesComponent
    ],
    imports: [
        CommonModule,
        FilesRoutingModule
    ],
    providers: []
})
export class FilesModule { }

