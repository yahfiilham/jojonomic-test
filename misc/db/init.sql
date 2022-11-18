CREATE TABLE IF NOT EXISTS public.user (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  nama varchar(100) NOT NULL,
  role varchar(15) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);

CREATE TABLE IF NOT EXISTS harga (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  admin_id VARCHAR (15) NOT NULL,
  harga_topup DECIMAL(12, 3) NOT NULL,
  harga_buyback DECIMAL(12, 3) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);  


CREATE TABLE IF NOT EXISTS rekening (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  no_rek varchar(20) NOT NULL,
  saldo DECIMAL(12, 3) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);

CREATE TABLE IF NOT EXISTS topup (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  harga varchar(100) NOT NULL,
  gram varchar(100) NOT NULL,
  norek varchar(20) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);

CREATE TABLE IF NOT EXISTS transaksi (
  reff_id VARCHAR (15) PRIMARY KEY NOT NULL,
  type varchar(20) NOT NULL,
  no_rek varchar(20) NOT NULL,
  saldo DECIMAL(12, 3) NOT NULL,
  gram DECIMAL(12, 3) NOT NULL,
  harga_topup DECIMAL(12, 3) NOT NULL,
  harga_buyback DECIMAL(12, 3) NOT NULL,
  created_at INT NULL DEFAULT EXTRACT(epoch FROM NOW()),
  updated_at INT NULL DEFAULT EXTRACT(epoch FROM NOW())
);

INSERT INTO public.user (reff_id, nama, role, created_at, updated_at) 
VALUES ('a001','admin1', 'admin', EXTRACT(epoch FROM NOW()), EXTRACT(epoch FROM NOW())),
('6pyXwTHVg','customer1', 'customer', EXTRACT(epoch FROM NOW()), EXTRACT(epoch FROM NOW()));

INSERT INTO rekening (reff_id, no_rek, saldo, created_at, updated_at) 
VALUES ('xzlcmANVR','r001', 0, EXTRACT(epoch FROM NOW()), EXTRACT(epoch FROM NOW()));