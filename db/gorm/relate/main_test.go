package relate

//var db *sql.DB
//
//var db *gorm.DB
//
//func TestMain(m *testing.M) {
//	dbName := "userserver"
//	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
//	pool, err := dockertest.NewPool("")
//	if err != nil {
//		log.Fatalf("Could not connect to docker: %s", err)
//	}
//
//	// pulls an image, creates a container based on it and runs it
//	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
//	if err != nil {
//		log.Fatalf("Could not start resource: %s", err)
//	}
//
//	defer func() {
//		// You can't defer this because os.Exit doesn't care for defer
//		if err := pool.Purge(resource); err != nil {
//			log.Fatalf("Could not purge resource: %s", err)
//		}
//	}()
//
//	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
//	if err := pool.Retry(func() error {
//		var err error
//		db, err = gorm.Open(mysql.Open(fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp"))))
//		if err != nil {
//			return err
//		}
//		return nil
//		//db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
//		//if err != nil {
//		//	return err
//		//}
//		//return db.Ping()
//	}); err != nil {
//		log.Fatalf("Could not connect to docker: %s", err)
//	}
//
//	//sqlReader, err := os.Open("./user-server.sql")
//	//if err != nil {
//	//	log.Fatal("open sql file err:", err)
//	//}
//
//	err = db.Exec("create database " + dbName).Error
//	if err != nil {
//		log.Fatal("create database err:", err)
//	}
//	//
//	//if err := pool.Retry(func() error {
//	//	var err error
//	//	db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/" + dbName, resource.GetPort("3306/tcp")))
//	//	if err != nil {
//	//		return err
//	//	}
//	//	return db.Ping()
//	//}); err != nil {
//	//	log.Fatalf("Could not connect to docker: %s", err)
//	//}
//	//
//	//sqlContent, err := io.ReadAll(sqlReader)
//	//sqls := strings.Split(string(sqlContent), ";")
//	//for _, s := range sqls {
//	//	_, err = db.Exec(s)
//	//	if err != nil {
//	//		log.Println("exec err :", err)
//	//	}
//	//}
//	m.Run()
//
//}
