CREATE TABLE tx
(  tx_hash TEXT  NOT NULL PRIMARY KEY,
  gas_limit BIGINT NOT NULL ,
  gas_price BIGINT NOT NULL ,
  gas_used BIGINT NOT NULL ,
  mini_block_hash TEXT  NOT NULL ,
  nonce BIGINT NOT NULL ,
  receiver TEXT  NOT NULL ,
  receiver_shard BIGINT NOT NULL ,
  round BIGINT NOT NULL ,
  sender TEXT  NOT NULL ,
  sender_shard BIGINT NOT NULL ,
  signature TEXT  NOT NULL ,
  status TEXT  NOT NULL ,
  value TEXT  NOT NULL ,
  fee TEXT  NOT NULL ,
  timestamp BIGINT NOT NULL ,
  data TEXT  NOT NULL 
  
);

CREATE TABLE smart_contract_result
(  tx_hash TEXT  NOT NULL ,
  hash TEXT  NOT NULL ,
  timestamp BIGINT NOT NULL ,
  nonce BIGINT NOT NULL ,
  gas_limit BIGINT NOT NULL ,
  gas_price BIGINT NOT NULL ,
  value TEXT  NOT NULL ,
  sender TEXT  NOT NULL ,
  receiver TEXT  NOT NULL ,
  relayed_value TEXT ,
  data TEXT  NOT NULL ,
  prev_tx_hash TEXT  NOT NULL ,
  original_tx_hash TEXT  NOT NULL ,
  call_type TEXT  NOT NULL ,
  logs JSONB NOT NULL DEFAULT '[]'::JSONB
);