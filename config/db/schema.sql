-- SYSTEM
CREATE TABLE IF NOT EXISTS "application"."users"
(
    "id" SERIAL PRIMARY KEY,
    "email" TEXT NOT NULL UNIQUE,
    "username" TEXT NOT NULL UNIQUE,
    "active" BOOLEAN NOT NULL DEFAULT true,
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT current_timestamp,
    "updated_at" TIMESTAMP DEFAULT current_timestamp
);

-- CREATE INDEX IF NOT EXISTS "users_email_idx" ON "users" ( "email" );


-- CREATE TABLE IF NOT EXISTS "application"."files"
-- (
--     "id" SERIAL,
--     "name" TEXT NOT NULL,
--     "mime_type" TEXT NOT NULL,
--     "size" INTEGER NOT NULL,
--     "system_path" TEXT,
--     "user_id" INTEGER NOT NULL,
--     "hash" TEXT NOT NULL,
--     "public" BOOLEAN NOT NULL DEFAULT true,
--     "folder" INTEGER,
--     "created_at" TIMESTAMP DEFAULT current_timestamp,
--     PRIMARY KEY (id),
--     CHECK ("size" > 0),
--     FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
-- );

-- CREATE INDEX IF NOT EXISTS "files_name_idx" ON "files" ( "name" );


-- CREATE TABLE IF NOT EXISTS "application"."folders"
-- (
--     "id" SERIAL,
--     "name" TEXT NOT NULL,
--     "color" TEXT NOT NULL,
--     "user_id" INTEGER NOT NULL,
--     "created_at" TIMESTAMP DEFAULT current_timestamp,
--     PRIMARY KEY (id),
--     FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
-- );

-- CREATE INDEX IF NOT EXISTS "folders_name_idx" ON "folders" ( "name" );

-- -- BOOKSHELF
-- -- CREATE TABLE IF NOT EXISTS "categories"
-- -- (
-- --     "category_name" CHAR(32) NOT NULL PRIMARY KEY,
-- --     "user_id" INTEGER NOT NULL,
-- --     "created_at" TIMESTAMP DEFAULT current_timestamp,
-- --     FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
-- -- );

-- -- CREATE INDEX IF NOT EXISTS "categories_category_name_idx" ON "categories" ( "category_name" );

-- -- CREATE TABLE IF NOT EXISTS "books"
-- -- (
-- --     "id" SERIAL,
-- --     "name" CHAR(256) NOT NULL,
-- --     "year" INTEGER,
-- --     "publisher_id" INTEGER,
-- --     "edition" SMALLINT,
-- --     "file_id" INTEGER,
-- --     "isbn" TEXT,
-- --     FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON DELETE SET NULL
-- -- );
-- -- CREATE INDEX IF NOT EXISTS "books_name_idx" ON "books" ( "name" );


-- -- CREATE TABLE IF NOT EXISTS "books_categories"
-- -- (
-- --     "book_id" INTEGER NOT NULL,
-- --     "category" CHAR(32) NOT NULL,
-- --     FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE,
-- --     FOREIGN KEY ("category") REFERENCES "categories" ("category_name") ON DELETE CASCADE
-- -- );


-- -- CREATE TABLE IF NOT EXISTS "authors"
-- -- (
-- --     "id" SERIAL,
-- --     "fullname" CHAR(256) NOT NULL,
-- --     "birthday" INTEGER,
-- --     "death" INTEGER,
-- --     "user_id" INTEGER NOT NULL,
-- --     "created_at" TIMESTAMP DEFAULT current_timestamp,
-- --     FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
-- -- );
-- -- CREATE INDEX IF NOT EXISTS "authors_name_idx" ON "authors" ("fullname");

-- -- CREATE TABLE IF NOT EXISTS "author_books"
-- -- (
-- --     "author_id" INTEGER NOT NULL,
-- --     "book_id" INTEGER NOT NULL,
-- --     FOREIGN KEY ("author_id") REFERENCES "authors" ("id") ON DELETE SET NULL,
-- --     FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE SET NULL
-- -- );

-- -- -- SYSTEM ACCESS RELATED
-- -- CREATE TABLE IF NOT EXISTS "access_log"
-- -- (
-- --     "file_id" INTEGER NOT NULL,
-- --     "user_id" INTEGER NOT NULL,
-- --     "access_time" TIMESTAMP DEFAULT current_timestamp,
-- --     FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON DELETE CASCADE,
-- --     FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
-- -- );
