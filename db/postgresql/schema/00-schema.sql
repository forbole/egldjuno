CREATE TABLE block
(  hash BIGINT  NOT NULL ,
  epoch TEXT NOT NULL ,
  nonce TEXT NOT NULL ,
  prev_hash BIGINT  NOT NULL ,
  proposer BIGINT  NOT NULL ,
  pub_key_bitmap BIGINT  NOT NULL ,
  round TEXT NOT NULL ,
  shard TEXT NOT NULL ,
  size TEXT NOT NULL ,
  size_txs TEXT NOT NULL ,
  state_root_hash BIGINT  NOT NULL ,
  time_stamp TEXT NOT NULL ,
  tx_count TEXT NOT NULL ,
  gas_consumed TEXT NOT NULL ,
  gas_refunded TEXT NOT NULL ,
  gas_penalized TEXT NOT NULL ,
  max_gas_limit TEXT NOT NULL
);
CREATE INDEX block_round_index ON block (round);