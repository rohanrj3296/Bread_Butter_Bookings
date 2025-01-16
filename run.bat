go build -o Bread_Butter.exe ./cmd/web/. 
Bread_Butter.exe -dbname=bread_butter_bookings -dbuser=postgres -dbpass=password -cache=false/true -production=false/true
