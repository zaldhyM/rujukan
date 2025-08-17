USE `master`;

CREATE TABLE `jenis_referensi` (
  `ID` tinyint(4) NOT NULL AUTO_INCREMENT,
  `DESKRIPSI` varchar(45) NOT NULL,
  `SINGKATAN` char(5) NOT NULL,
  `APLIKASI` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

INSERT INTO master.jenis_referensi (ID, DESKRIPSI, SINGKATAN, APLIKASI)
VALUES 
(1, 'Kelompok Fakses', '', 0),
(2, 'Jenis Kelamin', '', 0);