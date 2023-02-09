import { Component, OnInit } from '@angular/core';
import { StorageService } from '../../globals/files.service';
import { FileItem, FolderItem } from '../../globals/models';


@Component({
    selector: 'app-files',
    templateUrl: './files.component.html',
    styleUrls: ['./files.component.scss'],
})
export class FilesComponent implements OnInit {

    // items?: Array<FileItem|FolderItem>;

    constructor(private storageService: StorageService) {}

    ngOnInit(): void {
        
    }

}
