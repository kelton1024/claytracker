INSERT INTO ranges(name, address1, address2, city, state_id, zipcode, lat, lng) 
VALUES ($1, $2, $3, $4, (select state_id from states where name=UPPER($5)), $6, $7, $8);