-- Revert city_display_name back to city values

UPDATE stops SET city_display_name = city;
