## Document management

## Create database
```text
docker run --name docmgdb -e MYSQL_ROOT_PASSWORD=123456 mysql:latest

mysql -u root -p
mysql> create user 'docmguser'@'%' identified by '123456';
mysql> create database docmg;
mysql> grant all privileges on docmg.* to 'docmguser'@'%';
mysql> flush privileges;
```

## Aggregate pattern
- The pattern's real purpose is to help you ensure transactional consistency between the entities 
- Root aggregate is an entity, that is used by external object
- One aggregate per transaction
    - There's a good rule for working with aggregates 
    - Prevent failed concurrency transactions
- Keep the aggregate small
- Eventual consistency 
- Reference by ID 
## Technical Support or Questions
If you have questions or need help integrating the product please "thanh29695@gmail.com" instead of opening an issue.

