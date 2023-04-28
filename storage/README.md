## Filesystem module

### responsibilities

- save and retrieve files
- remove files if needed
- indexing and validating


Basicaly this service only responsible on saveing files in apropriate place and return files if needed.

Urls:

| Method    | Path                          | Description   |
| :-------- | :---------------------------- | :------------ |
| POST      | `/save`                       | Save file     |
| GET       | `/files/:y/:m/:d/:filename`   | return file   |


Not implemented yet.
| Method    | Path                          | Description   |
| :-------- | :---------------------------- | :------------ |
| GET       | `/files/:file_id`             | return file   |
| DELETE    | `/files/:file_id`             | delete file   |
