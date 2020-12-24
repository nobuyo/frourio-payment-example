CREATE TABLE "Purchase" (
  "id" SERIAL,
  "itemId" integer NOT NULL,
  "amount" integer NOT NULL,

  PRIMARY KEY ("id")
);
