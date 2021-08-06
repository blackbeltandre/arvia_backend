-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 06 Agu 2021 pada 20.28
-- Versi server: 8.0.25
-- Versi PHP: 7.4.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `arvia`
--

-- --------------------------------------------------------

--
-- Stand-in struktur untuk tampilan `last_stok`
-- (Lihat di bawah untuk tampilan aktual)
--
CREATE TABLE `last_stok` (
`sisa` bigint
,`laba_kotor` bigint
,`harga_jual` int
,`qty` int
,`penjualan` varchar(255)
,`lokasi` varchar(255)
,`kd_brg` varchar(10)
,`stok_awal` int
,`lokasi_awal` varchar(5)
,`modal_awal` bigint
);

-- --------------------------------------------------------

--
-- Struktur dari tabel `stoks`
--

CREATE TABLE `stoks` (
  `id_stok` int NOT NULL DEFAULT '0',
  `tgl_beli` varchar(10) NOT NULL,
  `kd_brg` varchar(10) NOT NULL,
  `qty` int NOT NULL,
  `harga` bigint NOT NULL,
  `lokasi` varchar(5) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `stoks`
--

INSERT INTO `stoks` (`id_stok`, `tgl_beli`, `kd_brg`, `qty`, `harga`, `lokasi`) VALUES
(1, '12/7/21', 'Barang A', 5, 1000, 'X'),
(2, '20/7/21', 'Barang A', 5, 1000, 'Y'),
(3, '12/7/21', 'Barang C', 5, 800, 'X'),
(4, '12/7/21', 'Barang B', 5, 500, 'Y'),
(5, '12/7/21', 'Barang D', 5, 1500, 'Y'),
(6, '15/7/21', 'Barang B', 5, 550, 'X'),
(7, '15/7/21', 'Barang A', 5, 1200, 'X'),
(8, '12/8/21', 'Barang B', 5, 500, 'X'),
(9, '12/8/21', 'Barang D', 5, 1500, 'X'),
(10, '10/7/21', 'Barang C', 5, 750, 'Y'),
(11, '20/7/21', 'Barang B', 5, 600, 'Y');

-- --------------------------------------------------------

--
-- Stand-in struktur untuk tampilan `temp`
-- (Lihat di bawah untuk tampilan aktual)
--
CREATE TABLE `temp` (
`quantity` decimal(32,0)
,`penjualan` varchar(255)
,`lokasi` varchar(255)
,`harga_jual` int
);

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `id` int NOT NULL,
  `tgl_transaksi` varchar(255) DEFAULT NULL,
  `penjualan` varchar(255) DEFAULT NULL,
  `qty` int DEFAULT NULL,
  `harga_jual` int DEFAULT NULL,
  `lokasi` varchar(255) DEFAULT NULL,
  `modal` int DEFAULT NULL,
  `laba_kotor` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`id`, `tgl_transaksi`, `penjualan`, `qty`, `harga_jual`, `lokasi`, `modal`, `laba_kotor`) VALUES
(1, '16/7/21', 'Barang A', 7, 1300, 'X', NULL, NULL),
(26, '11/7/21', 'Barang C', 8, 900, 'Y', NULL, NULL),
(27, '21/7/21', 'Barang B', 9, 700, 'X', NULL, NULL),
(28, '15/8/21', 'Barang D', 3, 1700, 'X', NULL, NULL),
(29, '22/7/21', 'Barang A', 2, 1300, 'X', NULL, NULL),
(34, '22/7/21', 'Barang B', 5, 500, 'Y', NULL, NULL);

-- --------------------------------------------------------

--
-- Struktur untuk view `last_stok`
--
DROP TABLE IF EXISTS `last_stok`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `last_stok`  AS SELECT (`a`.`qty` - `b`.`qty`) AS `sisa`, (`a`.`harga_jual` - `b`.`harga`) AS `laba_kotor`, `a`.`harga_jual` AS `harga_jual`, `a`.`qty` AS `qty`, `a`.`penjualan` AS `penjualan`, `a`.`lokasi` AS `lokasi`, `b`.`kd_brg` AS `kd_brg`, `b`.`qty` AS `stok_awal`, `b`.`lokasi` AS `lokasi_awal`, `b`.`harga` AS `modal_awal` FROM (`transactions` `a` left join `stoks` `b` on(((`a`.`penjualan` = `b`.`kd_brg`) and (`a`.`lokasi` = `b`.`lokasi`)))) GROUP BY `a`.`harga_jual`, `b`.`harga`, `a`.`qty`, `a`.`penjualan`, `a`.`lokasi`, `b`.`kd_brg`, `b`.`qty`, `b`.`lokasi`, `b`.`harga` ;

-- --------------------------------------------------------

--
-- Struktur untuk view `temp`
--
DROP TABLE IF EXISTS `temp`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `temp`  AS SELECT sum(`transactions`.`qty`) AS `quantity`, `transactions`.`penjualan` AS `penjualan`, `transactions`.`lokasi` AS `lokasi`, `transactions`.`harga_jual` AS `harga_jual` FROM `transactions` GROUP BY `transactions`.`penjualan`, `transactions`.`lokasi`, `transactions`.`harga_jual` ORDER BY `transactions`.`penjualan` ASC, `transactions`.`lokasi` ASC, `transactions`.`harga_jual` ASC ;

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
