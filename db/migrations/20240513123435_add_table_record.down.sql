ALTER TABLE "record" DROP CONSTRAINT IF EXISTS "record_identityNumber_fkey";
ALTER TABLE "record" DROP CONSTRAINT IF EXISTS "record_nip_fkey";

DROP TABLE IF EXISTS "record";