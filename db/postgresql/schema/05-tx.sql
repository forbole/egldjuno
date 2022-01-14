CREATE TABLE tx
(  tx_hash TEXT  NOT NULL ,
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
(  tx_hash BIGINT  NOT NULL ,
  hash BIGINT  NOT NULL ,
  timestamp TEXT NOT NULL ,
  nonce TEXT NOT NULL ,
  gas_limit TEXT NOT NULL ,
  gas_price TEXT NOT NULL ,
  value BIGINT  NOT NULL ,
  sender BIGINT  NOT NULL ,
  receiver BIGINT  NOT NULL ,
  relayed_value BIGINT  NOT NULL ,
  data BIGINT  NOT NULL ,
  prev_tx_hash BIGINT  NOT NULL ,
  original_tx_hash BIGINT  NOT NULL ,
  call_type BIGINT  NOT NULL ,
  logs BIGINT  NOT NULL
);