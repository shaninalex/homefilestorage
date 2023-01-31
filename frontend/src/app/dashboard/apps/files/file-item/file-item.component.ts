import { Component, Input } from '@angular/core';
import { FileItem } from 'src/app/dashboard/globals/models';
import { faLock, faLockOpen } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-file-item',
  templateUrl: './file-item.component.html',
  styleUrls: ['./file-item.component.scss']
})
export class FileItemComponent {
    @Input() file?: FileItem;
    notpublic = faLock;
    public = faLockOpen;
}
