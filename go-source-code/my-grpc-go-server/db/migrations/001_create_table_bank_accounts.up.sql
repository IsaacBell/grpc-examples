CREATE TABLE IF NOT EXISTS bank_accounts(
  account_uuid 			  UUID 			      PRIMARY KEY,
  account_number 		  VARCHAR(20) 	  UNIQUE NOT NULL,
  account_name 			  VARCHAR(100) 	  NOT NULL,
  currency            VARCHAR(5)      NOT NULL,    
  current_balance 		NUMERIC(15,2) 	NOT NULL,
  created_at 			    TIMESTAMPTZ,
  updated_at 			    TIMESTAMPTZ
);
