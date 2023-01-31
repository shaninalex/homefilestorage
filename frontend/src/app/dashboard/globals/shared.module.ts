import { NgModule } from "@angular/core";
import { FileIconComponent } from "./components/file-icon.component";
import { FileSizePipe } from "./filesize.pipe";
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

@NgModule({
    declarations: [
        FileSizePipe,
        FileIconComponent
    ],
    imports: [
        FontAwesomeModule
    ],
    exports: [
        FileSizePipe,
        FileIconComponent,
        FontAwesomeModule
    ]
})
export class SharedModule { }
