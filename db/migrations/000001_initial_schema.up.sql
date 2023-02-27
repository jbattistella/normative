CREATE TABLE "market" (
  "id" serial PRIMARY KEY,
  "name" varchar(50)
);

CREATE TABLE "intraday" (
  "id" bigserial,
  "datetime" timestamp with time zone unique,
  "open" float,
  "high" float,
  "low" float,
  "close" float,
  "volume" int,
  "bidvolume" int,
  "askvolume" int,
  "date" date,
  "market_id" int,
  PRIMARY KEY ("id", "datetime")
);

CREATE TABLE "end_of_day" (
  "id" bigserial,
  "date" date unique,
  "open" float,
  "high" float,
  "low" float,
  "close" float,
  "volume" int,
  "bidvolume" int,
  "askvolume" int,
  "market_id" int,
  PRIMARY KEY ("id", "date")
  
);

CREATE TABLE "events" (
  "id" serial PRIMARY KEY,
  "date" date NOT NULL,
  "time" time NOT NULL,
  "forecast" varchar(50) NOT NULL,
  "impact" varchar(50) NOT NULL,
  "last_update" int NOT NULL,
  "name" varchar(100) NOT NULL,
  "previous" varchar(50) NOT NULL,
  "region" varchar(50) NOT NULL
);

ALTER TABLE "intraday" ADD FOREIGN KEY ("market_id") REFERENCES "market" ("id");

ALTER TABLE "end_of_day" ADD FOREIGN KEY ("market_id") REFERENCES "market" ("id");

ALTER TABLE "intraday" ADD FOREIGN KEY ("date") REFERENCES "end_of_day" ("date");

