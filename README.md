# logger

### Feature

   * Info, Debug, Warn, Error, Trace, Fatal, Panic 
     
### Installation

       #GoModule
       $ go get -u github.com/<user>/<repo-name>
       
### Usage

##### Create instance
    
    //Create AppLogger instance
    `log := logger.NewDefaultLogger(userId) //userId can be nil`

##### Info level log

    //Create AppLogger instance
    `log := logger.NewDefaultLogger(userId) //userId can be nil`
     
     const TAG = "DATABASE_LAYER" //service name or package name
     log.Info(TAG,"log msg")    
##### Error level log

    //Create AppLogger instance
    `log := logger.NewDefaultLogger(userId) //userId can be nil`
     
     const TAG = "DATABASE_LAYER" //service name or package name
     err := errors.New("connection error")
     log.Error(TAG,"Error:",err)
##### Debug level log

    //Create AppLogger instance
    `log := logger.NewDefaultLogger(userId) //userId can be nil`
     
     const TAG = "DATABASE_LAYER" //service name or package name
     log.Debug(TAG,"debug log")

##### Warn level log

    //Create AppLogger instance
    `log := logger.NewDefaultLogger(userId) //userId can be nil`
     
     const TAG = "DATABASE_LAYER" //service name or package name
     log.Warn(TAG,"log msg")
##### Trace level log

    //Create AppLogger instance
    `log := logger.NewDefaultLogger(userId) //userId can be nil`
     
     const TAG = "DATABASE_LAYER" //service name or package name
     log.Trace(TAG,"log msg")
##### Fatal level log

    //Create AppLogger instance
    `log := logger.NewDefaultLogger(userId) //userId can be nil`
     
     const TAG = "DATABASE_LAYER" //service name or package name
     log.Fatal(TAG,"log msg")
#####Panic level log

    //Create AppLogger instance
    `log := logger.NewDefaultLogger(userId) //userId can be nil`
     
     const TAG = "DATABASE_LAYER" //service name or package name
     log.Panic(TAG,"log msg")
                       