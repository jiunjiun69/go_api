
# swagger
```
http://127.0.0.1:8000/swagger/index.html
```

# 執行
```
go run main.go 
```

# 修改swagger
```
swag init
```

# MySQL新增指令 fish_final_sql.txt
```
use service_api;

CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '標籤名稱',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '建立時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='標籤管理';

CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='認證管理';

CREATE TABLE `blog_fish_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `fish_id` int(11) NOT NULL COMMENT '魚隻ID',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '標籤ID',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '建立時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魚隻標籤關聯';

CREATE TABLE `blog_fish` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '魚隻名稱',
  `desc` varchar(255) DEFAULT '' COMMENT '魚隻介紹',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面圖片地址',
  `content` longtext COMMENT '魚隻習性',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魚隻管理';

CREATE TABLE `blog_fishpower` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '活動力狀態',
  `power` longtext COMMENT '活動力數值',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魚隻活動力管理';
```

# 資料庫直接匯入資料 service_api.sql
```
-- MySQL dump 10.13  Distrib 8.0.22, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: service_api
-- ------------------------------------------------------
-- Server version	8.0.22

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `blog_auth`
--

DROP TABLE IF EXISTS `blog_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blog_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
  `created_on` int unsigned DEFAULT '0' COMMENT '新建時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='認證管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_auth`
--

