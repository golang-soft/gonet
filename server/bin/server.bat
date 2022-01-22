start "center" server.exe "center"
start "worlddb" server.exe "worlddb"
start "world" server.exe "world"
timeout /nobreak /t 4
start "world-1" server.exe "world"
start "account" server.exe "account"

start "netgate" server.exe "netgate"

start "login" server.exe "login"
start "zone" server.exe "zone"
start "grpcserver" server.exe "grpcserver"


timeout /nobreak /t 4
start "account-1" server.exe "account"

timeout /nobreak /t 4
start "netgate-1" server.exe "netgate"
timeout /nobreak /t 4
start "netgate-2" server.exe "netgate"