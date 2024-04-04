-- ************ Start : Copy and paste in dbdiagram.io ***************

-- Table accounts as A {
--     id bigserial [pk] 
--     owner varchar [not null]
--     balance bigint [not null]
--     currency varchar [not null]
--     created_at timestamptz [not null , default : `now()`]

--     Indexes {
--         owner
--     }
-- }

-- Table entries {
--     id bigserial [pk]
--     account_id bigint [ref: > A.id] -- foreign key
--     amount bigint [not null , note : 'it can be negative or positive']
--     created_at timestamptz [not null , default : `now()`]

--     Indexes {
--         account_id
--     }
-- }

-- Table transfers {
--     id bigserial [pk]
--     from_account_id bigint [ref : > A.id]
--     to_account_id bigint [ref : > A.id]
--     amount bigint [not null , note : 'it must be positve']
--     created_at timestamptz [not null, default : `now()`]

--     Indexes {
--         from_account_id
--         to_account_id
--         (from_account_id , to_account_id)
--     }
-- }


-- ************ End : Copy and paste in dbdiagram.io ***************

-- Enum Currency { -- let the application handle validation
--     USD
--     EUR
-- }

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint,
  "to_account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'it can be negative or positive';

COMMENT ON COLUMN "transfers"."amount" IS 'it must be positve';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
