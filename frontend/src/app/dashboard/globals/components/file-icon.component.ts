import { Component, Input } from "@angular/core";
import { FILE_TYPES_ICONS } from '../constants';
import { faFile, faFilePdf, faFileArchive, faTable, faFilePowerpoint, faFileWord, faImage, faVideo, faHeadphones } from '@fortawesome/free-solid-svg-icons';

@Component({
    selector: "app-file-type-icon",
    template: `
        <fa-icon [icon]="icon()"></fa-icon>
    `
})
export class FileIconComponent {
    @Input() filetype?: string;
    iconToShow: any = faFile;

    file_to_icon: {[key: string]: any} = {
        "file-pdf": faFilePdf,
        "file-archive": faFileArchive,
        "file-spreadsheet": faTable,
        "file-powerpoint": faFilePowerpoint,
        "file-word": faFileWord,
        "image": faImage,
        "video": faVideo,
        "headphones": faHeadphones,
        "default": faFile
    }

    icon() {
        for (const key of Object.keys(FILE_TYPES_ICONS)) {
            const index = FILE_TYPES_ICONS[key].findIndex((t:string)=> t === this.filetype);
            if (index > -1) {
                return this.file_to_icon[key];
            }
        }
        return this.file_to_icon["default"];
    }
}
