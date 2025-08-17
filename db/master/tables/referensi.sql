-- master.referensi definition
USE `master`;

CREATE TABLE `referensi` (
  `JENIS` tinyint(4) NOT NULL,
  `ID` smallint NOT NULL AUTO_INCREMENT,
  `DESKRIPSI` varchar(150) NOT NULL,
  `STATUS` tinyint(4) NOT NULL DEFAULT 1,
  PRIMARY KEY (`JENIS`,`ID`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=latin1 ROW_FORMAT=DYNAMIC;

INSERT INTO master.referensi (JENIS, ID, DESKRIPSI, STATUS)
VALUES 
(1, 1, 'RSU', 1),
(1, 2, 'RSB', 1),
(1, 3, 'RSIA', 1),
(1, 4, 'Balai', 1),
(1, 5, 'Klinik', 1),
(1, 6, 'PKM', 1),
(1, 7, 'RSK', 1),
(1, 8, 'RSGM', 1),
(1, 9, 'PSC', 1),
(1, 10, 'DInas Kota/Kab', 1),
(1, 11, 'Dinas Propinsi', 1),
(1, 99, 'Kemkes', 1),
(2, 1, 'Laki - Laki', 1),
(2, 2, 'Perempuan', 1);