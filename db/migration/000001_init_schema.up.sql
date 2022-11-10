CREATE TABLE "Accounts" (
                            "id" bigserial PRIMARY KEY,
                            "balance" bigint NOT NULL DEFAULT 0,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Services" (
                            "id" bigserial PRIMARY KEY,
                            "name" varchar UNIQUE NOT NULL,
                            "price" bigint NOT NULL
);

CREATE TABLE "Orders" (
                          "id" bigserial PRIMARY KEY,
                          "id_account" bigserial NOT NULL,
                          "id_service" bigserial NOT NULL,
                          "price_service" bigint NOT NULL,
                          "status" varchar NOT NULL,
                          "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "History" (
                           "id" bigserial PRIMARY KEY,
                           "id_account" bigserial NOT NULL,
                           "amount" bigint NOT NULL,
                           "comment" varchar NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Accounts" ("id");

CREATE INDEX ON "Services" ("id");

CREATE INDEX ON "Orders" ("id");

CREATE INDEX ON "Orders" ("id_account");

CREATE INDEX ON "History" ("id_account");

ALTER TABLE "Orders" ADD FOREIGN KEY ("id_account") REFERENCES "Accounts" ("id");

ALTER TABLE "Orders" ADD FOREIGN KEY ("id_service") REFERENCES "Services" ("id");

ALTER TABLE "History" ADD FOREIGN KEY ("id_account") REFERENCES "Accounts" ("id");

INSERT INTO "Services" ("name", "price") VALUES ('Выделить цветом', 30);

INSERT INTO "Services" ("name", "price") VALUES ('XL-объявление', 30);

INSERT INTO "Services" ("name", "price") VALUES ('Увеличить количество просмотров', 220);
