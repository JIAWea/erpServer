/*
 Navicat Premium Data Transfer

 Source Server         : ray
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : localhost:3306
 Source Schema         : erp

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 30/07/2023 12:10:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for erp_menu
-- ----------------------------
DROP TABLE IF EXISTS `erp_menu`;
CREATE TABLE `erp_menu`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` int(10) UNSIGNED DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(10) UNSIGNED DEFAULT NULL COMMENT '更新时间',
  `deleted_at` int(10) UNSIGNED DEFAULT NULL COMMENT '删除时间',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `status` tinyint(1) DEFAULT 1 COMMENT '1正常, 2禁用',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单图标',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单访问路径',
  `redirect` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '重定向路径',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '前端组件',
  `sort` int(3) UNSIGNED NOT NULL DEFAULT 999 COMMENT '菜单顺序(1-999)',
  `is_hidden` tinyint(1) DEFAULT 0 COMMENT '是否隐藏',
  `type` tinyint(1) DEFAULT 1 COMMENT '1菜单, 2api',
  `parent_id` bigint(20) DEFAULT 0 COMMENT '父级',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of erp_menu
-- ----------------------------
REPLACE INTO `erp_menu` VALUES (1, NULL, NULL, 0, '财务管理', 1, 'el-icon-bank-card', '/finance', '/finance/account', 'Layout', 1, 0, 1, 0);
REPLACE INTO `erp_menu` VALUES (2, NULL, NULL, 0, '系统管理', 1, 'el-icon-lock', '/sys', '/sys/user', 'Layout', 2, 0, 1, 0);
REPLACE INTO `erp_menu` VALUES (3, NULL, NULL, 0, '账户', 1, 'el-icon-s-custom', '/finance/account', '', 'Account', 3, 0, 1, 1);
REPLACE INTO `erp_menu` VALUES (4, NULL, NULL, 0, '支出', 1, 'el-icon-s-finance', '/finance/expense', '', 'Expense', 3, 0, 1, 1);
REPLACE INTO `erp_menu` VALUES (5, NULL, NULL, 0, '收入', 1, 'el-icon-wallet', '/finance/income', '', 'Income', 3, 0, 1, 1);
REPLACE INTO `erp_menu` VALUES (6, NULL, NULL, 0, '应付', 1, 'el-icon-s-finance', '/finance/payable', '', 'Payable', 3, 0, 1, 1);
REPLACE INTO `erp_menu` VALUES (7, NULL, NULL, 0, '应收', 1, 'el-icon-wallet', '/finance/receivable', '', 'Receivable', 3, 0, 1, 1);
REPLACE INTO `erp_menu` VALUES (8, NULL, NULL, 0, '用户管理', 1, 'user', '/sys/user', '', 'User', 3, 0, 1, 2);
REPLACE INTO `erp_menu` VALUES (9, NULL, NULL, 0, '角色管理', 1, 'tree', '/sys/role', '', 'Role', 3, 0, 1, 2);
REPLACE INTO `erp_menu` VALUES (10, NULL, NULL, 0, '用户列表', 1, '', '/erp/ListUser', '', 'User', 3, 1, 2, 20);
REPLACE INTO `erp_menu` VALUES (11, NULL, NULL, 0, '添加用户', 1, '', '/erp/CreateUser', '', 'User', 3, 1, 2, 20);
REPLACE INTO `erp_menu` VALUES (12, NULL, NULL, 0, '删除用户', 1, '', '/erp/DeleteUser', '', 'User', 3, 1, 2, 20);
REPLACE INTO `erp_menu` VALUES (13, NULL, NULL, 0, '编辑用户', 1, '', '/erp/UpdateUser', '', 'User', 3, 1, 2, 20);
REPLACE INTO `erp_menu` VALUES (14, NULL, NULL, 0, '分配角色', 1, '', '/erp/UpdateUserRole', '', 'Role', 3, 1, 2, 21);
REPLACE INTO `erp_menu` VALUES (15, NULL, NULL, 0, '角色列表', 1, '', '/erp/ListRole', '', 'Role', 3, 1, 2, 21);
REPLACE INTO `erp_menu` VALUES (16, NULL, NULL, 0, '添加角色', 1, '', '/erp/CreateRole', '', 'Role', 3, 1, 2, 21);
REPLACE INTO `erp_menu` VALUES (17, NULL, NULL, 0, '删除角色', 1, '', '/erp/DeleteRole', '', 'Role', 3, 1, 2, 21);
REPLACE INTO `erp_menu` VALUES (18, NULL, NULL, 0, '编辑角色', 1, '', '/erp/UpdateRole', '', 'Role', 3, 1, 2, 21);
REPLACE INTO `erp_menu` VALUES (19, NULL, NULL, 0, '分配权限', 1, '', '/erp/UpdateRoleMenu', '', 'Role', 3, 1, 2, 21);
REPLACE INTO `erp_menu` VALUES (20, NULL, NULL, 0, '用户管理', 1, '', '', '', '', 3, 0, 2, 0);
REPLACE INTO `erp_menu` VALUES (21, NULL, NULL, 0, '角色管理', 1, '', '', '', '', 3, 0, 2, 0);

SET FOREIGN_KEY_CHECKS = 1;
