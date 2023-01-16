-- SYSTEM
CREATE TABLE "users"
(
    "id" SERIAL PRIMARY KEY,
    "email" CHAR(256) NOT NULL,
    "hashed_password" TEXT NOT NULL,
    "active" BOOLEAN NOT NULL DEFAULT true,
    "created_at" TIMESTAMP DEFAULT Now()
);

CREATE INDEX "users_email_idx" ON "users" ( "email" );


CREATE TABLE "files"
(
    "id" SERIAL PRIMARY KEY,
    "name" CHAR(256) NOT NULL UNIQUE,
    "mime_type" CHAR(128) NOT NULL,
    "size" INTEGER NOT NULL,
    "system_path" CHAR(256),
    "owner" INTEGER NOT NULL,
    "public" BOOLEAN NOT NULL DEFAULT true,
    "created_at" TIMESTAMP DEFAULT Now()
    CHECK ("size" > 0),
    FOREIGN KEY ("owner") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE INDEX "files_name_idx" ON "files" ( "name" );


-- BOOKSHELF
CREATE TABLE "categories"
(
    "category_name" CHAR(32) NOT NULL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT Now(),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);

CREATE INDEX "categories_category_name_idx" ON "categories" ( "category_name" );

CREATE TABLE "books"
(
    "id" SERIAL PRIMARY KEY,
    "name" CHAR(256) NOT NULL,
    "year" INTEGER,
    "publisher_id" INTEGER,
    "edition" SMALLINT,
    "file_id" INTEGER,
    "isbn" TEXT,
    FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON DELETE SET NULL
);
CREATE INDEX "books_name_idx" ON "books" ( "name" );


CREATE TABLE "books_categories"
(
    "book_id" INTEGER NOT NULL,
    "category" CHAR(32) NOT NULL,
    FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("category") REFERENCES "categories" ("category_name") ON DELETE CASCADE
);


CREATE TABLE "authors"
(
    "id" SERIAL PRIMARY KEY,
    "fullname" CHAR(256) NOT NULL,
    "birthday" INTEGER,
    "death" INTEGER,
    "user_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT Now(),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);
CREATE INDEX "authors_name_idx" ON "authors" ("fullname");

CREATE TABLE "author_books"
(
    "author_id" INTEGER NOT NULL,
    "book_id" INTEGER NOT NULL,
    FOREIGN KEY ("author_id") REFERENCES "authors" ("id") ON DELETE SET NULL,
    FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE SET NULL
);

-- SYSTEM ACCESS RELATED
CREATE TABLE "access_log"
(
    "file_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "access_time" TIMESTAMP DEFAULT Now(),
    FOREIGN KEY ("file_id") REFERENCES "files" ("id") ON DELETE CASCADE,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE
);