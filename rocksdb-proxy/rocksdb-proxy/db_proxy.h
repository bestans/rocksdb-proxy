#pragma once

#define ROCKSDB_LIBRARY_API __declspec(dllexport)
#include <string.h>

extern "C" {
	struct DataString {
		void* strPtr;
		const char* data;
		int len;
	};
	struct CustomString {
		char* str;
		int len;
	};
	struct PtrArray {
		void** ptr;
		int len;
	};

	extern ROCKSDB_LIBRARY_API void Test();
	extern ROCKSDB_LIBRARY_API int GetValue();
	extern ROCKSDB_LIBRARY_API void* CreateNewDB(const char* str, PtrArray* newcf, PtrArray* ptrArray, PtrArray* ptrNames);
	extern ROCKSDB_LIBRARY_API void* CreateDB(const char* str);
	extern ROCKSDB_LIBRARY_API bool	DBPut(void* db, CustomString* key, CustomString* value);
	extern ROCKSDB_LIBRARY_API bool DBColumnFamilyPut(void* pdb, void* phandle, CustomString* key, CustomString* value);
	extern ROCKSDB_LIBRARY_API DataString* DBGet(void* db, CustomString* key);
	extern ROCKSDB_LIBRARY_API DataString* DBColumnFamilyGet(void* pdb, void* phandle, const char* key);
	extern ROCKSDB_LIBRARY_API void ReleaseString(void* data);
}