import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FilesComponent } from './files.component';
import { FilesRoutingModule } from './files-routing.module';
import { FileItemComponent } from './file-item/file-item.component';
import { SharedModule } from '../../globals/shared.module';
import { FolderItemComponent } from './folder-item/folder-item.component';


@NgModule({
    declarations: [
        FilesComponent,
        FileItemComponent,
        FolderItemComponent
    ],
    imports: [
        CommonModule,
        FilesRoutingModule,
        SharedModule
    ],
    providers: []
})
export class FilesModule { }

