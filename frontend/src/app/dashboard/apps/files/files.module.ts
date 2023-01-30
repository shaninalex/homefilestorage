import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FilesComponent } from './files.component';
import { FilesRoutingModule } from './files-routing.module';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
// import { FileItemComponent } from './file-item/file-item.component';
import { FileSizePipe } from '../../globals/filesize.pipe';


@NgModule({
    declarations: [
        FilesComponent,
        // FileItemComponent,
        FileSizePipe
    ],
    imports: [
        CommonModule,
        FilesRoutingModule,
        FontAwesomeModule
    ],
    providers: []
})
export class FilesModule { }

