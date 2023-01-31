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
        this.files = [
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                created_at: new Date(),
                public: true,
                type: "application/pdf"
            },
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                created_at: new Date(),
                public: true,
                type: "video/x-ms-wmv"
            },
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                created_at: new Date(),
                public: true,
                type: "image/png"
            },
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                created_at: new Date(),
                public: true,
                type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
            },
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                created_at: new Date(),
                public: true,
                type: "application/vnd.ms-excel"
            },
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                created_at: new Date(),
                public: true,
                type: "application/zip"
            },
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                created_at: new Date(),
                public: true,
                type: "audio/mpeg"
            },
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                created_at: new Date(),
                public: true,
                type: "video/mp4"
            },
            {
                id: 1,
                name: "document.pdf",
                size: 2314123,
                created_at: new Date(),
                public: true,
                type: "video/x-msvideo"
            },
        ];
    }

}
