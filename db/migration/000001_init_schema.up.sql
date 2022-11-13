CREATE TABLE "accounts" (
                            "id" bigserial PRIMARY KEY,
                            "owner" varchar NOT NULL,
                            "balance" bigint NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "services" (
                            "id" bigserial PRIMARY KEY,
                            "name" varchar UNIQUE NOT NULL,
                            "price" bigint NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
                          "id" bigserial PRIMARY KEY,
                          "id_account" bigserial NOT NULL,
                          "id_service" bigserial NOT NULL,
                          "price_service" bigint NOT NULL,
                          "status" varchar NOT NULL,
                          "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "history" (
                           "id" bigserial PRIMARY KEY,
                           "id_account" bigserial NOT NULL,
                           "amount" bigint NOT NULL,
                           "comment" varchar NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("id");

CREATE INDEX ON "services" ("id");

CREATE INDEX ON "orders" ("id");

CREATE INDEX ON "orders" ("id_account");

CREATE INDEX ON "history" ("id_account");

ALTER TABLE "orders" ADD FOREIGN KEY ("id_account") REFERENCES "accounts" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("id_service") REFERENCES "services" ("id");

ALTER TABLE "history" ADD FOREIGN KEY ("id_account") REFERENCES "accounts" ("id");

INSERT INTO "services" ("name", "price") VALUES ('Выделить цветом', 30);

INSERT INTO "services" ("name", "price") VALUES ('XL-объявление', 30);

INSERT INTO "services" ("name", "price") VALUES ('Увеличить количество просмотров', 220);
