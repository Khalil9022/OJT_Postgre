select * from skala_rental_tabs;
select * from customer_data_tabs;
select * from branch_tabs;
select * from loan_data_tabs;
SELECT * FROM mst_company_tabs;
SELECT * FROM staging_customers;
SELECT * from staging_errors;
SELECT * FROM vehicle_data_tabs;

UPDATE customer_data_tabs
	SET approval_status='0'
	WHERE custcode='006A092023010000000003';

