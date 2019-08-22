Installation/SetUp

    Install the package: go get -u github.com/gorilla/mux
    This helps in creating go routers

    Installing go-jwt package:
    go get github.com/dgrijalva/jwt-go

    PQ package: which helps us connect to PostGre DB
    go get github.com/lib/pq

    BCRYPT package: This is part of the go installation
    or else : go get golang.org/x/crypto/bcrypt

    Go Spew: Could be used for printing detailed struct formats
    go get -u github.com/davecgh/go-spew/spew


JWT

    JSON  Web Token
    Means of exchanging information between 2 parties
    Digitally Signed

    Structure: {Base64 encoder Header}.{Base64 encoder Payload}.{Signature}
    Header: Consists of Alg &, Token Type --> {"alg": "HS256", "type":"JWT"}
    Payload: Payload carry the claims
        - User &, additional data such as expiry
        - Types of claims: Registered, Public &, Private
        - eg: {"email":"abc@abc.com", "Issuer": "course"}
    Signature: Computed from Header, Payload &, a secret 
        - Generated by the algo mentioned in the header
        - Digitally signed using a secret only known to the developer
        - Cannot be decrypted unless you have the access to secret

PostGres SQL

    For the tutorial we will PostGresSQL as a service.
    URL: https://www.elephantsql.com
    Create an instance [Free one] ->
    To create a user table:
        create table users (
        id serial primary key,
        email text not null unique,
        password text not null
        );
    To insert records:
        insert into users (email, password) values ('abc@gmail.com', '123456');
        insert into users (email, password) values ('xyz@gmail.com', '123456');
    To select records:
        SELECT * FROM "public"."users"
