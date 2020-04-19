#pragma once

#include <stdio.h>
#include <stdlib.h>

typedef struct {
    void* strPtr;
    const char* Data;
    int Len;
} DataString, *DataStringPtr;

typedef struct {
    char* str;
    int len;
} CustomString;
typedef struct {
    void* ptr;
    int len;
} PtrArray;
typedef struct {
    CustomString** ptr;
    int len;
} PtrStrArray;
void Test();
int GetValue();
void* CreateNewDB(const char* str, CustomString** newcf, int cfLen, PtrArray* ptrArray, PtrArray* ptrNames);
void* CreateDB(const char* str);
int DBPut(void* db, CustomString* key, CustomString* value);
int DBColumnFamilyPut(void* pdb, void* phandle, CustomString* key, CustomString* value);
DataString* DBGet(void* db, CustomString* key);
DataString* DBColumnFamilyGet(void* pdb, void* phandle, const char* key);
void ReleaseString(void* data);

typedef struct {
    int a;
    int b;
} Foo;

typedef struct {
    void** ptr;
    int len;
} PtrStrNewArray;
int pass_array(Foo **in);
int pass_str(CustomString** in, int len);
int pass_ptr(PtrArray* ptr);
int pass_cstringptr(PtrStrNewArray* ptr);

void* MyMalloc(int size);
void MyFeee(void* addr);
int pass_cstringptr2(CustomString** ptr, int len);
//{
//    int i;
//    int r = 0;
//
//    for(i = 0; i < 2; i++) {
//        r += in[i]->a;
//        r *= in[i]->b;
//    }
//    return r;
//}
//void Test();
//int GetValue();
//void* CreateNewDB(const char* str, const char* cf);
//void* CreateDB(const char* str);
//int DBPut(void* db, const char* key, const char* value);
//DataString* DBGet(void* db, const char* key);
//void ReleaseString(void* data);