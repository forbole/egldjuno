CREATE TABLE account
(  address TEXT  NOT NULL ,
  nonce BIGINT NOT NULL ,
  balance TEXT  NOT NULL ,
  root_hash TEXT  NOT NULL ,
  tx_count BIGINT NOT NULL ,
  scr_count BIGINT NOT NULL ,
  shard BIGINT NOT NULL ,
  developer_reward TEXT  NOT NULL
);