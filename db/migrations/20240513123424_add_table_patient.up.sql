CREATE TYPE gender AS ENUM (
    'MALE',
    'FEMALE'
);

CREATE TABLE IF NOT EXISTS "patient" (
    "id" uuid UNIQUE NOT NULL DEFAULT (gen_random_uuid()) PRIMARY KEY,
    "identityNumber" bigint UNIQUE NOT NULL,
    "name" varchar(50) NOT NULL,
    "phoneNumber" varchar(20),
    "birthDate" DATE NOT NULL,
    "gender" gender NOT NULL,
    "identityCardScanning" varchar(255) NOT NULL,
    "createdAt" timestamp NOT NULL
);