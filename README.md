[![Build Status](https://travis-ci.com/miguelramirez93/mutant_detector.svg?branch=master)](https://travis-ci.com/miguelramirez93/mutant_detector)[![Coverage Status](https://coveralls.io/repos/github/miguelramirez93/mutant_detector/badge.svg?branch=master)](https://coveralls.io/github/miguelramirez93/mutant_detector?branch=master)
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

NOTE: Dockefile will run test coverage for you at build step ;)

After your environment is ready, you can navigate to:
- http://localhost:8080/swagger/index.html to see swagger documentation


current deployed env:
- https://mutant-detector-ml-2.herokuapp.com/swagger/index.html
