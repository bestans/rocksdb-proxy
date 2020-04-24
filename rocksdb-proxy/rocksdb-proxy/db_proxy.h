#pragma once

#define ROCKSDB_LIBRARY_API __declspec(dllexport)
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

namespace rocksdb {
	class DB;
	class ColumnFamilyHandle;
}
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
	#pragma pack(1)
	struct PtrArray {
		void** ptr;
		int len;
	};
	#pragma pack()

	extern ROCKSDB_LIBRARY_API void Test();
	extern ROCKSDB_LIBRARY_API int GetValue();
	extern ROCKSDB_LIBRARY_API void* CreateNewDB(const char* str, char** cfArray, int* lenList, int cfLen, PtrArray* ptrArray, PtrArray* pNameArr);
	extern ROCKSDB_LIBRARY_API void* CreateDB(const char* str);
	extern ROCKSDB_LIBRARY_API bool	DBPut(void* db, CustomString* key, CustomString* value);
	extern ROCKSDB_LIBRARY_API bool DBColumnFamilyPutOld(void* pdb, void* phandle, CustomString* key, CustomString* value);
	extern ROCKSDB_LIBRARY_API bool DBColumnFamilyPut(rocksdb::DB* pdb, rocksdb::ColumnFamilyHandle* phandle, char* key, int keylen, char* value, int valuelen);
	extern ROCKSDB_LIBRARY_API DataString* DBGet(void* db, CustomString* key);
	extern ROCKSDB_LIBRARY_API DataString* DBColumnFamilyGetOld(void* pdb, void* phandle, CustomString* key);
	extern ROCKSDB_LIBRARY_API DataString* DBColumnFamilyGet(rocksdb::DB* pdb, rocksdb::ColumnFamilyHandle* phandle, char* key, int keylen);
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
	extern ROCKSDB_LIBRARY_API int pass_ptr(PtrArray* ptr, int value) {
		//printf("pass_str:len=%d\n", ptr->len);
		auto rptr = (CustomString**)ptr->ptr;
		for (int i = 0; i < ptr->len; i++) {
			if (value == 1) {
				printf("pass_ptr:%s\n", rptr[i]->str);
				fflush(stdout);
			}
		}
		fflush(stdout);
		return 1;

	}
	extern ROCKSDB_LIBRARY_API int pass_cstringptr(PtrArray* ptr) {
		auto cs = (CustomString**)ptr->ptr;
		return (*cs)->len;

	}

	int gvalue = 0;
	extern ROCKSDB_LIBRARY_API void SetValue(int value) {
		gvalue = value;
	}
	extern ROCKSDB_LIBRARY_API void* MyMalloc(int size) {
		auto addr = malloc(size);
		if (gvalue == 1) {
			printf("MyMalloc:%ld,%d,%d\n", addr, size, sizeof(PtrArray));
			fflush(stdout);
		}
		return addr;
	}
	extern ROCKSDB_LIBRARY_API void MyFeee(void* addr) {
		if (gvalue == 1) {
			printf("MyFeee:%ld\n", addr);
			fflush(stdout);
		}
		free(addr);
	}

	extern ROCKSDB_LIBRARY_API int pass_cstringptr2(CustomString** ptr, int len) {
		return (*ptr)->len;

	}
	extern ROCKSDB_LIBRARY_API int pass_goptrarr(PtrArray* ptr) {
		//auto cs = (char*)ptr->ptr;
		//printf("strlen=%d,len=%d\n", strlen(cs), ptr->len);
		//fflush(stdout);
		return ptr->len;
	}
	extern ROCKSDB_LIBRARY_API int pass_goptrarr2(PtrArray* ptr) {
		auto newptr = (PtrArray*)malloc(sizeof(PtrArray));
		newptr->len = 1;
		newptr->ptr = nullptr;
		ptr->ptr = (void**)newptr;
		ptr->len = 1;
		//auto cs = (char*)ptr->ptr;
		//printf("ptr=%ld\n", newptr->ptr);
		//fflush(stdout);
		return ptr->len;
	}
	extern ROCKSDB_LIBRARY_API int pass_goptrarr3(PtrArray* ptr) {
		auto newptr = (PtrArray*)malloc(sizeof(PtrArray));
		newptr->len = 2;
		auto str = (char*)malloc(2 * sizeof(char));
		str[0] = '1';
		str[1] = '2';
		newptr->ptr = (void**)str;
		ptr->ptr = (void**)newptr;
		ptr->len = 2;
		//auto cs = (char*)ptr->ptr;
		//printf("ptr=%ld\n", newptr->ptr);
		//fflush(stdout);
		return 1;
	}
	extern ROCKSDB_LIBRARY_API int pass_goptrarrarr(PtrArray* ptr) {
		auto cs = (PtrArray**)ptr->ptr;
		for (int i = 0; i < ptr->len; i++) {
			printf("str=%*.s\n", cs[i]->len, (char*)(cs[i]->ptr));
		}
		fflush(stdout);
		return ptr->len;
	}
	extern ROCKSDB_LIBRARY_API PtrArray* pass_gogetptr() {
		auto ptr = (PtrArray*)malloc(sizeof(PtrArray));
		ptr->ptr = (void**)malloc(100);
		ptr->len = 1;
		//auto cs = (char*)ptr->ptr;
		//printf("strlen=%d,len=%d\n", strlen(cs), ptr->len);
		//fflush(stdout);
		return ptr;
	}
	extern ROCKSDB_LIBRARY_API int pass_string(PtrArray* ptr, PtrArray* sstr, int value) {
		auto arr = (PtrArray**)ptr->ptr;
		for (int i = 0; i < ptr->len; i++) {
			if (value == 1) {
				printf("pass_string:%s,%d\n", (char*)(arr[i]->ptr), arr[i]->len);
				fflush(stdout);
			}
		}
		if (value == 1) {
			printf("pass_string_single:%s,%d\n", (char*)sstr->ptr, sstr->len);
			fflush(stdout);
		}
		return 1;
	}
	extern ROCKSDB_LIBRARY_API int pass_string2(char** ptr, int* lenList, int total, PtrArray* sstr, int value) {
		auto arr = ptr;
		for (int i = 0; i < total; i++) {
			if (value == 1) {
				printf("pass_string:%s,%d\n", arr[i], lenList[i]);
				fflush(stdout);
			}
		}
		if (value == 1) {
			printf("pass_string_single:%s,%d\n", (char*)sstr->ptr, sstr->len);
			fflush(stdout);
		}
		return 1;
	}
	extern ROCKSDB_LIBRARY_API int pass_str_direct(char* str, int len) {
		if (gvalue == 1) {
			printf("pass_str_direct:%s,%d\n", str, len);
			fflush(stdout);
			printf("pass_str_direct:%.*s\n", len, str);
			fflush(stdout);
		}
		return 1;
	}
}