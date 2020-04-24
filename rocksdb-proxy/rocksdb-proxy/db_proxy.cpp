#include "pch.h"
#include "db_proxy.h"
#include <stdio.h>
#include "rocksdb/db.h"
#include <iostream>
#pragma warning( disable : 4996)
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

	void* CreateNewDB(const char* str, char** cfArray, int* lenList, int cfLen, PtrArray* ptrArray, PtrArray* pNameArr) {
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
			column_families.push_back(ColumnFamilyDescriptor(
				std::string(cfArray[i]), ColumnFamilyOptions()));
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

		std::cout << "111111";
		fflush(stdout);
		ptrArray->ptr = (void**)malloc(sizeof(void*) * handles.size());
		ptrArray->len = (int)handles.size();
		std::cout << "22222222222";
		fflush(stdout);
		pNameArr->ptr = (void**)malloc(sizeof(void*) * handles.size());
		pNameArr->len = (int)handles.size();
		std::cout << "33333333";
		fflush(stdout);
		for (int i = 0; i < (int)handles.size(); i++) {
			ptrArray->ptr[i] = handles[i];

			std::cout << "444444444" <<handles[i]->GetName() << "," << handles[i] << std::endl;
			fflush(stdout);
			//handle名字
			//auto newStr = (CustomString*)malloc(sizeof(CustomString));
			//newStr->len = sizeof(char) * handles[i]->GetName().size();
			//newStr->str = (char*)malloc(newStr->len + 1);
			auto newStr = (char*)malloc(sizeof(char) * handles[i]->GetName().size() + 1);
			std::cout << "444444444aaa:" << newStr <<"," << handles[i]->GetName() << std::endl;
			fflush(stdout);
			std::strcpy(newStr, handles[i]->GetName().c_str());
			std::cout << "444444444bbbb" << handles[i]->GetName() << std::endl;
			fflush(stdout);
			pNameArr->ptr[i] = newStr;
			printf("cf:%s\n", handles[i]->GetName().c_str());
			fflush(stdout);
		}

		std::cout << "55555555555";
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
	bool DBColumnFamilyPutOld(void* pdb, void* phandle, CustomString* key, CustomString* value) {
		auto db = (DB*)pdb;
		auto handle = (ColumnFamilyHandle*)phandle;
		auto s = db->Put(WriteOptions(), handle, std::string(key->str, key->len), std::string(value->str, value->len));
		printf("DBColumnFamilyPutOld:key=%s,value=%s,ok=%d,ptr=%d,handle=%ld\n", key->str, value->str, s.ok(), pdb, phandle);
		fflush(stdout);
		return s.ok();
	}
	bool DBColumnFamilyPut(DB* pdb, ColumnFamilyHandle* phandle, char* key, int keylen, char* value, int valuelen) {
		auto s = pdb->Put(WriteOptions(), phandle,
			std::string(key, keylen), std::string(value, valuelen));
		if (gvalue == 1) {
			printf("DBColumnFamilyPut:key=%*.s,value=%*.s,ok=%d,ptr=%d,handle=%ld\n", keylen, key, valuelen, value, s.ok(), pdb, phandle);
			fflush(stdout);
		}
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
	DataString* DBColumnFamilyGetOld(void* pdb, void* phandle, CustomString* key) {
		auto db = (DB*)pdb;
		auto data = new DataString();
		auto value = new std::string();
		auto s = db->Get(ReadOptions(), (ColumnFamilyHandle*)phandle, std::string(key->str, key->len), value);
		printf("DBColumnFamilyGetOld:new:key=%s,value=%s,ok=%d,ptr=%d,handle=%d,err=%s\n", key->str, value->c_str(), s.ok(), pdb, phandle, s.ToString().c_str());
		fflush(stdout);
		data->strPtr = value;
		data->len = value->length();
		data->data = value->c_str();
		return data;
	}
	DataString* DBColumnFamilyGet(DB* pdb, ColumnFamilyHandle* phandle, char* key, int keylen) {
		auto data = new DataString();
		auto value = new std::string();
		auto s = pdb->Get(ReadOptions(), phandle, std::string(key, keylen), value);
		if (gvalue == 1) {
			printf("DBColumnFamilyGet:new:key=%*.s, key=%s,keylen=%d,value=%s,ok=%d,ptr=%d,handle=%d,err=%s\n", keylen, key, key, keylen, value->c_str(), s.ok(), pdb, phandle, s.ToString().c_str());
			fflush(stdout);
		}
		data->strPtr = value;
		data->len = value->length();
		data->data = value->c_str();
		return data;
	}
}