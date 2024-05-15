CREATE TABLE IF NOT EXISTS "user" (
    "id" uuid UNIQUE NOT NULL DEFAULT (gen_random_uuid()) PRIMARY KEY,
    "nip" bigint UNIQUE NOT NULL,
    "name" varchar(50) NOT NULL,
    "password" varchar(255),
    "identityCardScanning" varchar(255),
    "createdAt" timestamp NOT NULL,
    "updatedAt" timestamp NOT NULL
);