-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 30, 2022 at 03:30 PM
-- Server version: 10.4.18-MariaDB
-- PHP Version: 8.0.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `alta_online_shop`
--

-- --------------------------------------------------------

--
-- Table structure for table `operator`
--

CREATE TABLE `operator` (
  `id_operator` int(11) NOT NULL,
  `nama_operator` varchar(100) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `operator`
--

INSERT INTO `operator` (`id_operator`, `nama_operator`, `created_at`, `updated_at`) VALUES
(1111, 'AS', '2022-03-24', '2022-03-24 09:06:06'),
(1112, 'Simpati', '2022-03-24', '2022-03-24 09:07:06'),
(1113, 'Tri', '2022-03-24', '2022-03-24 09:08:24'),
(1114, 'Im3', '2022-03-24', '2022-03-24 09:08:43'),
(1115, 'By.U', '2022-03-24', '2022-03-24 09:09:09');

-- --------------------------------------------------------

--
-- Table structure for table `payment_method`
--

CREATE TABLE `payment_method` (
  `id_payment` int(11) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `payment_method`
--

INSERT INTO `payment_method` (`id_payment`, `nama`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Transfer via Bank', NULL, '2022-03-24', '2022-03-24 09:36:28'),
(2, 'Alfamart', NULL, '2022-03-24', '2022-03-24 09:36:28'),
(3, 'Indomaret', NULL, '2022-03-24', '2022-03-24 09:36:28');

-- --------------------------------------------------------

--
-- Table structure for table `product`
--

CREATE TABLE `product` (
  `id_product` int(11) NOT NULL,
  `id_type` int(11) DEFAULT NULL,
  `id_operator` int(11) DEFAULT NULL,
  `kode` varchar(100) DEFAULT NULL,
  `nama_produk` varchar(100) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product`
--

INSERT INTO `product` (`id_product`, `id_type`, `id_operator`, `kode`, `nama_produk`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 1113, 'ABC12345678', 'product dummy', NULL, '2022-03-24', '2022-03-29 15:00:42'),
(2, 1, 1113, 'EFG1234567', 'Pulsa 50rb', NULL, '2022-03-24', '2022-03-24 09:20:58'),
(3, 2, 1111, 'AB123CD456', 'Kuota Internet 1 Minggu 5GB', NULL, '2022-03-24', '2022-03-24 09:20:22'),
(4, 2, 1111, 'EF789GH012', 'Kuota Internet 24 Jam 10 GB', NULL, '2022-03-24', '2022-03-24 09:20:22'),
(5, 2, 1111, 'IJ345KL678', 'Kuota Nonton Unlimited 24 Jam', NULL, '2022-03-24', '2022-03-24 09:22:47'),
(6, 3, 1114, '050 680 730', 'Kuota Nelpon Sesama 30 Menit', NULL, '2022-03-24', '2022-03-24 09:25:41'),
(7, 3, 1114, '230 890 345', 'Kuota Nelpon 60 Menit', NULL, '2022-03-24', '2022-03-24 09:25:41'),
(8, 3, 1114, '560 455 230', 'Kuota Nelpon 1 Bulan', NULL, '2022-03-24', '2022-03-24 09:26:27');

-- --------------------------------------------------------

--
-- Table structure for table `product_desc`
--

CREATE TABLE `product_desc` (
  `id_product` int(11) NOT NULL,
  `description` text DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_desc`
--

INSERT INTO `product_desc` (`id_product`, `description`, `created_at`, `updated_at`) VALUES
(1, 'Pulsa 20rb ', '2022-03-24', '2022-03-24 09:28:27'),
(2, 'Pulsa 50rb', '2022-03-24', '2022-03-24 09:28:55'),
(3, 'Masa aktif 1 minggu dengan total kuota 5GB ', '2022-03-24', '2022-03-24 09:29:52'),
(4, 'Kuota internet 10GB bebas akses semua aplikasi masa aktif 24 jam, dapat diperpanjang (S&K)', '2022-03-24', '2022-03-24 09:31:25'),
(5, 'Kuota Youtube Unlimited bebas akses Youtube selama 24 jam', '2022-03-24', '2022-03-24 09:32:29'),
(6, 'Nelpon sesama operator selama 30 menit', '2022-03-24', '2022-03-24 09:33:07'),
(7, 'Nelpon bebas ke semua operator selama 60 menit', '2022-03-24', '2022-03-24 09:33:43'),
(8, 'Kuota nelpon 1 bulan bebas telpon ke operator mana saja ', '2022-03-24', '2022-03-24 09:34:28');

-- --------------------------------------------------------

--
-- Table structure for table `product_type`
--

CREATE TABLE `product_type` (
  `id_type` int(11) NOT NULL,
  `nama_produk` varchar(100) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_type`
--

INSERT INTO `product_type` (`id_type`, `nama_produk`, `created_at`, `updated_at`) VALUES
(1, 'Pulsa', '2022-03-24', '2022-03-24 09:10:27'),
(2, 'Kuota Internet', '2022-03-24', '2022-03-24 09:11:43'),
(3, 'Kuota Telpon', '2022-03-24', '2022-03-24 09:11:29');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id_transaksi` int(11) NOT NULL,
  `id_user` int(11) DEFAULT NULL,
  `id_payment` int(11) DEFAULT NULL,
  `status` varchar(100) DEFAULT NULL,
  `total_qty` int(11) DEFAULT NULL,
  `total_price` decimal(25,3) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id_transaksi`, `id_user`, `id_payment`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(1, 1, 3, 'SUKSES', 3, '200.000', '2022-03-24', '2022-03-24 09:51:09'),
(2, 1, 3, 'SUKSES', 3, '200.000', '2022-03-24', '2022-03-24 09:51:20'),
(3, 1, 3, 'SUKSES', 3, '250.000', '2022-03-24', '2022-03-24 09:51:29'),
(4, 2, 2, 'SUKSES', 3, '150.000', '2022-03-24', '2022-03-24 09:53:50'),
(5, 2, 2, 'SUKSES', 3, '200.000', '2022-03-24', '2022-03-24 09:53:50'),
(6, 2, 2, 'SUKSES', 3, '250.000', '2022-03-24', '2022-03-24 09:53:50'),
(7, 3, 1, 'SUKSES', 3, '200.000', '2022-03-24', '2022-03-24 09:55:41'),
(8, 3, 1, 'SUKSES', 3, '250.000', '2022-03-24', '2022-03-24 09:55:41'),
(9, 3, 1, 'SUKSES', 3, '150.000', '2022-03-24', '2022-03-24 09:55:41'),
(10, 4, 1, 'SUKSES', 3, '200.000', '2022-03-24', '2022-03-24 09:57:27'),
(11, 4, 1, 'SUKSES', 3, '250.000', '2022-03-24', '2022-03-24 09:57:27'),
(12, 4, 2, 'SUKSES', 3, '150.000', '2022-03-24', '2022-03-24 09:57:27'),
(13, 5, 3, 'SUKSES', 3, '250.000', '2022-03-24', '2022-03-24 09:59:37'),
(14, 5, 3, 'SUKSES', 3, '200.000', '2022-03-24', '2022-03-24 09:59:37'),
(15, 5, 1, 'PROCESS', 3, '150.000', '2022-03-24', '2022-03-24 09:59:37');

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `id_transaksi` int(11) NOT NULL,
  `id_product` int(11) NOT NULL,
  `status` varchar(100) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `price` decimal(25,3) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_details`
--

INSERT INTO `transaction_details` (`id_transaksi`, `id_product`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES
(1, 2, 'SUKSES', 2, '50.000', '2022-03-24', '2022-03-24 10:03:49'),
(1, 7, 'SUKSES', 1, '100.000', '2022-03-24', '2022-03-24 10:03:49'),
(2, 3, 'SUKSES', 2, '50.000', '2022-03-24', '2022-03-24 10:06:32'),
(2, 7, 'SUKSES', 1, '100.000', '2022-03-24', '2022-03-24 10:06:32'),
(3, 2, 'SUKSES', 2, '50.000', '2022-03-24', '2022-03-24 10:10:05'),
(3, 5, 'SUKSES', 1, '150.000', '2022-03-24', '2022-03-24 10:10:05'),
(4, 2, 'SUKSES', 2, '50.000', '2022-03-24', '2022-03-24 10:14:06'),
(4, 3, 'SUKSES', 1, '50.000', '2022-03-24', '2022-03-24 10:11:39'),
(5, 2, 'SUKSES', 2, '50.000', '2022-03-24', '2022-03-24 10:16:44'),
(5, 7, 'SUKSES', 1, '100.000', '2022-03-24', '2022-03-24 10:18:15'),
(6, 1, 'SUKSES', 3, '25.000', '2022-03-24', '2022-03-29 15:03:00'),
(6, 4, 'SUKSES', 1, '200.000', '2022-03-24', '2022-03-24 10:19:44'),
(7, 1, 'SUKSES', 3, '25.000', '2022-03-24', '2022-03-29 15:03:00'),
(7, 5, 'SUKSES', 1, '150.000', '2022-03-24', '2022-03-24 10:22:18'),
(8, 2, 'SUKSES', 2, '50.000', '2022-03-24', '2022-03-24 10:26:42'),
(8, 5, 'SUKSES', 1, '150.000', '2022-03-24', '2022-03-24 10:26:42'),
(9, 1, 'SUKSES', 3, '25.000', '2022-03-24', '2022-03-29 15:03:00'),
(9, 7, 'SUKSES', 1, '100.000', '2022-03-24', '2022-03-24 10:28:20'),
(10, 3, 'SUKSES', 2, '50.000', '2022-03-24', '2022-03-24 10:30:48'),
(10, 6, 'SUKSES', 1, '100.000', '2022-03-24', '2022-03-24 10:30:48'),
(11, 2, 'SUKSES', 1, '50.000', '2022-03-24', '2022-03-24 10:32:40'),
(11, 6, 'SUKSES', 1, '100.000', '2022-03-24', '2022-03-24 10:32:40'),
(11, 7, 'SUKSES', 1, '100.000', '2022-03-24', '2022-03-24 10:32:40'),
(12, 1, 'SUKSES', 3, '25.000', '2022-03-24', '2022-03-29 15:03:00'),
(12, 6, 'SUKSES', 1, '100.000', '2022-03-24', '2022-03-24 10:33:55'),
(13, 2, 'SUKSES', 1, '50.000', '2022-03-24', '2022-03-24 10:35:45'),
(13, 3, 'SUKSES', 1, '50.000', '2022-03-24', '2022-03-24 10:35:45'),
(13, 5, 'SUKSES', 1, '150.000', '2022-03-24', '2022-03-24 10:35:45'),
(14, 1, 'SUKSES', 3, '25.000', '2022-03-24', '2022-03-29 15:03:00'),
(14, 5, 'SUKSES', 1, '150.000', '2022-03-24', '2022-03-24 10:37:23'),
(15, 1, 'PROCESS', 3, '25.000', '2022-03-24', '2022-03-29 15:03:00'),
(15, 7, 'PROCESS', 1, '100.000', '2022-03-24', '2022-03-24 10:38:46');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id_user` int(11) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `alamat` varchar(100) DEFAULT NULL,
  `tgl_lahir` date DEFAULT NULL,
  `status_user` smallint(6) DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updates_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id_user`, `nama`, `alamat`, `tgl_lahir`, `status_user`, `gender`, `created_at`, `updates_at`) VALUES
(1, 'Via', 'Jl. Salak', '1999-01-11', NULL, 'P', '2022-03-24', '2022-03-24 09:42:25'),
(2, 'Josh', 'Jl. Buntu', '2001-03-18', NULL, 'L', '2022-03-24', '2022-03-24 09:42:25'),
(3, 'Steven', 'Jl. Merpati', '2001-06-21', NULL, 'L', '2022-03-24', '2022-03-24 09:42:25'),
(4, 'Lia', 'Jl. Baru', '1997-07-05', NULL, 'P', '2022-03-24', '2022-03-24 09:42:25'),
(5, 'Shein', 'Jl. Mawar', '2006-09-27', NULL, 'P', '2022-03-24', '2022-03-24 09:42:25');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `operator`
--
ALTER TABLE `operator`
  ADD PRIMARY KEY (`id_operator`);

--
-- Indexes for table `payment_method`
--
ALTER TABLE `payment_method`
  ADD PRIMARY KEY (`id_payment`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`id_product`),
  ADD KEY `id_type` (`id_type`),
  ADD KEY `id_operator` (`id_operator`);

--
-- Indexes for table `product_desc`
--
ALTER TABLE `product_desc`
  ADD PRIMARY KEY (`id_product`);

--
-- Indexes for table `product_type`
--
ALTER TABLE `product_type`
  ADD PRIMARY KEY (`id_type`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id_transaksi`),
  ADD KEY `id_user` (`id_user`),
  ADD KEY `id_payment` (`id_payment`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD PRIMARY KEY (`id_transaksi`,`id_product`),
  ADD KEY `id_product` (`id_product`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id_user`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `product`
--
ALTER TABLE `product`
  ADD CONSTRAINT `product_ibfk_1` FOREIGN KEY (`id_type`) REFERENCES `product_type` (`id_type`),
  ADD CONSTRAINT `product_ibfk_2` FOREIGN KEY (`id_operator`) REFERENCES `operator` (`id_operator`);

--
-- Constraints for table `product_desc`
--
ALTER TABLE `product_desc`
  ADD CONSTRAINT `product_desc_ibfk_1` FOREIGN KEY (`id_product`) REFERENCES `product` (`id_product`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `user` (`id_user`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`id_payment`) REFERENCES `payment_method` (`id_payment`);

--
-- Constraints for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`id_transaksi`) REFERENCES `transactions` (`id_transaksi`),
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`id_product`) REFERENCES `product` (`id_product`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
