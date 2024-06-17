-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Hôte : localhost
-- Généré le : mer. 20 mars 2024 à 17:51
-- Version du serveur : 10.4.32-MariaDB
-- Version de PHP : 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `digiJob`
--

-- --------------------------------------------------------

--
-- Structure de la table `barreprogressioncriteres`
--

CREATE TABLE `barreprogressioncriteres` (
  `id` int(10) UNSIGNED NOT NULL,
  `nbreCritere` int(20) DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `barreprogressioncriteres`
--

INSERT INTO `barreprogressioncriteres` (`id`, `nbreCritere`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(1, 10, 1, '2023-11-26 18:07:01', '2023-11-26 20:13:15'),
(7, 10, 22, '2023-11-29 18:00:59', '2023-11-29 18:00:59'),
(46, 10, 24, '2023-12-09 07:41:53', '2023-12-09 07:42:45'),
(52, 0, 25, '2023-12-09 10:48:01', '2023-12-09 10:48:01'),
(53, 0, 26, '2023-12-09 13:45:24', '2023-12-09 13:45:24'),
(67, 0, 28, '2024-01-27 19:39:31', '2024-01-27 19:39:31');

-- --------------------------------------------------------

--
-- Structure de la table `barreprogressioncvs`
--

CREATE TABLE `barreprogressioncvs` (
  `id` int(10) UNSIGNED NOT NULL,
  `nbreFormation` int(20) DEFAULT NULL,
  `nbreExpeProf` int(20) DEFAULT NULL,
  `nbreComptProf` int(20) DEFAULT NULL,
  `nbrePrixCertificat` int(20) DEFAULT NULL,
  `nbreLangue` int(20) DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `barreprogressioncvs`
--

INSERT INTO `barreprogressioncvs` (`id`, `nbreFormation`, `nbreExpeProf`, `nbreComptProf`, `nbrePrixCertificat`, `nbreLangue`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(1, 20, 18, 10, 10, 10, 1, '2023-11-26 17:38:06', '2023-12-07 15:49:32'),
(97, 20, 18, 10, 10, 10, 22, '2023-11-29 16:24:23', '2023-11-29 16:24:23'),
(293, 20, 18, 10, 10, 10, 24, '2023-12-09 07:26:58', '2023-12-09 07:44:36'),
(304, 20, 18, 10, 10, 0, 25, '2023-12-09 10:43:41', '2023-12-09 10:48:45'),
(307, 20, 18, 10, 10, 10, 26, '2023-12-09 13:43:06', '2023-12-09 14:06:29'),
(312, 0, 0, 0, 0, 0, 20, '2023-12-22 15:57:55', '2023-12-22 15:57:55'),
(320, 0, 0, 0, 0, 0, 28, '2024-01-27 19:39:36', '2024-01-27 19:39:36');

-- --------------------------------------------------------

--
-- Structure de la table `barreprogressionprofils`
--

CREATE TABLE `barreprogressionprofils` (
  `id` int(10) UNSIGNED NOT NULL,
  `nbreInfoPerso` int(22) DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `barreprogressionprofils`
--

INSERT INTO `barreprogressionprofils` (`id`, `nbreInfoPerso`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(1, 2, 1, '2023-11-26 17:11:52', '2023-12-07 16:01:31'),
(2, 21, 1, '2023-12-07 16:01:42', '2023-12-07 16:01:42'),
(3, 21, 1, '2023-12-08 10:21:47', '2023-12-08 10:21:47'),
(4, 20, 1, '2023-12-08 10:22:09', '2023-12-08 10:22:09'),
(5, 21, 1, '2023-12-08 10:22:56', '2023-12-08 10:22:56'),
(6, 21, 1, '2023-12-08 12:06:36', '2023-12-08 12:06:36'),
(7, 21, 1, '2023-12-08 12:08:18', '2023-12-08 12:08:18'),
(8, 21, 1, '2023-12-08 12:12:03', '2023-12-08 12:12:03'),
(9, 21, 1, '2023-12-08 12:12:40', '2023-12-08 12:12:40'),
(10, 19, 24, '2023-12-09 07:26:39', '2023-12-09 07:26:39'),
(11, 16, 25, '2023-12-09 10:43:27', '2023-12-09 10:43:27'),
(12, 17, 26, '2023-12-09 13:42:57', '2023-12-09 13:42:57'),
(13, 21, 1, '2024-01-09 14:56:14', '2024-01-09 14:56:14'),
(14, 21, 1, '2024-01-09 14:56:29', '2024-01-09 14:56:29'),
(15, 21, 1, '2024-01-09 14:56:59', '2024-01-09 14:56:59'),
(16, 21, 1, '2024-01-10 09:46:33', '2024-01-10 09:46:33'),
(17, 21, 1, '2024-01-10 09:46:56', '2024-01-10 09:46:56'),
(18, 20, 1, '2024-01-10 09:47:45', '2024-01-10 09:47:45'),
(19, 21, 1, '2024-01-10 09:47:53', '2024-01-10 09:47:53'),
(20, 21, 1, '2024-01-10 12:33:47', '2024-01-10 12:33:47'),
(21, 21, 1, '2024-01-16 14:29:32', '2024-01-16 14:29:32'),
(22, 21, 1, '2024-01-16 14:31:18', '2024-01-16 14:31:18'),
(23, 21, 1, '2024-01-22 10:25:33', '2024-01-22 10:25:33'),
(24, 19, 1, '2024-01-27 19:22:35', '2024-01-27 19:22:35'),
(25, 12, 27, '2024-01-27 19:30:39', '2024-01-27 19:30:39'),
(26, 10, 27, '2024-01-27 19:31:34', '2024-01-27 19:31:34');

-- --------------------------------------------------------

--
-- Structure de la table `candidatureOffreGenerales`
--

CREATE TABLE `candidatureOffreGenerales` (
  `id` int(10) UNSIGNED NOT NULL,
  `motivation` varchar(255) DEFAULT NULL,
  `cv` varchar(255) DEFAULT NULL,
  `offreGenerales_id` int(10) UNSIGNED DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `candidatureOffreGenerales`
--

INSERT INTO `candidatureOffreGenerales` (`id`, `motivation`, `cv`, `offreGenerales_id`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(4, 'mmop', '006189f686173da5c27a9dc8.pdf', 2, 1, '2024-01-10 19:44:49', '2024-01-24 19:34:29'),
(39, 'je fais un autre test ok', '852c3a5a5d11d939f343f167.pdf', 3, 1, '2024-01-22 13:04:50', '2024-01-22 13:55:48'),
(48, 'Les parents de Esther', 'f69d412454045aff4e66dfdb.pdf', 2, 20, '2024-01-24 19:33:11', '2024-01-24 19:37:42'),
(49, 'ok pour un test', '0a00b8f2a6d5a1aa7919e8a6.pdf', 3, 20, '2024-01-24 19:38:30', '2024-01-24 19:38:30');

-- --------------------------------------------------------

--
-- Structure de la table `competenceCles`
--

CREATE TABLE `competenceCles` (
  `id` int(10) UNSIGNED NOT NULL,
  `libelle` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `competenceCles`
--

INSERT INTO `competenceCles` (`id`, `libelle`, `created_at`, `updated_at`) VALUES
(1, 'SQL', '2023-09-28 19:21:59', '2023-09-28 19:21:59'),
(2, 'HTML', '2023-09-28 19:21:59', '2023-09-28 19:21:59'),
(3, 'DOCKER', '2023-11-03 15:10:23', '2023-11-03 15:10:23'),
(4, 'JAVASCRIPT', '2023-11-03 15:10:23', '2023-11-03 15:10:23');

-- --------------------------------------------------------

--
-- Structure de la table `competenceCle_offreGenerales`
--

CREATE TABLE `competenceCle_offreGenerales` (
  `id` int(10) UNSIGNED NOT NULL,
  `competenceCles_id` int(10) UNSIGNED DEFAULT NULL,
  `offreGenerales_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `competenceCle_offreGenerales`
--

INSERT INTO `competenceCle_offreGenerales` (`id`, `competenceCles_id`, `offreGenerales_id`, `created_at`, `updated_at`) VALUES
(1, 1, 2, '2024-01-15 14:55:12', '2024-01-15 14:55:12'),
(2, 4, 2, '2024-01-15 14:55:12', '2024-01-15 14:55:12'),
(3, 2, 3, '2024-01-15 14:55:53', '2024-01-15 14:55:53'),
(4, 3, 3, '2024-01-15 14:55:53', '2024-01-15 14:55:53'),
(5, 3, 2, '2024-01-16 09:33:24', '2024-01-16 09:33:24');

-- --------------------------------------------------------

--
-- Structure de la table `competenceCle_usersDigicodeuses`
--

CREATE TABLE `competenceCle_usersDigicodeuses` (
  `id` int(10) UNSIGNED NOT NULL,
  `competenceCles_id` int(10) UNSIGNED DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `competenceCle_usersDigicodeuses`
--

INSERT INTO `competenceCle_usersDigicodeuses` (`id`, `competenceCles_id`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(42, 3, 1, '2023-12-07 15:45:39', '2023-12-07 15:45:39'),
(44, 1, 1, '2023-12-07 15:50:20', '2023-12-07 15:50:20'),
(45, 1, 24, '2023-12-09 07:26:39', '2023-12-09 07:26:39'),
(46, 2, 24, '2023-12-09 07:26:39', '2023-12-09 07:26:39'),
(47, 2, 25, '2023-12-09 10:43:27', '2023-12-09 10:43:27'),
(48, 3, 25, '2023-12-09 10:43:27', '2023-12-09 10:43:27'),
(49, 4, 25, '2023-12-09 10:43:27', '2023-12-09 10:43:27'),
(50, 2, 26, '2023-12-09 13:42:57', '2023-12-09 13:42:57'),
(51, 4, 26, '2023-12-09 13:42:57', '2023-12-09 13:42:57'),
(52, 2, 1, '2024-01-16 14:29:32', '2024-01-16 14:29:32'),
(53, 4, 1, '2024-01-16 14:31:18', '2024-01-16 14:31:18');

-- --------------------------------------------------------

--
-- Structure de la table `critereDigicodeuses`
--

CREATE TABLE `critereDigicodeuses` (
  `id` int(10) UNSIGNED NOT NULL,
  `type_contrat` varchar(25) DEFAULT NULL,
  `localite` varchar(25) DEFAULT NULL,
  `duree_contrat` varchar(25) DEFAULT NULL,
  `modalite` varchar(25) DEFAULT NULL,
  `metier_recherche` varchar(25) DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `critereDigicodeuses`
--

INSERT INTO `critereDigicodeuses` (`id`, `type_contrat`, `localite`, `duree_contrat`, `modalite`, `metier_recherche`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(1, 'CDI', 'Abidjan', '1 ans', 'Télétravail', '', 1, '2023-11-25 13:48:33', '2024-01-25 18:09:42'),
(5, 'Stage', 'Abidjan', '3 mois', 'Télétravail', 'CDI', 22, '2023-11-29 16:30:52', '2023-11-29 16:30:52'),
(19, 'CDD', 'Abidjan', '1 ans', 'Télétravail', '', 24, '2023-12-09 07:42:15', '2023-12-09 07:42:51');

-- --------------------------------------------------------

--
-- Structure de la table `cvActiviteExtras`
--

CREATE TABLE `cvActiviteExtras` (
  `id` int(10) UNSIGNED NOT NULL,
  `libelleActivite` varchar(255) DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `cvCentreInterets`
--

CREATE TABLE `cvCentreInterets` (
  `id` int(10) UNSIGNED NOT NULL,
  `libelleCentreInteret` varchar(255) DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `cvCompetences`
--

CREATE TABLE `cvCompetences` (
  `id` int(10) UNSIGNED NOT NULL,
  `titre` varchar(150) DEFAULT NULL,
  `pourcentage` int(11) NOT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `cvCompetences`
--

INSERT INTO `cvCompetences` (`id`, `titre`, `pourcentage`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(6, 'HTML ', 60, 22, '2023-11-23 15:05:49', '2023-11-23 15:05:49'),
(8, 'HTML', 80, 1, '2023-11-27 08:23:09', '2023-11-27 13:18:03'),
(16, 'CSS', 60, 1, '2023-11-27 14:52:51', '2023-11-27 14:52:51'),
(17, 'HTML', 50, 24, '2023-12-09 07:30:16', '2023-12-09 07:30:16'),
(18, 'HTML', 50, 25, '2023-12-09 10:46:36', '2023-12-09 10:46:36'),
(19, 'javascript', 30, 25, '2023-12-09 10:46:55', '2023-12-09 10:46:55'),
(20, 'CSS', 90, 25, '2023-12-09 10:46:55', '2023-12-09 10:46:55'),
(21, 'HTML', 50, 26, '2023-12-09 13:44:14', '2023-12-09 13:44:14');

-- --------------------------------------------------------

--
-- Structure de la table `cvExperiences`
--

CREATE TABLE `cvExperiences` (
  `id` int(10) UNSIGNED NOT NULL,
  `moisDebut` varchar(50) DEFAULT NULL,
  `moisFin` varchar(50) DEFAULT NULL,
  `anneeDebut` varchar(30) DEFAULT NULL,
  `anneeFin` varchar(30) DEFAULT NULL,
  `now` varchar(30) DEFAULT NULL,
  `titreExperience` varchar(150) DEFAULT NULL,
  `nomEntreprise` varchar(150) DEFAULT NULL,
  `typeContrat` varchar(150) DEFAULT NULL,
  `descriptionExperience` varchar(255) DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `cvExperiences`
--

INSERT INTO `cvExperiences` (`id`, `moisDebut`, `moisFin`, `anneeDebut`, `anneeFin`, `now`, `titreExperience`, `nomEntreprise`, `typeContrat`, `descriptionExperience`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(2, 'janvier', 'juin', '2010', '2019', '', 'Test 2', 'Simplon CI', '', 'Deuxième test à effectuer', 1, '2023-11-16 13:07:48', '2023-12-05 18:43:02'),
(3, 'juillet', '', '2020', '', 'À aujourd\'hui', 'Test final', 'ENS', '', 'Faire le dernier test hhhhhhhhhhhhh', 1, '2023-11-17 10:11:39', '2023-12-05 18:42:44'),
(4, 'avril', '', '2023', '', 'À aujourd\'hui', 'fullstack developper junior', 'DigiFemmes', '', 'test', 22, '2023-11-23 15:05:20', '2023-11-23 15:05:20'),
(5, 'septembre', 'septembre', '2020', '2021', '', 'Test final', 'Simplon CI', '', 'Test d’expérience pro', 24, '2023-12-09 07:29:42', '2023-12-09 07:29:42'),
(6, 'avril', '', '2019', '', 'À aujourd\'hui', 'fullstack developper junior', 'ENS', '', 'developper', 25, '2023-12-09 10:45:45', '2023-12-09 10:45:45'),
(7, 'mars', 'juin', '2017', '2019', '', 'Test final', 'ENS', '', 'test', 25, '2023-12-09 10:46:17', '2023-12-09 10:46:17'),
(8, 'mars', '', '2021', '', 'À aujourd\'hui', 'Test final', 'Simplon CI', '', 'okok', 25, '2023-12-09 10:51:41', '2023-12-09 10:51:41'),
(9, 'avril', '', '2022', '', 'À aujourd\'hui', 'fullstack developper junior', 'DigiFemmes', '', 'test2', 26, '2023-12-09 13:44:00', '2023-12-09 13:44:00');

-- --------------------------------------------------------

--
-- Structure de la table `cvFormations`
--

CREATE TABLE `cvFormations` (
  `id` int(10) UNSIGNED NOT NULL,
  `moisDebut` varchar(50) DEFAULT NULL,
  `moisFin` varchar(50) DEFAULT NULL,
  `anneeDebut` varchar(30) DEFAULT NULL,
  `anneeFin` varchar(30) DEFAULT NULL,
  `now` varchar(30) DEFAULT NULL,
  `titreFormation` varchar(150) DEFAULT NULL,
  `nomEcole` varchar(150) DEFAULT NULL,
  `descriptionFormation` varchar(255) DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `cvFormations`
--

INSERT INTO `cvFormations` (`id`, `moisDebut`, `moisFin`, `anneeDebut`, `anneeFin`, `now`, `titreFormation`, `nomEcole`, `descriptionFormation`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(4, 'avril', '', '2021', '', 'now', 'Dernier essai', 'UVCI', 'Nouvelle description', 1, '2023-11-15 12:36:36', '2023-11-16 10:34:56'),
(7, 'janvier', '', '2008', '', 'À aujourd\'hui', 'Test', 'UPA', 'Je fais mon test final  jkkk', 1, '2023-11-16 11:32:28', '2023-12-05 18:41:57'),
(10, 'janvier', 'janvier', '2023', '2025', '', 'Developpeur', 'DigiFemmes Academy', 'un test', 22, '2023-11-23 15:02:32', '2023-11-23 15:02:32'),
(11, 'mars', '', '2015', '', 'À aujourd\'hui', 'networking', 'ESATIC', 'test 2', 22, '2023-11-23 15:04:25', '2023-11-23 15:04:25'),
(13, 'janvier', '', '2021', '', 'À aujourd\'hui', 'Test avant présentation', 'DigiFemmes', 'Je fais des test de présentation', 24, '2023-12-09 07:28:09', '2023-12-09 07:28:09'),
(14, 'mars', 'mai', '2023', '2023', '', 'Developpeur', 'DigiFemmes', 'j\'apprends le coding ok', 25, '2023-12-09 10:44:33', '2023-12-09 10:50:39'),
(15, 'avril', '', '2018', '', 'À aujourd\'hui', 'designer', 'ESATIC', 'ux designer', 25, '2023-12-09 10:45:13', '2023-12-09 10:45:13'),
(16, 'juin', 'mai', '2022', '2023', '', 'Developpeur', 'UPA', 'ok ok', 25, '2023-12-09 10:51:21', '2023-12-09 10:51:21'),
(17, 'avril', '', '2023', '', 'À aujourd\'hui', 'Developpeur', 'DigiFemmes', 'test1', 26, '2023-12-09 13:43:39', '2023-12-09 13:43:39');

-- --------------------------------------------------------

--
-- Structure de la table `cvPrixCertifications`
--

CREATE TABLE `cvPrixCertifications` (
  `id` int(10) UNSIGNED NOT NULL,
  `titrePrix` varchar(50) NOT NULL,
  `nomEtablissement` varchar(150) NOT NULL,
  `dateObtention` varchar(50) NOT NULL,
  `description` varchar(255) NOT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `cvPrixCertifications`
--

INSERT INTO `cvPrixCertifications` (`id`, `titrePrix`, `nomEtablissement`, `dateObtention`, `description`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(1, 'Test1 modif', 'DFemmes', '14/08/2022', 'Description of test modif', 1, '2023-11-21 17:12:39', '2023-11-21 19:21:13'),
(3, 'hackaton', 'DigiFemmes', '14/08/2023', 'test', 22, '2023-11-23 15:08:17', '2023-11-23 15:08:17'),
(4, 'hackaton', 'DigiJob', '14/08/2022', 'ok', 24, '2023-12-09 07:30:48', '2023-12-09 07:30:48'),
(5, 'hackaton', 'DigiFemmes', '14/08/2022', 'hackaton', 25, '2023-12-09 10:47:16', '2023-12-09 10:47:16'),
(6, 'hackaton', 'DigiFemmes', '14/08/2022', 'test3', 26, '2023-12-09 13:44:31', '2023-12-09 13:44:31');

-- --------------------------------------------------------

--
-- Structure de la table `cvReferences`
--

CREATE TABLE `cvReferences` (
  `id` int(10) UNSIGNED NOT NULL,
  `nom` varchar(150) DEFAULT NULL,
  `prenoms` varchar(150) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `structure` varchar(255) DEFAULT NULL,
  `numero` varchar(130) DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `entreprises`
--

CREATE TABLE `entreprises` (
  `id` int(10) UNSIGNED NOT NULL,
  `nomEntreprise` varchar(50) DEFAULT NULL,
  `secteurActivite` varchar(150) DEFAULT NULL,
  `adresse` varchar(150) DEFAULT NULL,
  `logo` varchar(255) DEFAULT NULL,
  `siteWeb` varchar(255) DEFAULT NULL,
  `descriptionEntreprise` varchar(255) DEFAULT NULL,
  `numeroEntreprise` varchar(50) DEFAULT NULL,
  `videoEntreprise` varchar(150) DEFAULT NULL,
  `tailleEntreprise` varchar(150) DEFAULT NULL,
  `facebook` varchar(255) DEFAULT NULL,
  `twitter` varchar(255) DEFAULT NULL,
  `linkedin` varchar(255) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `slogan` varchar(255) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `entreprises`
--

INSERT INTO `entreprises` (`id`, `nomEntreprise`, `secteurActivite`, `adresse`, `logo`, `siteWeb`, `descriptionEntreprise`, `numeroEntreprise`, `videoEntreprise`, `tailleEntreprise`, `facebook`, `twitter`, `linkedin`, `email`, `slogan`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Google', 'Informatique', 'Abidjan023695', 'cc217be512fb4de885612f27.png', NULL, 'Google est une entreprise au service de la technologie...', '00000000000', 'test1.mp4', 'Grande-entreprise ', 'https://academy.digifemmes.com/intra/abidjan/profile', NULL, NULL, 'google@gmail.com', NULL, '', '2024-01-05 12:02:57', '2024-01-27 14:16:06'),
(2, 'Nasa', 'Aeronautics', 'USA010235', 'ec03f74990a66c8189e00588.png', 'https://www.sqlshack.com/sql-order-by-clause-overview-and-examples/', 'La National Aeronautics and Space Administration, plus connue sous son acronyme NASA, est l\'agence fédérale responsable de la majeure partie du programme spatial civil des États-Unis. ', '+5201246', NULL, ' Grande-entreprise ', NULL, 'Acteur mondial des concessions, de la construction et de\n							l\'énergie', NULL, NULL, 'Acteur mondial des concessions, de la construction et de\n							l\'énergie', '', '2024-01-11 11:32:03', '2024-01-26 11:27:36');

-- --------------------------------------------------------

--
-- Structure de la table `evenements`
--

CREATE TABLE `evenements` (
  `id` int(10) UNSIGNED NOT NULL,
  `titreEvent` varchar(255) DEFAULT NULL,
  `descriptionEvent` varchar(255) DEFAULT NULL,
  `fileUpload` varchar(255) DEFAULT NULL,
  `dateEvent` varchar(30) DEFAULT NULL,
  `typeEvent` varchar(255) DEFAULT NULL,
  `heure` varchar(25) DEFAULT NULL,
  `fraisInscriptionEvent` varchar(255) DEFAULT 'Gratuit',
  `modaliteEvent` varchar(255) DEFAULT NULL,
  `lien` varchar(255) DEFAULT NULL,
  `lieuEvent` varchar(255) DEFAULT NULL,
  `dateLimite` varchar(30) DEFAULT NULL,
  `entreprises_id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `evenements`
--

INSERT INTO `evenements` (`id`, `titreEvent`, `descriptionEvent`, `fileUpload`, `dateEvent`, `typeEvent`, `heure`, `fraisInscriptionEvent`, `modaliteEvent`, `lien`, `lieuEvent`, `dateLimite`, `entreprises_id`, `created_at`, `updated_at`) VALUES
(1, 'Hackathon', 'Times New Roman est une police de caractère avec empattement connue et répandue, faisant partie de la famille des réales.', '3dc94d92742e920ee8764be4.jpe', '10/02/2024', NULL, NULL, '10000 fcfa', 'En ligne', 'test1', 'Abidjan/Cocody', '15/02/2024', 1, '2024-01-11 11:36:40', '2024-01-27 14:45:20'),
(2, 'Ckeckpoint', 'Times New Roman est une police de caractère avec empattement connue et répandue, faisant partie de la famille des réales.', '0d8f4d2bd6dbb4c8af5a9e92.jpe', '14/02/2024', NULL, NULL, '1000 fcfa', 'Présentiel', NULL, 'Pulman/Abidjan', NULL, 2, '2024-01-11 11:44:05', '2024-01-27 14:09:17'),
(3, 'Formation UX/UI', 'Times New Roman est une police de caractère avec empattement connue et répandue, faisant partie de la famille des réales.', '3dc94d92742e920ee8764be4.jpe', '10/02/2024', NULL, NULL, 'Gratuit', 'Présentiel', 'Lien', 'Angré', NULL, 1, '2024-01-11 18:55:16', '2024-01-27 14:45:27');

-- --------------------------------------------------------

--
-- Structure de la table `langues`
--

CREATE TABLE `langues` (
  `id` int(10) UNSIGNED NOT NULL,
  `libelle` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `langues`
--

INSERT INTO `langues` (`id`, `libelle`, `created_at`, `updated_at`) VALUES
(1, 'Français', '2023-09-23 23:42:36', '2023-09-23 23:42:36'),
(2, 'Anglais', '2023-11-10 11:38:47', '2023-11-10 11:38:47'),
(3, 'Espagnol', '2023-11-10 11:38:47', '2023-11-10 11:38:47'),
(4, 'Allemand', '2023-11-10 11:39:26', '2023-11-10 11:39:26'),
(5, 'Chinois', '2023-11-10 11:39:26', '2023-11-10 11:39:26');

-- --------------------------------------------------------

--
-- Structure de la table `langue_usersDigicodeuses`
--

CREATE TABLE `langue_usersDigicodeuses` (
  `id` int(10) UNSIGNED NOT NULL,
  `langues_id` int(10) UNSIGNED DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `langue_usersDigicodeuses`
--

INSERT INTO `langue_usersDigicodeuses` (`id`, `langues_id`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(7, 2, 22, '2023-11-23 15:00:44', '2023-11-23 15:00:44'),
(8, 1, 22, '2023-11-23 15:01:48', '2023-11-23 15:01:48'),
(47, 3, 1, '2023-12-07 15:49:32', '2023-12-07 15:49:32'),
(48, 1, 24, '2023-12-09 07:27:14', '2023-12-09 07:27:14'),
(49, 1, 25, '2023-12-09 10:50:27', '2023-12-09 10:50:27'),
(50, 3, 25, '2023-12-09 10:50:27', '2023-12-09 10:50:27'),
(51, 2, 26, '2023-12-09 13:43:17', '2023-12-09 13:43:17'),
(52, 3, 26, '2023-12-09 13:43:17', '2023-12-09 13:43:17'),
(53, 2, 1, '2023-12-22 16:01:00', '2023-12-22 16:01:00');

-- --------------------------------------------------------

--
-- Structure de la table `langue_usersEntreprises`
--

CREATE TABLE `langue_usersEntreprises` (
  `id` int(10) UNSIGNED NOT NULL,
  `langues_id` int(10) UNSIGNED DEFAULT NULL,
  `usersEntreprise_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `missionOffreGenerales`
--

CREATE TABLE `missionOffreGenerales` (
  `id` int(10) UNSIGNED NOT NULL,
  `missionPoste` varchar(255) DEFAULT NULL,
  `offreGenerales_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `offreGenerales`
--

CREATE TABLE `offreGenerales` (
  `id` int(10) UNSIGNED NOT NULL,
  `titrePoste` varchar(255) DEFAULT NULL,
  `descriptionPoste` varchar(255) DEFAULT NULL,
  `nombrePoste` int(11) DEFAULT NULL,
  `anneeExperience` varchar(255) DEFAULT NULL,
  `formation` varchar(255) DEFAULT NULL,
  `ville` varchar(255) DEFAULT NULL,
  `dateLimite` date DEFAULT NULL,
  `entreprises_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `typeContrat` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `offreGenerales`
--

INSERT INTO `offreGenerales` (`id`, `titrePoste`, `descriptionPoste`, `nombrePoste`, `anneeExperience`, `formation`, `ville`, `dateLimite`, `entreprises_id`, `created_at`, `updated_at`, `typeContrat`) VALUES
(2, 'Développeur Fullstack', 'Times New Roman est une police de caractère avec empattement connue et répandue, faisant partie de la famille des réales.', 3, '2', 'Informatique et science du numérique', 'Abidjan', '2023-10-25', 1, '2023-10-17 11:17:50', '2024-01-17 17:42:16', 'CDD'),
(3, 'Community Manager', 'Je fais un test pour le moment donc les éléments ne doivent pas être pareil', 2, '3', 'Multimédia du numérique', 'Korhogo', '2023-10-27', 2, '2023-10-17 11:20:41', '2024-01-17 17:45:13', 'CDI');

-- --------------------------------------------------------

--
-- Structure de la table `offreSpeciales`
--

CREATE TABLE `offreSpeciales` (
  `id` int(10) UNSIGNED NOT NULL,
  `titrePoste` varchar(255) DEFAULT NULL,
  `descriptionPoste` varchar(255) DEFAULT NULL,
  `ville` varchar(255) DEFAULT NULL,
  `dateLimite` date DEFAULT NULL,
  `message` varchar(255) DEFAULT NULL,
  `usersRecruteurs_id` int(10) UNSIGNED DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `participantEvenements`
--

CREATE TABLE `participantEvenements` (
  `id` int(10) UNSIGNED NOT NULL,
  `evenements_id` int(10) UNSIGNED DEFAULT NULL,
  `usersDigicodeuses_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `participantEvenements`
--

INSERT INTO `participantEvenements` (`id`, `evenements_id`, `usersDigicodeuses_id`, `created_at`, `updated_at`) VALUES
(1, 3, 1, '2024-01-27 17:22:44', '2024-01-27 17:22:44'),
(2, 3, 1, '2024-01-27 17:50:29', '2024-01-27 17:50:29'),
(3, 3, 1, '2024-01-27 17:52:03', '2024-01-27 17:52:03'),
(5, 3, 1, '2024-01-27 18:03:07', '2024-01-27 18:03:07'),
(6, 2, 1, '2024-01-27 18:10:39', '2024-01-27 18:10:39');

-- --------------------------------------------------------

--
-- Structure de la table `photoEntreprises`
--

CREATE TABLE `photoEntreprises` (
  `id` int(10) UNSIGNED NOT NULL,
  `nomPhoto` varchar(150) DEFAULT NULL,
  `photoUpload` varchar(150) DEFAULT NULL,
  `entreprises_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `photoEntreprises`
--

INSERT INTO `photoEntreprises` (`id`, `nomPhoto`, `photoUpload`, `entreprises_id`, `created_at`, `updated_at`) VALUES
(1, 'Seminaire', 'cc217be512fb4de885612f27.png', 2, '2024-01-27 12:47:49', '2024-01-27 12:47:49'),
(2, 'Equipe', 'cc217be512fb4de885612f27.png', 2, '2024-01-27 12:47:49', '2024-01-27 12:47:49'),
(3, 'Equipe Rec', 'ec03f74990a66c8189e00588.png', 1, '2024-01-27 12:49:33', '2024-01-27 12:49:33'),
(4, 'Autre', 'ec03f74990a66c8189e00588.png', 1, '2024-01-27 12:49:33', '2024-01-27 12:49:33');

-- --------------------------------------------------------

--
-- Structure de la table `prixEntreprises`
--

CREATE TABLE `prixEntreprises` (
  `id` int(10) UNSIGNED NOT NULL,
  `titrePrix` varchar(255) DEFAULT NULL,
  `fichierPrixUpload` varchar(255) DEFAULT NULL,
  `usersRecruteurs_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `profilRechercherOffreGenerales`
--

CREATE TABLE `profilRechercherOffreGenerales` (
  `id` int(10) UNSIGNED NOT NULL,
  `ProfilPoste` varchar(255) DEFAULT NULL,
  `offreGenerales_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `statutCandidatures`
--

CREATE TABLE `statutCandidatures` (
  `id` int(10) UNSIGNED NOT NULL,
  `satutAuto` varchar(30) DEFAULT NULL,
  `satutTalentManager` varchar(30) DEFAULT 'En attente',
  `satutEntreprise` varchar(30) DEFAULT 'En attente',
  `candidatureOffreGenerales_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `statutCandidatures`
--

INSERT INTO `statutCandidatures` (`id`, `satutAuto`, `satutTalentManager`, `satutEntreprise`, `candidatureOffreGenerales_id`, `created_at`, `updated_at`) VALUES
(6, 'Approuvée', 'En attente', 'En attente', 4, '2024-01-16 14:38:56', '2024-01-22 09:34:16'),
(15, 'Approuvée', 'En attente', 'En attente', 39, '2024-01-22 13:04:50', '2024-01-22 13:04:50'),
(24, 'Réjetée', 'Réjetée', 'Réjetée', 48, '2024-01-24 19:33:11', '2024-01-24 19:33:11'),
(25, 'Réjetée', 'Réjetée', 'Réjetée', 49, '2024-01-24 19:38:30', '2024-01-24 19:38:30');

-- --------------------------------------------------------

--
-- Structure de la table `temoignageEntreprises`
--

CREATE TABLE `temoignageEntreprises` (
  `id` int(10) UNSIGNED NOT NULL,
  `descriptionTemoignage` varchar(255) DEFAULT NULL,
  `usersRecruteurs_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Structure de la table `usersDigicodeuses`
--

CREATE TABLE `usersDigicodeuses` (
  `id` int(10) UNSIGNED NOT NULL,
  `nom` varchar(50) NOT NULL,
  `prenoms` varchar(225) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `pays` varchar(255) DEFAULT NULL,
  `ville` varchar(255) DEFAULT NULL,
  `date_naissance` varchar(30) DEFAULT NULL,
  `indicatif` varchar(30) DEFAULT NULL,
  `numero_telephone` varchar(255) DEFAULT NULL,
  `nationalite` varchar(255) DEFAULT NULL,
  `cv` varchar(255) DEFAULT NULL,
  `photo` varchar(255) DEFAULT NULL,
  `titrePosteCv` varchar(255) DEFAULT NULL,
  `portfolio` varchar(255) DEFAULT NULL,
  `facebook` varchar(255) DEFAULT NULL,
  `twitter` varchar(255) DEFAULT NULL,
  `linkedin` varchar(255) DEFAULT NULL,
  `temoignage` varchar(255) DEFAULT NULL,
  `autorisation` varchar(30) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `niveau_etude` varchar(50) DEFAULT NULL,
  `annee_experience` varchar(30) DEFAULT NULL,
  `adresse` varchar(50) DEFAULT NULL,
  `statut` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Déchargement des données de la table `usersDigicodeuses`
--

INSERT INTO `usersDigicodeuses` (`id`, `nom`, `prenoms`, `email`, `password`, `username`, `pays`, `ville`, `date_naissance`, `indicatif`, `numero_telephone`, `nationalite`, `cv`, `photo`, `titrePosteCv`, `portfolio`, `facebook`, `twitter`, `linkedin`, `temoignage`, `autorisation`, `description`, `niveau_etude`, `annee_experience`, `adresse`, `statut`, `created_at`, `updated_at`) VALUES
(1, 'Ouattara', 'Nafissah', 'nafissah@gmail.com', '$2a$10$97TbdLjuhRYgMByGz84qc.X727FDhk.A37Iy.CEjzxnp1I8Pa.pwm', 'Nafissah', 'DZ', 'abidjan', '12/12/2020', '+213', '0707326321', 'ivoirienne', '', 'ed98e0645304bbab23e0725b.png', 'Developpeur Fullstack', 'LienPortfolio', 'essaiFaceBook', 'LienTwitter', 'LienLinkedin', '', NULL, 'Le test a faire ok', '', '', 'Abidjan 01', 'À l\'écoute d\'opportunités', '2023-09-27 16:18:59', '2024-01-27 19:22:35'),
(2, 'Sangaré', 'Keita', 'keita@gmail.com', '$2a$10$jPxTr7yiFXuZCW78QWAtIux1AcpLxDudGAu.vJnDooAWsbhiVCIJK', 'test', 'CI', 'Abidjan', '2000-12-12', '+225', '0707326321', 'Ivoirienne', 'cvPdf', 'profileImagePdf', 'Developpeur Fullstack', NULL, 'lienFacebook', 'lienTwitter', 'lienLinkedin', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '2023-09-27 17:36:03', '2023-10-20 11:42:40'),
(20, 'Koné', 'Affou', 'affou@gmail.com', '$2a$10$yqzENl37Q5rZ2CT4VSjBnuNe/aW5z.L/tkcbXObFc6HwManJZK7d.', NULL, '', 'abidjan', '12/12/2020', '+53', '0707326321', '', 'http: no such file', 'http: no such file', '', '', 'essaiFaceBook', 'LienTwitterEssai', '', '', 'on', 'test', '', '', 'Abidjan 02', '', '2023-10-28 14:43:19', '2023-11-22 16:52:36'),
(21, 'bony', 'esther', 'bonyesther2002@gmail.com', '$2a$10$C1UBPxBlgvka5N1OG.CTvuogG0bBnA7JrHCgLwQmd/ZKIpQ3MnPI.', NULL, 'Cote_d_Ivoire', 'man', NULL, '+225', '0758500145', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'on', NULL, NULL, NULL, NULL, NULL, '2023-10-30 16:29:11', '2023-10-30 16:29:11'),
(22, 'Bony', 'N\'guessan grace aimée', 'bonyesther2002@icloud.com', '$2a$10$hWOpmgCgA4Bkrq1iYtxo8OQVi6q4QtWLy6AMAq/tpTghGKcP/ctyO', NULL, 'Cote_d_Ivoire', 'abidjan', '07/06/2001', '+225', '0101100970', 'ivoirienne', 'views/assets/img/uploadCv/upload-4290886006.pdf', 'views/assets/img/uploadPhotoProfil/upload-134591483.jpeg', 'Developpeur Fullstack', '', '', '', '', '', 'on', 'bonjour juste un test', 'licence', '2', '', '', '2023-11-23 14:54:13', '2023-11-23 14:59:07'),
(23, 'Kamaté', 'Mariam', 'mariam@gmail.com', '$2a$10$xlYSmsJP8.oG1kUz2DEBd..zhkvKfh/kvtC4CNSuShDdmsmXMQ7x6', NULL, 'Cote_d_Ivoire', 'Dabou', NULL, '+225', '102345687', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'on', NULL, NULL, NULL, NULL, NULL, '2023-12-04 11:45:25', '2023-12-04 11:45:25'),
(24, 'Kanté', 'Nana', 'nana@gmail.com', '$2a$10$i0u8fBucFICxKBOyeJWJcOh3haouGFdacRBWJFFh1CQew.tRD3zci', NULL, 'AL', 'Dabou', '12/30/1998', '+355', '0707326321', 'ivoirienne', 'de6521d05d3d9738b0b74866.pdf', '9e325856e2822ac72c43107c.jpe', 'Developpeur Fullstack', '', 'essaiFaceBook', 'LienTwitterP', '', '', '', 'Je fais un test de présentation', 'Bac+5', '3', 'Abidjan 01', 'À l\'écoute d\'opportunités', '2023-12-09 07:23:49', '2023-12-09 07:26:39'),
(25, 'tro', 'Grace', 'trograce@gmail.com', '$2a$10$XjLLXykKOQ.RrnjzZEt0f.vRvaWBGqN0Srh7i/uLatdP1WH6sKXe.', NULL, 'CI', 'abidjan', '12/12/2004', '+225', '0505050506', 'ivoirienne', '7f88237803bed2e11678f5a0.pdf', '99b03ad8ffc18988f231a757.jpe', 'Developpeur Fullstack', '', '', '', '', '', 'on', 'Times New Roman est une police de caractère avec empattement connue et répandue, faisant partie de la famille des réales.', 'licence', '3', '', 'À l\'écoute d\'opportunités', '2023-12-09 10:35:45', '2023-12-09 10:43:27'),
(26, 'tra', 'lou', 'tralou@gmail.com', '$2a$10$wcURCHiwPqym8OR7UG1Bn./aW8h5gapd0uim.6CN4FuKRuSv5uZ4q', NULL, 'CI', 'Abidjan', '12/12/2020', '+225', '0708090405', 'ivoirienne', '149f3b0c6557efc1742df638.pdf', 'd65c98d4780acad9b325b20d.jpe', 'Developpeur Fullstack', '', '', '', '', '', '', 'testTimes New Roman est une police de caractère avec empattement connue et répandue, faisant partie de la famille des réales.', 'licence', '3', 'Abidjan 01', 'À l\'écoute d\'opportunités', '2023-12-09 13:40:01', '2023-12-09 13:42:57'),
(27, 'nafi', 'tiapegue', 'nafi@gmail.com', '$2a$10$jxsglmUpuhtX7t2eUktQBeYCSbNAuxenjymPf0Pb5BeLzrOyy5ICG', NULL, 'CI', '', '11/15/2023', '+225', '0707326321', 'ivoirienne', '', '', 'Developpeur Fullstack', '', '', '', '', '', '', '', '', '', '', '', '2024-01-27 19:24:18', '2024-01-27 19:31:34'),
(28, 'Ouattara', 'Aïchata', 'aicha@gmail.com', '$2a$10$FIt9sGbrp1QR1FsYb46kF.ReCMQcpjVAGmf7doZAdfNhg4RD0G5ES', NULL, 'AL', 'Kobec', NULL, '+355', '0707326258', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, '2024-01-27 19:32:58', '2024-01-27 19:32:58'),
(29, 'KOuakou', 'Akissi', 'akissi.kouakou@uvci.edu.ci', '$2a$10$13C/fYds9fzG32CviOBcYeHkmhJjXOzBv5iQ6w65rGUu2Bdwl.eSC', NULL, 'CI', 'Abidjan', NULL, '+225', '0789126469', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, '2024-03-20 16:48:49', '2024-03-20 16:48:49');

-- --------------------------------------------------------

--
-- Structure de la table `usersEntreprises`
--

CREATE TABLE `usersEntreprises` (
  `id` int(10) UNSIGNED NOT NULL,
  `nom` varchar(50) DEFAULT NULL,
  `prenoms` varchar(255) DEFAULT NULL,
  `numero` varchar(50) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `poste` varchar(150) DEFAULT NULL,
  `photo` varchar(255) DEFAULT NULL,
  `entreprise_id` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Index pour les tables déchargées
--

--
-- Index pour la table `barreprogressioncriteres`
--
ALTER TABLE `barreprogressioncriteres`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `barreprogressioncvs`
--
ALTER TABLE `barreprogressioncvs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `barreprogressionprofils`
--
ALTER TABLE `barreprogressionprofils`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `candidatureOffreGenerales`
--
ALTER TABLE `candidatureOffreGenerales`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_candidatureOffreGenerales_usersDigicodeuses_id` (`usersDigicodeuses_id`),
  ADD KEY `offreGenerales_id` (`offreGenerales_id`) USING BTREE;

--
-- Index pour la table `competenceCles`
--
ALTER TABLE `competenceCles`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `competenceCle_offreGenerales`
--
ALTER TABLE `competenceCle_offreGenerales`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_competenceCle_offreGenerales_competenceCles_id` (`competenceCles_id`),
  ADD KEY `fk_competenceCle_offreGenerales_offreGenerales_id` (`offreGenerales_id`);

--
-- Index pour la table `competenceCle_usersDigicodeuses`
--
ALTER TABLE `competenceCle_usersDigicodeuses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_competenceCle_usersDigicodeuses_competenceCles_id` (`competenceCles_id`),
  ADD KEY `fk_competenceCle_usersDigicodeuses_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `critereDigicodeuses`
--
ALTER TABLE `critereDigicodeuses`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `fk_critereDigicodeuses_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `cvActiviteExtras`
--
ALTER TABLE `cvActiviteExtras`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_cvActiviteExtras_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `cvCentreInterets`
--
ALTER TABLE `cvCentreInterets`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_centreInterets_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `cvCompetences`
--
ALTER TABLE `cvCompetences`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_competences_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `cvExperiences`
--
ALTER TABLE `cvExperiences`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_experiences_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `cvFormations`
--
ALTER TABLE `cvFormations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_formations_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `cvPrixCertifications`
--
ALTER TABLE `cvPrixCertifications`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_cvPrixCertifications_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `cvReferences`
--
ALTER TABLE `cvReferences`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_referenceCV_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `entreprises`
--
ALTER TABLE `entreprises`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `evenements`
--
ALTER TABLE `evenements`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_evenements_entreprises_id` (`entreprises_id`) USING BTREE;

--
-- Index pour la table `langues`
--
ALTER TABLE `langues`
  ADD PRIMARY KEY (`id`);

--
-- Index pour la table `langue_usersDigicodeuses`
--
ALTER TABLE `langue_usersDigicodeuses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_langue_usersDigicodeuses_langues_id` (`langues_id`),
  ADD KEY `fk_langue_usersDigicodeuses_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `langue_usersEntreprises`
--
ALTER TABLE `langue_usersEntreprises`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_langue_usersEntreprises_langues_id` (`langues_id`) USING BTREE,
  ADD KEY `fk_langue_usersEntreprises_usersEntreprise_id` (`usersEntreprise_id`) USING BTREE;

--
-- Index pour la table `missionOffreGenerales`
--
ALTER TABLE `missionOffreGenerales`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_missionOffreGenerales_offreGenerales_id` (`offreGenerales_id`);

--
-- Index pour la table `offreGenerales`
--
ALTER TABLE `offreGenerales`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_offreGenerales_entreprises_id` (`entreprises_id`) USING BTREE;

--
-- Index pour la table `offreSpeciales`
--
ALTER TABLE `offreSpeciales`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_offreSpeciales_usersRecruteurs_id` (`usersRecruteurs_id`),
  ADD KEY `fk_offreSpeciales_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `participantEvenements`
--
ALTER TABLE `participantEvenements`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_participantEvenements_evenements_id` (`evenements_id`),
  ADD KEY `fk_participantEvenements_usersDigicodeuses_id` (`usersDigicodeuses_id`);

--
-- Index pour la table `photoEntreprises`
--
ALTER TABLE `photoEntreprises`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_photoEntreprises_entreprises_id` (`entreprises_id`) USING BTREE;

--
-- Index pour la table `prixEntreprises`
--
ALTER TABLE `prixEntreprises`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_prixEntreprises_usersRecruteurs_id` (`usersRecruteurs_id`);

--
-- Index pour la table `profilRechercherOffreGenerales`
--
ALTER TABLE `profilRechercherOffreGenerales`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_profilRechercherOffreGenerales_offreGenerales_id` (`offreGenerales_id`);

--
-- Index pour la table `statutCandidatures`
--
ALTER TABLE `statutCandidatures`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `candidatureOffreGenerales_id` (`candidatureOffreGenerales_id`) USING BTREE;

--
-- Index pour la table `temoignageEntreprises`
--
ALTER TABLE `temoignageEntreprises`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_temoignageEntreprises_usersRecruteurs_id` (`usersRecruteurs_id`);

--
-- Index pour la table `usersDigicodeuses`
--
ALTER TABLE `usersDigicodeuses`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Index pour la table `usersEntreprises`
--
ALTER TABLE `usersEntreprises`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_usersEntreprises_entreprise_id` (`entreprise_id`);

--
-- AUTO_INCREMENT pour les tables déchargées
--

--
-- AUTO_INCREMENT pour la table `barreprogressioncriteres`
--
ALTER TABLE `barreprogressioncriteres`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=68;

--
-- AUTO_INCREMENT pour la table `barreprogressioncvs`
--
ALTER TABLE `barreprogressioncvs`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=321;

--
-- AUTO_INCREMENT pour la table `barreprogressionprofils`
--
ALTER TABLE `barreprogressionprofils`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;

--
-- AUTO_INCREMENT pour la table `candidatureOffreGenerales`
--
ALTER TABLE `candidatureOffreGenerales`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=50;

--
-- AUTO_INCREMENT pour la table `competenceCles`
--
ALTER TABLE `competenceCles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT pour la table `competenceCle_offreGenerales`
--
ALTER TABLE `competenceCle_offreGenerales`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT pour la table `competenceCle_usersDigicodeuses`
--
ALTER TABLE `competenceCle_usersDigicodeuses`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=54;

--
-- AUTO_INCREMENT pour la table `critereDigicodeuses`
--
ALTER TABLE `critereDigicodeuses`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;

--
-- AUTO_INCREMENT pour la table `cvActiviteExtras`
--
ALTER TABLE `cvActiviteExtras`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `cvCentreInterets`
--
ALTER TABLE `cvCentreInterets`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `cvCompetences`
--
ALTER TABLE `cvCompetences`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- AUTO_INCREMENT pour la table `cvExperiences`
--
ALTER TABLE `cvExperiences`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT pour la table `cvFormations`
--
ALTER TABLE `cvFormations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT pour la table `cvPrixCertifications`
--
ALTER TABLE `cvPrixCertifications`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT pour la table `cvReferences`
--
ALTER TABLE `cvReferences`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `entreprises`
--
ALTER TABLE `entreprises`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT pour la table `evenements`
--
ALTER TABLE `evenements`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT pour la table `langues`
--
ALTER TABLE `langues`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT pour la table `langue_usersDigicodeuses`
--
ALTER TABLE `langue_usersDigicodeuses`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=54;

--
-- AUTO_INCREMENT pour la table `langue_usersEntreprises`
--
ALTER TABLE `langue_usersEntreprises`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `missionOffreGenerales`
--
ALTER TABLE `missionOffreGenerales`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `offreGenerales`
--
ALTER TABLE `offreGenerales`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT pour la table `offreSpeciales`
--
ALTER TABLE `offreSpeciales`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `participantEvenements`
--
ALTER TABLE `participantEvenements`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT pour la table `photoEntreprises`
--
ALTER TABLE `photoEntreprises`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT pour la table `prixEntreprises`
--
ALTER TABLE `prixEntreprises`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `profilRechercherOffreGenerales`
--
ALTER TABLE `profilRechercherOffreGenerales`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `statutCandidatures`
--
ALTER TABLE `statutCandidatures`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT pour la table `temoignageEntreprises`
--
ALTER TABLE `temoignageEntreprises`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT pour la table `usersDigicodeuses`
--
ALTER TABLE `usersDigicodeuses`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=30;

--
-- AUTO_INCREMENT pour la table `usersEntreprises`
--
ALTER TABLE `usersEntreprises`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `candidatureOffreGenerales`
--
ALTER TABLE `candidatureOffreGenerales`
  ADD CONSTRAINT `fk_candidatureOffreGenerales_offreGenerales_id` FOREIGN KEY (`offreGenerales_id`) REFERENCES `offreGenerales` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_candidatureOffreGenerales_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `competenceCle_offreGenerales`
--
ALTER TABLE `competenceCle_offreGenerales`
  ADD CONSTRAINT `fk_competenceCle_offreGenerales_competenceCles_id` FOREIGN KEY (`competenceCles_id`) REFERENCES `competenceCles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_competenceCle_offreGenerales_offreGenerales_id` FOREIGN KEY (`offreGenerales_id`) REFERENCES `offreGenerales` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `competenceCle_usersDigicodeuses`
--
ALTER TABLE `competenceCle_usersDigicodeuses`
  ADD CONSTRAINT `fk_competenceCle_usersDigicodeuses_competenceCles_id` FOREIGN KEY (`competenceCles_id`) REFERENCES `competenceCles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_competenceCle_usersDigicodeuses_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `critereDigicodeuses`
--
ALTER TABLE `critereDigicodeuses`
  ADD CONSTRAINT `critereDigicodeuses_ibfk_1` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `cvActiviteExtras`
--
ALTER TABLE `cvActiviteExtras`
  ADD CONSTRAINT `fk_cvActiviteExtras_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `cvCentreInterets`
--
ALTER TABLE `cvCentreInterets`
  ADD CONSTRAINT `fk_centreInterets_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `cvCompetences`
--
ALTER TABLE `cvCompetences`
  ADD CONSTRAINT `fk_competences_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `cvExperiences`
--
ALTER TABLE `cvExperiences`
  ADD CONSTRAINT `fk_experiences_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `cvFormations`
--
ALTER TABLE `cvFormations`
  ADD CONSTRAINT `fk_formations_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `cvPrixCertifications`
--
ALTER TABLE `cvPrixCertifications`
  ADD CONSTRAINT `fk_cvPrixCertifications_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `cvReferences`
--
ALTER TABLE `cvReferences`
  ADD CONSTRAINT `fk_referenceCV_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `evenements`
--
ALTER TABLE `evenements`
  ADD CONSTRAINT `fk_evenements_usersRecruteurs_id` FOREIGN KEY (`entreprises_id`) REFERENCES `entreprises` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `langue_usersDigicodeuses`
--
ALTER TABLE `langue_usersDigicodeuses`
  ADD CONSTRAINT `fk_langue_usersDigicodeuses_langues_id` FOREIGN KEY (`langues_id`) REFERENCES `langues` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_langue_usersDigicodeuses_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `langue_usersEntreprises`
--
ALTER TABLE `langue_usersEntreprises`
  ADD CONSTRAINT `fk_langue_usersRecruteurs_langues_id` FOREIGN KEY (`langues_id`) REFERENCES `langues` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_langue_usersRecruteurs_usersRecruteurs_id` FOREIGN KEY (`usersEntreprise_id`) REFERENCES `entreprises` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `missionOffreGenerales`
--
ALTER TABLE `missionOffreGenerales`
  ADD CONSTRAINT `fk_missionOffreGenerales_offreGenerales_id` FOREIGN KEY (`offreGenerales_id`) REFERENCES `offreGenerales` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `offreGenerales`
--
ALTER TABLE `offreGenerales`
  ADD CONSTRAINT `fk_offreGenerales_usersRecruteurs_id` FOREIGN KEY (`entreprises_id`) REFERENCES `entreprises` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `offreSpeciales`
--
ALTER TABLE `offreSpeciales`
  ADD CONSTRAINT `fk_offreSpeciales_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_offreSpeciales_usersRecruteurs_id` FOREIGN KEY (`usersRecruteurs_id`) REFERENCES `entreprises` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `participantEvenements`
--
ALTER TABLE `participantEvenements`
  ADD CONSTRAINT `fk_participantEvenements_evenements_id` FOREIGN KEY (`evenements_id`) REFERENCES `evenements` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `fk_participantEvenements_usersDigicodeuses_id` FOREIGN KEY (`usersDigicodeuses_id`) REFERENCES `usersDigicodeuses` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `photoEntreprises`
--
ALTER TABLE `photoEntreprises`
  ADD CONSTRAINT `fk_photoEntreprises_usersRecruteurs_id` FOREIGN KEY (`entreprises_id`) REFERENCES `entreprises` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `prixEntreprises`
--
ALTER TABLE `prixEntreprises`
  ADD CONSTRAINT `fk_prixEntreprises_usersRecruteurs_id` FOREIGN KEY (`usersRecruteurs_id`) REFERENCES `entreprises` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `profilRechercherOffreGenerales`
--
ALTER TABLE `profilRechercherOffreGenerales`
  ADD CONSTRAINT `fk_profilRechercherOffreGenerales_offreGenerales_id` FOREIGN KEY (`offreGenerales_id`) REFERENCES `offreGenerales` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `statutCandidatures`
--
ALTER TABLE `statutCandidatures`
  ADD CONSTRAINT `fk_statutCandidatures_candidatureOffreGenerales_id` FOREIGN KEY (`candidatureOffreGenerales_id`) REFERENCES `candidatureOffreGenerales` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `temoignageEntreprises`
--
ALTER TABLE `temoignageEntreprises`
  ADD CONSTRAINT `fk_temoignageEntreprises_usersRecruteurs_id` FOREIGN KEY (`usersRecruteurs_id`) REFERENCES `entreprises` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Contraintes pour la table `usersEntreprises`
--
ALTER TABLE `usersEntreprises`
  ADD CONSTRAINT `fk_usersEntreprises_entreprise_id` FOREIGN KEY (`entreprise_id`) REFERENCES `entreprises` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
