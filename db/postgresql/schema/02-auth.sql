CREATE TABLE account
(  address TEXT NOT NULL ,
  balance TEXT NOT NULL ,
  nonce TEXT NOT NULL ,
  shard BIGINT  NOT NULL ,
  scam_info JSONB ,
  code TEXT,
  code_hash TEXT ,
  root_hash TEXT NOT NULL ,
  tx_count BIGINT  NOT NULL ,
  scr_count BIGINT  NOT NULL ,
  username TEXT,
  developer_reward TEXT NOT NULL ,
  owner_address TEXT  ,
  deployed_at BIGINT  ,
  is_upgradeable TEXT  ,
  is_readable TEXT  ,
  is_payable TEXT  ,
  is_payable_by_smart_contract TEXT 
);

   CREATE TABLE token_balance
(  address TEXT NOT NULL ,
  identifier TEXT NOT NULL ,
  balance TEXT NOT NULL
);