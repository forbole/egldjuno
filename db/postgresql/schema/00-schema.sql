CREATE TABLE block
(  hash TEXT  NOT NULL ,
  epoch BIGINT NOT NULL ,
  nonce BIGINT NOT NULL ,
  prev_hash TEXT  NOT NULL ,
  proposer BIGINT  NOT NULL ,
  pub_key_bitmap TEXT  NOT NULL ,
  round BIGINT NOT NULL ,
  shard BIGINT NOT NULL ,
  size BIGINT NOT NULL ,
  size_txs BIGINT NOT NULL ,
  state_root_hash TEXT  NOT NULL ,
  time_stamp BIGINT NOT NULL ,
  tx_count BIGINT NOT NULL ,
  gas_consumed BIGINT NOT NULL ,
  gas_refunded BIGINT NOT NULL ,
  gas_penalized BIGINT NOT NULL ,
  max_gas_limit BIGINT NOT NULL
);
CREATE INDEX block_round_index ON block (round);
