#include "pch.h"
#include "db_proxy.h"
#include <stdio.h>
#include "rocksdb/db.h"
#include <iostream>

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

	void* CreateNewDB(const char* str, CustomString** cfArray, int cfLen, PtrArray* ptrArray, PtrArray* pNameArr) {
		printf("CreateNewDB:%s,%d\n", str, cfLen);
		fflush(stdout);
		DB* db = nullptr;
		Options options;
		// Optimize RocksDB. This is the easiest way to get RocksDB to perform well
		options.IncreaseParallelism();
		options.OptimizeLevelStyleCompaction();
		// create the DB if it's not already present
		options.create_if_missing = true;
		options.create_missing_column_families = true;

		// open DB
		std::vector<ColumnFamilyDescriptor> column_families;
		// have to open default column family
		column_families.push_back(ColumnFamilyDescriptor(
			kDefaultColumnFamilyName, ColumnFamilyOptions()));

		for (int i = 0; i < cfLen; i++) {
			std::cout << cfArray[i] << std::endl;
			//column_families.push_back(ColumnFamilyDescriptor(
			//	std::string(newcf), ColumnFamilyOptions()));
		}
		fflush(stdout);
		// open the new one, too
		std::vector<ColumnFamilyHandle*> handles;
		auto s = DB::Open(options, std::string(str), column_families, &handles, &db);
		std::cout << "handle size " << handles.size() << std::endl;
		if (!s.ok()) {
			printf("open failed:%s\n", s.ToString().c_str());
			fflush(stdout);
			return db;
		}
		fflush(stdout);
		
		ptrArray->ptr = (void**)malloc(sizeof(void*) * handles.size());
		ptrArray->len = (int)handles.size();
		pNameArr->ptr = (void**)malloc(sizeof(void*) * handles.size());
		pNameArr->len = (int)handles.size();
		for (int i = 0; i < (int)handles.size(); i++) {
			ptrArray->ptr[i] = handles[i];

			//handle名字
			auto newStr = (CustomString*)malloc(sizeof(CustomString));
			newStr->len = sizeof(char) * handles[i]->GetName().size();
			newStr->str = (char*)malloc(newStr->len);
			strcpy_s(newStr->str, newStr->len, handles[i]->GetName().c_str());
			pNameArr->ptr[i] = newStr;
			printf("cf:%s\n", handles[i]->GetName().c_str());
		}

		fflush(stdout);
		return db;
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
	bool DBPut(void* pdb, CustomString* key, CustomString* value) {
		auto db = (DB*)pdb;
		auto s = db->Put(WriteOptions(), std::string(key->str, key->len), std::string(value->str, value->len));
		printf("dbput:key=%s,value=%s,ok=%d,ptr=%d\n", key->str, value->str, s.ok(), pdb);
		fflush(stdout);
		return s.ok();
	}
	bool DBColumnFamilyPut(void* pdb, void* phandle, CustomString* key, CustomString* value) {
		auto db = (DB*)pdb;
		auto handle = (ColumnFamilyHandle*)phandle;
		auto s = db->Put(WriteOptions(), handle, std::string(key->str, key->len), std::string(value->str, value->len));
		printf("DBColumnFamilyPut:key=%s,value=%s,ok=%d,ptr=%d\n", key, value, s.ok(), pdb);
		fflush(stdout);
		return s.ok();
	}

	void ReleaseString(void* data) {
		printf("ReleaseString,ptr=%d\n", data);
		fflush(stdout);
		auto pData = (DataString*)data;
		delete (std::string*)pData->strPtr;
		delete pData;
	}
	DataString* DBGet(void* pdb, CustomString *key) {
		auto db = (DB*)pdb;
		auto data = new DataString();
		auto value = new std::string();
		auto s = db->Get(ReadOptions(), std::string(key->str, key->len), value);
		printf("DBGet:key=%s,value=%s,ok=%d,ptr=%d\n", key, value->c_str(), s.ok(), pdb);
		fflush(stdout);
		data->strPtr = value;
		data->len = value->length();
		data->data = value->c_str();
		return data;
	}
	DataString* DBColumnFamilyGet(void* pdb, void* phandle, const char* key) {
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