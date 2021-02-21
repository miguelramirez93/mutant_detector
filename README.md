[![Build Status](https://travis-ci.com/miguelramirez93/mutant_detector.svg?branch=master)](https://travis-ci.com/miguelramirez93/mutant_detector)

# mutant_detector

This is a sample project with clean architecture applied with go

go libs used:
- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://gorm.io/docs/index.html)
- [testify](https://github.com/stretchr/testify)

To deploy localy:
1. download and install [tilt](https://docs.tilt.dev/install.html)
2. Install docker and docker-compose
3. Run ```tilt up --hud=true``` at the root of this project

After your environment is ready, you can navigate to:
- http://localhost:8080/swagger/index.html to see swagger documentation


current deployed env:
- https://mutant-detector-v1.herokuapp.com/swagger/index.html