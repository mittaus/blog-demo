#wait for the SQL Server to come up
sleep 120s

# mkdir /var/opt/mssql/ReplData/
# chown mssql /var/opt/mssql/ReplData/
# chgrp mssql /var/opt/mssql/ReplData/

echo "running set up script"
#run the setup script to create the DB and the schema in the DB
/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P 123456Aa -d master -i db-init.sql