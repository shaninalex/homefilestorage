import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FilesComponent } from './files.component';
import { FilesRoutingModule } from './files-routing.module';
import { FileItemComponent } from './file-item/file-item.component';
import { SharedModule } from '../../globals/shared.module';
    


@NgModule({
    declarations: [
        FilesComponent,
        FileItemComponent,
    ],
    imports: [
        CommonModule,
        FilesRoutingModule,
        SharedModule
    ],
    providers: []
})
export class FilesModule { }

