CREATE TABLE "Item" (
  "id" SERIAL,
  "name" varchar(255) NOT NULL,
  "description" varchar(255),
  "amount" integer NOT NULL,

  PRIMARY KEY ("id")
);
