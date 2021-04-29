-- phpMyAdmin SQL Dump
-- version 5
-- https://www.phpmyadmin.net/
--
-- ホスト: localhost
-- 生成日時: 
-- サーバのバージョン： 

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

--
-- データベース: `sns_sample`
--

-- -------------------------------------------------------------------------

--
-- テーブルの構造 `users`
--

CREATE TABLE `users` (
    `id` int(10) UNSIGNED NOT NULL,
    `screen_name` text NOT NULL,
    `user_name` text NOT NULL
    `email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `password` varchar(255) NOT NULL,
    `create_at` int(10) UNSIGNED NOT NULL,
    `update_at` int(10) UNSIGNED NOT NULL,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- -------------------------------------------------------------------------

--
-- テーブルの構造 `user_tokens`
--

CREATE TABLE `user_tokens` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `token` varchar(255) NOT NULL,
  `token_expire_at` int(10) UNSIGNED NOT NULL,
  `refresh_token` varchar(255) NOT NULL,
  `refresh_token_expire_at` int(10) UNSIGNED NOT NULL,
  `created_at` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- -------------------------------------------------------------------------

--
-- テーブルの構造 `tweets`
--

CREATE TABLE `tweets` (
    `id` int(10) UNSIGNED NOT NULL,
    `userId` int(10) UNSIGNED NOT NULL,
    `contents` varchar(140) NOT NULL,
    `crate_at` int(10) UNSIGNED NOT NULL,
    `update_at`int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- -------------------------------------------------------------------------

--
-- ダンプしたテーブルのインデックス
--

--
-- テーブルのインデックス `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- テーブルのインデックス `user_token`
--
ALTER TABLE `user_token`
  ADD PRIMARY KEY (`id`);

--
-- テーブルのインデックス `tweets`
--
ALTER TABLE `tweets`
  ADD PRIMARY KEY (`id`);

-- -------------------------------------------------------------------------

--
-- ダンプしたテーブルのAUTO_INCREMENT
--

--
-- テーブルのAUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- テーブルのAUTO_INCREMENT `user_token`
--
ALTER TABLE `user_token`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- テーブルのAUTO_INCREMENT `tweets`
--
ALTER TABLE `tweets`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;