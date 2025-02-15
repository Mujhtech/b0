ALTER TABLE "users" ADD COLUMN "subscription_plan" TEXT NOT NULL DEFAULT 'free';

ALTER TABLE "users" ADD COLUMN "stripe_customer_id" TEXT NULL DEFAULT NULL;

ALTER TABLE "users" ADD COLUMN "stripe_subscription_id" TEXT NULL DEFAULT NULL;

ALTER TABLE "users" ADD COLUMN "stripe_subscription_status" TEXT NULL DEFAULT NULL;

