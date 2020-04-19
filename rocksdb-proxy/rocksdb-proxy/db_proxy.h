#pragma once

#define ROCKSDB_LIBRARY_API __declspec(dllexport)
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

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
	extern ROCKSDB_LIBRARY_API void* CreateNewDB(const char* str, CustomString** cfArray, int cfLen, PtrArray* ptrArray, PtrArray* pNameArr);
	extern ROCKSDB_LIBRARY_API void* CreateDB(const char* str);
	extern ROCKSDB_LIBRARY_API bool	DBPut(void* db, CustomString* key, CustomString* value);
	extern ROCKSDB_LIBRARY_API bool DBColumnFamilyPut(void* pdb, void* phandle, CustomString* key, CustomString* value);
	extern ROCKSDB_LIBRARY_API DataString* DBGet(void* db, CustomString* key);
	extern ROCKSDB_LIBRARY_API DataString* DBColumnFamilyGet(void* pdb, void* phandle, const char* key);
	extern ROCKSDB_LIBRARY_API void ReleaseString(void* data);

	typedef struct {
		int a;
		int b;
	} Foo;

	extern ROCKSDB_LIBRARY_API int pass_array(Foo** in) {
		int i;
		int r = 0;

		for (i = 0; i < 2; i++) {
			r += in[i]->a;
			r *= in[i]->b;
		}
		return r;
	}
	extern ROCKSDB_LIBRARY_API int pass_str(CustomString** in, int len) {
		printf("pass_str:len=%d\n", len);
		for (int i = 0; i < len; i++) {
			printf("pass_str:%s\n", in[i]->str);
		}
		fflush(stdout);
		return len;
	}
	extern ROCKSDB_LIBRARY_API int pass_ptr(PtrArray* ptr) {
		//printf("pass_str:len=%d\n", ptr->len);
		auto rptr = (CustomString**)ptr->ptr;
		for (int i = 0; i < ptr->len; i++) {
			//printf("pass_ptr:%s\n", rptr[i]->str);
		}
		//fflush(stdout);
		return 1;

	}
	extern ROCKSDB_LIBRARY_API int pass_cstringptr(PtrArray* ptr) {
		auto cs = (CustomString**)ptr->ptr;
		return (*cs)->len;

	}
	extern ROCKSDB_LIBRARY_API void* MyMalloc(int size) {
		auto addr = malloc(size);
		//printf("MyMalloc:%ld\n", addr);
		//fflush(stdout);
		return addr;
	}
	extern ROCKSDB_LIBRARY_API void MyFeee(void* addr) {
		//printf("MyFeee:%ld\n", addr);
		//fflush(stdout);
		free(addr);
	}

	extern ROCKSDB_LIBRARY_API int pass_cstringptr2(CustomString** ptr, int len) {
		return (*ptr)->len;

	}
}