ALTER TABLE IF EXISTS "MenuItem" 
ALTER COLUMN "status" DROP NOT NULL;

ALTER TABLE IF EXISTS "Menu" 
ALTER COLUMN "status" DROP NOT NULL;

ALTER TABLE IF EXISTS "Ingredient" 
ALTER COLUMN "status" DROP NOT NULL;

ALTER TABLE IF EXISTS "MenuItem_Ingredient" 
ALTER COLUMN "status" DROP NOT NULL;

ALTER TABLE IF EXISTS "Menu_MenuItem" 
ALTER COLUMN "status" DROP NOT NULL;