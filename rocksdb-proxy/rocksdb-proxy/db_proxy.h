#pragma once

#define ROCKSDB_LIBRARY_API __declspec(dllexport)


extern "C" {
	struct DataString {
		void* strPtr;
		const char* data;
		int len;
	};

	extern ROCKSDB_LIBRARY_API void Test();
	extern ROCKSDB_LIBRARY_API int GetValue();
	extern ROCKSDB_LIBRARY_API void* CreateDB(const char* str);
	extern ROCKSDB_LIBRARY_API bool	DBPut(void* db, const char* key, const char* value);
	extern ROCKSDB_LIBRARY_API void* DBGet(void* db, const char* key);
	extern ROCKSDB_LIBRARY_API void ReleaseString(void* data);
}