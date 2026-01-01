#include <iostream>
#include <string>
#include <cstdio>
#include <rocksdb/options.h>
#include <rocksdb/db.h>
#include <rocksdb/utilities/optimistic_transaction_db.h>
#include <rocksdb/utilities/transaction.h>
#include <rocksdb/slice.h>
#include "glue.h"

using ROCKSDB_NAMESPACE::DB;
using ROCKSDB_NAMESPACE::OptimisticTransactionDB;
using ROCKSDB_NAMESPACE::OptimisticTransactionOptions;
using ROCKSDB_NAMESPACE::Options;
using ROCKSDB_NAMESPACE::ReadOptions;
using ROCKSDB_NAMESPACE::Snapshot;
using ROCKSDB_NAMESPACE::Status;
using ROCKSDB_NAMESPACE::Transaction;
using ROCKSDB_NAMESPACE::WriteOptions;

///////////////////////////////
// Global variables
///////////////////////////////
Options options;
DB* db;
OptimisticTransactionDB* txn_db;
WriteOptions writeOptions;
ReadOptions readOptions;
OptimisticTransactionOptions txnOptions;

/**
 * @brief Creates/Initializes database
 * 
 * Basic setup used to create a DB if it doesn't exist and opens the database
 * 
 * @param dbPath Path to DB files
 * @return integer specifying status of init function 
 * 
 */
int init(const char* dbPath){
    options.create_if_missing = true;

    std::string str(dbPath);
    Status s = OptimisticTransactionDB::Open(options, str, &txn_db);
    if(!s.ok()){
        return 1; 
    }
    db = txn_db->GetBaseDB();
    return 0;
}

/**
 * @brief Inserts data into the database
 * 
 * Inserts the given value into the database using the supplied key
 * 
 * @param key The key that is associated with the value
 * @param value The value to be saved in the database
 * @return integer specifying the status of the insert 
 * 
 */
int insert(const char* key, const char* value){
    Transaction* txn = txn_db->BeginTransaction(writeOptions);
    assert(txn);

    // TODO: Need to generate random key (UUID)
    Status s = txn->Put(key, value);
    if(!s.ok()){
        return 1;
    }

    s = txn->Commit();
    delete txn;
    if(!s.ok() || s.IsBusy()){
        return 1;
    }

    return 0;
}

/**
 * @brief Queries data from the database
 * 
 * Queries based on the supplied key
 * 
 * @param key The key to query on
 * @return char* pointing to the value 
 * 
 */
char* query(const char* key){
    // TODO: Use iterators to filter data?
    std::string value; 
    Status s = db->Get(readOptions, key, &value);
    if(!s.ok()){
        return nullptr;
    }

    // TODO: Update so we don't have to use malloc
    char* output = (char*)malloc(10);
    std::strcpy(output, value.c_str());
    printf(value.c_str());
    return output;
}

// Uncomment the following if you want to test without building Go code.
// int main(){
//     std::string value = "MyKey";
//     init("/tmp/rocksdb_txn");
//     insert("TestKey", "Test");
//     std::cout << query(value.c_str()) << std::endl;
//     return 0;
// }
