@ echo off

SET THIS_DIR=%~dp0
SET TEST_DIR=%~dp0test-data

del readYmeta.exe
del *.pdf
del %TEST_DIR%\*.pdf

PAUSE

go build readYmeta.go



readYmeta.exe
readYmeta.exe %TEST_DIR%\yoda-metadata[blank].json
readYmeta.exe %TEST_DIR%\yoda-metadata[test].json
readYmeta.exe %TEST_DIR%\yoda-metadata[uu011].json
readYmeta.exe %TEST_DIR%\yoda-metadata[uu012].json
readYmeta.exe %TEST_DIR%\yoda-metadata[uu013].json

dir *.pdf
dir %TEST_DIR%\*.pdf

