import { Component } from '@angular/core';
// import { faFilePdf } from '@fortawesome/free-solid-svg-icons';
import { faFilePdf, faFolder } from '@fortawesome/free-solid-svg-icons';
import { FileItem } from '../../globals/models';


@Component({
    selector: 'app-files',
    templateUrl: './files.component.html',
})
export class FilesComponent {

    files: FileItem[];
    icon = faFilePdf;
    folder_icon = faFolder;

    constructor() { 
        this.files = [
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                modified: new Date(),
                public: true
            }
        ];
        console.log("Files component")
    }

}
