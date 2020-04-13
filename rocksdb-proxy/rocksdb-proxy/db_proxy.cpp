#include "pch.h"
#include "db_proxy.h"
#include <stdio.h>
#include "rocksdb/db.h"

using namespace rocksdb;
std::string kDBPath = "/tmp/rocksdb_simple_example";
extern "C" {
	// TODO: 这是一个库函数示例
	void fnrocksdbproxy()
	{
	}

	void Test() {
		printf("test xxxx\n");
		fflush(stdout);
	}
	int GetValue() {
		return 111;
	}

	void* CreateDB(const char* str) {
		printf("%s\n", str);
		fflush(stdout);
		DB* db = nullptr;
		Options options;
		// Optimize RocksDB. This is the easiest way to get RocksDB to perform well
		options.IncreaseParallelism();
		options.OptimizeLevelStyleCompaction();
		// create the DB if it's not already present
		options.create_if_missing = true;

		// open DB
		Status s = DB::Open(options, std::string(str), &db);
		if (!s.ok()) {
			return nullptr;
		}
		return db;
	}
	bool DBPut(void* pdb, const char* key, const char* value) {
		auto db = (DB*)pdb;
		auto s = db->Put(WriteOptions(), key, value);
		printf("dbput:key=%s,value=%s,ok=%d,ptr=%d\n", key, value, s.ok(), pdb);
		fflush(stdout);
		return s.ok();
	}

	void ReleaseString(void* data) {
		auto pData = (DataString*)data;
		delete (std::string*)pData->strPtr;
		delete pData;
	}
	void* DBGet(void* pdb, const char* key) {
		auto db = (DB*)pdb;
		auto data = new DataString();
		auto value = new std::string();
		auto s = db->Get(ReadOptions(), std::string(key), value);
		printf("DBGet:key=%s,value=%s,ok=%d,ptr=%d\n", key, value->c_str(), s.ok(), pdb);
		fflush(stdout);
		data->strPtr = value;
		data->len = value->length();
		data->data = value->c_str();
		return data;
	}
}