import { Controller, Get, Req } from '@nestjs/common';
import { Request } from '@nestjs/common';

class FileItem {
    constructor(
        public id: number,
        public name: string,
        public size: number,
        public created_at: string,
        public ispublic: boolean,
        public parrent: number | null,
        public type: string) { }
}

class FolderItem {
    constructor(
        public id: number,
        public name: string,
        public created_at: string,
        public ispublic: boolean,
        public parent: number | null,
        public child: any,
    ) { }
}

const FILETREE: Array<FolderItem | FileItem> = [
    new FolderItem(
        1,
        'main',
        new Date().toString(),
        false,
        null,
        [
            new FolderItem(
                5,
                'main',
                new Date().toString(),
                false,
                1,
                [
                    new FileItem(6, 'Book5', 123123, new Date().toString(), false, 5, 'application/pdf'),
                    new FileItem(7, 'Book6', 123123, new Date().toString(), false, 5, 'application/pdf'),
                    new FileItem(8, 'Book7', 123123, new Date().toString(), false, 5, 'application/pdf'),
                    new FileItem(9, 'Book8', 123123, new Date().toString(), false, 5, 'application/pdf'),
                ]
            ),
            new FileItem(2, 'Book2', 123123, new Date().toString(), false, 1, 'application/pdf'),
            new FileItem(3, 'Book3', 123123, new Date().toString(), false, 1, 'application/pdf'),
            new FileItem(4, 'Book4', 123123, new Date().toString(), false, 1, 'application/pdf'),
        ]
    ),
    new FileItem(10, 'Book2', 123123, new Date().toString(), false, null, 'application/pdf'),
    new FileItem(11, 'Book3', 123123, new Date().toString(), false, null, 'application/pdf'),
    new FileItem(12, 'Book4', 123123, new Date().toString(), false, null, 'application/pdf'),
]


@Controller('api/v1/files')
export class FilesController {
    @Get()
    getAllNotes(@Req() request: Request): any {
        return FILETREE;
    }
}
