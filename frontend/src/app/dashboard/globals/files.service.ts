import { Injectable } from "@angular/core";
import { Observable, of } from "rxjs";
import { ServerResponse } from "./models";
import { SharedModule } from "./shared.module";


@Injectable({
    providedIn: SharedModule,
})
export class StorageService {

    getServerResponse(): Observable<ServerResponse> {
        return of({
            items: [],
            limit: 40,
            offset: 0,
            parent: 0
        })
    }
}