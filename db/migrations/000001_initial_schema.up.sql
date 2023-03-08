CREATE TABLE "market" (
  "id" serial PRIMARY KEY,
  "name" varchar(50)
);

CREATE TABLE "market_days" (
  "id" bigserial,
  "date" date unique,
  "open" real NOT NULL,
  "high" real NOT NULL,
  "low" real NOT NULL,
  "last" real NOT NULL,
  "range" real  NOT NULL,
  "volume" real NOT NULL,
  "poc3yr" real NOT NULL,
  "poc1yr" real NOT NULL,
  "poc0yr" real NOT NULL,   
  "poc1wk" real NOT NULL,
  "poc1m" real NOT NULL,
  "market_id" int,
  PRIMARY KEY ("id", "date")
);

CREATE TABLE "day_info"(
  "id" bigserial, 
  "date" date unique,
  "day_range" real NOT NULL,
  "week_range" real NOT NULL,
  "month_range" real NOT NULL,
  "year_range" real NOT NULL,
  "market_day_date" int NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "open_prices"(
  "id" bigserial,
  "market" varchar(20) NOT NULL,
  "year_open" real NOT NULL,
  "month_open" real NOT NULL,
  "week_open" real NOT NULL,
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




