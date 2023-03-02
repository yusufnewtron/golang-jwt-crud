/*
PostgreSQL Backup
Database: tesapi/public
Backup Time: 2023-03-02 15:11:29
*/

DROP SEQUENCE IF EXISTS "public"."informasi_informasiid_seq";
DROP SEQUENCE IF EXISTS "public"."items_Id_seq";
DROP SEQUENCE IF EXISTS "public"."m_user_id_user_seq";
DROP SEQUENCE IF EXISTS "public"."sequence_id_books";
DROP SEQUENCE IF EXISTS "public"."tblbooks_id_seq";
DROP TABLE IF EXISTS "public"."items";
DROP TABLE IF EXISTS "public"."m_user";
CREATE SEQUENCE "informasi_informasiid_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
CREATE SEQUENCE "items_Id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
CREATE SEQUENCE "m_user_id_user_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 32767
START 1
CACHE 1;
CREATE SEQUENCE "sequence_id_books" 
INCREMENT 3
MINVALUE  1
MAXVALUE 100
START 3
CACHE 1;
CREATE SEQUENCE "tblbooks_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE TABLE "items" (
  "id" int4 NOT NULL DEFAULT nextval('"items_Id_seq"'::regclass),
  "itemname" text COLLATE "pg_catalog"."default",
  "price" int4,
  "createddate" timestamptz(6) DEFAULT now()
)
;
ALTER TABLE "items" OWNER TO "postgres";
CREATE TABLE "m_user" (
  "id_user" int2 NOT NULL GENERATED ALWAYS AS IDENTITY (
INCREMENT 1
MINVALUE  1
MAXVALUE 32767
START 1
),
  "username" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "fullname" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "password" text COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "m_user" OWNER TO "postgres";
BEGIN;
LOCK TABLE "public"."items" IN SHARE MODE;
DELETE FROM "public"."items";
INSERT INTO "public"."items" ("id","itemname","price","createddate") VALUES (2, 'Pencil', 2000, '2023-03-01 14:00:00+07'),(3, 'Bolpoin', 2500, '2023-03-01 14:00:00+07'),(5, 'Stabilo', 6000, '2023-03-01 14:00:00+07'),(6, 'Tipe X', 8000, '2023-03-01 14:00:00+07'),(7, 'Penggaris', 2000, '2023-03-01 14:00:00+07'),(8, 'Spidol', 7000, '2023-03-01 14:00:00+07'),(9, 'Kapur', 15000, '2023-03-01 14:00:00+07'),(10, 'Penghapus', 4000, '2023-03-01 14:00:00+07'),(1, 'Buku', 5000, '2023-03-01 14:00:00+07');
COMMIT;
BEGIN;
LOCK TABLE "public"."m_user" IN SHARE MODE;
DELETE FROM "public"."m_user";
INSERT INTO "public"."m_user" ("id_user","username","fullname","password") VALUES (2, 'user', 'User', '$2a$10$3wrHYiQWSJjWUZuqgw1.wemGvbyBBpPQ7CzSodWmkQ1Z8PMFl3TMu'),(5, 'andi', 'Andi', '$2a$10$OKCmjBy77BPoie6zx03RoO2v7QLffbb69HUg9WCrIVXzh4UxnJ8ca'),(6, 'usrok', 'Usrok', '$2a$10$Kyu/nXHXwb0M1oixw.WZgOVr6BiDnWjay5aOQvIQ7r/XCUr85E1RC'),(7, 'admin', 'Admin', '$2a$10$RPjWY7BkD4IpxecRR2EBPuoAD2fYasVHERvgaM7ieCj0UNGjSd8LS');
COMMIT;
ALTER TABLE "items" ADD CONSTRAINT "items_pkey" PRIMARY KEY ("id");
ALTER TABLE "m_user" ADD CONSTRAINT "m_user_pkey" PRIMARY KEY ("id_user");
SELECT setval('"informasi_informasiid_seq"', 2, false);
ALTER SEQUENCE "informasi_informasiid_seq" OWNER TO "postgres";
ALTER SEQUENCE "items_Id_seq"
OWNED BY "items"."id";
SELECT setval('"items_Id_seq"', 13, true);
ALTER SEQUENCE "items_Id_seq" OWNER TO "postgres";
ALTER SEQUENCE "m_user_id_user_seq"
OWNED BY "m_user"."id_user";
SELECT setval('"m_user_id_user_seq"', 8, true);
ALTER SEQUENCE "m_user_id_user_seq" OWNER TO "postgres";
SELECT setval('"sequence_id_books"', 6, false);
ALTER SEQUENCE "sequence_id_books" OWNER TO "postgres";
SELECT setval('"tblbooks_id_seq"', 5, true);
ALTER SEQUENCE "tblbooks_id_seq" OWNER TO "postgres";
