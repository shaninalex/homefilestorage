-- SYSTEM
CREATE TABLE IF NOT EXISTS "users"
(
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "email" CHAR(256) NOT NULL UNIQUE,
    "hashed_password" TEXT NOT NULL,
    "active" BOOLEAN NOT NULL DEFAULT true,
    "created_at" DATETIME DEFAULT current_timestamp
);

CREATE INDEX IF NOT EXISTS "users_email_idx" ON "users" ( "email" );


CREATE TABLE IF NOT EXISTS "files"
(
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "name" CHAR(256) NOT NULL UNIQUE,
    "mime_type" CHAR(128) NOT NULL,
    "size" INTEGER NOT NULL,
    "system_path" CHAR(256),
    "owner" INTEGER NOT NULL,
    "hash" TEXT NOT NULL,
    "public" BOOLEAN NOT NULL DEFAULT true,
    "created_at" DATETIME DEFAULT current_timestamp
    CHECK ("size" > 0),
    FOREIGN KEY ("owner") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "files_name_idx" ON "files" ( "name" );


CREATE TABLE IF NOT EXISTS "folders"
(
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "name" CHAR(256) NOT NULL,
    "color" CHAR(256) NOT NULL,
    "owner" INTEGER NOT NULL,
    "created_at" DATETIME DEFAULT current_timestamp,
    FOREIGN KEY ("owner") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "folders_name_idx" ON "folders" ( "name" );

-- BOOKSHELF
CREATE TABLE IF NOT EXISTS "categories"
(
    "category_name" CHAR(32) NOT NULL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "created_at" DATETIME DEFAULT current_timestamp,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "categories_category_name_idx" ON "categories" ( "category_name" );

CREATE TABLE IF NOT EXISTS "books"
(
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "name" CHAR(256) NOT NULL,
    "year" INTEGER,
    "publisher_id" INTEGER,
    "edition" SMALLINT,
    "file_id" INTEGER,
    "isbn" TEXT,
    FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON DELETE SET NULL
);
CREATE INDEX IF NOT EXISTS "books_name_idx" ON "books" ( "name" );


CREATE TABLE IF NOT EXISTS "books_categories"
(
    "book_id" INTEGER NOT NULL,
    "category" CHAR(32) NOT NULL,
    FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("category") REFERENCES "categories" ("category_name") ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS "authors"
(
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "fullname" CHAR(256) NOT NULL,
    "birthday" INTEGER,
    "death" INTEGER,
    "user_id" INTEGER NOT NULL,
    "created_at" DATETIME DEFAULT current_timestamp,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS "authors_name_idx" ON "authors" ("fullname");

CREATE TABLE IF NOT EXISTS "author_books"
(
    "author_id" INTEGER NOT NULL,
    "book_id" INTEGER NOT NULL,
    FOREIGN KEY ("author_id") REFERENCES "authors" ("id") ON DELETE SET NULL,
    FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE SET NULL
);

-- -- SYSTEM ACCESS RELATED
CREATE TABLE IF NOT EXISTS "access_log"
(
    "file_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "access_time" DATETIME DEFAULT current_timestamp,
    FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);


-- -- TOOD: videos
-- -- This app will be able to store not only books, but also videos. Why?
-- -- Because on my laptop I have 4-5 downloaded courses and it will be convinuent
-- -- to store them to. May be I can create some streaming service to watch them from 
-- -- device.

-- -- TODO:
-- --      - mock data
-- --      - develop and create functions