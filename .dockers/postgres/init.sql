SET @@SESSION.SQL_LOG_BIN = 0;

SET @@GLOBAL.GTID_PURGED=/*!80000 '+'*/ '';

CREATE DATABASE IF NOT EXISTS `gotest_db`;