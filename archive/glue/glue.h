#pragma once

#ifdef __cplusplus
extern "C" {
#endif

int init(const char* dbPath);
int insert(const char* key, const char* value);
char* query(const char* key);

#ifdef __cplusplus
}
#endif