LOCK TABLES `blog_auth` WRITE;
/*!40000 ALTER TABLE `blog_auth` DISABLE KEYS */;
/*!40000 ALTER TABLE `blog_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_fish`
--

DROP TABLE IF EXISTS `blog_fish`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blog_fish` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '魚隻名稱',
  `desc` varchar(255) DEFAULT '' COMMENT '魚隻介紹',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT '封面圖片地址',
  `content` longtext COMMENT '魚隻習性',
  `created_on` int unsigned DEFAULT '0' COMMENT '新建時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  `state` tinyint unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='魚隻管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_fish`
--

LOCK TABLES `blog_fish` WRITE;
/*!40000 ALTER TABLE `blog_fish` DISABLE KEYS */;
INSERT INTO `blog_fish` VALUES (1,'小丑魚','小丑魚（英語：Clownfish或anemonefish）是對雀鯛科底下的海葵魚亞科（Amphiprioninae）魚類的俗稱，是一種熱帶海水魚。已知有約30種，一種來自棘頰雀鯛屬（Premnas），其餘來自雙鋸魚屬（Amphiprion）。小丑魚大多生活在水質清澈、光線充足、水深50公尺以上之海域 為熱帶性觀賞魚，分布於太平洋與印度洋，不存在於大西洋。','https://img.ltn.com.tw/Upload/liveNews/BigPic/600_phpfHKaUl.jpg','小丑魚原生於印度洋和太平洋較溫暖的水中，包括大堡礁和紅海。雖然大部分種類的分布有其限制，但某些分布的十分廣泛。小丑魚生活在淺海的底部的淺潟湖或珊瑚礁的海葵中，與海葵存在共生關係，保護了小丑魚不被天敵吃掉。小丑魚的皮膚上有層黏液 ，也可使自己不被海葵蟄傷，不過沒有這層保護的話一樣有可能遭到海葵襲擊。在繁殖方面，小丑魚有先公後母的變性特質，稱作順序性雌雄同體。一群小丑魚中一般以體型最大者轉變為母，然而最終成為母魚後就無法變為公魚，生產時平常會將卵產在平坦的礁石表面，夜間孵化。小丑魚不存在於大西洋中。',1609597350,'JN',1609597355,'JN6969',0,0,1);
/*!40000 ALTER TABLE `blog_fish` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_fish_tag`
--

DROP TABLE IF EXISTS `blog_fish_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blog_fish_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `fish_id` int NOT NULL COMMENT '魚隻ID',
  `tag_id` int unsigned NOT NULL DEFAULT '0' COMMENT '標籤ID',
  `created_on` int unsigned DEFAULT '0' COMMENT '建立時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='魚隻標籤關聯';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_fish_tag`
--

LOCK TABLES `blog_fish_tag` WRITE;
/*!40000 ALTER TABLE `blog_fish_tag` DISABLE KEYS */;
INSERT INTO `blog_fish_tag` VALUES (1,1,1,1609597350,'JN',1609597355,'JN6969',0,0);
/*!40000 ALTER TABLE `blog_fish_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_fishpower`
--

DROP TABLE IF EXISTS `blog_fishpower`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blog_fishpower` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '活動力狀態',
  `power` longtext COMMENT '活動力數值',
  `created_on` int unsigned DEFAULT '0' COMMENT '新建時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  `state` tinyint unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='魚隻活動力管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_fishpower`
--

LOCK TABLES `blog_fishpower` WRITE;
/*!40000 ALTER TABLE `blog_fishpower` DISABLE KEYS */;
INSERT INTO `blog_fishpower` VALUES (1,'小丑魚','350',1609597362,'JN',1609597362,'',0,0,1),(2,'錦鯉','400',1609597383,'JN69',1609597383,'',0,0,1);
/*!40000 ALTER TABLE `blog_fishpower` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blog_tag`
--

DROP TABLE IF EXISTS `blog_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blog_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '標籤名稱',
  `created_on` int unsigned DEFAULT '0' COMMENT '建立時間',
  `created_by` varchar(100) DEFAULT '' COMMENT '建立人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改時間',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0' COMMENT '刪除時間',
  `is_del` tinyint unsigned DEFAULT '0' COMMENT '是否刪除 0為未刪除、1為已刪除',
  `state` tinyint unsigned DEFAULT '1' COMMENT '狀態 0為禁用、1為啟用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='標籤管理';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog_tag`
--

LOCK TABLES `blog_tag` WRITE;
/*!40000 ALTER TABLE `blog_tag` DISABLE KEYS */;
/*!40000 ALTER TABLE `blog_tag` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-01-02 22:29:26

```
# blog-service（部落格後端）

blog-service 部落格後端
是《Go 語言程式設計之旅：一起用 Go 做專案》中的專案，是第二章 [HTTP 應用：寫一個完整的部落格後端] 的附屬原始碼。

## 關於本書

本書涵蓋 Go 語言的各大經典實戰，不介紹 Go 語言的語法基礎，內容面向專案實踐，同時會針對核心細節進行分析。而在實際專案迭代中，常常會出現或多或少的事故，因此本書也針對 Go 語言的大殺器（分析工具）以及常見問題進行了全面講解。

本書適合已經大致學習了 Go 語言的基礎語法后，想要跨越到下一個階段的開發人員，可以填補該階段的空白和進一步拓展你的思維方向。

- 作者：陳劍煜（煎魚），GitHub：[@eddycjy](https://github.com/eddycjy)，微信公眾號：腦子進煎魚了。
- 作者：徐新華（polaris），GitHub：[@polaris](https://github.com/polaris1119)，微信公眾號：Go語言中文網。

## 購買鏈接

- 京東：https://item.jd.com/12685249.html
- 噹噹：http://product.dangdang.com/28982027.html
- 天貓：https://detail.tmall.com/item.htm?id=622185710833

## 目錄

第 2 章 HTTP應用：寫一個完整的部落格後端
- 2.1 部落格之旅
- 2.2 專案設計
- 2.3 公共元件
- 2.4 介面文件
- 2.5 介面校驗
- 2.6 模組開發：標籤管理
- 2.7 模組開發：文章管理
- 2.8 上傳圖片和檔案服務
- 2.9 API訪問控制
- 2.10 常見應用中介軟體
- 2.11 鏈路追蹤
- 2.12 應用配置問題
- 2.13 編譯程式應用
- 2.14 優雅重啟和停止
- 2.15 思考

## 功能清單

相對完整的後端應用實現，包含從 0 到 1 的階段，涉及大量實際開發中所使用的各類技巧和元件使用。

## 交流

如果你買了書，但還沒進讀者群或配套的免費星球，請關注我的公眾號並點選相應的自定義菜單加入：

![image](https://image.eddycjy.com/7074be90379a121746146bc4229819f8.jpg)

## 後話

- 如果本專案對你有所幫助歡迎進行 star，另外有任何的建議或意見歡迎隨時提交 [issues](https://github.com/go-programming-tour-book/blog-service/issues/new) 或 [pr](https://github.com/go-programming-tour-book/blog-service/pulls)。

- 如果你不是本書的讀者，我正式向你推薦 《Go 語言程式設計之旅：一起用 Go 做專案》這本書，本書針對常見的專案型別，主要細分為 5 + 1 板塊，分別是命令列、HTTP、RPC、Websocket 應用、程序內快取以及 Go 語言中的大殺器，希望你能夠喜歡。

