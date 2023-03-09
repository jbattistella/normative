CREATE TABLE "market" (
  "id" serial PRIMARY KEY,
  "name" varchar(50)
);

CREATE TABLE "market_days" (
  "id" bigserial,
  "date" date NOT NULL,
  "open" float NOT NULL,
  "high" float NOT NULL,
  "low" float NOT NULL,
  "last" float NOT NULL,
  "range" float  NOT NULL,
  "volume" float NOT NULL,
  "market" varchar(50) NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "open_prices"(
  "id" bigserial,
  "market" varchar(20) NOT NULL,
  "year_open" float NOT NULL,
  "month_open" float NOT NULL,
  "week_open" float NOT NULL,
  "updated" date NOT NULL,
  PRIMARY KEY ("id")
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




