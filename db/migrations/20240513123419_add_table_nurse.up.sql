CREATE TABLE IF NOT EXISTS "nurse" (
    "id" uuid UNIQUE NOT NULL DEFAULT (gen_random_uuid()) PRIMARY KEY,
    "nip" bigint UNIQUE NOT NULL,
    "name" varchar(50) NOT NULL,
    "password" varchar(255),
    "isGranted" boolean NOT NULL,
    "identityCardScanning" varchar(255) NOT NULL,
    "createdAt" timestamp NOT NULL,
    "updatedAt" timestamp NOT NULL,
    "deletedAt" timestamp
);