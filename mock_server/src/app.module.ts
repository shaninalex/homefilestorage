import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { NotesController } from './notes/notes.controller';
import { FilesController } from './files/files.controller';
import { AuthController } from './auth/auth.controller';
import { JwtModule } from '@nestjs/jwt';

@Module({
    imports: [
        JwtModule.register({
            secret: "VERY SECRET STRING!",
            signOptions: { expiresIn: '360s' },
        }),
    ],
    controllers: [AppController, NotesController, FilesController, AuthController],
    providers: [AppService],
})
export class AppModule { }
