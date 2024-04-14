CREATE TABLE IF NOT EXISTS bank_transfers(
    transfer_uuid           UUID            PRIMARY KEY,
    from_account_uuid       UUID            REFERENCES bank_accounts,
    to_account_uuid         UUID            REFERENCES bank_accounts,
    currency                VARCHAR(5)      NOT NULL,
    amount                  NUMERIC(15,2)   NOT NULL,
    transfer_timestamp      TIMESTAMPTZ     NOT NULL,
    transfer_success        BOOLEAN         NOT NULL DEFAULT FALSE,
    created_at 			        TIMESTAMPTZ,
    updated_at 			        TIMESTAMPTZ
);