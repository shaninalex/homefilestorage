import { Component } from '@angular/core';
// import { faFilePdf } from '@fortawesome/free-solid-svg-icons';
import { faFilePdf, faFolder } from '@fortawesome/free-solid-svg-icons';
import { FileItem } from '../../globals/models';


@Component({
    selector: 'app-files',
    templateUrl: './files.component.html',
    styleUrls: ['./files.component.scss'],
})
export class FilesComponent {

    files: FileItem[];
    icon = faFilePdf;
    folder_icon = faFolder;

    constructor() {

        this.files = []
    }

}
