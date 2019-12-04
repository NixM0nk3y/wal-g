package main

var Env = map[string]string{
	"STAGING_DIR":               "staging",
	"ENV_FILE":                  "staging/env.file",
	"CONTEXT_FILE":              "staging/context.file",
	"API_HOST":                  "api01.local",
	"API_HTTP_PORT":             "80",
	"WORKER_HOST_01":            "mongodb01",
	"WORKER_HOST_02":            "mongodb02",
	"TUNE_COMP_CONFIG":          "/staging/tune_config.json",
	"TUNE_COMP_CONFIG_STAGING":  "staging/tune_config.json",
	"DB_HOST":                   "database01.local",
	"DB_USER":                   "root",
	"DB_PASSWORD":               "rootpw",
	"DB_NAME":                   "loadtest",
	"DB_MONGO_PORT":             "27018",
	"TARGET_HOST":               "target01.local",
	"TARGET_PG_PORT":            "5432",
	"TARGET_PG_USER":            "admin",
	"TARGET_PG_PASSWORD":        "adminpw",
	"TARGET_PG_DBNAME":          "postgres",
	"TARGET_PGBENCH_INIT_ARGS":  "-i -s 10",
	"TARGET_SSH_PORT":           "22",
	"TARGET_SSH_USER_USER":      "comp_user",
	"TARGET_SSH_USER_PASSWORD":  "comp_pass",
	"TARGET_SSH_ADMIN_USER":     "root",
	"TARGET_SSH_ADMIN_PASSWORD": "rootpw",
	"TEST_ID":                   "13",
	"TEST_CLEANUP_DELAY":        "60",
	"MINIO_ACCESS_KEY": 		 "1NdvGnOio1ad3HFmWNae",
	"MINIO_SECRET_KEY": 		 "PZK1ZuHiVM7I8vzLfBeEf6yfElqrXrZdNfaPORIM",
	"WALG_S3_PREFIX": 		     "s3://dbaas/mongodb-backup/test_uuid/test_mongodb",
}