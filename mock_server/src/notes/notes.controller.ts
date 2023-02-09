import { Controller, Get, Req } from '@nestjs/common';
import { Request } from 'express';


interface Label {
    id: number
    name: string
}


interface Note {
    id: number, 
    title: string,
    content: string,
    created_at: string,
    color: string,
    labels: Label[]
}

const NOTES: Note[] = [
    {id: 1, title: 'test', content: 'test', created_at: new Date().toString(), color: 'default', labels: [{id: 1, name: 'learn'}, {id: 2, name: 'learn2'},]},
    {id: 2, title: 'test', content: 'test', created_at: new Date().toString(), color: 'default', labels: [{id: 1, name: 'learn'}, {id: 2, name: 'learn2'},]},
    {id: 3, title: 'test', content: 'test', created_at: new Date().toString(), color: 'default', labels: [{id: 1, name: 'learn'}, {id: 2, name: 'learn2'},]},
    {id: 4, title: 'test', content: 'test', created_at: new Date().toString(), color: 'default', labels: [{id: 1, name: 'learn'}, {id: 2, name: 'learn2'},]},
    {id: 5, title: 'test', content: 'test', created_at: new Date().toString(), color: 'default', labels: [{id: 1, name: 'learn'}, {id: 2, name: 'learn2'},]},
    {id: 6, title: 'test', content: 'test', created_at: new Date().toString(), color: 'default', labels: [{id: 1, name: 'learn'}, {id: 2, name: 'learn2'},]},
    {id: 7, title: 'test', content: 'test', created_at: new Date().toString(), color: 'default', labels: [{id: 1, name: 'learn'}, {id: 2, name: 'learn2'},]},
    {id: 8, title: 'test', content: 'test', created_at: new Date().toString(), color: 'default', labels: [{id: 1, name: 'learn'}, {id: 2, name: 'learn2'},]},
]


@Controller('api/v1/notes')
export class NotesController {

    @Get()
    getAllNotes(@Req() request: Request): Note[] {
        return NOTES;
    }
}
