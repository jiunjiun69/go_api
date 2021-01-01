# 執行
'''
go run main.go 
'''

# 修改swagger
'''
swag init
'''

# MySQL新增指令
'''
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

CREATE TABLE `blog_fishpower_tag` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='魚隻活動力標籤關聯';

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
'''

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

